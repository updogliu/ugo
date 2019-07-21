package ujson

import "encoding/json"

func GetIndentedJsonStr(data interface{}) string {
	marshalled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(marshalled)
}

var Beautify = GetIndentedJsonStr

func GetCompactJsonStr(data interface{}) string {
	marshalled, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return string(marshalled)
}
