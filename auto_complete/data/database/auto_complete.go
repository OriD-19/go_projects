package database

import (
	"log"

	"OriD19.com/auto_complete/data/structures"
)

var AutoCompleteTrie structures.Trie

func populateAutoComplete() {
	AutoCompleteTrie = structures.NewTrie()

	res, err := GetAllContacts()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, contact := range *res {
		// Populate our autoComplete trie with the current records
		AutoCompleteTrie.Insert(contact.Name)
	}
}
