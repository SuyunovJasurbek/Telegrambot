package postgres

import (
	"fmt"
	"telegram_bot/models"
	"telegram_bot/service"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	Db *sqlx.DB
}

var (
	userPaymentTable = "users_payment"
	userOpenBudjet   = "users_openbudjet"
	dataTable        = "openbudjet"
	userCheckTable   = "user_check"
	usersTable       = "users_data"
	tokenTable       = "token_check"
)
func NewData(db *sqlx.DB) *Data {
	return &Data{
		Db: db,
	}
}


// CreateUser implements storage.DataI
func (d *Data) CreateUser(data models.UsersOpenbudjet) {
	query := `INSERT INTO ` + userOpenBudjet + ` (chat_id,phone,full_name,time) VALUES ($1, $2, $3, $4)`
	_, err := d.Db.Exec(query, data.ChatID, data.Phone, data.FullName, data.Time)
	if err != nil {
		fmt.Println(err)
	}
}

// CreateUserPayme implements storage.DataI
func (d *Data) CreateUserPayme(data models.UsersPayment) {
	query := `INSERT INTO ` + userPaymentTable + ` (chat_id,phone,full_name,price,time) VALUES ($1, $2, $3, $4, $5)`
	_, err := d.Db.Exec(query, data.ChatID, data.Phone, data.FullName, data.Price, data.Time)
	if err != nil {
		fmt.Println(err)
	}
}

// GetUsers implements storage.DataI
func (s *Data) GetUsers() models.UsersOpenbudjetResponse {
	var users []models.UsersOpenbudjet
	query := `SELECT * FROM ` + usersTable
	err := s.Db.Select(&users, query)
	if err != nil {
		fmt.Println(err)
	}
	return models.UsersOpenbudjetResponse{Users: users}
}

// GetUsersPayme implements storage.DataI
func (s *Data) GetUsersPayme() models.UsersPaymentResponse {
	var users []models.UsersPayment
	query := `SELECT * FROM ` + userCheckTable
	err := s.Db.Select(&users, query)
	if err != nil {
		fmt.Println(err)
	}
	return models.UsersPaymentResponse{Users: users}
}

// -----------------------------------------------TOKEN CHECK-----------------------------------------------
// GetToken implements storage.DataI
func (s *Data) GetToken() string {
	var count int
	query1 := `SELECT count(*) FROM ` + tokenTable
	err := s.Db.QueryRow(query1).Scan(&count)
	if err != nil {
		fmt.Println("Topilamdi")
	}
	fmt.Println("shomrod", count)
	if count == 0 {
		d := service.LoginPayme()
		query := `INSERT INTO ` + tokenTable + ` (key,token) VALUES (1, $1)`
		_, err := s.Db.Exec(query, d)
		if err != nil {
			fmt.Println("Token create error ", err.Error())
		}
		return d
	} else if count == 1 {
		var token string
		query1 := `SELECT token FROM ` + tokenTable + ` WHERE key=1`
		err := s.Db.QueryRow(query1).Scan(&token)
		if err != nil {
			fmt.Println("Topilamdi")
		}
		return token
	}
	return ""
}

// UpdateToken implements storage.DataI
func (s *Data) UpdateToken(token string) error {
	query := `UPDATE ` + tokenTable + ` SET token=$1 WHERE key=1`
	_, err := s.Db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

// -----------------------------------------------USER CHECK-----------------------------------------------
// DeleteUser implements storage.DataI
func (d *Data) DeleteUser(chat_id int) error {
	// delete chat_id from database
	query := `DELETE FROM ` + dataTable + ` WHERE chat_id=$1`
	_, err := d.Db.Exec(query, chat_id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatus implements storage.DataI
func (d *Data) UpdateStatus(status int, chat_id int) error {
	query := `UPDATE ` + dataTable + ` SET status=$1 WHERE chat_id=$2`

	_, err := d.Db.Exec(query, status, chat_id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateType implements storage.DataI
func (d *Data) UpdateType(typ int, chat_id int) error {
	query := `UPDATE ` + dataTable + ` SET type=$1 WHERE chat_id=$2`

	_, err := d.Db.Exec(query, typ, chat_id)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePhone implements storage.DataI
func (s *Data) UpdatePhone(phone, chat_id int) (int, error) {
	query := `UPDATE ` + dataTable + ` SET phone=$1 WHERE chat_id=$2`
	_, err := s.Db.Exec(query, phone, chat_id)
	if err != nil {
		return 0, err
	}
	return phone, nil
}

// CreateData implements storage.DataI
func (s *Data) CreateData(data models.DataCreate) error {
	var phone = 0
	var chat_id int
	query1 := `SELECT chat_id FROM ` + dataTable + ` WHERE chat_id=$1`
	err := s.Db.QueryRow(query1, data.ChatID).Scan(&chat_id)
	if err != nil {
		fmt.Println("Topilamdi")
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	fmt.Println(chat_id)
	fmt.Println("++++++++++++++++++++++++++++++++++")
	if chat_id == 0 {
		query := `INSERT INTO ` + dataTable + ` (chat_id, phone , full_name,last_name , status, type, time ) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := s.Db.Exec(query, data.ChatID, data.Phone, data.FullName, data.LastName, data.Status, data.Type, data.Time)
		if err != nil {
			return err
		}
		return nil
	} else if chat_id == data.ChatID {
		query := `UPDATE ` + dataTable + ` SET phone=$1,  status=0  WHERE chat_id=$2`
		_, err := s.Db.Exec(query, phone, chat_id)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// GetData implements storage.DataI
func (s *Data) GetStatus(chat_id int) (int, error) {
	var status int
	query := `SELECT status FROM ` + dataTable + ` WHERE chat_id=$1`
	err := s.Db.Get(&status, query, chat_id)
	if err != nil {
		return 0, err
	}
	return status, nil
}

// GetTypes implements storage.DataI
func (s *Data) GetTypes(chat_id int) (int, error) {
	var types int
	query := `SELECT status FROM ` + dataTable + ` WHERE chat_id=$1`
	err := s.Db.Get(&types, query, chat_id)
	if err != nil {
		return 0, err
	}
	return types, nil
}

// CaptchaCheckCreate implements storage.DataI
func (s *Data) CaptchaCheckCreate(chat_id int, phone int, key string) error {
	query1 := `SELECT chat_id FROM ` + userCheckTable + ` WHERE chat_id=$1`
	var chat_id1 int
	err := s.Db.QueryRow(query1, chat_id).Scan(&chat_id1)
	if err != nil {
		fmt.Println("Topilamdi")
	}
	if chat_id1 == 0 {
		query := `INSERT INTO ` + userCheckTable + ` (chat_id,phone,key) VALUES ($1, $2, $3)`
		_, err := s.Db.Exec(query, chat_id, phone, key)
		if err != nil {
			return err
		}
		return nil
	} else if chat_id1 == chat_id {
		query := `UPDATE ` + userCheckTable + ` SET  key=$1 WHERE chat_id=$2`
		_, err := s.Db.Exec(query, key, chat_id)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// CaptchaCheckDelete implements storage.DataI
func (d *Data) CaptchaCheckDelete(chat_id int) error {
	query := `DELETE FROM ` + userCheckTable + ` WHERE chat_id=$1`
	_, err := d.Db.Exec(query, chat_id)
	if err != nil {
		return err
	}
	return nil
}

// CaptchaCheckGet implements storage.DataI
func (s *Data) CaptchaCheckGet(chat_id int) (models.CheckResponce, error) {
	var res models.CheckResponce
	query := `SELECT phone, key FROM ` + userCheckTable + ` WHERE chat_id=$1`
	d := s.Db.QueryRow(query, chat_id)
	err := d.Scan(&res.Phone, &res.OptKey)
	if err != nil {
		fmt.Println(err.Error())
		return models.CheckResponce{}, nil
	}
	return res, nil
}

// UpdateCheck implements storage.DataI
func (s *Data) UpdateCheck(chat_id int, key string) error {
	query := `UPDATE ` + userCheckTable + ` SET optkey=$1 WHERE chat_id=$2`
	_, err := s.Db.Exec(query, key, chat_id)
	if err != nil {
		fmt.Println("Errorr db :---")
		return err
	}
	return nil
}

// UpdateCheckPhoneSucces implements storage.DataI
func (s *Data) UpdateCheckPhoneSucces(chat_id int, OptKey string) error {
	query := `UPDATE ` + userCheckTable + ` SET optkey=$1 WHERE chat_id=$2`
	_, err := s.Db.Exec(query, OptKey, chat_id)
	if err != nil {
		fmt.Println("Errorr db :---")
		return err
	}
	return nil
}
// GetPhone implements storage.DataI
func (s *Data) GetPhone(chat_id int) (int, error) {
	var tel int
	query := `SELECT phone  FROM ` + userCheckTable + ` WHERE chat_id=$1`
	err := s.Db.Get(&tel, query, chat_id)
	if err != nil {
		fmt.Println("Errorr db :---_____________________________________________________________________________")
		fmt.Println(err.Error())
		fmt.Println("Errorr db :---_____________________________________________________________________________")
		return 0, nil
	}
	return tel, nil
}

// FinalSMSCodeGet implements storage.DataI
func (s *Data) FinalSMSCodeGet(chat_id int) (models.FinalSMSCodeRequest, error) {
	var res models.FinalSMSCodeRequest
	query := `SELECT key, optkey FROM ` + userCheckTable + ` WHERE chat_id=$1`
	d := s.Db.QueryRow(query, chat_id)
	err := d.Scan(&res.Key, &res.OptKey)
	if err != nil {
		fmt.Println(err.Error())
		return models.FinalSMSCodeRequest{}, nil
	}
	return res, nil
}

// ------------------------------------------------AMOUNT TABLE-------------------------------------------------------------------------------------
// UserAmountCreate implements storage.DataI
func (s *Data) UserAmountCreate(dat models.AmountCreateRequest) error {
	var count int
	quer1 := `SELECT count(*) FROM ` + usersTable + ` WHERE chat_id=$1`
	err := s.Db.QueryRow(quer1, dat.ChatID).Scan(&count)
	if err != nil {
		fmt.Println("Boshlanggich kirimdagi error:", dat.ChatID)
	}
	if count == 0 {
		fmt.Println("Boshlanggich chat id topilmadi::", dat.ChatID)
		quer2 := `INSERT INTO ` + usersTable + ` (chat_id,time, amount) VALUES ($1, $2, $3)`
		_, err := s.Db.Exec(quer2, dat.ChatID, dat.Time, 5000)
		if err != nil {
			fmt.Println("Boshlanggich kirimdagi ERROR INSSERDA :", err.Error())
		}
		return nil
	}
	quer3 := `UPDATE ` + usersTable + ` SET amount=amount+$1 WHERE chat_id=$2`
	_, err = s.Db.Exec(quer3, 5000, dat.ChatID)
	if err != nil {
		fmt.Println("Amount Update errorrrr: ", err.Error())
	}
	return nil
}

// UserAmountGet implements storage.DataI
func (s *Data) UserAmountGet(chat_id int) (int, error) {
	var amout int
	query := `SELECT amount FROM ` + usersTable + ` WHERE chat_id=$1`
	err := s.Db.Get(&amout, query, chat_id)
	if err != nil {
		return 0, nil
	}
	return amout, nil
}

// UserAmountSendPhone implements storage.DataI
func (s *Data) UserAmountSendPhone(chat_id int) error {
	quer3 := `UPDATE ` + usersTable + ` SET amount=$1 WHERE chat_id=$2`
	_, err := s.Db.Exec(quer3, 0, chat_id)
	if err != nil {
		fmt.Println("Amount Send Get  errorrrr: ", err.Error())
		return nil
	}
	return nil
}