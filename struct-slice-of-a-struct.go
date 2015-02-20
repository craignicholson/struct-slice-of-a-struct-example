package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Customer with one or more address
type Customer struct {
	FirstName string
	LastName  string
	Address   []Location
}

// Location
type Location struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Country  string
	Lat      string
	Long     string
}

func main() {

	//Example 1 - Dynamic Slice
	customer := SliceOfCustomers()
	PrintJSON(customer)

}

// Example showing how we can add multiple locations
// to one customer when you do not know the # of
// locations to be added
func SliceOfCustomers() []Customer {

	// Example of the known length of the slice
	customers := make([]Customer, 5)
	for i := 0; i < 5; i++ {
		customers[i].FirstName = "Craig"
		customers[i].LastName = "Nicholson"

		// Example of dynamic allocation to slice
		locations := make([]Location, 0)
		for i := 0; i < 2; i++ {
			data := Location{}
			data.Address1 = strconv.Itoa(i)
			data.Address2 = strconv.Itoa(i)
			data.City = strconv.Itoa(i)
			data.State = strconv.Itoa(i)
			data.Country = strconv.Itoa(i)

			locations = append(locations, data)
		}
		customers[i].Address = locations
	}
	return customers
}

// Print JSON
func PrintJSON(customer []Customer) {
	json, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the Results
	fmt.Println(string(json))
	fmt.Println("\n")
}
