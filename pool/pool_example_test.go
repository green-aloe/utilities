package pool_test

import (
	"fmt"
	"log"

	"github.com/green-aloe/utility/pool"
)

func ExamplePool_Get() {
	type Person struct {
		Name string
		Age  int

		internal int
	}

	pool := pool.Pool[Person]{
		NewItem: func() Person {
			return Person{
				Name: "Alice",
				Age:  74,
			}
		},
	}

	person := pool.Get()
	fmt.Println(person)

	// Output:
	// {Alice 74 0}
}

func ExamplePool_Store() {
	type Person struct {
		Name string
		Age  int

		internal int
	}

	pool := pool.Pool[Person]{
		NewItem: func() Person {
			return Person{
				Name: "Alice",
				Age:  74,
			}
		},
		Prestore: func(person Person) Person {
			log.Println("Storing", person.Name, "in pool")

			person.Name = ""
			person.Age = 0

			return person
		},
	}

	pool.Store(Person{
		Name:     "Bob",
		Age:      42,
		internal: 1,
	})

	person := pool.Get()
	fmt.Println(person)

	// Output:
	// { 0 1}
}

func ExamplePool_Count() {
	var pool pool.Pool[int]
	for i := 0; i < 10; i++ {
		pool.Store(i)
	}

	count := pool.Count()
	fmt.Println(count)

	// Output:
	// 10
}

func ExamplePool_Clear() {
	var pool pool.Pool[int]
	for i := 0; i < 10; i++ {
		pool.Store(i)
	}

	pool.Clear()

	count := pool.Count()
	fmt.Println(count)

	// Output:
	// 0
}
