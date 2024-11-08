package main

import (
	_ "fmt"
	"net/http"
	"github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello from Golang!"))
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	// Set up the logging configuraton
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	// Define the handler
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/helloworld", helloWorld)

	// Start the Server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8090", nil); err != nil {
			logrus.Fatalf("Failed to start the server: %v", err)
		}
	}()

	// Log that server has started
	logrus.Info("Started Server, listening on port 8090")

	// Block the main goroutine to keep the sever running
	select{}
}
