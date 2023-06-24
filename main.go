package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hihiii")
	runServer()
}

func runServer() {
	fileserver := http.FileServer(http.Dir("./pages"))
	http.Handle("/", fileserver)
	// http.HandleFunc("/generator.html", handleGenerate)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("BRUHHH err: couldn't open the server :((")
	}
}

// func handleGenerate(w http.ResponseWriter, r *http.Request) {
// 	// if r.URL.Path != "/generator.html" {
// 	// 	http.Error(w, "404 Page not found", http.StatusNotFound)
// 	// 	return
// 	// }
// 	// if r.Method != "GET" {
// 	// 	http.Error(w, "Method not GET", http.StatusNotFound)
// 	// 	return
// 	// }

// 	fmt.Println("Generator loaded <33")
// }
