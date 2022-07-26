package main


func RecursiveFind(list []int,f int,low int,high int)int{
	mid:=low+(high-low)/2
	if f==list[mid]{
		return mid
	}else if f>list[mid]{
		return RecursiveFind(list,f,mid+1,high)
	}else{
		return RecursiveFind(list,f,low,mid-1)
	}
}