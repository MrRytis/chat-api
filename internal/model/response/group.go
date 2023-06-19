package response

type Group struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Admin User   `json:"admin"`
	Users []User `json:"users"`
}

type User struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
