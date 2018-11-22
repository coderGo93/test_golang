package formulario

import (
	"database/sql"
	"fmt"
	jsonComponent "server/components/json"

	dbComponent "server/components/database"
	"strings"
)

//Constantes
const (
	ComponentName   = "formulario"
	ComponentParent = "model"
	componentDir    = "server/component/" + ComponentParent + "/components/" + ComponentName
)

var (
	db *sql.DB
)

//La funcion init() siempre corre primero en el archivo, es como el constructor
// Solo corre en programa genera, no cada vez en esta clase

func init() {
	fmt.Println("Success -> " + componentDir)
}

//Formulario estructura
type Formulario struct {
	ID                int64  `json:"id"` //es para mostrar como se veria y como debe encontrar el estilo de llave
	Nombre            string `json:"nombre"`
	PrimerApellido    string `json:"primer_apellido"`
	SegundoApellido   string `json:"segundo_apellido"`
	FechaNacimiento   string `json:"fecha_nacimiento"`
	LugarNacimiento   string `json:"lugar_nacimiento"`
	Celular           string `json:"celular"`
	CorreoElectronico string `json:"correo_electronico"`
	XML               string `json:"xml"`
}

//GetFormulariosPaging Obtengo formularios con paginado
func GetFormulariosPaging(page int, sizePage int) ([]string, error) {
	db = dbComponent.GetDatabase()
	var jsonArray []string
	stmt, err := db.Prepare(`SELECT * FROM formulario LIMIT $1, $2`)

	if err != nil {
		fmt.Println("No hay información")
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(page, sizePage)

	if err != nil {
		return nil, err
	}
	formulario := Formulario{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&formulario.ID, &formulario.Nombre, &formulario.PrimerApellido, &formulario.SegundoApellido,
			&formulario.FechaNacimiento, &formulario.LugarNacimiento, &formulario.Celular, &formulario.CorreoElectronico,
			&formulario.XML)
		json := jsonComponent.ToJSON(formulario) + ","
		jsonArray = append(jsonArray, json)
		if err != nil {
			return nil, err
		}
	}
	if len(jsonArray) > 0 {
		ultimo := jsonArray[len(jsonArray)-1]
		jsonArray = jsonArray[:len(jsonArray)-1]
		ultimo = strings.TrimRight(ultimo, ",")
		jsonArray = append(jsonArray, ultimo)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return jsonArray, nil
}

//InsertFormulario inserto el formulario con los datos correspondientes
func InsertFormulario(formulario Formulario) (int64, error) {
	db = dbComponent.GetDatabase()
	id := int64(0)
	stmt, err := db.Prepare(`insert into formulario (nombre, primer_apellido, segundo_apellido,
							 fecha_nacimiento, lugar_nacimiento, celular, correo_electronico, xml)
		values ($1, $2, $3, $4, $5, $6, $7, $8)`)

	if err != nil {
		return id, err
	}
	result, err := stmt.Exec(formulario.Nombre, formulario.PrimerApellido, formulario.SegundoApellido,
		formulario.FechaNacimiento, formulario.LugarNacimiento, formulario.Celular,
		formulario.CorreoElectronico, formulario.XML)

	if err != nil {
		return id, err
	}

	id, err = result.LastInsertId() //agarro el id que se insertó
	if err != nil {
		return id, err
	}

	return id, nil

}

//UpdateFormulario Modifico los datos correspondientes
func UpdateFormulario(formulario Formulario) (bool, error) {
	db = dbComponent.GetDatabase()
	stmt, err := db.Prepare(`UPDATE formulario SET correo_electronico = $1, celular = $2
			WHERE id = $3`)

	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(formulario.CorreoElectronico, formulario.Celular, formulario.ID)

	if err != nil {
		return false, err
	}

	return true, nil

}
