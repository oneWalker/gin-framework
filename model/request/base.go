package request

type Login struct {
	Username string   `json:"username"`
	Phone    string   `json:"phone"`
	Password string   `json:"password"`
	Hobbys   []string `json:"hobbys"`
}
