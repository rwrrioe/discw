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
	var parsed []map[string]interface{}
	log.Println("entering parser.go parse")
	if err := json.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("from parser.go parse, error:%v", err)
	}

	p.JSON = parsed
	return &p.JSON, nil
}
