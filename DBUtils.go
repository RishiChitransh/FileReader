package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

var fileMetaData = []map[string]string {}

/** Basic DB setup details **/
const (
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "test_db"
	sslmode = "disable"
)
/**
Function to return connection string for the PostGres connection

Uses - Environment variable DB_URI for getting database URL

**/
func getConnectionString() string{
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_URI"), port, user, password, dbname, sslmode)

	return connStr
}

/**
Function to return DB connection object that can be used for database calls

**/
func returnDbObject() *sql.DB {
	connStr := getConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

/**
Function to return specific checksum for each file meta-data to be inserted in DB.
Parses the global file check sum list created at the start of the program

**/
func getCheckSum(fileName string) string {
	var retFileCheckSum = ""
	for _, val :=  range fileCheckSum {
		if val["fileName"] == fileName {
			retFileCheckSum = val["checkSum"]
		}
	}
	return retFileCheckSum
}

/**
Get a list of pre-stored check sum in the DB.
This is to confirm whether the file under inspection is redundant without having to store full file contents in the DB

**/
func getImportedCheckSums() []map[string]string {
	listOfCheckSum := []map[string]string {}

	dbObj := returnDbObject()
	defer dbObj.Close()

	rows, err := dbObj.Query("SELECT \"FILENAME\", \"CHECKSUM\" FROM file_log")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		fileName string
		checkSum string
	)
	for rows.Next() {
		err = rows.Scan(&fileName, &checkSum)
		if err != nil {
			panic(err)
		}
		checkSumRecord := map[string]string {
			"fileName" : fileName,
			"checkSum" : checkSum,
		}
		listOfCheckSum = append(listOfCheckSum, checkSumRecord)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return listOfCheckSum
}

/**
Function for massaging data  -

Reject in case mandatory fields are missing
Supply blank values for consistency to missing keys, in case all mandatory keys are present.

If complete, insert the row in DB successfully
**/
func curateSetForStorage(fileMap []map[string]string, source string) {
	fileMeta := map[string]string{
		"file_name" : source,
		"import_time" : time.Now().String(),
		"total_count" : strconv.Itoa(len(fileMap)),
		"checksum" : getCheckSum(source),
	}

	db := returnDbObject()
	defer db.Close()

	count := 0
	for _, record := range fileMap {
		if !checkForMandatory(record) {
			fmt.Println("Inappropriate record. Missing critical information!!")
			continue
		} else {
			count++
			record = supplyMissingValues(record)
			sqlStatement := "INSERT INTO call_log (\"CALL_DATE\", \"CALL_DISPOSITION\", \"PHONE_NUM\", \"FIRST_NAME\", \"LAST_NAME\", \"ADDRESS1\", \"ADDRESS2\", \"CITY\", \"STATE\", \"ZIP\") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
			_, err := db.Exec(sqlStatement, record["call_date"], record["call_disposition"], record["phone"], record["first"], record["last"], record["address1"], record["address2"], record["city"], record["state"], record["zip"])
			if err != nil {
				panic(err)
			} else {
				fmt.Println("\nRow inserted successfully!")
			}
		}
	}

	fileMeta["import_count"] = strconv.Itoa(count)
	fileMetaData = append(fileMetaData, fileMeta)
}

/**
Function to feed File meta data in the DB

**/

func updateFileMeta() {

	db := returnDbObject()
	defer db.Close()

	for _,val := range fileMetaData {

		sqlStatement := "INSERT INTO file_log (\"IMPORT_DATE_TIME\", \"FILENAME\", \"TOTAL_ROWS\", \"IMPORTED_ROWS\", \"CHECKSUM\") VALUES ($1,$2,$3,$4,$5)"
		_, err := db.Exec(sqlStatement, val["import_time"], val["file_name"], val["total_count"], val["import_count"], val["checksum"])

		if err != nil {
			panic(err)
		} else {
			fmt.Println("\nFile info inserted successfully!")
		}
	}
}
