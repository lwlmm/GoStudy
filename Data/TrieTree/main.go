package main

import (
	"fmt"
)

type Trie struct {
	Next [26]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie {}
}

func (this *Trie) Insert(word string) {
	for _, c := range word {
		if this.Next[c-'a'] == nil {
			next := Constructor()
			this.Next[c-'a'] = &next
		}
		this = this.Next[c-'a']
	}
	this.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.Next[c-'a'] == nil {
			return false
		}
		this = this.Next[c-'a']
	}
	return this.IsEnd
}

// Returns if there is any word in the trie starts with the given prefix
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.Next[c-'a'] == nil {
			return false
		}
		this = this.Next[c-'a']
	}
	return true
}

func main() {
	trie := Constructor()

	trie.Insert("apple")

	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))

	trie.Insert("app")
	
	fmt.Println(trie.Search("app"))
}