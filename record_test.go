package main

import "testing"

// Tests that a known valid record is valid.
func TestRecordIsValid(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}

	if !record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, the Record.IsValid was %v", false, record.IsValid())
	}
}

// Tests that a record with a missing name is invalid.
func TestRecordIsValidName(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "20500",
	}

	if record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, the Record.IsValid was %v", false, record.IsValid())
	}
}

// Tests that a record with a missing address is invalid.
func TestRecordIsValidAddress(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "",
		Zip:     "20500",
	}

	if record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, but got %v", false, record.IsValid())
	}
}

// Tests that a record with a missing zip is invalid.
func TestRecordIsValidZip(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "",
	}

	if record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, but got %v", false, record.IsValid())
	}
}

// Tests that a record with a non-five-digit zip is invalid.
func TestRecordIsValidZipInvalidLength(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "00000000000000",
	}

	if record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, but got %v", false, record.IsValid())
	}
}

// Tests that a record with an out of range zip is invalid.
func TestRecordIsValidZipInvalidValue(t *testing.T) {
	record := Record{
		Id:      "abc123",
		Name:    "Adam Comer",
		Address: "1600 Pennsylvania Avenue",
		Zip:     "999999",
	}

	if record.IsValid() {
		t.Fatalf("Record.IsValid should be %v, but got %v", false, record.IsValid())
	}
}
