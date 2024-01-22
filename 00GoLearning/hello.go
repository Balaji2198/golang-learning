package main

import (
	"errors"
	"fmt"
)

func calculateSum(a, b int) int {
	return a + b
}

func printArray(arr ...int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d - %d\n", i, arr[i])
	}
}

func sliceArray() {
	arr := [5]int{1, 3, 5, 7}
	fmt.Printf("arr: %v\n", arr)
	mySlice := []int{1, 3, 5, 7}
	fmt.Printf("mySlice: %v\n", mySlice)
	mySlice = append(mySlice, 9, 11)
	fmt.Printf("updated mySlice: %v\n", mySlice)
	fmt.Println(mySlice[2])
}

func mapOperations() {
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println("m:", m)
	val, isPresent := m["a"]
	fmt.Printf("val:%d, isPresent: %v", val, isPresent)
}

func rangeOperations(s1 []int) {
	fmt.Println()
	for i, val := range s1 {
		fmt.Printf("index: %d, val: %d ", i, val)
	}
	fmt.Println()
}

func addSub(x, y int) (int, int, int, int) {
	return x + y, x - y, x, y
}

func calculateSum2(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

func closureIntSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func ptrTest(ptr *int) {
	*ptr = 0
}

// Interfaces
type person struct {
	name string
	age  int
}

type Animal interface {
	getName() string
	speak() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) getName() string {
	return d.name
}

func (d Dog) speak() string {
	return "Woof!"
}

func (c Cat) getName() string {
	return c.name
}

func (c Cat) speak() string {
	return "Meow!"
}

func animal_speak(a Animal) {
	fmt.Printf("%s says %s\n", a.getName(), a.speak())
}

// END Interfaces

// Struct Embedding
type base struct {
	num  int
	bstr string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v str=%v", b.num, b.bstr)
}

type container struct {
	base
	cstr string
}

// End Struct Embedding

// Generics
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

type node[T any] struct {
	next *node[T]
	val  T
}

type List[T any] struct {
	head, tail *node[T]
}

func (lst *List[T]) Push(v T) {
	if lst.head == nil {
		lst.head = &node[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &node[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() []T {
	var nodes []T
	for n := lst.head; n != nil; n = n.next {
		nodes = append(nodes, n.val)
	}
	return nodes
}

// END Generics

// Error handling
func incrementOne(val int) (int, error) {
	if val == 404 {
		return -1, errors.New("Error 404!")
	}
	return val + 1, nil
}

type argError struct {
	arg         int
	errorString string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errorString)
}

func incrementTwo(val int) (int, error) {
	if val == 404 {
		return -1, argError{val, "Error 404!!"}
	}
	return val + 2, nil
}

// END Error handling

func main() {
	fmt.Println("Hello world.")

	sum := calculateSum(1, 7)
	fmt.Println("sum =", sum)

	arr := []int{1, 3, 5, 7}
	printArray(arr...)
	sliceArray()
	mapOperations()
	rangeOperations(arr)
	sum1, sub1, x, y := addSub(5, 2)
	fmt.Println(sum1, sub1, x, y)
	calculateSum2(3, 5, 6)

	nextInt1 := closureIntSeq()
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
	nextInt2 := closureIntSeq()
	fmt.Println(nextInt2())
	fmt.Println(nextInt1())

	ptrVal := 5
	fmt.Println("ptrVal", ptrVal)
	ptrTest(&ptrVal)
	fmt.Println("ptrVal", ptrVal)

	person1 := person{name: "bal", age: 40}
	fmt.Println(person1)

	// Interfaces
	animals := []Animal{Dog{"doggy"}, Cat{"catty"}, Dog{"doggy2"}}
	for _, animal := range animals {
		animal_speak(animal)
	}

	// Struct Embedding
	base1 := base{num: 1, bstr: "base string"}
	container1 := container{base: base1, cstr: "container string"}
	fmt.Printf("container: %+v\n", container1)
	fmt.Println("describe: ", container1.describe())

	// Generics
	var m = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("keys: ", MapKeys(m))
	fmt.Println("keys: ", MapKeys[int, string](m))

	lst := List[int]{}
	lst.Push(101)
	lst.Push(102)
	fmt.Println("all values: ", lst.getAll())
	fmt.Println("head: ", lst.head)
	fmt.Println("tail: ", lst.tail)

	if val, e := incrementOne(404); e != nil {
		fmt.Println("incrementOne failed:", e)
	} else {
		fmt.Println("incremented value:", val)
	}

	if val, e := incrementTwo(404); e != nil {
		fmt.Println("incrementTwo failed:", e)
	} else {
		fmt.Println("incremented value:", val)
	}

}
