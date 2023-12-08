package domain

type BankMethod struct {
	Id       uint
	Title    string
	BankCode string
	Image    string
}

type WalletMethod struct {
	Id         uint
	Title      string
	WalletCode string
	Image      string
}
