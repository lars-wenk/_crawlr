package broker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/lars-wenk/_crawlr/internal/config"
	"github.com/lars-wenk/_crawlr/pkg/model"
)

const authURL = "https://api.comdirect.de/oauth/token"
const sessionURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions"
const sessionTanActivationURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions/"
const tokenURL = "https://api.comdirect.de/api/oauth/token"
const balanceURL = "https://api.comdirect.de/api/banking/clients/user/v2/accounts/balances?"
const depotURL = "https://api.comdirect.de/api/brokerage/clients/user/v3/depots"

type ComdirectCrawler interface {
	GetAuth() error
	GetToken() (model.ComdirectAuthResponse, error)
	CheckSession(model.ComdirectAuthResponse) (model.ComdirectSession, error)
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

func (c comdirectCrawler) GetAuth() error {

	rToken, err := c.GetToken()
	fmt.Println("-----rToken: ----")
	fmt.Println(rToken)
	rCheck, err := c.CheckSession(rToken)
	fmt.Println("-----rCheck: ----")
	fmt.Println(rCheck)

	if err != nil {
		return err
	}
	return nil
}

func (c comdirectCrawler) GetStocks() {
	return
}

func (c comdirectCrawler) GetToken() (model.ComdirectAuthResponse, error) {
	var respSession = model.ComdirectAuthResponse{}
	t := "POST"
	reqH := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/json",
	}

	reqB := map[string]string{
		"client_id":     c.conf.ComdirectClientID,
		"client_secret": c.conf.ComdirectSecretID,
		"grant_type":    "password",
		"username":      c.conf.ComdirectZugangsnummer,
		"password":      c.conf.ComdirectTAN,
	}

	sessionBody, err := c.makeRequest(authURL, t, reqH, reqB)

	if err != nil {
		//@todo - error handling
	}

	err = json.Unmarshal(sessionBody, &respSession)
	if err != nil {
		//@todo - error handling
	}

	return respSession, nil
}

func (c comdirectCrawler) CheckSession(authResp model.ComdirectAuthResponse) (model.ComdirectSession, error) {
	t := "GET"

	r := model.ComdirectRequestInformation{}
	r.ClientRequestID.SessionID = c.generateRandomChars(32)
	r.ClientRequestID.RequestID = strconv.Itoa(c.generateRandomNumbers())

	b, err := json.Marshal(r)
	if err != nil {
		//@todo - error handling
	}

	reqH := map[string]string{
		"Content-Type":        "application/json",
		"Accept":              "application/json",
		"Authorization":       "Bearer " + authResp.AccessToken,
		"x-http-request-info": string(b),
	}

	sessionBody, err := c.makeRequest(sessionURL, t, reqH, nil)

	if err != nil {
		//@todo - error handling
	}
	var tmp []interface{}
	if err := json.Unmarshal(sessionBody, &tmp); err != nil {
		//@todo - error handling
	}

	md, ok := tmp[0].(map[string]interface{})

	if ok != true {
		//@todo - error handling
	}

	var rc = model.ComdirectSession{}
	rc.Identifier = md["identifier"].(string)
	rc.SessionTanActive = md["sessionTanActive"].(bool)
	rc.Activated2FA = md["activated2FA"].(bool)

	//fmt.Println(rc)

	if err != nil {
		//@todo - error handling
		fmt.Println("Fehler bei Unmarshal")
		fmt.Println(err)
	}

	return rc, nil
}

func (c comdirectCrawler) makeRequest(rURL string, t string, reqH map[string]string, reqB map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: httpRequestTimeout,
	}

	d := url.Values{}
	if len(reqB) > 0 && reqB != nil {
		for kB, vB := range reqB {
			d.Set(kB, vB)
		}
	}

	req, err := http.NewRequest(t, rURL, strings.NewReader(d.Encode()))
	if err != nil {
		return nil, err
	}

	for kH, vH := range reqH {
		req.Header.Set(kH, vH)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c comdirectCrawler) generateRandomChars(length int) string {
	const charset = "-abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func (c comdirectCrawler) generateRandomNumbers() int {
	rand.Seed(time.Now().UnixNano())
	min := 111111111
	max := 999999999

	return min + rand.Intn(max-min)
}
