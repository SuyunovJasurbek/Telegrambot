package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"telegram_bot/models"

	"github.com/google/uuid"
)

func RandomNumberGet() (models.RandomNumberResponse, error) {

	url := "https://openbudget.uz/api/v2/vote/captcha-2"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "openbudget.uz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("access-captcha", "czM2ZTQ4azk2cjEwOGUxMDV0")
	req.Header.Add("hl", "uz_cyr")
	req.Header.Add("referer", "https://openbudget.uz/boards-list/1/52984948-43a3-49e0-9dc2-839e6d35e130")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var result models.RandomNumberResponse
	_ = json.Unmarshal(body, &result)
	fmt.Println("natija")
	return result, nil
}
func PhoneCheck(dos models.PhoneCheckRequest) (int, models.PhoneCheckResponse, error) {
	uuid := uuid.New().String()
	url := "https://openbudget.uz/api/v2/vote/check"
	method := "POST"
	dat := fmt.Sprintf(`{"captchaKey":"%s","captchaResult":%d,"phoneNumber":"%d","boardId":1}`, dos.CaptchaKey, dos.CaptchaResult, dos.Phone)
	payload := strings.NewReader(dat)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return -1, models.PhoneCheckResponse{}, err
	}
	req.Header.Add("authority", "openbudget.uz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "XSRF-TOKEN="+uuid+";")
	req.Header.Add("hl", "uz_cyr")
	req.Header.Add("origin", "https://openbudget.uz")
	req.Header.Add("referer", "https://openbudget.uz/boards-list/1/52984948-43a3-49e0-9dc2-839e6d35e130")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Add("x-xsrf-token", uuid)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return -1, models.PhoneCheckResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return -1, models.PhoneCheckResponse{}, err
	}
	if res.StatusCode == 200 {
		var resp models.PhoneCheckResponse
		err = json.Unmarshal(body, &resp)
		if err != nil {
			fmt.Println(err)
			return -1, models.PhoneCheckResponse{}, err
		}
		return 0, resp, nil
	}

	if res.StatusCode == 500 {
		return -1, models.PhoneCheckResponse{}, nil
	}

	if res.StatusCode == 400 {

		var resp models.PhoneCaptchaErrorResponse

		_ = json.Unmarshal(body, &resp)

		if resp.Code == 112 {
			return -1, models.PhoneCheckResponse{}, nil
		}
		if resp.Code == 103 {
			return 1, models.PhoneCheckResponse{}, nil
		}
	}
	return -2, models.PhoneCheckResponse{}, nil
}
func SmsCheck(da models.FinalOpenbudjetRequest) (int, error) {
	uuid := uuid.New().String()
	url := "https://openbudget.uz/api/v2/vote/verify"
	method := "POST"
	dat := fmt.Sprintf(`{"otpKey":"%s","otpCode":"%s","initiativeId":"52984948-43a3-49e0-9dc2-839e6d35e130","subinitiativesId":[]}`, da.OtpKey, da.OtpCode)
	payload := strings.NewReader(dat)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return -2, err
	}
	req.Header.Add("authority", "openbudget.uz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "XSRF-TOKEN="+uuid+";")
	req.Header.Add("hl", "uz_cyr")
	req.Header.Add("origin", "https://openbudget.uz")
	req.Header.Add("referer", "https://openbudget.uz/boards-list/1/52984948-43a3-49e0-9dc2-839e6d35e130")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Add("x-xsrf-token", uuid)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 2, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 2, err
	}
	if res.StatusCode == 200 {
		fmt.Println("OK")
		return 0, nil
	} else if res.StatusCode == 400 {
		var res models.SMSBadrequest
		_ = json.Unmarshal(body, &res)
		if res.Code == 108 {
			fmt.Println("SMS kod xato kiritildi")
			return 1, nil

		} else if res.Code == 109 {
			return -1, nil
		} else if res.Code==116 {
			return -2, nil
		}

	}
	return -1, nil
}