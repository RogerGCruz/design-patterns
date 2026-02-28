package main

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func (f *Folder) print(prefix string) {
	fmt.Printf("%s+ %s\n", prefix, f.name)
	for _, child := range f.children {
		child.print(prefix + prefix)
	}
}

func (f *Folder) clone() Inode {
	cloneFolder := &Folder{
		name: f.name + " - clone"}
	var children []Inode

	for _, child := range f.children {
		children = append(children, child.clone())
	}
	cloneFolder.children = children

	return cloneFolder
}
