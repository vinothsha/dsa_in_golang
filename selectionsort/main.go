package main
import "fmt"
func main(){
	arr:=[]int32{88,97,34,62,345,12,456,43,765}
	AscendingOrder(arr)
}

func AscendingOrder(arr []int32){
	
	 for i:=0;i<len(arr)-1;i++{
		min:=i
		for j:=i+1;j<len(arr);j++{
			if arr[min]>arr[j]{
				min=j
			}	
		}
		arr[min],arr[i]=arr[i],arr[min]
	 }
		fmt.Println(arr)
}

