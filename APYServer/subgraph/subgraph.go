package subgraph

import (
	"encoding/json"
	"net/http"
	"strings"
)

func NewSuperFormInterface(url string, apiKey string) *SuperformInterface {
	return &SuperformInterface{
		Client:  http.Client{},
		URL:     url,
		API_KEY: apiKey,
	}
}

func (s *SuperformInterface) GetDepositsWithdrawals(user string) ([]Deposit, []Withdrawal, error) {
	payload := strings.NewReader(superformreq(user))
	req, err := http.NewRequest("POST", s.URL, payload)
	if err != nil {
		return nil, nil, err
	}
	res, err := s.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	builderResponse := new(UserDepositsWithdrawals)
	if err := json.NewDecoder(res.Body).Decode(&builderResponse); err != nil {
		return nil, nil, err
	}
	return builderResponse.Data.Deposits, builderResponse.Data.Withdrawals, nil
}
