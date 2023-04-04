package main

import "fmt"

var (
	actorNmae string = "Elizabeth"
	companion string = "Smith"
	season int = 12
)

func main() {
	var declare int = 78;
	newDeclare := 99
	//----variable within the innermost scope takes the precedence-----------
	fmt.Println("Hello, excited with go")
	fmt.Print("Am runiing -------")
	fmt.Println(declare, newDeclare)
}
//-----Variables in go must be used