package main

import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello World")

	data, err := os.ReadFile("dummy.txt")
	checkErr(err)
	fmt.Printf("whole file data: %s\n", string(data))

	f, err := os.Open("dummy.txt")
	checkErr(err)

	byteData := make([]byte, 5)
	n1, err := f.Read(byteData)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(byteData[:n1]))

	f.Close()
}
