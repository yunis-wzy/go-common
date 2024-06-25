package ajson

import jsoniter "github.com/json-iterator/go"

var iter = jsoniter.Config{
	EscapeHTML:             false,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

func Unmarshal(data []byte, v interface{}) error {
	return iter.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return iter.Marshal(v)
}
