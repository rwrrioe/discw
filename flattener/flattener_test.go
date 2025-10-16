package flattener_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rwrrioe/Discworld/flattener"
)

func TestFlattenerDefault(t *testing.T) {
	jsonContent := `[
		{"Value": "Running", "Text": "Проводится"},
		{"Value": "Success", "Text": "Состоялся"},
		{"Value": "Failed", "Text": "Не состоялся"}
	]`

	tempfile := filepath.Join(os.TempDir(), "test_data.json")
	if err := os.WriteFile(tempfile, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("Failed to write temp JSON file : %v", err)
	}

	f := flattener.NewFlattener()
	result, err := f.Flatten(tempfile, "Value", "Text")
	if err != nil {
		t.Fatal(err)
	}

	expected := map[string]string{
		"Running": "Проводится",
		"Success": "Состоялся",
		"Failed":  "Не состоялся",
	}

	for k, v := range expected {
		if (*result)[k] != v {
			t.Errorf("Expected %s -> %s, got %s", k, v, (*result)[k])
		}
	}
}
