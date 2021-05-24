package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

func _goListCtxLists() []GoList {
	var slice = make([]GoList, _goListCtx.lists.Len())
	var i int = 0
	for e := _goListCtx.lists.Front(); e != nil; e = e.Next() {
		slice[i] = e.Value.(GoList)
		i++
	}
	return slice
}

func _goListCtxStore() {
	listSlice := _goListCtxLists()
	f, err := os.Create("GoLists.json")
	if err != nil {
		panic(err)
	}
	json.NewEncoder(f).Encode(listSlice)
	f.Close()
}

func listClear(l *list.List) {
	for e := l.Front(); e != nil; e = l.Front() {
		l.Remove(e)
	}
}

func _goListCtxLoad() bool {
	data, err := ioutil.ReadFile("GoLists.json")
	if err != nil {
		return false
	}
	var goLists []GoList
	err = json.Unmarshal(data, &goLists)
	if err != nil {
		panic(err)
	}
	_goListCtx.available = 1
	listClear(&_goListCtx.lists)
	for i := 0; i < len(goLists); i++ {
		_goListCtx.lists.PushBack(goLists[i])
		if goLists[i].ID > _goListCtx.available {
			_goListCtx.available = goLists[i].ID + 1
		}
	}
	return true
}

func _goListCtxInit() {
	if _goListCtxLoad() == false {
		_goListCtx.available = 1
		listClear(&_goListCtx.lists)
	}
}

func storeState(w http.ResponseWriter, r *http.Request) {
	_goListCtxStore()
}

func loadState(w http.ResponseWriter, r *http.Request) {
	_goListCtxLoad()
}

func resetState(w http.ResponseWriter, r *http.Request) {
}

func _goListCtxNew() int {
	ID := _goListCtx.available
	newList := GoList{ ID: ID, Items: make([]GoItem, 0)}
	_goListCtx.available++
	_goListCtx.lists.PushBack(newList)
	return ID
}

func _goListCtxDelete(ID int) bool {
	for e := _goListCtx.lists.Front(); e != nil; e= e.Next() {
		goList := e.Value.(GoList)
		if goList.ID == ID {
			_goListCtx.lists.Remove(e)
			return true
		}
	}
	return false
}

func _goListCtxUpdate(goList GoList) bool {
	for e := _goListCtx.lists.Front(); e != nil; e= e.Next() {
		i := e.Value.(GoList)
		if i.ID == goList.ID {
			_goListCtx.lists.Remove(e)
			_goListCtx.lists.PushBack(goList)
			return true
		}
	}
	return false
}

func updateList(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var goList GoList
	err := json.Unmarshal(reqBody, &goList)
	if err != nil { log.Println(err) }
	_goListCtxUpdate(goList)
}

func deleteList(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	body := string(reqBody)
	ID, _ := strconv.Atoi(body)
	if _goListCtxDelete(ID) {
		fmt.Printf("Deleted list %d\n", ID)
	} else {
		fmt.Printf("Failed to delete list %d\n", ID)
	}
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

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		fn(w, r)
	}
}

func handleRequest() {
	http.HandleFunc("/", addDefaultHeaders(rootPage))
	http.HandleFunc("/lists", addDefaultHeaders(getLists))
	http.HandleFunc("/create", addDefaultHeaders(createList))
	http.HandleFunc("/delete", addDefaultHeaders(deleteList))
	http.HandleFunc("/update", addDefaultHeaders(updateList))
	http.HandleFunc("/store", addDefaultHeaders(storeState))
	http.HandleFunc("/load", addDefaultHeaders(loadState))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	_goListCtxInit()
	handleRequest()
}
