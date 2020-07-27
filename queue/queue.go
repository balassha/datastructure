package main

import (
	"fmt"
	"sync"
)

type queue struct{
	elements []interface{}
	start,end int
	lock sync.RWMutex
}

func (q *queue)enque(el interface{}){
	q.lock.Lock()
	q.elements=append(q.elements,el)
	q.lock.Unlock()
}

func (q *queue)size()int{
	return len(q.elements)
}

func (q *queue)deque() interface{}{
	if len(q.elements)==0 {
		return nil
	}
	q.lock.Lock()
	el := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	q.lock.Unlock()
	return el
}

func (q *queue)front() interface{}{
	if len(q.elements)==0 {
		return nil
	}
	return q.elements[0]
}

func (q *queue)print(){
	q.lock.Lock()
	for _,v := range q.elements{
		fmt.Println(v)
	}
	q.lock.Unlock()
}


