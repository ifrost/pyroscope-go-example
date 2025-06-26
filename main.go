package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/pyroscope-io/client/pyroscope"
)

func busyWork() {
	for i := 0; i < 10000000; i++ {
		_ = rand.Float64() * rand.Float64()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	busyWork()
	fmt.Fprintf(w, "Hello from Go with Pyroscope!\n")
}

func main() {
	// Start profiling
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "go.app.demo",
		ServerAddress:   "http://pyroscope:4040",
		Logger:          pyroscope.StandardLogger,

		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Serving on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

