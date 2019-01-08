package trie

import (
	"fmt"
	"testing"
)

func TestTrie_Words(t *testing.T) {
	trie := NewTrie(26)
	trie.Insert("apple")
	trie.Search("apple")   // return true
	trie.Search("app")     // return false
	trie.StartsWith("app") // return true
	trie.Insert("app")
	trie.Search("app") // return true
	fmt.Println(trie.Words("app"))
}
