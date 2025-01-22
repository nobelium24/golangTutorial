package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {}

type ElementMap map[string]int

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}
type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Mapping() ElementMap {
	newMap := make(ElementMap)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	if doc.Type == html.ElementNode {
		newMap[doc.Data]++
	}
	return newMap
}
