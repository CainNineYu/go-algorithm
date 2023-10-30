package main

import "fmt"

// 注意切片是值传递，go没有传引用的说法，但是会有传递指针代替，文章https://juejin.cn/post/6888117219213967368，
// 传递的是复制的指针，但是存指针的地址是不一样的，函数中可能修改行惨会影响实参，也可能不会，要根据切片中指向的数组是否修改巨顶
func bubbleSort(arr []int) {
	//完成第一轮排序，外层1次
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				//
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}

		}

	}

	fmt.Print(arr)
}

// 冒泡排序：1.长度-1次轮数进行比较，每一轮将会确定一个数的位置，2.每一轮比较次数在逐渐减少，3.当前面的数比后面的数大时，交换位置

func main() {
	//定义数组
	arr := []int{1, 2, 3, 10, 5, -1, 10, 1}
	bubbleSort(arr)
}
