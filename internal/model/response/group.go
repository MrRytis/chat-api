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

type GroupList struct {
	PageNumber int32   `json:"pageNumber"`
	PageSize   int32   `json:"pageSize"`
	ItemsCount int32   `json:"itemsCount"`
	Items      []Group `json:"items"`
}
