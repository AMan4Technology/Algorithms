package AlienDictionary

import (
	"strings"

	"Algorithms"
)

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 {
		words = append(words, "")
	}
	dictionary := make(map[byte]*graphNode, 26)
	for length, i := len(words), 0; i < length-1; i++ {
		var (
			len0 = len(words[i])
			len1 = len(words[i+1])
			j    = 0
		)
		for lenOfWord := Algorithms.MinOfTwo(len0, len1); j < lenOfWord; j++ {
			node0, node1 := dictionary[words[i][j]], dictionary[words[i+1][j]]
			if node0 == nil {
				node0 = &graphNode{Next: make(map[byte]*graphNode, 1)}
				dictionary[words[i][j]] = node0
			}
			if node1 == nil {
				node1 = &graphNode{Next: make(map[byte]*graphNode, 0)}
				dictionary[words[i+1][j]] = node1
			}
			if words[i][j] != words[i+1][j] {
				if node1.Next[words[i][j]] != nil {
					return ""
				}
				if node0.Next[words[i+1][j]] == nil {
					node0.Next[words[i+1][j]] = node1
					node1.inDegree++
				}
				j++
				break
			}
		}
		for k := j; k < len0; k++ {
			if dictionary[words[i][k]] == nil {
				dictionary[words[i][k]] = &graphNode{Next: make(map[byte]*graphNode, 0)}
			}
		}
		for k := j; k < len1; k++ {
			if dictionary[words[i+1][k]] == nil {
				dictionary[words[i+1][k]] = &graphNode{Next: make(map[byte]*graphNode, 0)}
			}
		}
	}
	// 其余字符按照默认字典顺序输出
	// queue := NewMin(len(dictionary))
	// for key, node := range dictionary {
	//     if node.inDegree == 0 {
	//         queue.Push(int(key), int(key))
	//     }
	// }
	// var sb strings.Builder
	// for !queue.Empty() {
	//     var (
	//         key  = queue.Pop().Value
	//         node = dictionary[byte(key)]
	//     )
	//     sb.WriteByte(byte(key))
	//     for k, next := range node.Next {
	//         next.inDegree--
	//         if next.inDegree == 0 {
	//             queue.Push(int(k), int(k))
	//         }
	//     }
	// }
	// 其余字符无序输出
	var queue []byte
	for key, node := range dictionary {
		if node.inDegree == 0 {
			queue = append(queue, key)
		}
	}
	var sb strings.Builder
	for len(queue) != 0 {
		var (
			key  = queue[0]
			node = dictionary[key]
		)
		queue = queue[1:]
		sb.WriteByte(key)
		for k, next := range node.Next {
			next.inDegree--
			if next.inDegree == 0 {
				queue = append(queue, k)
			}
		}
	}
	if sb.Len() != len(dictionary) {
		return ""
	}
	return sb.String()
}

type graphNode struct {
	inDegree int
	Next     map[byte]*graphNode
}
