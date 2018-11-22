package json

import (
	"encoding/json"
	"fmt"
	"io"
)

//Constantes
const (
	CompomentName = "json"
	componentDir  = "server/components/" + CompomentName
)

//La funcion init() siempre corre primero en el archivo, es como el constructor
// Solo corre en programa genera, no cada vez en esta clase

func init() {
	fmt.Println("Success -> " + componentDir)
}

//ToJSON Convierto el objeto a json string
func ToJSON(obj interface{}) string {
	b, _ := json.Marshal(obj)

	return string(b)
}

//DecodeJSON se usa para decodear json a objeto que corresponda
func DecodeJSON(jsonString string, object interface{}) error {
	bytes := []byte(jsonString)

	return json.Unmarshal(bytes, object)
}

//DecodeJSONFromRequest se usa para decodear json a objeto que corresponde con la informacion de request.Body
func DecodeJSONFromRequest(jsonReader io.Reader, object interface{}) error {

	return json.NewDecoder(jsonReader).Decode(object)
}
