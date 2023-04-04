// // import ("fmt"; "path")
// package main

// import (
// 	"fmt"
// 	"path"
// )

// func main ()  {
// 	var dir, file String
// 	dir, file = path.Split("css/main.css")
// 	_, file = path.Split("css/main.css")
// 	_, file := path.Split("css/main.css")
// 	fmt.Println("file:", file)
// }

// func Split(path String) (dir, file String)
package main

import (
	"fmt"
	"time"
)

func hello(){
	fmt.Println("Trying out a goroutine")
}
func main(){
	go hello()
	fmt.Println("main function")
	time.Sleep(1 * time.Second)
	fmt.Println("Another Main fxn")
	fmt.Println("Hello World!")
}