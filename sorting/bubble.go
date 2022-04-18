package sorting

import (
	"sort"
)

func BubbleSort(data sort.Interface) {
	length := data.Len()
	for i := 0; i < length-1; i++ {
		isChanged := false
		for j := 0; j < length-1-i; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

func BubbleSort2(data []int) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		isChanged := false
		for j := 0; j < length-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
}

func BubbleSort3(data []int) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
