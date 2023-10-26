package main

import "fmt"

// 注意切片是值传递，go没有传引用的说法，但是会有传递指针代替，文章https://juejin.cn/post/6888117219213967368，
// 传递的是复制的指针，但是存指针的地址是不一样的，函数中可能修改行惨会影响实参，也可能不会，要根据切片中指向的数组是否修改巨顶
func insertingSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		//完成第一次，给第二个元素找到合适的位置并插入
		insertVal := arr[i]
		insertIndex := i - 1
		//从大到小，需要后移一位
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--
		}
		//插入数据
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}

	}

	fmt.Print(arr)
}

// 插入排序（比选择排序效率高），也属于内部排序法（要把数据全部加载到内存之中），是从预排序的数据中，按指定的规则选出某一元素，经过和其他元素重整，
// 再依照原则交换位置后，达到排序的目的（例如先选出最大的放在第一个，再从剩下的选出第最大的放在第二个）
// 注意：先把数据放在内存中然后进行比较，交换是很费时间的，尽量少交换，比较没有问题
func main() {
	//定义数组
	arr := []int{1, 2, 3, 10, 5, -1, 10, 1}
	insertingSort(arr)
}
