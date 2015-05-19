package jgl

import (
	// "fmt"
	"testing"
)

func TestReadMaterials(t *testing.T) {
	m, ok := DeserializeMaterials("testResources/materials.json")
	if ok != true {
		t.Error("Unable to deserialize materials")
	}
	SerializeMaterials("testResources/readMaterialsOutput.json", m)
	// for k, v := range m {
	//     fmt.Println(k, v)
	// }
}
