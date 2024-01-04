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
	// Retrieve the values of the HOSTNAME and IPADDR environment variables
	pod_name := os.Getenv("HOSTNAME")
	node_name := os.Getenv("NODE_NAME")
	namespace := os.Getenv("POD_NAMESPACE")
	ip_address := os.Getenv("POD_IP_ADDRESS")

	// Create a PageVariables instance with the environment variable values
	pageVariables := PageVariables{
		Podname:   pod_name,
		Nodename:  node_name,
		Namespace: namespace,
		Ipaddress: ip_address,
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("index.html")
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
