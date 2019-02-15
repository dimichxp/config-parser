// Code generated by go generate; DO NOT EDIT.
package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestNbThreadNormal0(t *testing.T) {
	parser := &parsers.NbThread{}
	line := strings.TrimSpace("nbthread 4")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestNbThreadNormal1(t *testing.T) {
	parser := &parsers.NbThread{}
	line := strings.TrimSpace("nbthread 4 # comment")
	err := ProcessLine(line, parser)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := parser.Result(true)
	if err != nil {
		t.Errorf(err.Error())
	}
	var returnLine string
	if result[0].Comment == "" {
		returnLine = fmt.Sprintf("%s", result[0].Data)
	} else {
		returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
	}
	if line != returnLine {
		t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, line))
	}
}
func TestNbThreadFail0(t *testing.T) {
	parser := &parsers.NbThread{}
	line := strings.TrimSpace("nbthread")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestNbThreadFail1(t *testing.T) {
	parser := &parsers.NbThread{}
	line := strings.TrimSpace("---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
func TestNbThreadFail2(t *testing.T) {
	parser := &parsers.NbThread{}
	line := strings.TrimSpace("--- ---")
	err := ProcessLine(line, parser)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
	}
	_, err = parser.Result(true)
	if err == nil {
		t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
	}
}
