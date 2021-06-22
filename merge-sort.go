package main

import (
	"fmt"
)

func main() {
	a := []int{10, 2, 99, 56, 78, 1, 7}
	fmt.Println(mergeSort(a)) // [1 2 7 10 56 78 99]
}

func mergeSort(arr []int) []int {
	var mergeArr []int

	inputArrLen := len(arr)
	halfLenInputArr := inputArrLen / 2

	if inputArrLen >= 2 {
		leftSideArr := mergeSort(arr[:halfLenInputArr])
		rightSideArr := mergeSort(arr[halfLenInputArr:])
		lenLeftSideArr := len(leftSideArr)
		lenRightSideArr := len(rightSideArr)

		j, k := 0, 0
		for i := 0; i < inputArrLen; i++ {
			if lenLeftSideArr > j && lenRightSideArr > k && leftSideArr[j] < rightSideArr[k] {
				mergeArr = append(mergeArr, leftSideArr[j])
				j++
			} else {
				if lenRightSideArr == k {
					mergeArr = append(mergeArr, leftSideArr[j])
					j++
				} else {
					mergeArr = append(mergeArr, rightSideArr[k])
				}

				if lenRightSideArr > k {
					k++
				}
			}
		}

		return mergeArr
	} else {
		return arr
	}
}
