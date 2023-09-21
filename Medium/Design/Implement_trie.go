package Design

type Node struct {
	Value    rune
	Children map[rune]*Node
	TheEnd   bool
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{
		Root: &Node{
			TheEnd:   true,
			Children: make(map[rune]*Node),
		},
	}
}

func (this *Trie) Insert(word string) {
	temp := this.Root

	for _, el := range word {
		if _, ok := temp.Children[el]; !ok {
			temp.Children[el] = &Node{
				Value:    el,
				TheEnd:   false,
				Children: make(map[rune]*Node),
			}
		}

		temp = temp.Children[el]
	}

	temp.TheEnd = true
}

func (this *Trie) Search(word string) bool {
	temp := this.Root

	for _, el := range word {
		if _, ok := temp.Children[el]; !ok {
			return false
		}

		temp = temp.Children[el]
	}

	return temp.TheEnd == true
}

func (this *Trie) StartsWith(prefix string) bool {
	temp := this.Root

	for _, el := range prefix {
		if _, ok := temp.Children[el]; !ok {
			return false
		}

		temp = temp.Children[el]
	}

	return true
}
