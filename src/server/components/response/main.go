package response

//Constantes
const (
	ComponentName = "response"
	componentDir  = "server/component/" + ComponentName
)

//ID estructura de id
type ID struct {
	ID int64 `json:"id"`
}

//Success estructura de success
type Success struct {
	Success int `json:"success"`
}

//Info estructura de Info
type Info struct {
	Info string `json:"info"`
}

//Token estructura del token
type Token struct {
	Token string `json:"token"`
	ID    int64  `json:"id"`
}

//Error estructura
type Error struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

//GetError to get error with number
func GetError(num int) Error {
	var jsonArray []string
	jsonArray = append(jsonArray, "")                                       //0
	jsonArray = append(jsonArray, "Token already exists.")                  //1
	jsonArray = append(jsonArray, "Database failure.")                      //2
	jsonArray = append(jsonArray, "Bad session. User or password is wrong") //3
	jsonArray = append(jsonArray, "No token or user is inexistent")         //4
	jsonArray = append(jsonArray, "Bad token")                              //5
	jsonArray = append(jsonArray, "object/json error conversion")           //6
	jsonArray = append(jsonArray, "update data error")                      //8
	jsonArray = append(jsonArray, "upload data error")                      //9
	jsonArray = append(jsonArray, "pdf error")                              //10
	jsonArray = append(jsonArray, "email error")                            //11

	error := &Error{Error: num, Message: jsonArray[num]}

	return *error
}
