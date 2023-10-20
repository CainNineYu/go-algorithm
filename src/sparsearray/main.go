package main

import "fmt"

// 稀疏数组
type valNode struct {
	row int
	col int
	val int //可以用interface{}
}

func main() {
	//1.创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子
	//2.输出看看原始的数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	// 3.转成稀疏数组
	/*
			1.遍历chessMap，如果发现有一个元素的值不为0，创建一个node结构体
		2.将其放入到对用的切片即可
	*/
	var sparseArr []valNode
	//标准的一个稀疏数组应该还有一个 记录元素的二维数组的规模（行和列和默认值）
	valNodeV := valNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNodeV)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个valNode 值节点
				valNodeV := valNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNodeV)
			}

		}
	}
	//输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：：")
	for i, node := range sparseArr {
		fmt.Printf("%d: %d %d %d \n", i, node.row, node.col, node.val)
	}
	//将这个稀疏数组存盘
	/*1.打开d:/chessmap.data =>恢复原始数组
	2.使用稀疏数组恢复
	*/
	// 恢复原始的数组
	var chessMap2 [11][11]int
	//遍历sparseArr[遍历文件每一行]
	for i, v := range sparseArr {
		if i != 0 { //跳过第一行
			chessMap2[v.row][v.col] = v.val
		}

	}
	//看看chessMap2 是不是恢复
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
