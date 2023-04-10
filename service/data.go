package service

import (
	"fmt"
	"strconv"
	"strings"
	"telegram_bot/helper"
	"telegram_bot/models"
	"time"
)

func (s *Service) CreateData(data models.DataCreate) error {
	var dat = models.DataCreate{
		ChatID:   data.ChatID,
		Phone:    data.Phone,
		FullName: data.FullName,
		LastName: data.LastName,
		Status:   0,
		Type:     0,
		Time:     int(time.Now().Unix()),
	}
	err := s.repository.Data().CreateData(dat)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *Service) GetStatus(chat_id int) int {
	status, err := s.repository.Data().GetStatus(chat_id)
	if err != nil {
		fmt.Println(err)
	}
	return status
}
func (s *Service) GetTypes(chat_id int) int {
	types, err := s.repository.Data().GetTypes(chat_id)
	if err != nil {
		fmt.Println(err)
	}
	return types
}
func (s *Service) UpdateStatus(status, chat_id int) error {
	err := s.repository.Data().UpdateStatus(status, chat_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *Service) UpdatePhone(phone, chat_id int) (int, error) {
	phon, err := s.repository.Data().UpdatePhone(phone, chat_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return phon, nil
}
func (s *Service) UpdateType(typ, chat_id int) error {
	err := s.repository.Data().UpdateType(typ, chat_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (s *Service) OpenbudgetCheck(phone, chat_id int) (string, error) {
	resp, err := RandomNumberGet()

	if err != nil {
		fmt.Println(err)
	}
	err = s.repository.Data().CaptchaCheckCreate(chat_id, phone, resp.CaptchaKey)
	if err != nil {
		fmt.Println(err)
	}
	base64 := helper.Base64Parse(resp.Image)
	helper.ImageParse(base64, strconv.Itoa(chat_id))
	return "", nil
}
func (s *Service) OpenbudgetPhoneCheck(chat_id, answer int) (int, error) {
	res, err := s.repository.Data().CaptchaCheckGet(chat_id)
	if err != nil {
		fmt.Println(err)
	}
	phone := strconv.Itoa(res.Phone)
	if phone[0] == '9' {
		phone = strings.Replace(phone, "9", "9989", 1)
	} else if phone[0] == '8' {
		phone = strings.Replace(phone, "8", "9988", 1)
	}
	phone_int, _ := strconv.Atoi(phone)
	var req = models.PhoneCheckRequest{
		CaptchaKey:    res.OptKey,
		Phone:         phone_int,
		CaptchaResult: answer,
	}
	fmt.Println(req)

	ress, ress2, err := PhoneCheck(req)
	if ress == 0 {
		err = s.repository.Data().UpdateCheck(chat_id, ress2.OtpKey)
		if err != nil {
			fmt.Println(err)
		}
		return 0, nil
	}

	if err != nil {
		return 2, nil
	}
	return ress, nil
}
func (s *Service) DeleteChatId(chat_id int) error {
	err := s.repository.Data().CaptchaCheckDelete(chat_id)
	if err != nil {
		fmt.Println("+-=-=-=-=-=")
	}
	return nil
}
func (s *Service) CheckResult(chat_id int) (string, error) {
	phon, err := s.repository.Data().GetPhone(chat_id)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := RandomNumberGet()
	if err != nil {
		fmt.Println(err)
	}
	err = s.repository.Data().CaptchaCheckCreate(chat_id, phon, resp.CaptchaKey)
	base64 := helper.Base64Parse(resp.Image)
	helper.ImageParse(base64, strconv.Itoa(chat_id))
	if err != nil {
		fmt.Println(err)
	}
	return "Suratdagi matematik amalni bajarib javobni kiriting :", nil
}
func (s *Service) FinalSMSCode(chat_id, answer int, fullname string) (int, error, int) {
	res, err := s.repository.Data().FinalSMSCodeGet(chat_id)
	if err != nil {
		fmt.Println(err)
	}

	var req = models.FinalOpenbudjetRequest{
		OtpKey:  res.OptKey,
		OtpCode: strconv.Itoa(answer),
	}

	ress, err := SmsCheck(req)
	if err != nil {
		fmt.Println(err)
	}
	if ress == 0 {
		// sms kod tugri
		var req = models.AmountCreateRequest{
			ChatID: chat_id,
			Time:   time.Now().String(),
		}
		phone, _ := s.repository.Data().GetPhone(chat_id)
		var data = models.UsersOpenbudjet{
			ChatID:   chat_id,
			Phone:    phone,
			FullName: fullname,
			Time:     time.Now().String(),
		}
		s.repository.Data().CreateUser(data)
		err := s.repository.Data().UserAmountCreate(req)
		if err != nil {
			fmt.Println("AmountCreateRequest error")
			fmt.Println(err)
		}
		s.repository.Data().CaptchaCheckDelete(chat_id)
		return 0, nil, phone
	} else if ress == 1 {
		// sms xato kiritildi
		return 1, nil, 0
	} else if ress == -1 {
		// urunishlar soni qolmadi. qaytadan urinib kuring
		s.repository.Data().CaptchaCheckDelete(chat_id)
		return -1, nil, 0
	} else if ress == -2 {
		// muddat utib ketgan
		return -2, nil , 0
	}
	return 3, nil , 0
}
