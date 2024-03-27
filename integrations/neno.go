package integrations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hyperbot/configs"
	"io"
	"net/http"
	"strings"
)

type nenoRequestBody struct {
	Message    string
	Recipients []string
}

func sendViaNeno(phone, message string) error {
	url := "https://useneno.online/v1"
	method := "POST"

	recipient := fmt.Sprintf("250%s", strings.TrimPrefix(phone, "0"))

	payload := nenoRequestBody{
		Message:    message,
		Recipients: []string{recipient},
	}

	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("x-api-key", configs.EnvNenoApiKey())
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(rbody))
	return nil
}

func SendOtpViaSMS(phone, code string) error {
	message := fmt.Sprintf("Your verification code is %s", code)
	return sendViaNeno(phone, message)
}
