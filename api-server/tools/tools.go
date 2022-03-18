package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/xid"
)

func GenId() string {
	guid := xid.New()
	return guid.String()
}

type Summary struct {
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Pubdate   string `json:"pubdate"`
	Cover     string `json:"cover"`
	Author    string `json:"author"`
}

type SummaryWrapper struct {
	Summary Summary `json:"summary"`
}

func FetchBooksData(isbn string) (Summary, error) {
	baseUrl := "https://api.openbd.jp/v1/get"
	baseQuery := "?isbn="

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s%s%s", baseUrl, baseQuery, isbn), nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return Summary{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Summary{}, err
	}

	var b []SummaryWrapper
	if err := json.Unmarshal(body, &b); err != nil {
		panic(err)
	}

	return b[0].Summary, err
}
