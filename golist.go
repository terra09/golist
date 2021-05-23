package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GoItem struct {
	Name string `json:"name"`
	Check bool `json:"check"`
}

type GoList struct {
	ID int `json:"id"`
	Items []GoItem `json:"items"`
}

func _dummieLists() []GoList {
	lists := []GoList{
		GoList{
			ID: 1, Items: []GoItem{
				GoItem{Name: "Banana", Check: false},
				GoItem{Name: "Chuchu", Check: true},
			},
		},
		GoList{
			ID: 2, Items: []GoItem{
				GoItem{Name: "chocolate", Check: false},
				GoItem{Name: "miojo", Check: true},
			},
		},
	}
	return lists
}

func getLists(w http.ResponseWriter, r *http.Request) {
	lists := _dummieLists()
	json.NewEncoder(w).Encode(lists)
	fmt.Println("getList has been accessed, result:")
	jsonData, err := json.MarshalIndent(lists, "", "\t")
	if err != nil { log.Println(err) }
	fmt.Println(string(jsonData))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "golist server bootstrap!")
	fmt.Println("rootPage has been accessed")
}

func handleRequest() {
	http.HandleFunc("/", rootPage)
	http.HandleFunc("/lists", getLists)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequest()
}
