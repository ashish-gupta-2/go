package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type User struct {
	Id    int
	Email string
}

type CachedUser struct {
	User
	expiresInTime int64
}

type LocalCahce struct {
	users           map[int]CachedUser

	mutex sync.RWMutex
	wg    sync.WaitGroup
	stop chan struct{}
}

func NewLocalCache(cleanUpInterval time.Duration) *LocalCahce {

	lc := &LocalCahce{
		users:           make(map[int]CachedUser),
		stop:  make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanUpInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupCache(cleanUpInterval)
	}(cleanUpInterval)

	return lc

}

func (lc *LocalCahce) cleanupCache(cleanUpInterval time.Duration) {
	ticker := time.NewTicker(cleanUpInterval)
	for {
		select {
		case <-lc.stop:
			fmt.Println("Stopping the caching...")
			return
		case <-ticker.C:
			lc.mutex.Lock()
			for id, user := range lc.users {
				if user.expiresInTime <= time.Now().Unix() {
					delete(lc.users, id)
				}
			}
			lc.mutex.Unlock()
		}
	}
}

func (lc *LocalCahce) stopCache() {
	close(lc.stop)
	lc.wg.Wait()
}


func (lc *LocalCahce) readUser(id int) (User, error) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	cu, ok := lc.users[id]
	if !ok {
		return User{}, errors.New(fmt.Sprintf("User is not available in cache id %v", id))
	}

	return cu.User, nil

}

func (lc *LocalCahce) delete(id int) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	delete(lc.users, id)
}

func (lc *LocalCahce) updateCache(u User, expiresInTime int64) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	cUser := CachedUser{
		User:              u,
		expiresInTime: expiresInTime,
	}

	lc.users[u.Id] = cUser

}


func main() {
	// It will check every second in the the cache
	localCache := NewLocalCache(time.Second)

	user1 := User{
		Id:    1,
		Email: "test@gmail.com",
	}

	user2 := User{
		Id:    2,
		Email: "test1@gmail.com",
	}

	localCache.updateCache(user1, int64(time.Second*1))
	localCache.updateCache(user2, int64(time.Second*4))

	u1, err := localCache.readUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user 1 ", u1)
	}

	u2, err := localCache.readUser(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user 2 ", u2)
	}

	time.Sleep(time.Second * 3)

	fmt.Println("After Sleep......")

	u1, err = localCache.readUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user 1 ", u1)
	}

	u2, err = localCache.readUser(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user 2 ", u2)
	}

	localCache.stopCache()

}
