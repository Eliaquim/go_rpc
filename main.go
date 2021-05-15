package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

// for RPC, functions need to sutisfy:
//		functions need to be method
//		functions need to be expoerted
//		functions need to have two arguments, both of them exported types, including built in types
//		functions second argument must be a pointer
//		functions must return an error type

// allows to elevate our functions to methods
type API int

var database []Item

// allows client to get the database
// Title argument is here only to satisfy rpc library
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

// add receiver for the API pointer to turn into method
// the first argument is passed by the caller, the second is the result being returned ti the client
// if error is returned, RPC will not send any data to the caller
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}
	// equal reply pointer to the getItemn
	*reply = getItem
	// nil error
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = database[idx]
		}
	}
	*reply = changed
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item

	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:] ...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {

	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}


	//fmt.Println("database inicial: ", database)
	//a := Item{"primairo", "um item de test"}
	//b := Item{"segundo", "um segundo item"}
	//c := Item{"terceiro", "e mais um item"}
	//
	//AddItem(a)
	//AddItem(b)
	//AddItem(c)
	//fmt.Println("segunda base de dados: ", database)
	//
	//DeleteItem(b)
	//fmt.Println("terceira base de dados: ", database)
	//
	//EditItem("terceiro", Item{"quarto", "um novo item"})
	//fmt.Println("quarta base de dados: ", database)
	//
	//x := GetByName("quarto")
	//y := GetByName("terceiro")
	//fmt.Println(x, y)


}
