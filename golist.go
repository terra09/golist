package main

import (
	"container/list"
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

type GoListCtx struct {
	available int
	lists list.List
}

// global variable that will store the lists
var _goListCtx GoListCtx

func _goListCtxInit() {
	dummy := _dummieLists()
	_goListCtx.lists.PushBack(dummy[0])
	_goListCtx.lists.PushBack(dummy[1])
	_goListCtx.available = 3
}

func _goListCtxLists() []GoList {
	var slice = make([]GoList, _goListCtx.lists.Len())
	var i int = 0
	for e := _goListCtx.lists.Front(); e != nil; e = e.Next() {
		slice[i] = e.Value.(GoList)
		i++
	}
	return slice
}

func _goListCtxNew() int {
	ID := _goListCtx.available
	newList := GoList{ ID: ID, Items: make([]GoItem, 0)}
	_goListCtx.available++
	_goListCtx.lists.PushBack(newList)
	return ID
}

func createList(w http.ResponseWriter, r *http.Request) {
	ID := _goListCtxNew()
	fmt.Fprintf(w, "%d", ID)
	fmt.Printf("Created list %d\n", ID)
}

func getLists(w http.ResponseWriter, r *http.Request) {
	lists := _goListCtxLists()
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
	http.HandleFunc("/create", createList)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	_goListCtxInit()
	handleRequest()
}
