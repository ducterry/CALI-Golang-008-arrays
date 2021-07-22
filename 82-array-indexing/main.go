package main

import "fmt"
/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo mảng
	var myArray [5]int

	myArray[0] = 100
	myArray[1] = 101
	myArray[3] = 103
	myArray[4] = 105

	// 02. In các phần tử trong mảng
	fmt.Println("===================")
	fmt.Printf(".myArray[0] = %d\n.myArray[1] = %d\n.myArray[2] = %d\n", myArray[0], myArray[1], myArray[2])
	fmt.Println("myArray = ", myArray)
	fmt.Println("===================")


	// 03. In theo vòng for
	for i := 0; i < len(myArray); i++ {
		fmt.Printf(".myArr[%d] = %d\n",i,myArray[i])
	}
	fmt.Println("myArray = ", myArray)
}
