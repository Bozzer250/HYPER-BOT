package integrations

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func sendViaNeno(phone, message string) error {
	url := "https://useneno.online/v1"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
    "message": %s,
    "recipients": [
        "25%s"
    ]
}`, message, phone))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("x-api-key", "4205bf4ca91043c6f779f3e133fcfb6d")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))
	return nil
}

func SendOtpViaSMS(phone, code string) error {
	// send OTP via Neno
	return sendViaNeno(phone, code)
}
