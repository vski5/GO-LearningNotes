package main

import "fmt"

func UpdateSlice(s []int) { //此处方括号内没写长度，s就是一个int的slice
	s[0] = 100
}
func main() {
	arr4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr4[2:6]
	fmt.Println(s)
	fmt.Println(arr4[2:6]) //这四个arr4定义的方括号下标区间都是一种slice
	fmt.Println(arr4[:6])  //slice不是值类型，是一种视图view，此处是看最开头到6这个数。
	fmt.Println(arr4[2:])
	fmt.Println(arr4[:])
	s1 := arr4[2:]
	fmt.Println("s1=", s1)
	s2 := arr4[:]
	fmt.Println("s2=", s2)

	UpdateSlice(s1)
	//s1 := arr4[2:]，从2这个数开始，因为2是s1这个切片的第一个数，UpdateSlice(s1) 就是改第一个数为100，所以2就被改成100了
	fmt.Println("改过的arr4=", arr4)            // [0 1 100 3 4 5 6 7]
	fmt.Println("被UpdateSlice(s1)后的s1=", s1) // [100 3 4 5 6 7]
	//此处s1是对上面的view,所以是从上面第三个开始。
	UpdateSlice(s2) //s2 := arr4[:]， 完全的arr4
	//这里的UpdateSlice(s2) 就是改第一个数为100，所以0就被改成100了
	fmt.Println("s2=", s2)
	fmt.Println("因为这里的s2是上面改过的arr4的view，所以2变成了100。view读的是上面一个arr4，不是第一个")
	fmt.Println("arr4=", arr4)

	arr5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	x := arr5[2:6]
	fmt.Println(x)
	//还可以对slice取slice，这被称作re slice
	x = x[:5] //从第一个切片开始向右数5个
	fmt.Println(x)
	x = x[:6] //从第一个切片开始向右数六个
	fmt.Println(x)
	x = x[2:]
	fmt.Println(x) //但是这里就变成了切片的切片，被进一步裁剪。从零开始数，0.1.2然后就数到了4这个数。
	x = x[2:]
	fmt.Println(x) //但是这里就变成了切片的切片，被进一步裁剪。从零开始数，0.1.2然后就数到了4这个数。
}
