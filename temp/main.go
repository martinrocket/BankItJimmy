package main

import "fmt"

type player struct {
	name  string
	score int
}

type playerList struct {
	list []player
}

func main() {
	m := new(playerList)
	fmt.Println(*m)
	nameOne := player{name: "John"}
	nameTwo := player{name: "Jane"}
	m.list = append(m.list, nameOne)
	m.list = append(m.list, nameTwo)
	for i := 0; i < len(m.list); i++ {
		fmt.Println(m.list[i])
	}
}
