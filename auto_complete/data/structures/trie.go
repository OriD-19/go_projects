package structures

import "fmt"

type Trie struct {
	Root *Node
}

func NewTrie() Trie {
	return Trie{
		Root: NewNode(), // Empty node representing the root
	}
}

// Insert inserts a new word into the trie.
func (t Trie) Insert(s string) {
	curr := t.Root

	for _, c := range s {
		if curr.Children[c] != nil {
			curr = curr.Children[c]
		} else {
			toInsert := NewNode()
			toInsert.Character = c
			curr.Children[c] = toInsert
			curr = curr.Children[c]
		}
	}

	// Mark the end of the word in the last insert node
	curr.IsWord = true
}

func dfs(n *Node, prefix string, words []string) []string {
	if n == nil {
		return words
	}

	// If the current word is a node
	if n.IsWord {
		words = append(words, prefix+string(n.Character))
	}

	// Iterate over the children
	for _, value := range n.Children {
		words = dfs(value, prefix+string(n.Character), words)
	}

	return words
}

// SearchPattern takes a pattern and returns a list of possible matches
func (t Trie) SearchPattern(substr string) []string {

	if len(substr) == 0 {
		return []string{}
	}

	curr := t.Root
	for _, c := range substr {
		if curr.Children[c] == nil {
			// This means that the current pattern does not exist
			return []string{}
		}

		curr = curr.Children[c]
	}

	// Now current holds the node in which we have to perform the DFS
	words := []string{}

	// The search pattern except the last character (starting node of the DFS)
	words = dfs(curr, substr[:len(substr)-1], words)

	return words
}

// countChildrenNodes takes a map of Nodes and returns the count of
// values that are not nil
func countChildrenNodes(a map[rune]*Node) int {
	count := 0

	for _, val := range a {
		if val != nil {
			count++
		}
	}

	return count
}

// Delete takes a word and tries to delete it from the trie.
func (t Trie) Delete(s string) error {
	curr := t.Root

	var prefix *Node = nil
	var prefixChar char = 0

	for _, c := range s {
		if curr.Children[c] == nil {
			return fmt.Errorf("Delete: The word is not registered")
		}

		count := countChildrenNodes(curr.Children)

		if count > 1 {
			prefix = curr
			prefixChar = c
		}

		curr = curr.Children[c]
	}

	// We are at the node corresponding to the last character of the word
	if !curr.IsWord {
		return fmt.Errorf("Delete: The word is not registered")
	}

	countCurr := countChildrenNodes(curr.Children)

	// First case: word is a prefix of another word(s)
	if countCurr > 0 {
		curr.IsWord = false
		return nil
	}

	// Second case: word shares prefix with other word(s)
	if prefix != nil {
		prefix.Children[prefixChar] = nil
		return nil
	}

	// Third case: word does not share any prefix
	t.Root.Children[rune(s[0])] = nil
	return nil
}
