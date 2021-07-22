package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo biến
	var (
		myArray1 = [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
		myArray2 = myArray1
	)

	// 02. In màn hình
	fmt.Println("myArray1= ", myArray1)

	myArray2[1] = "Việt Nam"
	fmt.Println("myArray2= ", myArray2)
}
