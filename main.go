package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadRecords loads and parses the Records from the filesystem.
func LoadRecords(filename string) ([]Record, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var records []Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// RecordKey is a struct used to identify the duplicate records.
//
// The RecordKey's primary use is to be a key in a map. A separate struct is necessary because Golang doesn't support
// hashing on a subset of a struct's fields.
type RecordKey struct {
	Name    string
	Address string
	Zip     string
}

// GetInvalidRecords filters out the invalid records and the duplicate records.
//
//	This function runs in O(N) where N is the number of records.
func GetInvalidRecords(records []Record) []Record {
	recordsMap := make(map[RecordKey]Record)
	seenDuplicateIds := make(map[string]bool)
	var invalidRecords []Record
	for _, r := range records {
		if !r.IsValid() {
			invalidRecords = append(invalidRecords, r)
			continue
		}

		key := RecordKey{
			Name:    r.Name,
			Address: r.Address,
			Zip:     r.Zip,
		}
		if dupRecord, ok := recordsMap[key]; ok {
			// Prevents repeating the duplicate record id that has more than 1 duplicate
			if _, ok := seenDuplicateIds[dupRecord.Id]; ok {
				invalidRecords = append(invalidRecords, r)
			} else {
				invalidRecords = append(invalidRecords, dupRecord, r)
				seenDuplicateIds[dupRecord.Id] = true
			}
		} else {
			recordsMap[key] = r
		}
	}

	return invalidRecords
}

// Application Entrypoint.
//
// Application Steps:
// 1. Loads and parses the records from disk.
// 2. Filters out the invalid records.
// 3. Prints the id of invalid records in the console.
func main() {
	records, err := LoadRecords("./data.json")
	if err != nil {
		panic(err)
	}

	invalidRecords := GetInvalidRecords(records)
	for _, r := range invalidRecords {
		fmt.Printf("%s\n", r.Id)
	}
}
