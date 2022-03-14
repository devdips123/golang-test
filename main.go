package main

import (
	"fmt"

	dao "dsahoo.com/golang/dao"
	web "dsahoo.com/golang/webservice"
)

func main() {
	fmt.Println("Inside the main function")
	dao.PrintAsJson(dao.GetCustomers())
	web.Expose()

}
