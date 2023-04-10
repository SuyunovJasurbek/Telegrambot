package storage

import "telegram_bot/models"

type StorageI interface {
	Data() DataI
	CloseDb()
}

type DataI interface {
	CreateData(models.DataCreate) error
	GetStatus(chat_id int) (int, error)
	GetTypes(chat_id int) (int, error)
	GetPhone(chat_id int )(int,  error)
	UpdateStatus(status, chat_id int) error
	UpdateType(typ, chat_id int) error
	DeleteUser(chat_id int) error
	UpdatePhone(phone, chat_id int) (int, error)
	//-------------------USER CHECK-------------------
	CaptchaCheckCreate(chat_id, phone int, key string) error
	CaptchaCheckGet(chat_id int) (models.CheckResponce , error)
	FinalSMSCodeGet(chat_id int) (models.FinalSMSCodeRequest , error)
	CaptchaCheckDelete(chat_id int) ( error)
	UpdateCheck(chat_id int,key string ) (error)
	UpdateCheckPhoneSucces(chat_id int,OptKey string ) (error)
	//-----------------USERS AMOUNT --------------------
	UserAmountCreate(models.AmountCreateRequest) error
	UserAmountGet(chat_id int)(int, error ) 
	UserAmountSendPhone(chat_id int ) error
	//-----------------TOKEN-------------------
	UpdateToken(token string) error
	GetToken() string
	//-----------------USERS OPENBUDJET --------------------
	CreateUser(data models.UsersOpenbudjet) 
	GetUsers() (models.UsersOpenbudjetResponse)
	//-----------------USERS PAYME -------------------
	CreateUserPayme(models.UsersPayment) 
	GetUsersPayme() (models.UsersPaymentResponse)
}