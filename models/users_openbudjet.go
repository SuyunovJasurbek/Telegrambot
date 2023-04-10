package models

type UsersPayment struct {
	ChatID   int    `json:"chat_id"`
	Phone    int    `json:"phone"`
	FullName string `json:"full_name"`
	Price    int    `json:"price"`
	Time     string `json:"time"`
}

type UsersOpenbudjet struct {
	ChatID   int    `json:"chat_id"`
	Phone    int    `json:"phone"`
	FullName string `json:"full_name"`
	Time     string `json:"time"`
}

type UsersOpenbudjetResponse struct {
	Users []UsersOpenbudjet `json:"users"`
}

type UsersPaymentResponse struct {
	Users []UsersPayment `json:"users"`
}