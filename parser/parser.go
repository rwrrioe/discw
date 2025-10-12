package parser

import (
	"encoding/json"
	"os"
)

type ParsedJSON map[string]interface{}

func NewParser() *Parser {
	return &Parser{
		JSON: make(map[string]interface{}),
	}
}

type Parser struct {
	JSON ParsedJSON
}

func (p *Parser) Parse(filename string) (*ParsedJSON, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(b, &p.JSON)
	return &p.JSON, nil
}
