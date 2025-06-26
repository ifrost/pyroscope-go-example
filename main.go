package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/pyroscope-io/client/pyroscope"
	_ "runtime/cgo" 
)

func doHeavyCalculation() {
	// Simulated heavy calculation
	for i := 0; i < 10_000_000; i++ {
		_ = rand.Float64() * rand.Float64()
	}
}

func doWork() {
	for i := 0; i < 5; i++ {
		doHeavyCalculation()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	doWork()
	fmt.Fprintf(w, "Nested profiling example with Pyroscope!\n")
}

func main() {
	// Start profiling
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "go.app.example",
		ServerAddress:   "http://pyroscope:4040",
		Logger:          pyroscope.StandardLogger,

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Serving on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

