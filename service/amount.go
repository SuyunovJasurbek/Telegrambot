package service

import (
	"fmt"
	"strconv"
	"telegram_bot/models"
	"time"
)

const (
	Beeline   = "545c7ecd5ae5eca82d1b462f"
	Ucell     = "545e1b1e5ae5eca82d1b4630"
	Perfectum = "545e1cae5ae5eca82d1b4631"
	UMS       = "549981c05ae5eca82d1b4661"
	UzMobile  = "55478199d2c4830936e6c832"
)

// -----------------------------------------------USER TOKEN-----------------------------------------------
func (s *Service) UpdateToken() string {
	token := LoginPayme()
	s.repository.Data().UpdateToken(token)
	return token
}
func (s *Service) GetToken() string {
	token := s.repository.Data().GetToken()
	return token
}
// -----------------------------------------------USER AMOUNT-----------------------------------------------
func (s *Service) GetAmount(chat_id int) (int, error) {
	amun, err := s.repository.Data().UserAmountGet(chat_id)
	if err != nil {
		fmt.Println("Amount get Error.... ")
		fmt.Println(err)
	}
	return amun, nil
}
func (s *Service) PriceFunding(phone_comp, phone int, chat_id int, fullname string)(string, int ){
	token := s.GetToken()
	var MerchantID string
	amount, err := s.GetAmount(chat_id)
	amount_int := amount * 100
	amount_string := fmt.Sprintf("%d", amount)
	if err != nil {
		fmt.Println("Amount get Error.... ")
		fmt.Println(err)
	}
	if phone_comp == 90 || phone_comp == 91 {
		MerchantID = Beeline
	} else if phone_comp == 93 || phone_comp == 94 {
		MerchantID = Ucell
	} else if phone_comp == 95 || phone_comp == 99 {
		MerchantID = UzMobile
	} else if phone_comp == 97 || phone_comp == 88 {
		MerchantID = UMS
	} else if phone_comp == 98 || phone_comp == 96 {
		MerchantID = Perfectum
	}
	var req = models.CreatePaymeRequest{
		Phone:        strconv.Itoa(phone),
		AmountString: amount_string,
		AmountInt:    amount_int,
		Merchant_ID:  MerchantID,
	}
	id := CreatePayme(req)
	if id == "phone" {
		return "Telefon nomer xato kiritildi .", 0
	}
	fmt.Println("ID: ", id)
	r := PayPayme(id, token)
	if r == 1 {
		var user = models.UsersPayment{
			ChatID:   chat_id,
			Phone:    phone,
			FullName: fullname,
			Price:    amount,
			Time:     time.Now().String(),
		}
		s.repository.Data().UserAmountSendPhone(chat_id)
		s.repository.Data().CreateUserPayme(user)
		re := fmt.Sprintf("https://payme.uz/checkout/%s", id)
		return re,amount
	} else if r == -1 {
		token2 := s.UpdateToken()
		res := PayPayme(id, token2)
		if res == 1 {
			var user = models.UsersPayment{
				ChatID:   chat_id,
				Phone:    phone,
				FullName: fullname,
				Price:    amount,
				Time:     time.Now().String(),
			}
			s.repository.Data().UserAmountSendPhone(chat_id)
			s.repository.Data().CreateUserPayme(user)
			re := fmt.Sprintf("https://payme.uz/checkout/%s", id)
			return re,amount
		}
	} else if r == 0 {
		return "Xozirda bizning balanc da mablag' yetarli emas. tulishi bilan yechib olishingiz mumkin.",0
	}
	return "Tulov mobaynida xatolik yuz berdi.",0
}