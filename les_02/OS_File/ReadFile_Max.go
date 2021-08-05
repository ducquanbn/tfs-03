package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

// Convert String Array -> Int Array
func cvStrToArr(str string) []int {
	strArray := strings.Fields(string(str))
	slice := []int{}
	for i, _ := range strArray {
		if _, err := fmt.Sscan(strArray[i], &i); err == nil {
			slice = append(slice, i)
		}
	}
	return slice
}

// tim Max, Min, AVG
func cal(slice []int) (int, int, float32) {
	max := slice[0]
	min := slice[0]
	sum := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		sum += v
	}
	avg := float32(sum) / float32(len(slice))
	return max, min, avg
}

// Bunlle sort
func bbSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i] ^= slice[j]
				slice[j] ^= slice[i]
				slice[i] ^= slice[j]
			}
		}
	}
}

// So nguyen to
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	var a float64 = math.Sqrt(float64(n))
	for i := 2; i <= int(a); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//Search
func search(slice []int, a int) bool {
	check := false
	for _, v := range slice {
		if a == v {
			return true
		}
	}
	return check
}
func main() {
	//Open file
	data, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	// byte to String
	str := string(data)
	slice := cvStrToArr(str)
	fmt.Println("File: ", slice)

	//Tim max, min, avg
	max, min, avg := cal(slice)
	fmt.Println("MAX of Slice: ", max)
	fmt.Println("MIN of Slice: ", min)
	fmt.Printf("AVG of Slice: %.2f\n", avg)

	//Sap xep tang dan
	bbSort(slice)
	fmt.Println("Sap xep tang dan: ", slice)

	//Tim so nguyen to
	fmt.Print("So nguyen to co trong file: ")
	for _, v := range slice {
		if isPrime(v) {
			fmt.Print(v, "\t")
		}
	}

	//Tim kiem trong file, tim so 5 trong file
	a := 5
	if search(slice, a) {
		fmt.Println("\nCo ton tai ", a, "trong file")
	}

	// Xuat lai vao file
	var strToByte string
	for _, v := range slice {
		strToByte += strconv.Itoa(v) + " "
	}
	// String to Byte, export to file
	var mydata []byte = []byte(strToByte)
	errs := ioutil.WriteFile("text.txt", mydata, 0777)
	if err != nil {
		fmt.Println(errs)
	} else {
		fmt.Println("Xuat file thanh cong")
	}
}
