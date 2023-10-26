package main

import "fmt"

// 注意切片是值传递，go没有传引用的说法，但是会有传递指针代替，文章https://juejin.cn/post/6888117219213967368，
// 传递的是复制的指针，但是存指针的地址是不一样的，函数中可能修改行惨会影响实参，也可能不会，要根据切片中指向的数组是否修改巨顶
// 注释：left最左边的下标，right最右边的下标
func quickSort(left int, right int, arr []int) {
	l := left
	r := right
	//中轴，支点（也可以找最后一个数），中间的话效果高点
	pivot := arr[(l+r)/2]
	//temp := 0
	//目标，将比pivot数小的放在左边，大的放在右边
	//for循环结束后，左右两边就搞定了
	for l < r {
		//从pivot左边找大于等于pivot的值
		for arr[l] < pivot {
			l++
		}
		//从pivot右边找小于等于pivot的值
		for arr[r] > pivot {
			r--
		}
		//表面此次分解任务已经完成
		if l >= r {
			break
		}
		//交换
		arr[l], arr[r] = arr[r], arr[l]
		//优化，相等则不交换
		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	//此次不用再比较，再移动一位就行了，没有这个则卡死在此处
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		quickSort(left, r, arr)
	}
	//向右递归
	if right > l {
		quickSort(l, right, arr)
	}
	fmt.Print(arr)
}

// 快速排序是对冒泡排序的一种改进：通过一趟排序将要排序的数据分割成独立的两部分，
// 其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列

func main() {
	//定义数组
	arr := []int{1, 2, 3, 10, 5, -1, 10, 1, 0, 999}
	quickSort(0, len(arr)-1, arr)
}
