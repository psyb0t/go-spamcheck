// wrapper for postmark spamcheck api
package spamcheck

import (
	"encoding/json"
	"net/http"
	"bytes"
)

// this type contains a checked rule's score and description
type Rule struct {
	Score       string `json:"score"`
	Description string `json:"description"`
}

// this type contains the response items of the spam check
type Response struct {
	Success bool    `json:"success"`
	Score   string  `json:"score"`
	Rules   []*Rule `json:"rules"`
	Report  string  `json:"report"`
}

// check if email is spam
//  - optional getLong param - if true get report details
func Check(email string, getLong ...bool) *Response {
	options := "short"
	if len(getLong) > 0 && getLong[0] {
		options = "long"
	}

	postData, err := json.Marshal(map[string]string{
		"email":   email,
		"options": options,
	})

	if err != nil {
		panic(err)
	}

	resp, err := http.Post(
		"https://spamcheck.postmarkapp.com/filter",
		"application/json",
		bytes.NewBuffer(postData),
	)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	response := &Response{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		panic(err)
	}

	return response
}
