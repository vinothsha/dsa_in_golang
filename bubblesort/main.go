package main

import "fmt"

func main() {
	arr := []int{99, 65, 32, 66, 455, 120, 654}
	Ascending(arr)
	Descending(arr)
}
func Ascending(arr []int) {
	for i:=0;i<len(arr)-1;i++{
		// fmt.Println("1 st ",len(arr)-1)
		for j:=0;j<len(arr)-i-1;j++{
			// fmt.Println("2 nd ",len(arr)-i-j)
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
func Descending(arr []int){
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]<arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}