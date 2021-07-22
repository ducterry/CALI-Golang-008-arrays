package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo mảng 2 chiều
	var (
		myArray = [2][2]int{
			{3, 5},
			{7, 9},
		}
	)


	// 02. In màn hình
	fmt.Println(myArray)
}
