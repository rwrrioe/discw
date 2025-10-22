package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ParsedJSON []map[string]interface{}

func NewParser() *Parser {
	return &Parser{
		JSON: make([]map[string]interface{}, 0),
	}
}

type Parser struct {
	JSON ParsedJSON
}

func (p *Parser) Parse(filename string) (*ParsedJSON, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var dto interface{}
	var parsed []map[string]interface{}
	log.Println("entering parser.go parse")
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, fmt.Errorf("from parser.go parse, error:%v", err)
	}

	switch v := dto.(type) {
	case []interface{}:
		for _, item := range v {
			if m, ok := item.(map[string]interface{}); ok {
				parsed = append(parsed, m)
			}
		}
	case map[string]interface{}:
		parsed = append(parsed, v)
	default:
		log.Printf("unexpected type: %#v", dto)
	}

	p.JSON = parsed
	return &p.JSON, nil
}
