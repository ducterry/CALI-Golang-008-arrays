package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 0.1 Khai báo mảng
	var (
		myArray1 = [5]int{2, 4, 6, 8, 10}
		myArray2 = [5]int{2}
		myArray3 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	)

	// 02. In mảng
	fmt.Println(myArray1)
	fmt.Println(myArray2)
	fmt.Println(myArray3)
}
