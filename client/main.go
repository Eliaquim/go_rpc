package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}


func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Erro de conexao: ", err)
	}

	a := Item{"Primeiro", "um item de test"}
	b := Item{"Segundo", "um segundo item"}
	c := Item{"Terceiro", "e mais um item, o terceiro"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)

	client.Call("API.GetDB", "", &db)

	fmt.Println("Database in client: ", db)

	client.Call("API.EditItem", Item{"Segundo", "Segundo item editado"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database after delete item: ", db)

	client.Call("API.GetByName", "Primeiro", &reply)
	fmt.Println("Primeiro item recebido: ", reply)

}
