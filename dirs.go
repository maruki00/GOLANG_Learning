package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// var newFile os.File
	// var errorFile error
	//Check if file already exists
	fileInfo, errFileExists := os.Stat("newfile.bin")
	if os.IsNotExist(errFileExists) {
		fmt.Println("File Doesn't Exists! ")
		// Create New File
		newFile, errorFile := os.Create("newFile.txt")
		fmt.Println("File was Created! ")
	}

	fmt.Printf("File Doesn't exists! %d", -1)
	os.Exit(0)
	if fileInfo != nil {
		fmt.Printf("the name is : %s", fileInfo.Name())
	}
	//Create new file
	emptyFile, err := os.Create("newfile.bin")
	if err != nil {
		log.Fatal("Could not make new file")
	}
	fmt.Print("New Empty File has been created! %s", emptyFile.Name())

	//Make new Directory
	errDir := os.Mkdir("new dir", 750)
	if errDir != nil {
		log.Fatal("Could not make new dir!")
	}

	fmt.Printf("New Dir has been created!")

	//Rename The file or Directory
	errRenameDir := os.Rename("new Dir", "testDir")
	errRenameFile := os.Rename("newfile.bin", "file.txt")

	if nil != errRenameDir {
		log.Fatal("Could not rename the dir")
	}

	if nil != errRenameFile {
		log.Fatal("Could Not rename the file!")
	}

}
