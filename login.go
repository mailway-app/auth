package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mailway-app/config"

	"github.com/pkg/errors"
)

type LoginResponse struct {
	Data struct {
		Login bool `json:"login"`
	} `json:"data"`
}

func login(user, pass string) (bool, error) {
	url := fmt.Sprintf("https://apiv1.mailway.app/instance/%s/responder/login", config.CurrConfig.ServerId)

	var jsonStr = []byte(fmt.Sprintf(`{"address": "%s","password": "%s"}`, user, pass))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return false, errors.Wrap(err, "could not create request")
	}

	req.Header.Add("Authorization", "Bearer "+config.CurrConfig.ServerJWT)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.Wrap(err, "could not send request")
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == 200 {
		var data LoginResponse
		if err := json.Unmarshal(body, &data); err != nil {
			return false, errors.Wrap(err, "could not parse response")
		}

		return data.Data.Login, nil
	} else {
		return false, errors.Errorf("login returned %d: %s", resp.StatusCode, body)
	}
}
