package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo biến
	var (
		dayOfWeek         = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
		myFloat           = [4]float64{3.5, 7.2, 4.8, 9.5}
		sumFloat  float64 = 0
	)

	// 02. Xử lý Array
	for index, value := range dayOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}


	for _, value := range myFloat {
		sumFloat = sumFloat + value
	}

	fmt.Printf("Sum of all the elements in array %v = %0.2f\n", myFloat, sumFloat)
}
