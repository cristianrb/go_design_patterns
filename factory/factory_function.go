package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// validation ...
	}
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("Cristian", 28)
	p.EyeCount = 1
}
