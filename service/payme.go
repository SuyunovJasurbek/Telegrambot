package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"telegram_bot/models"
)

func LoginPayme() string {

	url := "https://payme.uz/api/users.log_in"
	method := "POST"

	payload := strings.NewReader(`{"method":"users.log_in","params":{"login":"********","password":"*********"}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "payme.uz")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("app-version", "10.86.1491")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "ngx-device=6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461; _ym_uid=1676636425401500418; _ym_d=1676636425; _ga=GA1.2.1321590121.1676636425; _fbp=fb.1.1676636425148.551641001; _gid=GA1.2.1595986526.1679967894; _ym_isad=2")
	req.Header.Add("device", "63fdb80e71a4913049abb571; R@CZR%ZjiU@AHHGcrHEvTuiF7pjp=98%I0b07rtdwv1Cf3bwQxODOcbJO6M5ibph;")
	req.Header.Add("origin", "https://payme.uz")
	req.Header.Add("referer", "https://payme.uz/login")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("track-id", "6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Add("x-accept-language", "uz")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	api_session := res.Header.Get("API-SESSION")
	return api_session
}
func CreatePayme(da models.CreatePaymeRequest) string {

	url := "https://payme.uz/api/cheque.create"
	method := "POST"
	dat := fmt.Sprintf(`{"method":"cheque.create","params":{"account":{"phone":"%s","amount":"%s"},"amount":%d,"merchant_id":"%s"}}`, da.Phone, da.AmountString, da.AmountInt, da.Merchant_ID)
	payload := strings.NewReader(dat)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("authority", "payme.uz")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("app-theme", "light")
	req.Header.Add("app-version", "10.86.1491")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "ngx-device=6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461; _ym_uid=1676636425401500418; _ym_d=1676636425; _ga=GA1.2.1321590121.1676636425; _fbp=fb.1.1676636425148.551641001; _gid=GA1.2.1595986526.1679967894; _ym_isad=2")
	req.Header.Add("device", "63fdb80e71a4913049abb571; R@CZR%ZjiU@AHHGcrHEvTuiF7pjp=98%I0b07rtdwv1Cf3bwQxODOcbJO6M5ibph;")
	req.Header.Add("origin", "https://payme.uz")
	req.Header.Add("referer", "https://payme.uz/cabinet/main")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("track-id", "6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Add("x-accept-language", "uz")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var resp models.PaymeCreateResponce
	var phoneError models.PhoneErrorResponse
	_ = json.Unmarshal(body, &phoneError)
	_ = json.Unmarshal(body, &resp)
	if phoneError.Error.Origin=="receipt.create" {
		return "phone"
	}
	return resp.Result.Cheque.ID
}
func PayPayme(id, token string) int {

	url := "https://payme.uz/api/cheque.pay"
	method := "POST"
	dat := fmt.Sprintf(`{"method":"cheque.pay","params":{"id":"%s","card_id":"63ef5995f9b3d2b5a8256f2b"}}`, id)
	payload := strings.NewReader(dat)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	req.Header.Add("authority", "payme.uz")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,uz;q=0.6")
	req.Header.Add("api-session", token)
	req.Header.Add("app-theme", "light")
	req.Header.Add("app-version", "10.86.1491")
	req.Header.Add("content-type", "application/json; charset=UTF-8")
	req.Header.Add("cookie", "ngx-device=6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461; _ym_uid=1676636425401500418; _ym_d=1676636425; _ga=GA1.2.1321590121.1676636425; _fbp=fb.1.1676636425148.551641001; _gid=GA1.2.1595986526.1679967894; _ym_isad=2")
	req.Header.Add("device", "63fdb80e71a4913049abb571; R@CZR%ZjiU@AHHGcrHEvTuiF7pjp=98%I0b07rtdwv1Cf3bwQxODOcbJO6M5ibph;")
	req.Header.Add("origin", "https://payme.uz")
	req.Header.Add("referer", "https://payme.uz/checkout/6422a82f0395597ae7fd193f")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Linux\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("track-id", "6471c47996ac1e8c1eff9a076aae6a42bbf6b73ea319623fea6904afa3954461")
	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	req.Header.Add("x-accept-language", "uz")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	var tokenExpire models.TokenExpireResponse
	var paymeSuccsesfule models.PayPaymeSuccsesResponse
	var pricenot models.PriceNotResponse
	_ = json.Unmarshal(body, &tokenExpire)
	_ = json.Unmarshal(body, &paymeSuccsesfule)
	_ = json.Unmarshal(body, &pricenot)
	if pricenot.Error.Data == "insufficient_funds" {
		return 0
	}else if tokenExpire.Error.Message == "" && len(paymeSuccsesfule.Result.Loyalty.ExperienceTitle) > 0 {
		return 1
	} else if tokenExpire.Error.Message == "Access denied." && len(paymeSuccsesfule.Result.Loyalty.ExperienceTitle) == 0 {
		return -1
	} 
	return 2
}