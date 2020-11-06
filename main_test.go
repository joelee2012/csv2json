// main_test
package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestReadCsv2Dict(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "csv2json")
	check(err)
	csvFile.WriteString("field1,field2\n")
	csvFile.WriteString("value1,value2\n")
	defer os.Remove(csvFile.Name())

	expected := map[string]string{"field1": "value1", "field2": "value2"}
	json := ReadCsv2Dict(csvFile.Name())
	if len(json) != 1 || !reflect.DeepEqual(json[0], expected) {
		t.Errorf("expect: [%s], but got: %s", expected, json)
	}
}
