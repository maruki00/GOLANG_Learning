package testjson

import "encoding/json"

// type Person struct {
// }

// func JsonEncode(obj interface{}) string {
// 	jsonData, err := json.Marshal(obj)
// 	if err != nil {
// 		log.Fatal("Error: ", err.Error())
// 	}
// 	return string(jsonData)
// }

func JsonDecode(data []byte) map[string]interface{} {
	d := make(map[string]interface{})
	json.Unmarshal(data, &d)
	// if err != nil {
	// 	log.Fatal("Error: ", err.Error())
	// }
	return d
}
