package _5_Trie

type TrieNode struct {
	data         interface{}
	children     [26]*TrieNode
	isEndingChar bool
}

var root = NewTrieNode("/")

func NewTrieNode(data interface{}) *TrieNode {
	return &TrieNode{
		data: data,
	}
}

func Insert(text string) {
	p := root
	for i := 0; i < len(text); i++ {
		index := rune(text[i]) - 97
		if p.children[index] == nil {
			newNode := NewTrieNode(text[i])
			p.children[index] = newNode
		}
		p = p.children[index]
	}
	p.isEndingChar = true
}

func Find(text string) bool {
	p := root
	for i := 0; i < len(text); i++ {
		index := text[i] - 97
		if p.children[index] == nil {
			return false
		}
		p = p.children[index]
	}
	if !p.isEndingChar {
		return false
	}
	return true
}
