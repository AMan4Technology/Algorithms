package WordSearch

func findWords(board [][]byte, words []string) (results []string) {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return nil
	}
	trie := newNode(26)
	for _, word := range words {
		trie.Insert(word)
	}
	var (
		rowNum    = len(board)
		columnNum = len(board[0])
	)
	for i, row := range board {
		for j, c := range row {
			hash := c - 'a'
			if trie.next[hash] != nil {
				board[i][j] = '#'
				accessNode(string(c), board, trie.next[hash], i, j, rowNum, columnNum, &results)
				board[i][j] = c
			}
		}
	}
	return
}

var (
	dxs = []int{-1, 0, 1, 0}
	dys = []int{0, 1, 0, -1}
)

func accessNode(path string, board [][]byte, current *node, i, j, rowNum, columnNum int, results *[]string) {
	if current.isWord && !current.written {
		*results = append(*results, path)
		current.written = true
	}
	for k, dx := range dxs {
		x := i + dx
		if x == -1 || x == rowNum {
			continue
		}
		y := j + dys[k]
		if y == -1 || y == columnNum {
			continue
		}
		c := board[x][y]
		if c == '#' {
			continue
		}
		hash := c - 'a'
		if current.next[hash] != nil {
			board[x][y] = '#'
			accessNode(path+string(c), board, current.next[hash], x, y, rowNum, columnNum, results)
			board[x][y] = c
		}
	}
	return
}

func newNode(length int) *node {
	return &node{next: make([]*node, length)}
}

type node struct {
	next            []*node
	isWord, written bool
}

func (n *node) Insert(word string) {
	for _, c := range word {
		hash := c - 'a'
		if n.next[hash] == nil {
			n.next[hash] = newNode(26)
		}
		n = n.next[hash]
	}
	n.isWord = true
}
