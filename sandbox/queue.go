package main

import "fmt"


func main(){
	q1 := New()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q1.Enqueue(4)
	q1.Enqueue(5)
	

	fmt.Println(q1.size())

	q1.Dequeue()

	fmt.Println(q1.size())

}

type Queue struct{
	data []interface{}
}

func New() *Queue{
	return &Queue{}
}


func(q *Queue) Enqueue(value interface{}){
	q.data = append(q.data, value)
}

func(q *Queue) Dequeue() interface{} {
	 var value interface{}
	 if len(q.data) > 0 {
		value = q.data[0]
		q.data[0] = nil
		q.data = q.data[1:]
	 }
	 return value
}

func (q *Queue) size() int {
	return len((q.data))
}





