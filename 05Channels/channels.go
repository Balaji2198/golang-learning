package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// channels example
func generateValues(ch chan int) {
	ch <- 1
	ch <- 2
	// close(ch)
}

// END channels example

// Channel Directions example
// pings channel can only receive the values (chan<-)
func ping(pings chan<- string, message string) {
	pings <- message
}

// pings channel can only send the values (<-chan)
// pongs channel can only receive the values (chan<-)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

// END Channel Directions example

func searchFiles(directory string, fileTosearch string) string {
	fmt.Println("Searching: ", directory, fileTosearch)
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// fmt.Println(directory+"/"+file.Name(), file.IsDir())
		if file.Name() == fileTosearch {
			return "FOUND: " + filepath.Join(directory, file.Name())
		}
	}
	return "NOT FOUND: " + directory
}

func searchFilesAsync(directory string, fileToSearch string, ch chan string) {
	fmt.Println("Searching: ", directory, fileToSearch)
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == fileToSearch {
			ch <- "FOUND: " + filepath.Join(directory, file.Name())
			return
		}
	}
	ch <- "NOT FOUND: " + directory
}

func main() {
	fmt.Println("heyyy")
	// channels example
	ch := make(chan int)
	go generateValues(ch)

	var channelVal int
	channelVal = <-ch
	fmt.Println("channelVal: ", channelVal)

	fmt.Println("number of running goroutines: ", runtime.NumGoroutine())

	channelVal = <-ch
	fmt.Println("channelVal: ", channelVal)

	// Channel Directions
	fmt.Println("----Channel Directions example")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "This is pings pong")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	fmt.Println("----END Channel Directions example")

	// second time generate values -> example for select statement
	fmt.Println("----Select statement example")
	go generateValues(ch)
	for i := 0; i < 2; i++ {
		select {
		case val, ok := <-ch:
			if ok {
				fmt.Println("value: ", val)
			} else {
				fmt.Println("channel closed")
			}
		}
	}
	fmt.Println("----END Select statement example")

	// search files sync
	fmt.Println("----Search files sync")
	result := make([]string, 0)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	result = append(result, searchFiles(pwd, "hello.go"))
	result = append(result, searchFiles(pwd, "goRoutines.go"))
	result = append(result, searchFiles(pwd, "channels.go"))
	// result = append(result, searchFiles("./tmp4", "myfile.txt"))

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	fmt.Println("----END search files sync")

	fmt.Println("----Search files Async")
	fileChannel := make(chan string)
	// go searchFilesAsync(pwd, "hello.go", fileChannel)
	// go searchFilesAsync(pwd, "goRoutines.go", fileChannel)
	// go searchFilesAsync(pwd, "channels.go", fileChannel)
	fileNames := [3]string{"hello.go", "goRoutines.go", "channels.go"}
	for _, fileName := range fileNames {
		go searchFilesAsync(pwd, fileName, fileChannel)
	}

	for i := 0; i < 3; i++ {
		res := <-fileChannel
		fmt.Println(res)
	}
	fmt.Println("----END search files Async")
}
