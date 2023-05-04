package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) Talk() {
	fmt.Println("Name of Person: ", person.Name)
}

type Android struct {
	Person Person
	Model string
}

func main()  {
	var android Android
	android.Person.Talk()
}