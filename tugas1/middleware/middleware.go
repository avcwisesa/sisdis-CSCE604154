package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// define healthcheck middleware
func quorumHealthCheck(ctx *gin.Context) {
	select {
	case <-ctx.Request.Context().Done():
		return 0, ctx.Err()
	default:
	}

	jsonFile, err := os.Open("quorum.json")
	if err != nil {
		return 0, err
	}
	defer jsonFile.Close()

	buf, _ := ioutil.ReadAll(jsonFile)

	var quorum map[string]string
	json.Unmarshal(buf, &quorum)

	quorum := 1

	for id, host := range quorum {
		if id == "1506731561" continue
		pingReturn, err := ping(host)
		if err != nil {
			log.Println(err.Error())
		}

		if pingReturn == 1 {
			quorum += pingReturn
		}
	}

	c.Set("quorum", quorum)
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