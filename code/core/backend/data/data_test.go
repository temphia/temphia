package data

import (
	"testing"
)

func TestData(t *testing.T) {

	ass := DefaultNew()

	_, err := ass.GetAsset("build/not_found_file_xyz")
	if err == nil {
		t.Fatal("should not find this file")
	}

	_, err = ass.GetAsset("build/sample.txt")
	if err != nil {
		t.Fatal(err)
	}
}
