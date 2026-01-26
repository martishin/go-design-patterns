package main

import (
	"fmt"

	"github.com/martishin/go-design-patterns/patterns/creational/prototype/internal/fs"
)

func main() {
	file1 := fs.NewFile("File1")
	file2 := fs.NewFile("File2")
	file3 := fs.NewFile("File3")

	folder1 := fs.NewFolder("Folder1", file1)
	folder2 := fs.NewFolder("Folder2", folder1, file2, file3)

	fmt.Println("Printing hierarchy for Folder 2")
	fmt.Println(folder2)

	fmt.Println("\nPrinting hierarchy for cloned Folder 2")
	clonedFolder := folder2.Clone()
	fmt.Println(clonedFolder)

	fmt.Println("\nPrinting File 1")
	fmt.Println(file1)

	fmt.Println("\nPrinting cloned File 1")
	clonedFile := file1.Clone()
	fmt.Println(clonedFile)
}
