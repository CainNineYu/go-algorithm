package main

import "fmt"

// 指针要保证时同一个地图，使用引用，i j表示对地图的那个点进行测试
func SetWay(myMap *[8][7]int, i int, j int) bool {
	//结束是6，5
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 { //可以探测的点
			//假设这个点是可以通，但是需要探测，上下左右，假设错误则为false
			myMap[i][j] = 2
			//顺序必须是下右上左，不然路径全是2，不是最短
			if SetWay(myMap, i+1, j) { //向下
				return true
			} else if SetWay(myMap, i, j+1) { //向右
				return true
			} else if SetWay(myMap, i-1, j) { //向上
				return true
			} else if SetWay(myMap, i, j-1) { //向左
				return true
			} else { //都不通则为死路
				myMap[i][j] = 3
				return false
			}
		} else { //这个点不能进行探测，为1是墙
			return false

		}
	}
}

// 递归（逻辑重复需要用到递归）：1.创建一个函数时，就创建一个新的受保护的独立空间（新函数栈）
//
//	2.函数的局部变量是独立的，不会相互影响  3.递归必须向退出递归的条件逼近，否则就是无限递归
//	4.当一个函数执行完毕，或者遇到return就会返回，遵守谁调用，就将结果返回给谁，同时当函数执行完毕或者返回时，该函数本身也会被系统摧毁
func main() {
	//二维数组模拟迷宫
	//规则；1.为墙   0.没有走过的点  2.一个通路  3.走过的点，但是走不通
	var myMap [8][7]int //8行7列
	//最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//最左和最右设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	//设置路障
	myMap[3][1] = 1
	myMap[3][2] = 1

	SetWay(&myMap, 1, 1)
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}
}
