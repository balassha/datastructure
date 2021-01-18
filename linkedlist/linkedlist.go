package main

import(
	"fmt"
)

type node struct{
	next *node
	data interface{}
}

type linkedList struct{
	head *node
	len int
}

func (l *linkedList) add(n *node){
	if l.head == nil {
		l.head = n
		l.len++
	}else{
		t:= l.head
		for t.next != nil {
			t = t.next
		}
		t.next = n
		l.len++
	}
}

func (l *linkedList)print(){
	t:=l.head
	for t!=nil{
		fmt.Printf("Data : %v\n",t.data)
		t = t.next
	}
}

func (l *linkedList)middle(){
	var s,f *node
	if l != nil && l.head != nil{
		s = l.head
		f = l.head
	}
	for f !=nil && f.next != nil {
		s=s.next
		f=f.next.next
	}
	if f.next != nil {
		fmt.Printf("Middle: %v & %v\n",s.data,s.next.data)
	}else{
		fmt.Printf("Middle: %v\n",s.data)
	}
}

func (l *linkedList) reverse(){
	var c,p,n *node
	c = l.head
	for c!=nil{
		n = c.next
		if n==nil{
			l.head = c
		}
		c.next=p
		p=c
		c=n
	}
}

func (l *List) HasCycles(){
	m := make(map[*node]bool) 
	for c := l.head;c !=nil; c = c.next {
		if m[c] {
			fmt.Println("Cycles Detected")
			return true
		}
		m[c] = true
	}
	return false
}

// Reverse the contiguous even numbers and reassign to same list 
(HackerEarth - https://www.hackerearth.com/practice/data-structures/linked-list/singly-linked-list/practice-problems/algorithm/reversed-linked-list-01b722df/)
func (l *linkedList) restore(){
	if l == nil {
        return
    }
    n := l.head
    for n != nil {
        if n.data.(int) % 2 == 0 {
            pt := n
            var t linkedList
            for n.data.(int) %2 == 0 {
				t.add(&node{data:n.data,})
				if n.next==nil{
					break
				}else{
					n=n.next
				}
            }
            t.reverse()
            tp:=t.head
            for tp != nil && pt != nil{
                pt.data = tp.data
                tp = tp.next
                pt = pt.next
			}
			if n.next==nil{
				break
			}
		}else{
			if n.next==nil{
				break
			}else{
				n=n.next
			}
		}

		}
}

func (l *linkedList) delete(n *node){
	t:=l.head
	for t!=nil && t.next!=nil{
		if t.next.data == n.data {
			t.next = t.next.next
		}else{
			t=t.next
		}
	}
}

func main(){
	fmt.Println("In Main")

	in := [9]int{2,18,24,3,5,7,9,6,12}
	var l linkedList
	for _,v := range in{
		l.add(&node{data:v,})
	}
	fmt.Println(l.len)
	l.print()
	l.middle()
	l.reverse()
	l.restore()
	l.print()
}
