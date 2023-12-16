package resources

type MidtransInvoice struct {
	StatusMessage     string              `json:"status_message,omitempty"`
	TransactionId     string              `json:"transaction_id,omitempty"`
	OrderId           string              `json:"order_id,omitempty"`
	GrossAmount       string              `json:"gross_amount,omitempty"`
	PaymentType       string              `json:"payment_type,omitempty"`
	TransactionTime   string              `json:"transaction_time,omitempty"`
	TransactionStatus string              `json:"transaction_status,omitempty"`
	FraudStatus       string              `json:"fraud_status,omitempty"`
	Currency          string              `json:"currency,omitempty"`
	VaNumbers         []MidtransVaNumbers `json:"va_numbers,omitempty"`
	ExpiryTime        string              `json:"expiry_time,omitempty"`
	Actions           []MidtransActions   `json:"actions,omitempty"`
}

type MidtransVaNumbers struct {
	Bank     string `json:"bank,omitempty"`
	VaNumber string `json:"va_number,omitempty"`
}

type MidtransActions struct {
	Name   string `json:"name,omitempty"`
	Method string `json:"method,omitempty"`
	Url    string `json:"url,omitempty"`
}
