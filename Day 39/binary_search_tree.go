/*
In computer science, a binary search tree (BST), also called an ordered or sorted binary tree, is a rooted binary tree data structure whose internal nodes
each store a key greater than all the keys in the node’s left subtree and less than those in its right subtree. The time complexity of operations on the
binary search tree is directly proportional to the height of the tree.

Binary search trees allow binary search for fast lookup, addition, and removal of data items, and can be used to implement dynamic sets and lookup tables.
Since the nodes in a BST are laid in such as way that each comparison skips about half of the remaining tree, the lookup performance is proportional to
that of binary logarithm.

The performance of a binary search tree is dependent on the order of insertion of the nodes into the tree; several variations of the binary search tree can
be built with guaranteed worst-case performance. The basic operations include: search, traversal, insert and delete. BSTs with guaranteed worst-case
complexities perform better than an unsorted array, which would require linear search time.

The complexity analysis of BST shows that, on average, the insert, delete and search takes {\displaystyle O(\log n)}O(\log n) for {\displaystyle n}n nodes.
In the worst case, they degrade to that of a singly linked list: {\displaystyle O(n)}O(n).
*/

package main

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.key)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.key)
		printInOrder(n.right)
	}
}

func main() {
	var t Tree

	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('I')
	t.insert('H')

	fmt.Printf("Pre Order: ")
	printPreOrder(t.root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	printPostOrder(t.root)
	fmt.Println()
	fmt.Printf("In Order: ")
	printInOrder(t.root)
	fmt.Println()
}
