package server

//Policy Policy
type Policy struct {
	User   string `json:"user" form:"user" query:"user"`
	Path   string `json:"path" form:"path" query:"path"`
	Method string `json:"method" form:"method" query:"method"`
}

//UserRole UserRole
type UserRole struct {
	User string `json:"user" form:"user" query:"user"`
	Role string `json:"role" form:"role" query:"role"`
}

//User User
type User struct {
	Username    string `json:"username" form:"username" query:"username"`
	Password    string `json:"password" form:"password" query:"password"`
	NewPassword string `json:"new_password" form:"new_password" query:"new_password"`
}

//SuccessMessage SuccessMessage
type SuccessMessage struct {
	Status  int  `json:"status" form:"status" query:"status"`
	Success bool `json:"success" form:"success" query:"success"`
}

//DataMessage DataMessage
type DataMessage struct {
	Status int         `json:"status" form:"status" query:"status"`
	Data   interface{} `json:"data" form:"data" query:"data"`
}

//Message Message
type Message struct {
	Status  int         `json:"status" form:"status" query:"status"`
	Message interface{} `json:"message" form:"message" query:"message"`
}
