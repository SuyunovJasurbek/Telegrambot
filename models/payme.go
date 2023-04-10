package models

type PaymeCreateResponce struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Cheque struct {
			ID         string `json:"_id"`
			CreateTime int64  `json:"create_time"`
			PayTime    int    `json:"pay_time"`
			CancelTime int    `json:"cancel_time"`
			State      int    `json:"state"`
			Type       int    `json:"type"`
			External   bool   `json:"external"`
			Operation  int    `json:"operation"`
			Category   struct {
				ID        string `json:"_id"`
				Title     string `json:"title"`
				Color     string `json:"color"`
				Sort      int    `json:"sort"`
				Operation int    `json:"operation"`
				Indoor    bool   `json:"indoor"`
				Mcc       struct {
					Visa []string `json:"visa"`
				} `json:"mcc"`
			} `json:"category"`
			Error       interface{} `json:"error"`
			Description string      `json:"description"`
			Detail      interface{} `json:"detail"`
			Amount      int         `json:"amount"`
			Currency    int         `json:"currency"`
			Commission  int         `json:"commission"`
			Account     []struct {
				Name  string `json:"name"`
				Title string `json:"title"`
				Value string `json:"value"`
				Main  bool   `json:"main,omitempty"`
			} `json:"account"`
			Card     interface{} `json:"card"`
			Payer    interface{} `json:"payer"`
			Merchant struct {
				ID           string `json:"_id"`
				Name         string `json:"name"`
				Organization string `json:"organization"`
				Address      string `json:"address"`
				BusinessID   string `json:"business_id"`
				Epos         struct {
					MerchantID string `json:"merchantId"`
					TerminalID string `json:"terminalId"`
				} `json:"epos"`
				Date   int64       `json:"date"`
				Logo   string      `json:"logo"`
				Type   string      `json:"type"`
				Terms  interface{} `json:"terms"`
				Myhome bool        `json:"myhome"`
			} `json:"merchant"`
			Meta struct {
				Source     string      `json:"source"`
				Calc       interface{} `json:"calc"`
				HasInvoice bool        `json:"has_invoice"`
			} `json:"meta"`
			ProcessingID    interface{} `json:"processing_id"`
			CanRepeat       bool        `json:"can_repeat"`
			CanSave         bool        `json:"can_save"`
			CanCancel       bool        `json:"can_cancel"`
			CurrencyDetails struct {
				NumericCode int    `json:"numeric_code"`
				AlphaCode   string `json:"alpha_code"`
				Title       string `json:"title"`
			} `json:"currency_details"`
		} `json:"cheque"`
	} `json:"result"`
}
type PayPaymeSuccsesResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Cheque struct {
			ID         string `json:"_id"`
			CreateTime int64  `json:"create_time"`
			PayTime    int64  `json:"pay_time"`
			CancelTime int    `json:"cancel_time"`
			State      int    `json:"state"`
			Type       int    `json:"type"`
			External   bool   `json:"external"`
			Operation  int    `json:"operation"`
			Category   struct {
				ID        string `json:"_id"`
				Title     string `json:"title"`
				Color     string `json:"color"`
				Sort      int    `json:"sort"`
				Operation int    `json:"operation"`
				Indoor    bool   `json:"indoor"`
				Mcc       struct {
					Visa []string `json:"visa"`
				} `json:"mcc"`
			} `json:"category"`
			Error       interface{} `json:"error"`
			Description string      `json:"description"`
			Detail      interface{} `json:"detail"`
			Amount      int         `json:"amount"`
			Currency    int         `json:"currency"`
			Commission  int         `json:"commission"`
			Account     []struct {
				Name  string `json:"name"`
				Title string `json:"title"`
				Value string `json:"value"`
				Main  bool   `json:"main,omitempty"`
			} `json:"account"`
			Card struct {
				Number string `json:"number"`
				Expire string `json:"expire"`
				ID     string `json:"_id"`
				Name   string `json:"name"`
				Color  int    `json:"color"`
				Main   bool   `json:"main"`
			} `json:"card"`
			Payer struct {
				Phone string `json:"phone"`
				Lang  string `json:"lang"`
				IP    string `json:"ip"`
			} `json:"payer"`
			Merchant struct {
				ID           string `json:"_id"`
				Name         string `json:"name"`
				Organization string `json:"organization"`
				Address      string `json:"address"`
				BusinessID   string `json:"business_id"`
				Epos         struct {
					MerchantID string `json:"merchantId"`
					TerminalID string `json:"terminalId"`
				} `json:"epos"`
				Date   int64       `json:"date"`
				Logo   string      `json:"logo"`
				Type   string      `json:"type"`
				Terms  interface{} `json:"terms"`
				Myhome bool        `json:"myhome"`
			} `json:"merchant"`
			Meta struct {
				Source     string      `json:"source"`
				Calc       interface{} `json:"calc"`
				HasInvoice bool        `json:"has_invoice"`
			} `json:"meta"`
			ProcessingID    int  `json:"processing_id"`
			CanRepeat       bool `json:"can_repeat"`
			CanSave         bool `json:"can_save"`
			CanCancel       bool `json:"can_cancel"`
			CurrencyDetails struct {
				NumericCode int    `json:"numeric_code"`
				AlphaCode   string `json:"alpha_code"`
				Title       string `json:"title"`
			} `json:"currency_details"`
		} `json:"cheque"`
		Loyalty struct {
			Experience      int    `json:"experience"`
			ExperienceTitle string `json:"experience_title"`
		} `json:"loyalty"`
	} `json:"result"`
}
type CreatePaymeRequest struct {
	Phone        string
	AmountString string
	AmountInt    int
	Merchant_ID  string
}
type TokenExpireResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Method  string `json:"method"`
	} `json:"error"`
}
type TokenSuccsesResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Method string      `json:"method"`
		Data   interface{} `json:"data"`
	} `json:"result"`
}
type PhoneErrorResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    string `json:"data"`
		Origin  string `json:"origin"`
		Method  string `json:"method"`
	} `json:"error"`
}
type PriceNotResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Error   struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    string `json:"data"`
		Origin  string `json:"origin"`
		Method  string `json:"method"`
	} `json:"error"`
}