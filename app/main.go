package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// PageVariables holds the data to be passed to the HTML template
type PageVariables struct {
	Podname   string
	Nodename  string
	Namespace string
	Ipaddress string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the values of the environment variables and store in PageVariables
	pageVariables := PageVariables{
		Podname:   os.Getenv("HOSTNAME"),
		Nodename:  os.Getenv("NODE_NAME"),
		Namespace: os.Getenv("POD_NAMESPACE"),
		Ipaddress: os.Getenv("POD_IP_ADDRESS"),
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with the PageVariables and write the result to the response
	err = tmpl.Execute(w, pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Serve the / path with the handler function
	http.HandleFunc("/", handler)

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
