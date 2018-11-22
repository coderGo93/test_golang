package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	dbComponent "server/components/database"
	jsonComponent "server/components/json"
	jwtComponent "server/components/jwt"
	formularioComponent "server/components/model/formulario"
	pdfComponent "server/components/pdf"
	responseComponent "server/components/response"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	//ServerPort Port listening
	ServerPort = "9000"
)

func main() { // handle the request
	dbComponent.OpenDB() //open database
	//Create mux server

	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/formulario/obtener_formularios/{token}/{page}/{size_page}", getFormulariosPaging).Methods("GET")             // localhost:9000/formulario/obtener_formularios/0/10
	r.HandleFunc("/formulario/actualizar/{token}", updateFormulario).Methods("PUT").Headers("Content-Type", "application/json") // localhost:9000/formulario/actualizar
	r.HandleFunc("/formulario/subir/{token}", uploadFormulario).Methods("POST").Headers("Content-Type", "application/json")     // localhost:9000/formulario/subir
	r.HandleFunc("/test/pdf", testShowPDF).Methods("GET")                                                                       //localhost:9000/test/pdf

	// Initialize Go server
	fmt.Println("\n initializing golang test ...\nPORT = " + ServerPort)
	fmt.Println("Tiempo: " + time.Now().Format("02/01/2006 15:04:05"))
	log.Fatal(http.ListenAndServe(":"+ServerPort, r))
}

//GET methods
func index(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	fmt.Fprint(response, `{"success":1}`)
}

func getFormulariosPaging(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)      //creo para obtener los parametros de browser
	page := params["page"]           //pagina
	sizePage := params["size_page"]  //tamaño pagina
	pageInt, _ := strconv.Atoi(page) //convierto string a int
	sizePageInt, _ := strconv.Atoi(sizePage)
	json, err := formularioComponent.GetFormulariosPaging(pageInt, sizePageInt)
	if err != nil {
		fmt.Fprint(response, `{"error":"no se pudo obtener datos"}`)
	}
	fmt.Fprint(response, json)

}

//testShowPDF para mostrar pdf
func testShowPDF(response http.ResponseWriter, request *http.Request) {
	//Create pdf
	pdfComponent.GeneratePDFHelloWorld()
	// Open file
	dir, err := os.Getwd() //Get current directory
	f, err := os.Open(fmt.Sprintf("%s\\%s", dir, "hello.pdf"))
	if err != nil {
		fmt.Println(err)
		response.WriteHeader(500)
		return
	}
	defer f.Close()

	//Set header
	response.Header().Set("Content-type", "application/pdf")

	//Stream to response
	if _, err := io.Copy(response, f); err != nil {
		fmt.Println(err)
		response.WriteHeader(500)
	}
}

//PUT methods
func updateFormulario(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	token := params["token"]
	var formulario formularioComponent.Formulario
	body, _ := ioutil.ReadAll(request.Body)
	err := jsonComponent.DecodeJSON(string(body), &formulario)
	if err != nil {
		error := responseComponent.GetError(6)
		res := jsonComponent.ToJSON(&error)
		fmt.Fprint(response, res)
	}
	if jwtComponent.IsJWTValid(token, string("key")) {
		status, err := formularioComponent.UpdateFormulario(formulario)
		if status {
			success := responseComponent.Success{Success: 1}
			res := jsonComponent.ToJSON(&success)
			fmt.Fprint(response, res)
		}
		if err != nil {
			error := responseComponent.GetError(8)
			res := jsonComponent.ToJSON(&error)
			fmt.Fprint(response, res)
		}
	} else {
		error := responseComponent.GetError(5)
		res := jsonComponent.ToJSON(&error)
		fmt.Fprint(response, res)
	}

}

//POST methods
func uploadFormulario(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	token := params["token"]
	var formulario formularioComponent.Formulario
	body, _ := ioutil.ReadAll(request.Body)
	err := jsonComponent.DecodeJSON(string(body), &formulario) //convierto el request a objeto formulario
	if err != nil {
		error := responseComponent.GetError(6)
		res := jsonComponent.ToJSON(&error)
		fmt.Fprint(response, res)
	}
	if jwtComponent.IsJWTValid(token, string("key")) {
		idFormulario, err := formularioComponent.InsertFormulario(formulario)
		if idFormulario > 0 {
			success := responseComponent.Success{Success: 1}
			res := jsonComponent.ToJSON(&success)
			fmt.Fprint(response, res)
			return
		}
		if err != nil {
			error := responseComponent.GetError(8)
			res := jsonComponent.ToJSON(&error)
			fmt.Fprint(response, res)
		}
	} else {
		error := responseComponent.GetError(5)
		res := jsonComponent.ToJSON(&error)
		fmt.Fprint(response, res)
	}
}
