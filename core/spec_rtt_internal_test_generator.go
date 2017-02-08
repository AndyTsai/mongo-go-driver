// +build ignore

package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	outFileName = "spec_rtt_internal_test.go"
	testsDir    = "../specifications/source/server-selection/tests/rtt/"
)

func main() {
	tests, err := loadTests()
	if err != nil {
		log.Fatalf("Failed to load tests: %s", err)
		return
	}

	tmpl, err := getTemplate()
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
		return
	}

	file, err := os.Create(outFileName)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, tests)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
		return
	}
}

type rttTest struct {
	Name                   string
	AverageRTTSet          bool
	AverageRTT             float64
	AverageRTTUnmarshalled interface{} `yaml:"avg_rtt_ms"`
	NewRTT                 float64     `yaml:"new_rtt_ms"`
	NewAverageRTT          float64     `yaml:"new_avg_rtt"`
}

func loadTests() ([]rttTest, error) {
	testFiles, err := ioutil.ReadDir(testsDir)
	if err != nil {
		return nil, err
	}

	tests := make([]rttTest, 0, 7)

	for _, file := range testFiles {
		filename := file.Name()
		ext := path.Ext(filename)
		if ext != ".yml" {
			continue
		}

		fullpath := path.Join(testsDir, filename)
		content, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return nil, err
		}

		var test rttTest
		err = yaml.Unmarshal(content, &test)
		if err != nil {
			return nil, err
		}

		switch typed := test.AverageRTTUnmarshalled.(type) {
		case string:
			if typed != "NULL" {
				return nil, errors.New("unexpected string in AverageRTTUnmarshalled field")
			}
			test.AverageRTT = 0
			test.AverageRTTSet = false
		case int:
			test.AverageRTT = float64(typed)
			test.AverageRTTSet = true
		case float32:
			test.AverageRTT = float64(typed)
			test.AverageRTTSet = true
		case float64:
			test.AverageRTT = typed
			test.AverageRTTSet = true
		default:
			return nil, fmt.Errorf(
				"AverageRTTUnmarshalled contained unexpected type %T",
				test.AverageRTTUnmarshalled,
			)
		}

		test.Name = file.Name()[0 : len(filename)-len(ext)]

		tests = append(tests, test)
	}

	return tests, nil
}

func getTemplate() (*template.Template, error) {
	content := `package core

import (
	"testing"
	"time"
)
{{range .}}
func TestRTT_{{.Name}}(t *testing.T) {
	t.Parallel()

	var (
		startingAverageRTT = time.Duration({{.AverageRTT}}*float64(time.Millisecond))
		newRTT = time.Duration({{.NewRTT}}*float64(time.Millisecond))
		newAverageRTT = time.Duration({{.NewAverageRTT}}*float64(time.Millisecond))
	)

	mon := ServerMonitor{
		averageRTT: startingAverageRTT,
		averageRTTSet: {{.AverageRTTSet}},
	}

	expected := newAverageRTT
	actual := mon.updateAverageRTT(newRTT)

	if expected != actual {
		t.Errorf("\nStarting Avg RTT: %v\nNew RTT: %v\nExpected New Avg RTT: %v\nActual New Avg RTT: %v",
			startingAverageRTT.Seconds()*1000,
			newRTT.Seconds()*1000,
			expected.Seconds()*1000,
			actual.Seconds()*1000,
		)
	}
}
{{end}}
`
	tmpl, err := template.New("").Parse(content)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
