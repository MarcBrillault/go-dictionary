package dictionary

import (
	"flag"
	"fmt"
)

func ManageActions(d *Dictionary) {
	action := flag.String("action", "list", "Action to perform on the dictionary")
	flag.Parse()

	switch *action {
	case "add":
		cliActionAdd(d, flag.Args())
	case "define":
		cliActionDefine(d, flag.Arg(0))
	case "list":
		cliActionList(d)
	case "remove":
		cliActionRemove(d, flag.Arg(0))
	default:
		fmt.Printf("Unkonwn action: %v\n", *action)
	}
}

func cliActionAdd(d *Dictionary, args []string) {
	word := args[0]
	definition := args[1]

	err := d.Add(word, definition)
	HandleErr(err)
	fmt.Printf("Word %v has been added to the database\n", word)
}

func cliActionDefine(d *Dictionary, word string) {
	entry, err := d.Get(word)
	HandleErr(err)
	fmt.Println(entry)
}

func cliActionList(d *Dictionary) {
	words, entries, err := d.List()
	HandleErr(err)

	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func cliActionRemove(d *Dictionary, word string) {
	err := d.Remove(word)
	HandleErr(err)
	fmt.Printf("Word %v has been removed from the database\n", word)
}
