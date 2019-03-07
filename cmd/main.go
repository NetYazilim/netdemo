package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	n "netdemo"
	"os"

	"runtime"
)

const Ver string = "v1.0"

var GoVersion string

func init() {
	GoVersion = runtime.Version()
	fmt.Println(" _   _      _  ______                     ")
	fmt.Println("| \\ | |    | | |  _  \\                    ")
	fmt.Println("|  \\| | ___| |_| | | |___ _ __ ___   ___  ")
	fmt.Println("| . ` |/ _ \\ __| | | / _ \\ '_ ` _ \\ / _ \\ ")
	fmt.Println("| |\\  |  __/ |_| |/ /  __/ | | | | | (_) |")
	fmt.Printf("\\_| \\_/\\___|\\__|___/ \\___|_| |_| |_|\\___/ Ver.: %s\n", Ver)
	fmt.Println()

	fmt.Printf("Go   : %s\n", GoVersion)
}

func main() {
	log.Println("Server starting...")
	http.HandleFunc("/", n.HIndex)

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, e := range os.Environ() {
			fmt.Fprintf(w, "%s \n", e)
		}
	})

	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Fprintf(w, "Ip : %s \n", ip)

		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Fprintf(w, "%s : %s \n", name, value)

			}
		}
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
