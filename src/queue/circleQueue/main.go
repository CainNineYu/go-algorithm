package main

import (
	"errors"
	"fmt"
	"os"
)

// 环形队列
type CircleQueue struct {
	maxSize int    //4
	array   [4]int //数组=》环形队列
	head    int    //指向队列队首 默认0
	tail    int    //指向队列尾部 默认0
}

// 添加数据到队列
func (this *CircleQueue) Push(val int) (err error) {
	//先判断队列是否已满
	if this.IsFull() { //rear是队列尾部（含最后元素）
		return errors.New("queue full")
	}
	//注意：this.tail在队列的尾部，但是不包含最后的元素

	this.array[this.tail] = val                //把值留给尾部
	this.tail = (this.tail + 1) % this.maxSize //rear 后移
	return
}

// 从队列中取出数据
func (this *CircleQueue) Pop() (val int, err error) {
	//先判断队列是否为空
	if this.IsEmpty() { //队列为空
		return 0, errors.New("queue empty")
	}
	//取出，head指向队首，并且包含队首元素
	val = this.array[this.head]
	//注意是环形队列，不能this.head++
	this.head = (this.head + 1) % this.maxSize
	return val, err
}

// 判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	//设计一个辅助的变量，指向head
	tempHead := this.head
	for i := this.head; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	// 这是一个重要的环节
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

/*
环形队列：1.尾索引的下一个为头索引时表示队列满，即将队列容量空出一个作为约定，这个在做判断队列满的时候需要注意（tail+1）%maxSize==head 满  2.tail==head【空】

注意：1.尾部tail是没有包含元素的，默认环形队列的容量：maxSize-1

思路分析：
1.完成队列的基本操作：AddQueue//加入数据到队列(判断是否数据已满)；GetQueue//从队列取出数据（判断数据是否为空）；ShowQueue//显示队列
2.什么时候表示队列满，（tail+1）%maxSize=head
3.tail==head表示空
4.初始化时，tail=0，head=0
5.怎么统计该队列总共有多少元素：（tail+maxSize-head）%maxSize
*/
func main() {
	//1.创建一个队列
	queue := &CircleQueue{
		maxSize: 4,
		head:    0,
		tail:    0,
	}
	var key string
	var val int
	for {
		fmt.Println("1. 输入Add表示添加数据到队列")
		fmt.Println("2. 输入Get表示从队列获取数据")
		fmt.Println("3. 输入Show表示显示队列")
		fmt.Println("4. 输入Exit表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "Add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列OK")
			}
		case "Get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列取出一个数=", val)
			}

		case "Show":
			queue.ListQueue()
		case "Exit":
			os.Exit(0)
		}
	}
}
