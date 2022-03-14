package main

import (
	dao "first-program/dao"
	web "first-program/webservice"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	dao.ReadFile()
	// fmt.Println(getCustomers())
	// for _, cust := range getCustomers() {
	// 	fmt.Printf("Name: %s\n", cust.firstName)
	// }
	//custList := dao.GetCustomers()
	//fmt.Println(custList)
	//newCust := dao.Customer{FirstName: "Ellora", LastName: "Nath", Id: 3}
	//dao.CreateCustomer(newCust)
	//custList = dao.GetCustomers()
	//fmt.Println(custList)
	web.Expose()
	//dao.DeleteCustomerById(1)
	fmt.Println(dao.GetCustomers())
}
