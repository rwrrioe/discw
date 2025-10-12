package flattener

import (
	"github.com/rwrrioe/Discworld/parser"
)

type NestedJSON map[string]interface{}

type flatJSON map[string]interface{}

type Flattener struct {
	flatJSON flatJSON
}

func NewFlattener() *Flattener {
	return &Flattener{
		flatJSON: make(map[string]interface{}),
	}
}

func (n *Flattener) Flatten(filename string, key string, value string) (*flatJSON, error) {
	parser := parser.NewParser()
	jsonMap, err := parser.Parse(filename)
	if err != nil {
		return nil, err
	}
	wrapped := map[string]interface{}{
		"root": *jsonMap,
	}

	flatJSON := FlattenJSON(wrapped, key, value)

	return flatJSON, nil
}

func FlattenJSON(json map[string]interface{}, key string, value string) *flatJSON {
	flatJS := make(flatJSON)
	var resultJSON map[string]interface{}

	var extract func(interface{})
	extract = func(data interface{}) {
		switch d := data.(type) {
		case map[string]interface{}:
			val, valOk := d[key]
			txt, txtOk := d[value]

			if valOk && txtOk {
				if keyStr, ok := val.(string); ok {
					if txtStr, ok := txt.(string); ok {
						resultJSON[keyStr] = txtStr
					}
				}
			}

			for _, v := range d {
				extract(v)
			}
		case []interface{}:
			for _, v := range d {
				extract(v)
			}
		}
	}

	flatJS = flatJSON(json)
	return &flatJS
}
