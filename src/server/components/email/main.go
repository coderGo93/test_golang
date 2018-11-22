package email

import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

//Constantes
const (
	ComponentName = "email"
	componentDir  = "server/component/" + ComponentName
	host          = "host"
	port          = 587
	user          = "user@user.com"
	password      = "password"
	hostGmail     = "smtp.gmail.com"
)

//La funcion init() siempre corre primero en el archivo, es como el constructor
// Solo corre en programa genera, no cada vez en esta clase

func init() {
	fmt.Println("Success -> " + componentDir)
}

//SendPDFEMail Send message with pdf attached
func SendPDFEMail() (bool, error) {

	m := gomail.NewMessage()
	mantenimiento := fmt.Sprintf("Mantenimiento del sucursal %s", sucursal)
	m.SetAddressHeader("From", userError, fmt.Sprintf(`"%s"`, mantenimiento))
	m.SetHeader("To", "recipient@gmail.com")

	dir, err := os.Getwd() //Get current directory
	file := fmt.Sprintf("%s\\%s", dir, "hello.pdf")
	m.SetHeader("Subject", "PDF hello world")
	m.SetBody("text/html", "")
	m.Attach(file)

	d := gomail.NewDialer(host, port, userError, passwordError)

	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
