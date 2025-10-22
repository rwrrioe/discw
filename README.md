[![Go Report Card](https://goreportcard.com/badge/github.com/rwrrioe/Discworld)](https://goreportcard.com/report/github.com/rwrrioe/Discworld)
 [![Go Reference](https://pkg.go.dev/badge/github.com/rwrrioe/diskw.svg)](https://pkg.go.dev/github.com/rwrrioe/diskw)
# Discworld (diskw)

**Discworld (diskw)** is a CLI tool and a lightweight Go library that converts nested JSON **dictionaries** (objects containing key→value pairs) into a single-level **flat key‑value Go map** (`map[string]interface{}`).

The core idea: if your JSON is built from dictionaries (objects) where every node is a set of key→value pairs, Discworld will produce a flat map where each key is a path describing the nesting and the value is the original leaf value. This makes backend mapping and data normalization straightforward — you supply keys and values, give the tool the object, and it returns a flat map automatically.

---

## What it does
- Takes nested JSON dictionaries (objects made of key→value pairs) and flattens them into `map[string]interface{}`.
- Each nested path is converted into a single composite key (configurable separator, e.g. `.`).  
- Leaves (actual values) become values in the resulting flat map.
- Designed for backends that need to map nested structures into flat key-value storage, mapping rules, or simplified processing logic.

---

## Key behaviors
- **Input shape:** nested JSON objects (dictionaries). Each object contains key→value pairs; values can be primitives or further dictionaries.  
- **Output shape:** a flat `map[string]interface{}` where every key is a path to a leaf value.  
- **Configuration:** you can set the separator (default `.`), array handling mode (index/join/ignore), prefix for keys, and empty-value policy.
- **Simplicity:** if you have `key => value` pairs at leaves, Discworld will preserve the values and produce predictable flat keys for mapping on the backend.

---

## Installation

### CLI
```bash
go install github.com/you/diskw/cmd/diskw@latest
```

### Library
```bash
go get github.com/you/diskw@latest
```

---

## CLI usage
```
diskw -i input.json -o output.json
```

Common options:
```
  -i, --input string        Input JSON file (default: stdin)
  -o, --output string       Output JSON file (default: stdout)
  -s, --sep string          Key separator (default: ".")
  -a, --arrays string       Array handling mode: index|join|ignore (default: "index")
  -p, --prefix string       Optional prefix for all keys
  -e, --empty string        Empty value handling: omit|null|empty (default: "omit")
```
---

## Why use Discworld
- Removes boilerplate when converting nested dictionaries to flat key-value maps.  
- Makes backend mapping, ETL, and configuration handling simpler and more predictable.  
- Small, dependency-free, easy to embed in Go services.

---

## License
MIT 
