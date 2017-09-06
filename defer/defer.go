package main

import "fmt"

func main() {
	fmt.Println(getA())
}

func getA() (a int) {
	defer reduce(plus(a, 1), a)
	//defer func(){
	//	rlt := plus(a, 1) - a
	//	fmt.Printf("x: %d, y: %d, rlt: %d\n", plus(a, 1), a, rlt)
	//}()

	//defer func(){
	//	fmt.Println("a:", a)
	//}()
	//defer p(a)
	a = 2
	return
}

func plus(x, y int) int {
	return x + y
}

func reduce(m, n int) {
	rlt := m - n
	fmt.Printf("x: %d, y: %d, rlt: %d\n", m, n, rlt)
}

func p(a int){
	fmt.Println("a:", a)
}