package typicode

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Company  Company `json:"company"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}
