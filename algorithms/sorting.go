package algorithms

import "sort"

// 快速排序实现
func QuickSort(arr []int) []int {
	sort.Ints(arr)
	return arr
	// TODO: fix it
	// if len(arr) <= 1 {
	// 	return arr
	// }

	// pivot := arr[len(arr)/2]
	// var less, greater []int

	// for _, num := range arr[1:] {
	// 	if num <= pivot {
	// 		less = append(less, num)
	// 	} else {
	// 		greater = append(greater, num)
	// 	}
	// }

	// less = QuickSort(less)
	// greater = QuickSort(greater)

	// return append(append(less, pivot), greater...)
}
