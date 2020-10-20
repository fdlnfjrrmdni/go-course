package main

import ("fmt")

type segitiga interface {
	rataKiri() string
	rataKanan()	string
}

type num struct {
	number int
}

func (z num) rataKanan() string {
	n:=z.number
	i:=0
	for i < n {
		k:=0
		for k < i {
			fmt.Print(" ")
			k++	
		}
		j:=i
		for j < n {
			fmt.Print("*")
			j++
		}
		fmt.Print("\n")
		i++
	}
	return ""
}

func (z num) rataKiri() string {
	n:=z.number
	i:=0
	for i < n {
		j:=0
		for j <= i {
			fmt.Print("*")
			j++
		}
		fmt.Print("\n")
		i++
	}
	return ""
}

func main() {
	var pola segitiga
	pola = num{5}
	fmt.Println(pola.rataKiri())
	fmt.Println(pola.rataKanan())
}