package src

// 插入排序
func InsertSort(arr []interface{}) []interface{} {
	l := len(arr)
	for i := 1; i < l; i++ {
		flag := true
		for j := i; j > 0; j-- {
			if flag && insertSortCompare(arr[j-1], arr[j]) {
				insertSortExchange(arr, j-1, j)
			}else{
				flag = false
			}
		}
	}
	return arr
}

// 比较
func insertSortCompare(a interface{}, b interface{}) bool {
	return a.(int) > b.(int)
}

// 交换
func insertSortExchange(arr []interface{}, i1 int, i2 int) {
	arr[i1], arr[i2] = arr[i2], arr[i1]
}
