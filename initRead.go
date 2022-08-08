package main

import (
	"fmt"
	"log"
	"os"
)

/**
Entry point for the application.

Target the directory containing the data files and initiate the flow.
**/
func main() {
	dirName := "data/"
	dataDir, err := os.ReadDir(dirName)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dataDir {
		if !checkFileDuplicity(dirName+f.Name()) {
			readEachFile(dirName+f.Name())
			fmt.Println("End of File!!")
		} else {
			fmt.Println("Duplicate file found. Skipping!!")
		}
	}
	updateFileMeta()
	fmt.Println("Data processing completed. Thank You!!")

	/*
	Enable below to read from storage.
	Commenting out as currently free available data is ~124 MB in form of a .tar file
	*/
	//readFromStorage()
}





