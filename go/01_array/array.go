package _1_array

import (
	"github.com/pkg/errors"
	"fmt"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 *
 * Author: shuxun
 */

type Array struct {
	data []int
	len  uint
}

func NewArray(cap uint) *Array {
	if cap == 0 {
		return nil
	}
	array := &Array{
		data: make([]int, cap, cap),
		len:  0,
	}
	return array
}

func (this *Array) isIndexOutOfRange(index uint) bool {
	if index >= uint(cap(this.data)) {
		return true
	}
	return false
}

//插入数值到索引index上
func (this *Array) Insert(index uint, val int) error {
	if this.len == uint(cap(this.data)) {
		return errors.New("full array")
	}
	if this.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}
	if index > this.len {
		return errors.New("index is unsuitable")
	}
	//1 2 3 4 5 6
	for i := this.len; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	this.data[index] = val
	this.len++
	return nil
}

func (this *Array) InsertToTail(val int) error {
	return this.Insert(this.len, val)
}

func (this *Array) Find(index uint) (int,error) {
	if this.isIndexOutOfRange(index) {
		return 0,errors.New("out of index range")
	}
	return this.data[index],nil
}

func (this *Array) Delete(index uint) (int, error) {
	if this.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	if this.len == 0 {
		return 0, errors.New("array is nil")
	}
	if index > this.len-1 {
		return 0, errors.New("index beyond range")
	}
	val := this.data[index]
	for i := index; i < this.len-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.len--
	return val, nil
}

func (this *Array) Print() {
	var format string
	for i := uint(0); i < this.len; i++ {
		format += fmt.Sprintf("|+%v", this.data[i])
	}
	fmt.Println(format)
}
