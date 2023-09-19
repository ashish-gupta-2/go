package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type user struct {
	id    int
	email string
}

type cachedUser struct {
	user       user
	expireTime int64
}

type localcache struct {
	users map[int]cachedUser
	mutex sync.RWMutex
	wg    sync.WaitGroup
	stop  chan struct{}
}

func NewLocalCache(refreshTime time.Duration) *localcache {
	lc := &localcache{
		users: make(map[int]cachedUser),
		stop:  make(chan struct{}),
	}
	lc.wg.Add(1)
	go func(refreshTime time.Duration) {
		defer lc.wg.Done()
		lc.refreshLocalCache(refreshTime)
	}(refreshTime)

	return lc
}

func (lc *localcache) refreshLocalCache(refreshTime time.Duration) {
	ticker := time.NewTimer(refreshTime)

	for {
		select {
		case <-lc.stop:
			fmt.Println("Stopping the caching...")
			return
		case <-ticker.C:
			lc.mutex.Lock()
			for i, user := range lc.users {
				if user.expireTime <= time.Now().Unix() {
					delete(lc.users, i)
				}
			}
			lc.mutex.Unlock()
		}
	}
}

func (lc *localcache) addInCache(cu cachedUser) {
	lc.users[cu.user.id] = cu
}

func (lc *localcache) stopCache() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *localcache) readUser(id int) (user, error) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	cu, ok := lc.users[id]
	if !ok {
		return user{}, errors.New(fmt.Sprintf("User is not available in cache id %v", id))
	}

	return cu.user, nil

}

func (lc *localcache) delete(id int) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	delete(lc.users, id)
}
