package trie

func NewTrie(charNum int) Trie {
	return Trie{
		root:    newNode(charNum),
		charNum: charNum,
	}
}

type Trie struct {
	root    *node
	charNum int
}

func (t Trie) Insert(word string) {
	pre := t.root
	for _, c := range word {
		hash := c - 'a'
		if pre.next[hash] == nil {
			pre.next[hash] = newNode(t.charNum)
		}
		pre = pre.next[hash]
	}
	pre.isWord = true
}

func (t Trie) Search(word string) bool {
	pre := t.endNode(word)
	return pre != nil && pre.isWord
}

func (t Trie) StartsWith(prefix string) bool {
	pre := t.endNode(prefix)
	return pre != nil
}

func (t Trie) Words(prefix string) (results []string) {
	start := t.endNode(prefix)
	return start.words(prefix, results)
}

func (t Trie) endNode(prefix string) (pre *node) {
	pre = t.root
	for _, c := range prefix {
		hash := c - 'a'
		if pre.next[hash] == nil {
			return nil
		}
		pre = pre.next[hash]
	}
	return
}
