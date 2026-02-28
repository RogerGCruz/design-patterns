package main

func main() {
	file1 := &File{name: "File1.txt"}
	file2 := &File{name: "File2.txt"}
	file3 := &File{name: "File3.txt"}

	folder1 := &Folder{name: "Folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "Folder2"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("Design Patterns")
}
