package main

import (
	"fmt"
)

// 环形单向链表
type Node struct {
	no   int
	next *Node
}

// 添加数据到链表，num表示数量，返回为第一个指针，后续好标记
func Add(num int) *Node {
	first := &Node{} //空结点
	temp := &Node{}  //空结点
	if num < 1 {
		fmt.Println("数量不对")
	}
	//循环构建链表
	for i := 1; i <= num; i++ {
		boy := &Node{
			no: i,
		}
		//如果是第一个,构成循环链表需要一个辅助指针
		if i == 1 {
			first = boy
			temp = boy
			temp.next = first
		} else {
			temp.next = boy
			temp = boy
			temp.next = first
		}
	}
	return first
}

// 删除数据,头指针，k的数，以及m的值
func DelNode(head *Node, startNo int, countNum int) {

	helper := head
	//空链表
	if head.next == nil {
		fmt.Println("空链表无法删除")
		return
	}
	//遍历查看startNo要小于总数大于多少

	////就只有一个结点
	//if temp.next == head {
	//	if temp.no == id {
	//		temp.next = nil
	//	} else {
	//		fmt.Println("无该结点")
	//	}
	//	return head
	//}

	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//让first移动到startNo，删除则以first为准
	for i := 1; i <= startNo-1; i++ {
		head = head.next
		helper = helper.next
	}
	for {
		for i := 1; i <= countNum-1; i++ {
			head = head.next
			helper = helper.next
		}
		fmt.Printf("删除%d", head.no)
		head = head.next
		helper.next = head
		//退出条件，最后一个删除
		if helper == head {
			break
		}
	}
	fmt.Printf("最后删除%d", head.no)

	return
}

// 显示链表
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
环形单向链表：实际应用：Josephu约瑟夫问题：n个人围成一圈，约定编号为k（1<=k<=n）的人从1开始报数，数到m的那个人出列，他的下一位又从1开始报数，数到m的人又出列，循环，直到所有人出列为止，由此产生一个出队编号的序列
注意：一定要考虑删除的是头结点，那么就要更新头结点，并且返回头结点

思路分析：
1.5个人，n=5
2.第二个人开始报数（本身需要数一下），k=2，数三下，m=3
3.使用单向环形链表解决
4.
*/
func main() {
	//1.创建一个队列
	//head := &Node{}
	//test1 := &Node{
	//	no:   1,
	//}
	//test2 := &Node{
	//	no:   2,
	//	name: "test2",
	//}

	head := Add(5)
	ListSingleCircleLink(head)
	//注意
	DelNode(head, 2, 10)
	//ListSingleCircleLink(head)
}
