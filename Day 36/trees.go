/*
A tree is a data structure made up of nodes or vertices and edges without having any cycle. The tree with no nodes is called the null or empty tree.
A tree that is not empty consists of a root node and potentially many levels of additional nodes that form a hierarchy.

Browsers have this thing baked in, called the DOM - a cross-platform and language-independent application programming
interface, which treats internet documents as a tree structure wherein each node is an object representing a part of the document.
This means that when the browser reads your document’s HTML code it will load it and create a DOM out of it.

Now, we could treat the text nodes as separate Nodes, but we can make our lives simpler by assuming that any HTML element can have text in it.

*/

package main

type Node struct {
	tag      string
	text     string
	children []*Node
}

func main() {
	p := Node{
		tag:  "p",
		text: "This is a simple HTML document.",
	}

	h1 := Node{
		tag:  "h1",
		text: "Hello, World!",
	}

	html := Node{
		tag:      "html",
		children: []*Node{&p, &h1},
	}
}
