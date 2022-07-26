package main

import "fmt"

func main() {
	arr := []int{78, 87, 98, 99, 120, 141, 157, 178, 187, 198, 205, 307}
	find := 307
	a := IterFind(arr, find, 0, len(arr)-1)
	b := RecursiveFind(arr, find, 0, len(arr)-1)
	fmt.Println(a,b)
}
func RecursiveFind(lis []int, f, low, high int) int {
	mid := low + (high-low)/2
	if lis[mid] == f {
		return mid
	} else if f > lis[mid] {
		return RecursiveFind(lis, f, mid+1, high)
	} else {
		return RecursiveFind(lis, f, low, mid-1)
	}
}
func IterFind(arr []int, find, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		// fmt.Println(mid)
		if arr[mid] == find {
			return mid
		} else if find > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}