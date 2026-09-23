type TrieNode struct {
	children [26]*TrieNode
	word bool
}

func NewTrieNode() *TrieNode{
	return &TrieNode{}
}

type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{root: NewTrieNode()}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
	for _, c := range word{
		index := c-'a'
		if curr.children[index] == nil{
			curr.children[index] = NewTrieNode()
		}
		curr = curr.children[index]
	}
	curr.word = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool {
	curr := root
	for i := j; i < len(word); i++ {
		c := word[i]
		if c == '.'{
			for _, child := range curr.children{
				if child != nil && this.dfs(word, i+1, child){
					return true
				}
			}
			return false
		} else{
			index := c-'a'
			if curr.children[index] == nil{
				return false
			}
			curr = curr.children[index]
		}
	}
	return curr.word
}
