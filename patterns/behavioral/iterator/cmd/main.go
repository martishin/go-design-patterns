package main

import (
	"fmt"
	"log"

	internalfs "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/internal/filesystem"
	iterator "github.com/martishin/go-design-patterns/patterns/behavioral/iterator/pkg/filesystem"
)

func main() {
	tree, err := buildFileTree()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Depth-first traversal:")
	printPaths(tree.CreateDepthFirstIterator())

	fmt.Println()
	fmt.Println("Breadth-first traversal:")
	printPaths(tree.CreateBreadthFirstIterator())
}

func buildFileTree() (*internalfs.FileTree, error) {
	mainGo, err := internalfs.NewFile("main.go")
	if err != nil {
		return nil, err
	}

	serverGo, err := internalfs.NewFile("server.go")
	if err != nil {
		return nil, err
	}

	handlerGo, err := internalfs.NewFile("handler.go")
	if err != nil {
		return nil, err
	}

	readme, err := internalfs.NewFile("README.md")
	if err != nil {
		return nil, err
	}

	cmd, err := internalfs.NewDirectory("cmd", mainGo)
	if err != nil {
		return nil, err
	}

	internal, err := internalfs.NewDirectory("internal", serverGo, handlerGo)
	if err != nil {
		return nil, err
	}

	app, err := internalfs.NewDirectory("app", cmd, internal, readme)
	if err != nil {
		return nil, err
	}

	return internalfs.NewFileTree(app)
}

func printPaths(iterator iterator.Iterator) {
	for iterator.HasNext() {
		entry := iterator.Next()
		fmt.Printf("- %s\n", entry.Path())
	}
}
