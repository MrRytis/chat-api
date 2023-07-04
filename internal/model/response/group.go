package response

type Group struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Admin User   `json:"admin"`
	Users []User `json:"users"`
}

type GroupList struct {
	PageNumber int32   `json:"pageNumber"`
	PageSize   int32   `json:"pageSize"`
	ItemsCount int32   `json:"itemsCount"`
	Items      []Group `json:"items"`
}

type GroupUserAdded struct {
	Uuid    string `json:"uuid"`
	User    User   `json:"user"`
	Message string `json:"message"`
}

type SimpleGroup struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
