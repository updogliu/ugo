package ujson

import "encoding/json"

func Beautify(v interface{}) string {
	return GetIndentedJsonStr(v)
}

func Compact(v interface{}) string {
	return GetCompactJsonStr(v)
}

func GetIndentedJsonStr(v interface{}) string {
	marshalled, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(marshalled)
}

func GetCompactJsonStr(v interface{}) string {
	marshalled, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(marshalled)
}

func UnmarshalStr(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), v)
}
