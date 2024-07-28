package main

import "fmt"

// Define a struct chamada Person
type Person struct {
	Name string
	Age  int
}

// Define um método para a struct Person
// Este método retorna uma string que descreve a pessoa
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func main() {
	// Cria uma instância da struct Person
	person := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "John", Age: 57}

	// Chama o método Greet da instância person
	greeting := person.Greet()
	// Imprime a saudação
	fmt.Println(greeting)

	// Imprime a saudação
	fmt.Println(person2.Greet())
}
