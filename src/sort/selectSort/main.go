package main

import "fmt"

func selectSort(arr *[5]int) {
	for j := 0; j < len(arr)-1; j++ {
		max := arr[j]
		maxIndex := j
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			//go的交换特有的
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}
	fmt.Print(arr)
}

// 选择排序，也属于内部排序法，是从预排序的数据中，按指定的规则选出某一元素，经过和其他元素重整，
// 再依照原则交换位置后，达到排序的目的（例如先选出最大的放在第一个，再从剩下的选出第最大的放在第二个）
// 注意：先把数据放在内存中然后进行比较，交换是很费时间的，尽量少交换，比较没有问题
func main() {
	//定义数组
	arr := [5]int{1, 2, 3, 10, 5}
	selectSort(&arr)
}
