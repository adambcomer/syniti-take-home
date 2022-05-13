package main

import "strconv"

// Record is an individual record found in the dataset.
//
// Note: If a field in json object is missing or null, then the field will default to an empty string.
type Record struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Zip     string `json:"zip"`
}

// IsValid checks if a record has values for all the fields and the zip code is a valid US zip code.
func (r *Record) IsValid() bool {
	if r.Name == "" {
		return false
	}

	if r.Address == "" {
		return false
	}

	if len(r.Zip) != 5 {
		return false
	}
	i, err := strconv.Atoi(r.Zip)
	if err != nil {
		return false
	}
	if i < 1 || i > 99950 {
		return false
	}

	return true
}
