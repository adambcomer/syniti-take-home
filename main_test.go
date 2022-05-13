package main

import "testing"

// Tests that records that are independently invalid are selected.
func TestGetInvalidRecordsInvalid(t *testing.T) {
	records := []Record{{
		Id:      "abc123",
		Name:    "",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}, {
		Id:      "abc124",
		Name:    "Adam Comer",
		Address: "",
		Zip:     "20500",
	}, {
		Id:      "abc125",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "",
	}}

	invalidRecords := GetInvalidRecords(records)

	if len(invalidRecords) != 3 {
		t.Fatalf("Number of Invalid Record should be %d, but got %d", 3, len(invalidRecords))
	}

	for i, r := range invalidRecords {
		if records[i] != r {
			t.Fatalf("Invalid Record at %d should be %v, but got %v", i, records[i], r)
		}
	}
}

// Tests that a record with a single duplicate is marked as invalid along with the first seen record.
func TestGetInvalidRecordsDuplicates(t *testing.T) {
	records := []Record{{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}, {
		Id:      "abc124",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}}

	invalidRecords := GetInvalidRecords(records)

	if len(invalidRecords) != 2 {
		t.Fatalf("Number of Invalid Record should be %d, but got %d", 2, len(invalidRecords))
	}

	for i, r := range invalidRecords {
		if records[i] != r {
			t.Fatalf("Invalid Record at %d should be %v, but got %v", i, records[i], r)
		}
	}
}

// Tests that a record with multiple duplicates is marked as invalid along with the first seen record.
//
// This test makes checks that the first seen record is put in the invalid records list once.
func TestGetInvalidRecordsDuplicatesMultiple(t *testing.T) {
	records := []Record{{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}, {
		Id:      "abc124",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}, {
		Id:      "abc125",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}}

	invalidRecords := GetInvalidRecords(records)

	if len(invalidRecords) != 3 {
		t.Fatalf("Number of Invalid Record should be %d, but got %d", 3, len(invalidRecords))
	}

	for i, r := range invalidRecords {
		if records[i] != r {
			t.Fatalf("Invalid Record at %d should be %v, but got %v", i, records[i], r)
		}
	}
}
