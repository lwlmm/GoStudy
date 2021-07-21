package main

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	maxHeap  []int
	heapSize int
	//记录了堆中实际的节点数，最后一个元素为this.maxHeap[this.realSize]
	realSize int
}

func Constructor(capacity int) MaxHeap {
	// 堆中的第一个下标将被空出，以便操作和记忆
	heap := make([]int, capacity+1)
	maxheap := MaxHeap{
		maxHeap:  heap,
		heapSize: capacity,
	}
	return maxheap
}

func (this *MaxHeap) Add(element int) {
	this.realSize++
	if this.realSize > this.heapSize {
		fmt.Println("Full heap, cannot add more elements.")
		this.realSize--
		return
	}
	// 将最后一个节点赋值
	this.maxHeap[this.realSize] = element
	// 随后不断地将节点与其父节点比较，直到小于父节点的值或成为堆顶元素（index == 1）
	index := this.realSize
	parent := index / 2
	for this.maxHeap[index] > this.maxHeap[parent] && index > 1 {
		this.maxHeap[index], this.maxHeap[parent] = this.maxHeap[parent], this.maxHeap[index]
		index = parent
		parent /= 2
	}
}

func (this *MaxHeap) Peak() int {
	return this.maxHeap[1]
}

func (this *MaxHeap) Pop() int {
	if this.realSize < 1 {
		fmt.Println("No element in this heap.")
		return math.MinInt64
	}
	// 将堆顶换成最后一个元素，并删除最后一个元素
	peak := this.maxHeap[1]
	this.maxHeap[1] = this.maxHeap[this.realSize]
	this.realSize--
	index := 1
	// 从堆顶开始不断与左右节点比较，与较大的节点交换
	// 注意这里的退出条件是成为一个元素或者成为子节点（即index > this.realSize/2）
	for index < this.realSize && index <= this.realSize/2 {
		left := index * 2
		right := left + 1
		// 当左右子节点中有一个大于index时，选取较大的进行交换
		if this.maxHeap[index] < this.maxHeap[left] || this.maxHeap[index] < this.maxHeap[right] {
			if this.maxHeap[left] > this.maxHeap[right] {
				this.maxHeap[left], this.maxHeap[index] = this.maxHeap[index], this.maxHeap[left]
				index = left
			} else {
				this.maxHeap[right], this.maxHeap[index] = this.maxHeap[index], this.maxHeap[right]
				index = right
			}
		} else {
			break
		}
	}
	return peak
}

func HeapSort(nums []int) []int {
	maxheap := Constructor(len(nums))
	for _, n := range nums {
		maxheap.Add(n)
	}
	for i := len(nums)-1; i >= 0; i-- {
		nums[i] = maxheap.Pop()
	}
	return nums
}

func main() {
	nums := []int {6,5,4,3,2,1}
	fmt.Println(HeapSort(nums))
}