package main

import (
	"fmt"
)

// 环形单向链表
type Node struct {
	no   int
	name string
	next *Node
}

// 添加数据到链表
func Add(head *Node, new *Node) {
	//判断是否是首结点
	if head.next == nil {
		head.no = new.no
		head.name = new.name
		head.next = head //形成一个环形
		return
	}
	//定义一个临时变量，帮忙找到环形的最后一个结点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = new
	new.next = head

}

// 删除数据
func DelNode(head *Node, id int) *Node {
	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("空链表无法删除")
		return head
	}
	//就只有一个结点
	if temp.next == head {
		if temp.no == id {
			temp.next = nil
		} else {
			fmt.Println("无该结点")
		}
		return head
	}

	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//包含两个以上的结点
	flag := true
	for {
		if temp.next == head { //如果到这，说明比较到最后一个，但是最后一个还未进行比较
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next
				//
			}
			helper.next = temp.next
			flag = false
			break
		}
		temp = temp.next     //移动
		helper = helper.next //移动【一旦找到，helper发生作用】
	}
	//还需要比较一次
	if flag {
		if temp.no == id {
			helper.next = temp.next
		} else {
			fmt.Println("无该结点")
		}
	}
	return head
}

// 显示队列
func ListSingleCircleLink(head *Node) {
	temp := head
	//空链表退出
	if temp.next == nil {
		return
	}
	for {
		fmt.Println(temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

/*
环形单向链表：
注意：一定要考虑删除的是头结点，那么就要更新头结点，并且返回头结点

思路分析：
1.先让temp指向head
2.让helper指向环形链表的最后
3.让temp和要删除的id进行比较，如果相同，则通过helper进行删除，【必须要考虑如果删除的就是头结点】
4.
*/
func main() {
	//1.创建一个队列
	head := &Node{}
	test1 := &Node{
		no:   1,
		name: "test1",
	}
	test2 := &Node{
		no:   2,
		name: "test2",
	}

	Add(head, test1)
	Add(head, test2)
	ListSingleCircleLink(head)
	//注意
	head = DelNode(head, 10)
	ListSingleCircleLink(head)
}
