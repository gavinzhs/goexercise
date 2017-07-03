package main

import (
	"sort"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type PersonList []*Person

func (p PersonList) Len() int {
	return len(p)
}

func (p PersonList) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p PersonList) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}

type PersonSortByAge struct {
	PersonList
}

func (pl PersonSortByAge) Less(i, j int) bool {
	return pl.PersonList[i].Age > pl.PersonList[j].Age
}

type PersonWrapper struct {
	PersonList []*Person
	SortBy func (p, q *Person) bool
}

func (pw PersonWrapper) Len() int {
	return len(pw.PersonList)
}

func (pw PersonWrapper) Less(i, j int) bool {
	return pw.SortBy(pw.PersonList[i], pw.PersonList[j])
}

func (pw PersonWrapper) Swap(i, j int) {
	pw.PersonList[i], pw.PersonList[j] = pw.PersonList[j], pw.PersonList[i]
}

func main(){
	pl := PersonList{&Person{"a", 1},&Person{"a", 2},&Person{"a", 0},&Person{"a", 3},&Person{"a", 4}}
	sort.Sort(pl)
	for _, p := range pl{
		fmt.Println(p)
	}

	sort.Sort(PersonSortByAge{pl})
	for _, p := range pl{
		fmt.Println(p)
	}

	sort.Sort(PersonWrapper{pl, func(p, q *Person) bool {
		return p.Age < q.Age
	}})

	for _, p := range pl{
		fmt.Println(p)
	}
}


