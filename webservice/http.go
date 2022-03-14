package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"dsahoo.com/golang/dao"
)

func Expose() {
	fmt.Println("Server started at port 8080")
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/all", getAllCustomers)
	http.HandleFunc("/create", createCustomer)
	http.HandleFunc("/delete", deleteCustomer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func helloWorld(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello World!")
}

func getAllCustomers(res http.ResponseWriter, req *http.Request) {
	customers := dao.GetCustomers()
	data, _ := json.Marshal(customers)
	res.Header().Add("Content-Type", "application/json")
	io.WriteString(res, string(data))
}

func createCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("The request body: " + string(body))
		var newCustomer dao.Customer
		err = json.Unmarshal(body, &newCustomer)
		dao.CreateCustomer(newCustomer)
		if err != nil {
			panic(err)
		}
		res.WriteHeader(http.StatusCreated)
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func deleteCustomer(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		query := req.URL.Query()
		id := query.Get("Id")
		var responseText string
		fmt.Println("Request parameter: " + id)
		if id != "" {
			var n int
			idInt, _ := fmt.Sscan(id, &n)
			status := dao.DeleteCustomerById(idInt)
			if status {
				responseText = "Delete successful"
			} else {
				responseText = "nothing to delete"
			}
		}
		res.WriteHeader(http.StatusOK)
		response := make(map[string]string)
		response["message"] = responseText
		jsonResponse, _ := json.Marshal(response)
		res.Write(jsonResponse)
		res.Header().Add("Content-Type", "application/json")
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
