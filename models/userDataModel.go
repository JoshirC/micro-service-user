package models

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingUpData struct {
	Name     string `json:"name"`
	Rut      string `json:"rut"`
	Password string `json:"password"`
	Email    string `json:"email"`
	City     string `json:"city"`
}
