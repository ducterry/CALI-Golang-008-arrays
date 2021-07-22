package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai báo biến
	var (
		myArray1         = [3]string{"Mark Zuckerberg", "Bill Gates", "Larrt Page"}
		myArray2         = [4]float64{3.5, 7.2, 4.8, 9.5}
		sumFloat float64 = 0
	)

	// 02. In màn hình
	fmt.Println(myArray1)
	fmt.Println("========================")

	for i := 0; i < len(myArray2); i++ {
		sumFloat += myArray2[i]
	}
	fmt.Printf("Tổng của myArray2 là %0.2f", sumFloat)
}
