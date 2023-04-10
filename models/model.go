package models

type DataCreate struct {
	ChatID   int    `json:"chat_id"`
	Phone    int    `json:"phone"`
	FullName string `json:"full_name"`
	LastName string `json:"last_name"`
	Status   int    `json:"status"`
	Type     int    `json:"type"`
	Time     int    `json:"time"`
}

type RandomNumberResponse struct {
	Image      string `json:"image"`
	CaptchaKey string `json:"captchaKey"`
}
type PhoneCheckResponse struct {
	OtpKey     string `json:"otpKey"`
	RetryAfter int    `json:"retryAfter"`
}
	
type PhoneCaptchaErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type UserAmountRequest struct {
	Image      string `json:"image"`
	CaptchaKey string `json:"captchaKey"`
}

type CheckResponce struct {
	Phone  int    `json:"phone"`
	OptKey string `json:"otpKey"`
}

type PhoneCheckRequest struct {
	CaptchaKey    string `json:"captchaKey"`
	Phone         int    `json:"phone"`
	CaptchaResult int    `json:"captchaResult"`
}

type FinalSMSCodeRequest struct {
	OptKey string `json:"otpKey"`
	Key   string `json:"key"`
}

type FinalOpenbudjetRequest struct {
	OtpKey       string `json:"otpKey"`
	OtpCode      string `json:"otpCode"`
	InitiativeID string `json:"initiativeId"`
}

type SMSBadrequest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AmountCreateRequest struct {
	ChatID int  `json:"chat_id"`
	Time string `json:"time"`
}
	
