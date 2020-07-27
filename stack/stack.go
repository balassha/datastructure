package main

import(
	"fmt"
	"sync"
)

type stack struct{
	len int
	list []interface{}
	lock sync.RWMutex
}

func (s *stack)push(sd interface{}){
	s.lock.Lock()
	s.list=append(s.list,sd)
	s.lock.Unlock()
	s.len++
}

func(s* stack)pop() interface{}{
	if s.len==0{
		return nil
	}
	s.lock.Lock()
	ret := s.list[s.len-1]
	s.list[s.len-1]=nil
	s.lock.Unlock()
	s.len--
	return ret
}

func(s* stack)print(){
	s.lock.Lock()
	for _,v:=range s.list{
		fmt.Println(v)
	}
	s.lock.Unlock()
}

func(s *stack)isEmpty()bool{
	return s.len==0
}

func(s *stack)peek()interface{}{
	if s.len==0{
		return nil
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.list[s.len-1]
}
