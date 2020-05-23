package domain

// User dto with json for display in web response
type User struct {
	ID        int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
