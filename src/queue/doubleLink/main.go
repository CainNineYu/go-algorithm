package main

import "fmt"

//定义一个结点

type HeroNode struct {
	no   int
	name string    //信息
	pre  *HeroNode //表示指向前一个结点
	next *HeroNode //表示指向下一个结点
}

// 在末尾插入结点
func InsertHeroNode(head *HeroNode, newHead *HeroNode) {
	//先找到该链表的最后一个结点，再创建一个辅助结点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHead
	newHead.pre = temp
}

// 按照no排序插入结点
func InsertHeroNode2(head *HeroNode, newHead *HeroNode) {
	//先找到该链表的最后一个结点，再创建一个辅助结点
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHead.no {
			break
		} else if temp.next.no == newHead.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("having no")
		return
	} else {
		newHead.next = temp.next
		newHead.pre = temp
		if temp.next != nil {
			temp.next.pre = newHead
		}
		temp.next = newHead
	}

}

// 删除结点
func DelHeroNode(head *HeroNode, id int) {
	//先找到该链表的最后一个结点，再创建一个辅助结点
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}

	} else {
		fmt.Println("No id")
	}

}

// 显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Print("Node is null")
		return
	}
	//遍历链表，不知道链表大小，用for
	for {
		fmt.Printf("[%d]", temp.next.no)
		//判断链表是否是末尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

// 反向显示链表的所有结点信息
func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Print("Node is null")
		return
	}
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//遍历链表，不知道链表大小，用for
	for {
		fmt.Printf("[%d]", temp.no)
		//判断链表是否到了链表头部
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

/*
链表：链表是有序的列表
单向链表：为了对单向链表比较好操作增删改查，是带有头结点，作用是标识链表的头，本身这个结点不存放数据
注意：缺陷，1.查找只能是单向查找，
2.单向列表不能自我删除，只能依靠辅助结点

思路分析：
*/
func main() {
	//先创建一个头结点
	head := &HeroNode{}
	//创建一个新的结点
	head1 := &HeroNode{
		no:   1,
		name: "测试1", //信息
	}
	head2 := &HeroNode{
		no:   1,
		name: "测试1", //信息
	}
	//添加
	InsertHeroNode(head, head1)
	InsertHeroNode2(head, head2)
	DelHeroNode(head, 0)
	//显示
	ListHeroNode(head)
}
