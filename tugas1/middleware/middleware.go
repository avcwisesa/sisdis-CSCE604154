package middleware

import (
	"log"
	"os"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// define healthcheck middleware
func QuorumHealthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jsonFile, err := os.Open("quorum.json")
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()

		buf, _ := ioutil.ReadAll(jsonFile)

		var quorum map[string]string
		json.Unmarshal(buf, &quorum)

		quorumAmount := 1

		for id, host := range quorum {
			if id == "1506731561" {
				continue
			}
			pingReturn, err := ping(host)
			if err != nil {
				log.Println(err.Error())
			}

			if pingReturn == 1 {
				quorumAmount += pingReturn
			}
		}

		ctx.Set("quorum", quorumAmount)
	}
}

func ping(branch string) (int, error) {
	body := map[string]string{}

	jsonBody, err := json.Marshal(body)

	resp, err := http.Post(branch + "/ewallet/ping", "application/json",  bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var pingResponse map[string]int

	buf, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	json.Unmarshal(buf, &pingResponse)

	return pingResponse["pingReturn"], nil
}