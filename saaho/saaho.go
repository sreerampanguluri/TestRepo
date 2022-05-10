package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func whoAmI(w http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Sreeram Panguluri",
			Title: "Golang developer",
			State: "TX"},
	}

	json.NewEncoder(w).Encode(who)
	fmt.Println("endpoint hit", who)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to go web api")
	fmt.Println("Endpoint hit: homePage")
}
func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := " SreeramPanguluri "
	fmt.Fprintf(w, "little bit about Sreeram Panguluri ")
	fmt.Println("Endpoint hit:", who)

}
func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":9000", nil))

}
