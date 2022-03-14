package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Fields need to be starting with Cap letters to be marshalled
type Customer struct {
	FirstName   string
	LastName    string
	Id          int
	Address     string `json:"-"`
	PhoneNumber string
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile() {
	content, err := ioutil.ReadFile("demo.txt")
	handleError(err)
	fmt.Printf("File contents: %s\n", content)
}

func GetCustomers() (customers []Customer) {
	data, err := ioutil.ReadFile("./customers.json")
	handleError(err)
	if len(data) > 0 {
		err = json.Unmarshal(data, &customers)
		handleError(err)
	}

	// customers = append(customers, Customer{FirstName: "Debasish", LastName: "Sahoo", Id: 11, PhoneNumber: "216-225-4546"})
	// customers = append(customers, Customer{FirstName: "Amit", LastName: "Gupta", Id: 13})

	return customers
}

func CreateCustomer(customer Customer) {
	existingCustomer := GetCustomerById(customer.Id)
	customers := GetCustomers()
	if (existingCustomer == Customer{}) {

		customers = append(customers, customer)
		writeAsJson(customers)
	} else {
		fmt.Println("Customer already exists")
	}
	printAsJson(customers)
}

func GetCustomerById(id int) (cust Customer) {
	customers := GetCustomers()
	for _, cust := range customers {
		if cust.Id == id {
			return cust
		}
	}
	return cust
}

func DeleteCustomerById(id int) bool {
	var toDeleteCustomer int = -1
	customers := GetCustomers()
	for index, customer := range customers {
		if id == customer.Id {
			toDeleteCustomer = index
			break
		}
	}

	if toDeleteCustomer >= 0 {
		fmt.Printf("Customer id: %d to be deleted at index: %d\n", id, toDeleteCustomer)
		writeAsJson(append(customers[:toDeleteCustomer], customers[toDeleteCustomer+1:]...))
		return true
	} else {
		fmt.Printf("No customer matching Id: %d\n", id)
		return false
	}

}

func writeAsJson(customers []Customer) {
	//names := []string{"Debasish", "Ellora"}
	data, err := json.Marshal(customers)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("customers.json", data, 0644)
}

func printAsJson(customers []Customer) {
	data, err := json.Marshal(customers)
	handleError(err)
	fmt.Println(string(data))
}
