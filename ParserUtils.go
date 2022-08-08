package main

import (
	"fmt"
	"strings"
)


/**
Function to format imported data into viable format

Checks multiple combinations of name for key fields:

for e.g.
for CALL DATE - Call, call, Call & Date, created_at

Uses string lower function to remove case dependency

TODO: can be optimized by using RegExp??
**/
func formatRecord(name string, file []string) {

	var fullSet = []map[string]string {}
	fmt.Println(file)
	titleField := file[0]
	file = append(file[:0],file[1:]...)

	titleSet := strings.Split(titleField, ",")
	for _, line := range file {
		splitSet := strings.Split(line,",")
		var recordMap = map[string]string {}
		for index, val := range titleSet {
			//storing only required information. Ignoring rest
			if (strings.Contains(strings.ToLower(val), "call") || strings.Contains(val, "create")) && !strings.Contains(val, "count"){
				recordMap["call_date"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "disp") || strings.Contains(strings.ToLower(val), "status") {
				recordMap["call_disposition"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "phone") {
				recordMap["phone"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "first") {
				recordMap["first"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "last") {
				recordMap["last"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "address1") {
				recordMap["address1"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "address2") {
				recordMap["address2"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "city") {
				recordMap["city"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "state") {
				recordMap["state"] = splitSet[index]
			} else if strings.Contains(strings.ToLower(val), "zip") {
				recordMap["zip"] = splitSet[index]
			}
		}
		fullSet = append(fullSet, recordMap)
	}
	curateSetForStorage(fullSet, name)
}

/**
Function to check if mandatory fields are present in the record or not
**/
func checkForMandatory(record map[string]string) bool{
	var mandatory = []string { "call_date","call_disposition","phone"}
	var isMandatoryPresent bool = true
	for _,field := range mandatory {
		if record[field] == "" {
			isMandatoryPresent = false
		}
	}
	return isMandatoryPresent
}

/**
Function to provide add non-mandatory absent keys with blank values
in case all mandatory fields are present
**/
func supplyMissingValues(record map[string]string) map[string]string {
	infoFields := []string {"first", "last", "address1", "address2", "state", "city", "zip"}
	for _,val := range infoFields {
		if _, ok := record[val]; !ok {
			record[val] = ""
		}
	}
	return record
}
