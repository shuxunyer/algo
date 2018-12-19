package _4_stack

import "fmt"

type Browser struct {
	forward Stack
	back    Stack
}

func NewBrowser() *Browser {
	return &Browser{
		forward: NewLinkedListStack(),
		back:    NewArrayStack(),
	}
}

func (this *Browser) CanBack() bool {
	if this.back.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) CanForward() bool {
	if this.forward.IsEmpty() {
		return false
	}
	return true
}

func (this *Browser) Open(addr string) {
	fmt.Printf("Open new addr %+v\n", addr)
	this.forward.Flush()
}

func (this *Browser) PushBack(add string){
	this.back.Push(add)
}


func (this *Browser) Forward(){
	if this.forward.IsEmpty(){
		return
	}
	top :=this.forward.Pop()
	this.back.Push(top)
	fmt.Printf("forward to %+v\n",top)
}

func (this *Browser) Back(){
	if this.back.IsEmpty(){
		return
	}
	top :=this.back.Pop()
	this.forward.Push(top)
	fmt.Printf("back to %+v\n",top)
}
