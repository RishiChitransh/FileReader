package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

var fileCheckSum = []map[string]string {}

/**
Check if a file is duplicate:

Get a list of pre-calculated checksums present in DB and store as in-memory list for faster tallying.

For each file to be read, calculate checksum & tally the existing list:

if checksum present - reject as redundant
if checksum absent - add checksum to list of unique checksums

**/
func checkFileDuplicity(fileName string) bool {

	if len(fileCheckSum) == 0 {
		fileCheckSum = getImportedCheckSums()
	}

	isFileDuplicate := false
	content, err := os.Open(fileName)
	if err != nil{
		log.Fatal(err)
	}

	defer content.Close()

	copyBuf := make([]byte, 1024 * 1024)

	h := sha256.New()
	if _, err := io.CopyBuffer(h, content, copyBuf); err != nil {
		log.Fatal(err)
	}
	checkSum := hex.EncodeToString(h.Sum(nil))
	if len(fileCheckSum) != 0 {
		for _, val := range fileCheckSum {
			for _, value := range val {
				if value == checkSum {
				isFileDuplicate = true
				break
				}
			}
		}
		addCheckSumToList(fileName, checkSum)
	} else {
		addCheckSumToList(fileName, checkSum)
	}
	return isFileDuplicate
}

/**
Create a record of unique checksum and add the curated map to list
**/

func addCheckSumToList(fileName string, checkSum string) {
	fileRecordMap := map[string]string {
		"fileName" : fileName,
		"checkSum" : checkSum,
	}
	fileCheckSum = append(fileCheckSum, fileRecordMap)
}

/**
In case file is not duplicate, read the file line by line and pass teh records to the formatting function
**/
func readEachFile(fileName string) {
	newRead, err := os.Open(fileName)
	if err != nil{
		log.Fatal(err)
	}
	defer newRead.Close()

	var fileContent []string
	// read the file line by line using scanner
	scanner := bufio.NewScanner(newRead)

	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	formatRecord(fileName, fileContent)
}
