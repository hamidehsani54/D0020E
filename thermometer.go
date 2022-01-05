package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//man kan se på den som client
func main() {
	// What to execute for various page requests
	//så skapar man en tråd
	go http.HandleFunc("/", getTemperature)

	// Listens for incoming connections
	//valde en random port
	http.ListenAndServe(":8091", nil)
}

// Home page that includes a link to a subpage
func getTemperature(w http.ResponseWriter, req *http.Request) {
	//läser temp och sedan convertera den till string
	fmt.Fprintf(w, strconv.Itoa(readTemperature()))
}

// Returns a temperature
// TODO: Should be connected to a sensor
func readTemperature() int {
	// Sends a random number between 0 and 50 (for now)
	/* 	rand.Seed(time.Now().UnixNano())
	   	var randomNum = rand.Intn(50)

	   	return randomNum */

	return 28
}

// Register IP and port data to the Service Registry
/* func registerServiceToSR() {

} */

//printer ut allt på terminalen
