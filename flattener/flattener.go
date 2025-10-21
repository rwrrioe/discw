package flattener

import (
	"log"

	"github.com/rwrrioe/Discworld/parser"
)

type flatJSON map[string]string

type Flattener struct {
	flatJSON flatJSON
}

func NewFlattener() *Flattener {
	return &Flattener{
		flatJSON: make(map[string]string),
	}
}

func (n *Flattener) Flatten(filename string, key string, value string) (*flatJSON, error) {
	parsed, err := parser.NewParser().Parse(filename)
	if err != nil {
		return nil, err
	}

	log.Println("Parsed JSON:", parsed)
	flat := make(flatJSON)
	generic := make([]interface{}, len(*parsed))
	for i, m := range *parsed {
		generic[i] = m
	}

	FlattenJSON(generic, key, value, &flat)
	log.Println("Flatted JSON result:", flat)

	return &flat, nil
}

func FlattenJSON(data interface{}, key string, value string, flat *flatJSON) {
	var extract func(interface{})
	extract = func(item interface{}) {
		switch v := item.(type) {
		case map[string]interface{}:
			val, valOk := v[key]
			txt, txtOk := v[value]

			if valOk && txtOk {
				valStr, valOk := val.(string)
				txtStr, txtOk := txt.(string)

				if valOk && txtOk {
					(*flat)[valStr] = txtStr
				} else {
				}
			}
			for _, child := range v {
				extract(child)
			}

		case []interface{}:
			for _, child := range v {
				extract(child)
			}
		default:
		}
	}

	switch top := data.(type) {
	case []interface{}:
		for _, item := range top {
			extract(item)
			log.Printf("Extracted from array: %+v\n", item)
		}
	case []map[string]interface{}:
		for _, item := range top {
			extract(item)
			log.Printf("Extracted from []map: %+v\n", item)
		}
	default:
		extract(data)
		log.Printf("Extracted from default: %+v\n", data)
	}
	log.Println("extracted to flat", flat)
	log.Println("end flattenJSON")
}
