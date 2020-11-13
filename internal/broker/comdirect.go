package broker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lars-wenk/_crawlr/internal/config"
	"github.com/lars-wenk/_crawlr/pkg/model"
)

const authURL = "https://api.comdirect.de/oauth/token"
const sessionURL = "https://api.comdirect.de/api/url/session/clients/user/v1/sessions"
const sessionTanActivationURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions/"
const tokenURL = "https://api.comdirect.de/api/oauth/token"
const balanceURL = "https://api.comdirect.de/api/banking/clients/user/v2/accounts/balances?"
const depotURL = "https://api.comdirect.de/api/brokerage/clients/user/v3/depots"

type ComdirectCrawler interface {
	GetAuth()
	getSession() (model.ComdirectSession, error)
	//GetStocks()
}

type comdirectCrawler struct {
	conf config.Config
	//db database.DB

}

func NewComdirectCrawler(conf config.Config) ComdirectCrawler {
	return &comdirectCrawler{
		conf: conf,
	}
}

func (c comdirectCrawler) GetAuth() {
	sessionID, err := c.getSession()
}

func (c comdirectCrawler) GetStocks() {
	return
}

func (c comdirectCrawler) getSession() (model.ComdirectSession, error) {
	var respSession = model.ComdirectSession{}
	t := "GET"
	reqH := map[string]string{
		"client_id":  c.conf.ComdirectClientID,
		"secret_id":  c.conf.ComdirectSecretID,
		"grant_type": "password",
		"username":   c.conf.ComdirectZugangsnummer,
		"password":   c.conf.ComdirectTAN,
	}
	reqB := map[string]string{}
	sessionBody, err := c.makeRequest(authURL, t, reqH, reqB)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(sessionBody, &respSession)
	if err != nil {
		return nil, err
	}

	return respSession, nil
}

func (c comdirectCrawler) makeRequest(url string, t string, reqH map[string]string, reqB map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: httpRequestTimeout,
	}

	for kH, vH := range reqH {
		req.Header.Set(kH, vH)
	}

	req, err := http.NewRequest(t, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
