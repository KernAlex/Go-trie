package trie

/*
maps each element to the
tree
*/
type tNode struct {
	letters map[string]*tNode
	end bool
}

/*
Basic trieTree structure
*/
type trieTree struct {
	root *tNode
}
func NewTrieTree() trieTree{
	var tree = trieTree{}
	tree.root = &tNode{make(map[string]*tNode), false}
	return tree
}

/*
insert function
*/
func (t *trieTree) Insert(s string) {
	temp := t.root
	i := 0
	for ; i < len(s) - 1; i++ {
		_, ok := (*temp).letters[string(s[i])]
		if ok {
			temp = temp.letters[string(s[i])]
		} else {
			temp.letters[string(s[i])] = &tNode{make(map[string]*tNode), false}
			temp = temp.letters[string(s[i])]
		}
	}
	temp.end = true
	_, ok := (*temp).letters[string(s[i])]
	if ok {
		temp = temp.letters[string(s[i])]
	} else {
		temp.letters[string(s[i])] = &tNode{make(map[string]*tNode), false}
		temp = temp.letters[string(s[i])]
	}
}