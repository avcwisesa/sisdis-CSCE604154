package communicator

import (
	"io/ioutil"
	"bytes"
	"encoding/json"
	"net/http"
	"errors"
)

func GetSaldo(branch , userID string) (int, error) {
	body := map[string]string{
		"user_id": userID,
	}

	jsonBody, err := json.Marshal(body)

	resp, err := http.Post(branch + "/ewallet/getSaldo", "application/json",  bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var saldoResponse map[string]int

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(buf, &saldoResponse)

	if saldoResponse["saldo"] == -1 {
		registerStatus, err := Register(branch, userID, userID)
		if err != nil {
			return 0, err
		}

		if registerStatus != 1{
			return 0, errors.New("Error occured in branch office <register>")
		}

		resp, err := http.Post(branch + "/ewallet/getSaldo", "application/json",  bytes.NewBuffer(jsonBody))
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		var saldoResponse map[string]int

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		json.Unmarshal(buf, &saldoResponse)

		if saldoResponse["saldo"] < 0 {
			return 0, errors.New("Error occured in branch office <getSaldo>")
		}

	} else if saldoResponse["saldo"] < 0 {
		return 0, errors.New("Error occured in branch office <getSaldo>")
	}

	return saldoResponse["saldo"], nil
}

func GetTotalSaldo(branch , userID string) (int, error) {
	body := map[string]string{
		"user_id": userID,
	}

	jsonBody, err := json.Marshal(body)

	resp, err := http.Post(branch + "/ewallet/getTotalSaldo", "application/json",  bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var saldoResponse map[string]int

	buf, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(buf, &saldoResponse)

	if saldoResponse["saldo"] < 0 {
		return 0, errors.New("Error occured in branch office <getSaldo>")
	}

	return saldoResponse["saldo"], nil
}

func Register(branch , name, userID string) (int, error) {
	body := map[string]string{
		"nama": userID,
		"user_id": userID,
	}

	jsonBody, err := json.Marshal(body)

	resp, err := http.Post(branch + "/ewallet/register", "application/json",  bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var registerResponse map[string]int

	buf, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(buf, &registerResponse)

	return registerResponse["registerReturn"], nil
}