package parser_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/rwrrioe/Discworld/parser"
)

func toString(m map[string]interface{}) *map[string]string {
	strMap := make(map[string]string)
	for k, v := range m {
		kStr := fmt.Sprint(k)
		vStr := fmt.Sprint(v)
		strMap[kStr] = vStr
	}

	return &strMap
}

func TestParser(t *testing.T) {
	toParse := `[
  {
    "Value": "AcceptingApplications",
    "Text": "Прием заявок",
    "Items": []
  },
  {
    "Value": "AcceptingAppUserEnter",
    "Text": "Прием заявок; Регистрация участников в аукционном зале",
    "Items": []
  }
]`

	expected := []map[string]string{
		{
			"Value": "AcceptingApplications",
			"Text":  "Прием заявок",
			"Items": "[]",
		},

		{
			"Value": "AcceptingAppUserEnter",
			"Text":  "Прием заявок; Регистрация участников в аукционном зале",
			"Items": "[]",
		},
	}

	tempfile := filepath.Join(os.TempDir(), "test_parsedata.json")
	if err := os.WriteFile(tempfile, []byte(toParse), 0644); err != nil {
		t.Fatalf("Failed to write temp JSON file: %v", err)
	}

	p := parser.NewParser()
	parsed, err := p.Parse(tempfile)
	if err != nil {
		t.Fatal(err)
	}

	for k, v := range expected {
		if toString((*parsed)[k]) != &v {
			t.Errorf("Expected %v -> %s, got %s", k, v, (*parsed)[k])
		}
	}
}
