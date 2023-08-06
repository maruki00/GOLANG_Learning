package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	var errorFile error
	//Check if file already exists
	fileInfo, errFileExists := os.Stat("newfile.bin")
	if os.IsNotExist(errFileExists) {
		fmt.Println("File Doesn't Exists! ")
		// Create New File
		newFile, errorFile = os.Create("newFile.txt")
		if errorFile != nil {
			log.Fatal("Could not make new file")
		}
		fmt.Println("File was Created! ")
	}
	defer newFile.Close()

	//Write infoirmation to file!
	newFile.WriteString("hello world")
	//Write Bytes
	newFile.Write([]byte("kjfhgskjdfghlsd"))

	//Read From file
	scanner := bufio.NewScanner(newFile)
	for scanner.Scan() {
		fmt.Printf("Data : %s", scanner.Text())
	}

	//Get File info
	fileInfo, fileInfoError := os.Stat("newFile.txt")
	if fileInfoError != nil {
		log.Fatal("Something wrong !")
	}
	fmt.Printf("File Name : %s \n", fileInfo.Name())
	fmt.Printf("File Mode : %s \n", fileInfo.Mode())
	fmt.Printf("File Size : %s \n", fileInfo.Size())
	fmt.Printf("File ModTime : %s \n", fileInfo.ModTime())

	if fileInfo != nil {
		fmt.Printf("the name is : %s", fileInfo.Name())
	}

	//Rename The file or Directory
	errRenameFile := os.Rename("newfile.txt", "file.txt")

	if nil != errRenameFile {
		log.Fatal("Could Not rename the file!")
	}

	newFileTest, errTest := os.Create("newFile.test")
	if errTest != nil {
		log.Fatalf("Could not create file! %s ", newFileTest.Name())
		os.Exit(-1)
	}
	//Delete File
	errDelFile := os.Remove("newFile.test")
	if errDelFile != nil {
		log.Fatal("Could Not Delete the file ")
	}

}
