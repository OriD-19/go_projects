package structures

// Just cause I'm still a C programmer lmao
type char = rune

type Node struct {
	Character char
	Children  map[char]*Node
	IsWord    bool
}

func NewNode() *Node {
	return &Node{
		Character: -1, // This represents an empty node
		Children:  make(map[char]*Node),
		IsWord:    false,
	}
}
