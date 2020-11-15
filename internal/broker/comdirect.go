package broker

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/lars-wenk/_crawlr/internal/config"
	"github.com/lars-wenk/_crawlr/internal/utils"
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
	CheckSession(model.ComdirectAuthResponse) (model.ComdirectSession, string, error)
	TwoFactorAuth(model.ComdirectAuthResponse, model.ComdirectSession, string) (model.ComdirectSession, error)
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
	rCheck, requestInfo, err := c.CheckSession(rToken)
	fmt.Println("-----rCheck: ----")
	fmt.Println(rCheck)
	r2FA, err := c.TwoFactorAuth(rToken, rCheck, requestInfo)
	fmt.Println("-----r2FA: ----")
	fmt.Println(r2FA)

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

	reqH := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Accept":       "application/json",
	}

	d := url.Values{}
	d.Set("client_id", c.conf.ComdirectClientID)
	d.Set("client_secret", c.conf.ComdirectSecretID)
	d.Set("grant_type", "password")
	d.Set("username", c.conf.ComdirectZugangsnummer)
	d.Set("password", c.conf.ComdirectTAN)

	cr := utils.NewHttpRequest()
	sessionBody, err := cr.MakePostRequest(authURL, reqH, d)

	if err != nil {
		//@todo - error handling
	}

	err = json.Unmarshal(sessionBody, &respSession)
	if err != nil {
		//@todo - error handling
	}

	return respSession, nil
}

func (c comdirectCrawler) CheckSession(authResp model.ComdirectAuthResponse) (model.ComdirectSession, string, error) {

	rIStrg := string(c.getRequestInformation())
	reqH := map[string]string{
		"Content-Type":        "application/json",
		"Accept":              "application/json",
		"Authorization":       "Bearer " + authResp.AccessToken,
		"x-http-request-info": rIStrg,
	}

	cr := utils.NewHttpRequest()
	sessionBody, err := cr.MakeGetRequest(sessionURL, reqH)

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

	return rc, rIStrg, nil
}

func (c comdirectCrawler) TwoFactorAuth(authResp model.ComdirectAuthResponse, session model.ComdirectSession, requestInfo string) (model.ComdirectSession, error) {

	reqH := map[string]string{
		"Content-Type":        "application/json",
		"Accept":              "application/json",
		"Authorization":       "Bearer " + authResp.AccessToken,
		"x-http-request-info": requestInfo,
	}

	session.Activated2FA = true
	session.SessionTanActive = true

	rBJSON, err := json.Marshal(session)
	rBJSONtoStrg := string(rBJSON)
	sessionURLValidate := sessionURL + "/" + session.Identifier + "/validate"

	cr := utils.NewHttpRequest()
	sessionBody, err := cr.MakePostJSONRequest(sessionURLValidate, reqH, rBJSONtoStrg)

	var rc = model.ComdirectSession{}
	err = json.Unmarshal(sessionBody, &rc)
	if err != nil {
		//@todo - error handling
	}

	return rc, nil

}

func (c comdirectCrawler) getRequestInformation() []byte {
	r := model.ComdirectRequestInformation{}
	r.ClientRequestID.SessionID = c.generateRandomChars(32)
	r.ClientRequestID.RequestID = strconv.Itoa(c.generateRandomNumbers())

	b, err := json.Marshal(r)
	if err != nil {
		//@todo - error handling
	}

	return b
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
