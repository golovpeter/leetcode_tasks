package Trie

type Node struct {
	Val      rune
	Children map[rune]*Node
	TheEnd   bool
}

type Trie struct {
	Root *Node
}

func NewPrefixTrie() *Trie {
	return &Trie{
		Root: &Node{
			Children: make(map[rune]*Node),
			TheEnd:   true,
		},
	}
}

func (t *Trie) Insert(word string) {
	temp := t.Root

	for _, el := range word {
		if _, ok := temp.Children[el]; !ok {
			temp.Children[el] = &Node{
				Val:      el,
				TheEnd:   false,
				Children: make(map[rune]*Node),
			}
		}

		temp = temp.Children[el]
	}

	temp.TheEnd = true
}

func (t *Trie) Search(word string) bool {
	temp := t.Root

	for _, el := range word {
		if _, ok := temp.Children[el]; !ok {
			return false
		}

		temp = temp.Children[el]
	}

	return temp.TheEnd
}

func longestWord(words []string) string {
	trie := NewPrefixTrie()

	for _, el := range words {
		trie.Insert(el)
	}

	var maxWord string
	for _, el := range words {
		if len(el) > len(maxWord) || len(el) == len(maxWord) && el < maxWord {
			if allPrefixesExist(trie, el) {
				maxWord = el
			}
		}
	}

	return maxWord
}

func allPrefixesExist(t *Trie, w string) bool {
	for i := 1; i <= len(w); i++ {
		if !t.Search(w[:i]) {
			return false
		}
	}

	return true
}
