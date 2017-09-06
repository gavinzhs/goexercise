package main

import (
	"fmt"
)

type Persion struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]string)
	m["a1"] = "a"
	m["b1"] = "b"
	m["c1"] = "c"
	m["d1"] = "d"
	m["e1"] = "e"
	m["f1"] = "f"
	m["g1"] = "g"
	m["h1"] = "h"
	m["i1"] = "i"
	fmt.Println("m:", m)
	for k := range m {
		fmt.Println(k)
		if k == "e1" {
			delete(m, k)
		}
	}






	var m map[string]string
	fmt.Println("m addr: ", m)

	var s []string
	fmt.Println("s addr: ", s)

	var n = 1
	fmt.Println("n addr: ", n)
	fmt.Println("n addr: ", &n)

	var i *int
	fmt.Println("i addr: ", i)
	i = &n
	fmt.Println("i addr: ", i)
	fmt.Println("i addr: ", *i)
	fmt.Println("i addr: ", &i)
	testMapSingle()
	//testMapMulti()
}

func testMapSingle() {
	ps := make(map[string]Persion)
	ps["zhangsan"] = Persion{"zhangsan", 3}
	ps["lisi"] = Persion{"lisi", 4}
	p := ps["zhangsan"]
	p.Age = 5
	for _, v := range ps {
		fmt.Println(v)
	}
}

func testMapMulti() {
	ps := make(map[string][]*Persion)
	ps["zhangsan"] = []*Persion{&Persion{"zhangsan", 3}, &Persion{"zhangsan1", 31}}
	ps["lisi"] = []*Persion{&Persion{"lisi", 4}, &Persion{"lisi1", 41}}
	p, _ := ps["zhangsan"]
	p = append(p, &Persion{"zhangsan2", 32})
	for k, v := range ps {
		fmt.Println(k, v)
	}
}
