package jgl

import (
	"encoding/json"
	"fmt"
	"os"
)

// Create a dictionary of materials.(key = string, value = Material)
// Read the material data from the specified JSON file.
// No error handling on this.
func DeserializeMaterials(filename string) (map[string]Material, bool) {
	out := make(map[string]Material)

	fin, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file : ", filename)
		return out, false
	}
	defer fin.Close()

	decoder := json.NewDecoder(fin)
	if err := decoder.Decode(&out); err != nil {
		fmt.Println("Unable to decode: ", filename)
		return out, false
	}

	return out, true
}

// Given a dictionary [string]Material, serialize the data into a JSON file.
// Note there is not error handling for any operations.
func SerializeMaterials(filename string, materials map[string]Material) {
	fout, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Unable to open file for writing :", filename, err)
		return
	}
	defer fout.Close()

	data, err := json.MarshalIndent(materials, "", "  ")
	if err != nil {
		fmt.Println("Unable to encode materials to ", filename)
		return
	}
	fout.Write(data)
}
