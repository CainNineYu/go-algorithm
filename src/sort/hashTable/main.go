package main

import (
	"fmt"
)

// 定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 不带表头，第一个结点就会存东西
type EmpLink struct {
	Head *Emp
}

// 添加数据结构的添加方法（添加员工的方法），保证编号是从小到大，
func (this *EmpLink) EmpLinkInsert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil //在cur前面
	//如果当前EmpLink就是一个空链表
	if this.Head == nil {
		this.Head = emp
		return
	}
	//如果不是一个空链表，给emp找到对应的位置并插入
	//让cur和emp比较，pre保持在cur前面
	for {
		if cur != nil {
			//要考虑id是否能够重复
			if cur.Id > emp.Id {
				//找到位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}
	//退出，我们将emp是否添加
	//if cur == nil {
	pre.Next = emp
	emp.Next = cur
	//}
}

// 显示当前链表信息
func (this *EmpLink) EmpLinkShowAll(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}
	//遍历当前链表，并显示数据，
	cur := this.Head
	for {
		if cur != nil {
			fmt.Println(cur.Id)
			cur = cur.Next
		} else {
			break
		}
	}
}

// 查找
func (this *EmpLink) EmpLinkFindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	//就是哈希函数
	fmt.Printf("链表号：%d\n", this.Id%7)
}

// 给hashTable添加
func (this *HashTable) HashTableInsert(emp *Emp) {
	//使用散列函数（通过id或者名字或者（id加名字）进行散列），确定该数据添加到哪个列表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表进行添加
	this.LinkArr[linkNo].EmpLinkInsert(emp)
}

// 显示HashTable所有信息
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].EmpLinkShowAll(i)
	}
}

// 显示HashTable所有信息
func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	//使用对应的链表进行添加
	return this.LinkArr[linkNo].EmpLinkFindById(id)
}

// 编写散列方法
func (this *HashTable) HashFun(id int) int {
	//得到的值就是链表的下标
	return id % 7
	//return id % 7 % 3//也有可能是二级链表
}

// hashTable：哈希表（散列表）----根据关键码值（key-value）而直接进行访问的数据结构，就是说，
// 他通过吧关键码值映射到表中一个位置来访问记录，以加快查找的速度，这个映射函数叫做散列函数，存放记录的数组叫做散列表
//注：1.存放的id先取模，按照哈希表的长度，再存放到哈希表中

// 思路分析：1.使用链表来实现哈希表，该链表不带表头，即链表的第一个结点就存放雇员信息
func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("===============菜单===========")
		fmt.Println("input输入菜单")
		fmt.Println("show显示")
		fmt.Println("find查找")
		fmt.Println("exit退出")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入id")
			fmt.Scanln(&id)
			fmt.Println("请输入名字")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.HashTableInsert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":

			fmt.Println("请输入id")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Println("不存在%d\n", id)
			} else {
				emp.ShowMe()
			}
		case "exit":
		default:
			fmt.Println("输入错误")
		}
	}
}
