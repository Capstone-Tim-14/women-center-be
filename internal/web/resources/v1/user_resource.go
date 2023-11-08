package resources

type UserResource struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Phone_number int    `json:"phone_number"`
	Address      string `json:"address"`
	Status       string `json:"status"`
}
