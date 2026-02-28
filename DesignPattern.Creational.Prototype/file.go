package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(prefix string) {
	fmt.Printf("%s- %s\n", prefix, f.name)
}

func (f *File) clone() Inode {
	return &File{
		name: f.name,
	}
}
