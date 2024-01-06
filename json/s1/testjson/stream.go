package testjson

import "Json"

func encode(data *interface{}) {
	d := Json.Marshal(&data)
}
