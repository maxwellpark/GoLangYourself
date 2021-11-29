package data

// Export a struct for parsing a response
type Data struct {
	User User `json:"data"` // JSON tags match schema 
}

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}
