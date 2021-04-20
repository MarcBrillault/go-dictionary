package main

import (
	"example.com/dictionary/dictionary"
)

func main() {

	d, err := dictionary.New("./badger")
	dictionary.HandleErr(err)
	defer d.Close()

	dictionary.ManageActions(d)
}
