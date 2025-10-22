package flattener_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/rwrrioe/discw/flattener"
)

func TestFlattenerArray(t *testing.T) {
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

func TestFlattenerDefault(t *testing.T) {
	jsonContent := `
		{"Value": "Running", "Text": "Проводится"}`

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
	}

	for k, v := range expected {
		if (*result)[k] != v {
			t.Errorf("Expected %s -> %s, got %s", k, v, (*result)[k])
		}
	}
}

func TestFlannerToFile(t *testing.T) {
	var flattened map[string]string
	jsonContent := `[
		{"Value": "Running", "Text": "Проводится"},
		{"Value": "Success", "Text": "Состоялся"},
		{"Value": "Failed", "Text": "Не состоялся"}
	]`

	expected := map[string]string{
		"Running": "Проводится",
		"Success": "Состоялся",
		"Failed":  "Не состоялся",
	}

	parsefile := filepath.Join(os.TempDir(), "test_data.json")
	tempfile := filepath.Join(os.TempDir(), "test_result.json")

	if err := os.WriteFile(parsefile, []byte(jsonContent), 0644); err != nil {
		t.Fatalf("Failed to write temp JSON file : %v", err)
	}

	if err := flattener.NewFlattener().FlattenToFile(tempfile, parsefile, "Value", "Text"); err != nil {
		t.Fatal(err)
	}

	file, err := os.Open(tempfile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&flattened); err != nil {
		t.Error(err)
	}

	for k, v := range expected {
		if v != flattened[k] {
			t.Errorf("Expected %s -> %s, got %s", k, v, flattened[k])
		}
	}
}
