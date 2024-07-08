package code_test

import (
	"fmt"
	"gmr/gmr-x/ds"
	"gmr/gmr-x/offer/code"
	"testing"
)

func TestNum03(t *testing.T) {
	var linkNode ds.LinkNode
	head := linkNode.CreateHead(0)
	node0 := linkNode.Insert(1, head)
	node1 := linkNode.Insert(3, node0)
	node2 := linkNode.Insert(5, node1)
	node3 := linkNode.Insert(7, node2)
	_ = linkNode.Insert(9, node3)

	ln := code.Num03MethodOne(head)
	n := ln
	for n.NextNode != nil {
		fmt.Println(n.Data)
		n = n.NextNode
	}
}
