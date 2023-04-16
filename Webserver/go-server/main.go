package main

import (
	"fmt"
	"log"
	"net/http"
)

// create a handler function for the form
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err:= r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}


	fmt.Fprintf(w, "Post Request Successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
	fmt.Fprintf(w,"Email = %s\n",email)

}


// create a handler function for the hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return

	}

	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello, %s!", name)
}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf(("Starting server at port 8080"))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
