package src

// 冒泡排序
func BubbleSort(arr []interface{}) []interface{} {
	l := len(arr)
	for i := 1; i < l; i++ {
		flag := false // false:表示没有交换
		for j := 1; j < l; j++ {
			if bubbleSortCompare(arr[j-1], arr[j]) {
				flag = true
				bubbleSortExchange(arr, j-1, j)
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// 比较
func bubbleSortCompare(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

// 交换
func bubbleSortExchange(arr []interface{}, i1 int, i2 int) {
	arr[i1], arr[i2] = arr[i2], arr[i1]
}
