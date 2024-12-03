package userstruct

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Type     int    `json:"type"`
}
