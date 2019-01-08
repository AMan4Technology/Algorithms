package WordSearch

import "testing"

func TestFindWords(t *testing.T) {
	println(findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain"}))
}

func BenchmarkFindWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findWords([][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		}, []string{"oath", "pea", "eat", "rain"})
	}
}
