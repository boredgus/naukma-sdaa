package src

import (
	"context"
	"fmt"
	"github.com/goccy/go-graphviz"
)

func DrawGraph(tree *RBTree, name string, iteration int) {
	ctx := context.Background()
	g, err := graphviz.New(ctx)
	if err != nil {
		panic(err)
	}

	graph, err := g.Graph()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := graph.Close(); err != nil {
			panic(err)
		}
		if err := g.Close(); err != nil {
			panic(err)
		}
	}()

	if tree.root != nil {
		head, err := graph.CreateNodeByName(fmt.Sprintf("%v", tree.root.Value))
		if err != nil {
			panic(err)
		}
		addNode(graph, head, tree.root)
	}

	path := fmt.Sprintf("lab3_tree_%s_%v.png", name, iteration)
	if err := g.RenderFilename(ctx, graph, graphviz.PNG, path); err != nil {
		panic(err)
	}
}

func addNode(graph *graphviz.Graph, parentNode *graphviz.Node, node *RBTreeNode) {
	if node.Left != nil {
		left, err := graph.CreateNodeByName(fmt.Sprintf("%v", node.Left.Value))
		if err != nil {
			panic(err)
		}
		if node.Left.IsRed {
			left.SetColor("red")
		}

		edge, err := graph.CreateEdgeByName("e", parentNode, left)
		if err != nil {
			panic(err)
		}
		edge.SetLabel("L")

		addNode(graph, left, node.Left)
	}

	if node.Right != nil {
		right, err := graph.CreateNodeByName(fmt.Sprintf("%v", node.Right.Value))
		if err != nil {
			panic(err)
		}
		if node.Right.IsRed {
			right.SetColor("red")
		}

		edge, err := graph.CreateEdgeByName("e", parentNode, right)
		if err != nil {
			panic(err)
		}
		edge.SetLabel("R")

		addNode(graph, right, node.Right)
	}
}
