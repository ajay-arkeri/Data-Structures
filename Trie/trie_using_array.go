package trie

//trie - a tree data structure that takes advantages of strings that have same prefix to store it exactly once
//use ---> auto complete feature
const AlphabetLength = 26

type Trie struct {
	root *Node
}

type Node struct {
	children [AlphabetLength]*Node
	isEnd    bool
}

func Init() *Trie {
	t := &Trie{root: &Node{}}
	return t
}

func (t *Trie) Insert(w string) {

	currentNode := t.root
	for _, ch := range w {
		index := ch - 'a'
		if currentNode.children[index] == nil {
			currentNode.children[index] = &Node{}
		}
		currentNode = currentNode.children[index]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	currentNode := t.root

	for _, ch := range w {
		index := ch - 'a'
		if currentNode.children[index] == nil {
			return false
		}
		currentNode = currentNode.children[index]
	}
	if currentNode.isEnd {
		return true
	} else {
		return false
	}
}

//Not sure if its the optimal delete , but it works
func (t *Trie) Delete(w string) {
	if !t.Search(w) {
		return
	}
	currentNode := t.root
	for _, ch := range w {
		index := ch - 'a'
		currentNode = currentNode.children[index]
	}
	currentNode.isEnd = false
}
