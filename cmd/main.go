package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	n "netdemo"
	"os"
	u "os/user"
	"runtime"
	"strings"
)

const Ver string = "v1.2.1"

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

	user, err := u.Current()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("User  : ", user.Username)
	grpIds, _ := user.GroupIds()

	var groups []string
	for _, grpId := range grpIds {
		//		fmt.Printf("grpId : %s\n", grpId)
		grp, err := u.LookupGroupId(grpId)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		groups = append(groups, grp.Name)
	}

	fmt.Println("Groups: ", strings.Join(groups, ", "))

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

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "User  : %s \n", user.Username)
		fmt.Fprintf(w, "Groups: %s \n", strings.Join(groups, ", "))

	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
