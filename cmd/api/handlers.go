package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (app *config) sendOTP(w http.ResponseWriter, r *http.Request) {

	url := "https://control.msg91.com/api/v5/flow/"

	payload := strings.NewReader("{\"template_id\":\"EntertemplateID\",\"short_url\":\"1 (On) or 0 (Off)\",\"recipients\":[{\"mobiles\":\"919XXXXXXXXX\",\"VAR1\":\"VALUE1\",\"VAR2\":\"VALUE2\"}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authkey", "Enter your MSG91 authkey")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
