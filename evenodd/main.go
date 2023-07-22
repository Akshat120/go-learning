package main

import "fmt"

func main() {

	arr := []int{}
	i := 0
	for i <= 10 {
		arr = append(arr, i)
		i = i + 1
	}

	for i := range arr {
		if arr[i]%2 == 0 {
			fmt.Println(arr[i], " is Even")
		} else {
			fmt.Println(arr[i], " is Odd")
		}
	}
}
