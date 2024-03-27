package integrations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hyperbot/configs"
	"net/http"
)

type PollResponseParams struct {
	Amount    int     `json:"amount"`
	Client    string  `json:"client"`
	Fee       float64 `json:"fee"`
	Kind      string  `json:"kind"`
	Merchant  string  `json:"merchant"`
	Ref       string  `json:"ref"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
}

type authResp struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
	Expires int    `json:"expires"`
}

type cashInRequest struct {
	Amount int    `json:"amount"`
	Number string `json:"number"`
}

type CashinResponseParams struct {
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
	Kind      string `json:"kind"`
	Ref       string `json:"ref"`
	Status    string `json:"status"`
}

func Authenticate() (string, error) {
	url := "https://payments.paypack.rw/api/auth/agents/authorize"
	method := "POST"
	var response *authResp
	data := map[string]string{"client_id": configs.GetPaypackId(), "client_secret": configs.GetPaypackSecret()}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return response.Access, nil
}

func PaypackCashIn(amount int, number string) (CashinResponseParams, error) {
	var response *CashinResponseParams
	url := "https://payments.paypack.rw/api/transactions/cashin"
	method := "POST"
	token, _ := Authenticate()

	data := cashInRequest{
		Amount: amount,
		Number: number,
	}
	body := new(bytes.Buffer)
	_ = json.NewEncoder(body).Encode(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}

	return *response, nil
}

func PollTransactionStatus(ref string) (PollResponseParams, error) {
	var response *PollResponseParams
	url := fmt.Sprintf("https://payments.paypack.rw/api/transactions/find/%s", ref)
	method := "GET"
	token, _ := Authenticate()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return *response, err
	}

	return *response, nil
}
