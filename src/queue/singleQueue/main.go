package main

import (
	"errors"
	"fmt"
	"os"
)

// 队列
type Queue struct {
	maxSize int
	array   [4]int //数组=》模拟队列
	front   int    //表示指向队列首,但是并不是首位数据，需要加1才是
	rear    int    //表示指向队列尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxSize-1 { //rear是队列尾部（含最后元素）
		return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val
	return
}

// 从队列中取出数据
func (this *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if this.rear == this.front { //队列为空
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

// 显示队列，找到队首，然后到遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列的当前情况是：")
	// this.front 不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("arrary[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

/*
使用数组模拟队列：本身是有序列表，数据结构声明：maxSize为最大容量，front为队列前，rear为后的下标

非环形数组队列：存在的问题，取出数据之后，无法再循环添加数据，没有有效的利用数组空间
思路分析：1.创建一个数组array，作为队列的一个字段
2.front和rear初始化为-1
3.完成队列的基本操作：AddQueue//加入数据到队列(判断是否数据已满)；GetQueue//从队列取出数据（判断数据是否为空）；ShowQueue//显示队列
*/
func main() {
	//1.创建一个队列
	queue := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列OK")
			}
		case "Get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列取出一个数=", val)
			}

		case "Show":
			queue.ShowQueue()
		case "Exit":
			os.Exit(0)
		}
	}
}
