# FileReader
Code for reading/uploading file data


**Pre-Requisite**: Go, Docker

## Structure:

### initRead.go:

Is the entry point for the application. On running it will initiate a file read on the data directory to read input files

### DBUtils.go:

File for DB related activities:

	- Initiate DB Connection (using ENV Variables)
	- Retrieve Checksums
	- Commit file records in the database
	- Commit File meta data

	

### ParserUtils.go:

File for data parsing and data massaging:

	- Format Records for consistency
	- Check for mandatory fields
	- Supply non-mandatory fields


### fileUtils.go:

File for input file processing:

	- Calculate & verify check sum
	- Read file line by line

### docker-compose.yml

File for initiating PostGres service via Docker container

### Database Scripts.txt

Scripts for creating the default DB Structure

### data:

Folder contains the files for data input:

##NOTES:

	- Environment Variable for DB Connection - DB_URI



