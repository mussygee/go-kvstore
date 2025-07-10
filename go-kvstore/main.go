package main

import (
	"fmt"
	"go-database/storage"
	"log"
)

func main() {
	idx, err := storage.LoadIndex("data.db")
	if err != nil {
		log.Fatal(err)
	}

	err = idx.Set("name", "Mussy")
	if err != nil {
		log.Fatal("Set failed: ", err)

	}

	val, err := idx.Get("name")
	if err != nil {
		log.Fatal("Get failed: ", err)
	}

	fmt.Println("Got value for 'name':", val)

}
