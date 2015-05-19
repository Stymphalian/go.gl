package jgl

import (
	// "fmt"
	"testing"
)

func TestReadMaterials(t *testing.T) {
	m := DeserializeMaterials("testResources/materials.json")
	SerializeMaterials("testResources/readMaterialsOutput.json", m)
	// for k, v := range m {
	//     fmt.Println(k, v)
	// }
}
