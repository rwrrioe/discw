package parser_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/rwrrioe/Discworld/parser"
)

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

	expected := []map[string]interface{}{
		{
			"Items": []interface{}{},
			"Value": "AcceptingApplications",
			"Text":  "Прием заявок",
		},

		{
			"Items": []interface{}{},
			"Value": "AcceptingAppUserEnter",
			"Text":  "Прием заявок; Регистрация участников в аукционном зале",
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

	convertMap := []map[string]interface{}(*parsed)
	if !reflect.DeepEqual(convertMap, expected) {
		t.Errorf("Expected: %v, got %v", expected, *parsed)
		t.Logf("Parsed value %#v", parsed)
		t.Logf("Expected value %#v", expected)
	}
}
