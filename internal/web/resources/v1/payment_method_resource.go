package resources

type BankMethodResource struct {
	Title    string `json:"title,omitempty"`
	BankCode string `json:"bank_code,omitempty"`
	Image    string `json:"image,omitempty"`
}
