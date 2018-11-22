package pdf

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

//Constantes
const (
	ComponentName = "pdf"
	componentDir  = "server/component/" + ComponentName
	countCol      = 2
)

//La funcion init() siempre corre primero en el archivo, es como el constructor
// Solo corre en programa genera, no cada vez en esta clase

func init() {
	fmt.Println("Success -> " + componentDir)
}

//GeneratePDFHelloWorld Generate pdf
func GeneratePDFHelloWorld() (bool, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}
