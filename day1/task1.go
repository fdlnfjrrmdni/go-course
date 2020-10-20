package main

import ("fmt")

func printImage(n int) {
	if n % 2 == 1 {
		i := 0
		for i < n {
			j := 0
			for j < n {
				if j == 0 || i == n/2 || j == n-1 {
					fmt.Print("*    ")
				}else {
					fmt.Print("=    ")
				}
				j++
			}
			fmt.Print("\n\n")
			i++
		}
	}else {
		fmt.Println("error: the value must be an odd number")
	}
}

func main() {
	printImage(5)
}