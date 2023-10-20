# go-algorithm
golang algorithm and design mode

## sparsearray 稀疏数组
 实际需求，存储大量相同数据，只有部分数据不同；例如：五子棋存盘退出和续上盘的功能
 ### 处理方法
1.记录数组一共有几行几列，有多少个不同的值
2.把具有不同值的元素的行列记录在一个小规模的数组中，从而缩小程序的规模