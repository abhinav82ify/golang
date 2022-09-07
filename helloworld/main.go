package main

// This is an import block. It allows multiple packages to be imported without repeating import statements
import (
	"fmt"
	"net/http"

	"github.com/abhinav82ify/golang/helloworld/controllers"
	"github.com/abhinav82ify/golang/helloworld/models"
)

/*
Main function when part of the main package identifies the entry
point of the application.
*/
func main() {
	fmt.Println("Hello world")

	// variables()
	// pointers()
	// constants()
	// arrays()
	// maps()
	// structs()
	// packages()
	webserver()
	// loopTillCondition()
	// loopOverCollections()
	// panics()
	// ifs()
	// switchs()
}

func variables() {
	// variable declaration and initialization separately
	var i int
	i = 32
	fmt.Println(i)

	// variable declaration and initialization together
	var f float32 = 3.14
	fmt.Println(f)

	// shorthand
	firstname := "Abhinav"
	fmt.Println(firstname)

	// boolean
	b := true
	fmt.Println(b)

	// imaginary numbers
	c := complex(3, 4)
	fmt.Println(c)

	// extract from complex. multiple initization
	real, imag := real(c), imag(c)
	fmt.Println(real)
	fmt.Println(imag)
}

func pointers() {
	// use * to reference and dereference a variable
	var firstname *string = new(string)
	*firstname = "Abhinav"

	fmt.Println(firstname)
	fmt.Println(*firstname)

	// create pointer variables
	fn := "Abhinav"
	ptr := &fn
	fmt.Println(ptr, *ptr)

	fn = "Ab"
	fmt.Println(ptr, *ptr)
}

func constants() {
	const pi = 3.1415
	fmt.Println(pi)
}

func arrays() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println(a)

	a2 := []int{1, 2, 3}
	fmt.Println(a2)

	slice := a2[:]
	fmt.Println(slice)

	s2 := a2[1:]
	s3 := a2[:2]
	s4 := a2[1:2]

	fmt.Println(s2, s3, s4)
}

func maps() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}

func structs() {
	type user struct {
		id        int
		firstName string
		lastName  string
	}

	u := user{id: 1, firstName: "Abhinav", lastName: "Sharma"}
	fmt.Println(u)
}

func packages() {
	u := models.User{Id: 1, FirstName: "Abhinav", LastName: "Sharma"}
	fmt.Println(u)

	users := models.GetUsers()
	fmt.Println(users)

	nu, err := models.AddUser(models.User{FirstName: "Nikita", LastName: "Pareek"})
	fmt.Println(nu)
	fmt.Println(err)

	models.AddUser(models.User{FirstName: "Abhinav", LastName: "Sharma"})

	users = models.GetUsers()
	fmt.Println(users)
}

func webserver() {
	controllers.RegisterHandlers()
	http.ListenAndServe(":3000", nil)
}

func loopTillCondition() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}

	for j := 0; j < 5; j++ {
		println(j)
	}

	// infinite loop
	k := 0
	for {
		if k == 5 {
			break
		}
		println(k)
		k++
	}
}

func loopOverCollections() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for j, v := range slice {
		println(j, v)
	}

	ports := map[string]int{"http": 80, "https": 443}
	for k, v := range ports {
		println(k, v)
	}

	for k := range ports {
		println(k, ports[k])
	}

	for _, v := range ports {
		println(v)
	}
}

// panics are similar to exceptions
// panics act as an exception is thrown and statements after this are not executed
func panics() {
	println("Web server starting")
	panic("Something bad has happened")
	println("Web server starting failed")
}

func ifs() {
	u1 := models.User{
		Id:        1,
		FirstName: "Abhinav",
		LastName:  "Sharma",
	}

	u2 := models.User{
		Id:        2,
		FirstName: "Nikita",
		LastName:  "Pareek",
	}

	if u1.Id == u2.Id {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user")
	} else {
		println("Different user")
	}
}

func switchs() {
	type HTTPMethod struct {
		Method string
	}
	r := HTTPMethod{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	case "DELETE":
		println("Delete request")
	default:
		println("Unhandled method")
	}
}
