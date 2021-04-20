package main

import (
	"fmt"
	"os"

	"example.com/dictionary/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)

	d.Remove("PZHP")

	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}

	defer d.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v", err)
		os.Exit(1)
	}
}
