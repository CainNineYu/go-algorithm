package main

func binaryFinde(arr []int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		println("No value")
		return
	}
	middle := (leftIndex + rightIndex) / 2
	if arr[middle] > findVal {
		binaryFinde(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		binaryFinde(arr, middle+1, rightIndex, findVal)
	} else {
		println("值为", middle)
	}

}

// 二分法查找：1.必须是一个有序数组，2.先找到中间的下标middle=（leftIndex+rightIndex）/2，然后让中间下标的值和需要查找的值进行比较
//3.中间值大于所求值，就向leftIndex----middleIndex-1，还有小于以及相等情况
//4.逻辑用递归进行（退出递归条件：leftIndex>rightIndex）

func main() {
	//定义数组
	arr := []int{1, 2, 3, 10, 5, -1, 10, 1}
	binaryFinde(arr, 0, len(arr)-1, 100)
}
