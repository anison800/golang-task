package main

import "fmt"

func main() {

	eployee := Employee{Person: Person{Name: "kakak", Age: 23}, EmployeeID: "9981"}
	eployee.PrintInfo()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("打印信息：%s,%s,%d\n", e.EmployeeID, e.Person.Name, e.Person.Age)
}
