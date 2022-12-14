package login

import (
	"errors"
	"github.com/Asuka999/szpt-login/utils"
	cookiejar "github.com/juju/persistent-cookiejar"
	"net/http"
	"net/url"
	"strings"
)

type loginReply interface {
	GetCookiesString() string
	GetCookiesMap() (cookies []*http.Cookie)
	GetClinet() *http.Client
	GetJar() *cookiejar.Jar
	GetInstance() *User
}

func (U *User) GetCookiesString() string {
	Json, _ := U.Jar.MarshalJSON()

	return string(Json)
}

func (U *User) GetCookiesMap() (cookies []*http.Cookie) {
	cookies = U.Jar.AllCookies()
	return cookies
}

func (U *User) GetClinet() *http.Client {
	return U.Clinet
}

func Login(Account string, Passwrod string) (loginReply, error) {
	user := &User{
		account:  Account,
		passwrod: Passwrod,
	}
	user.newHttpClinet()
	user.getEncryInfo()
	err := user.login()

	if err != nil {
		return nil, err
	}
	return user, nil
}

type User struct {
	account        string
	passwrod       string
	lt             string
	pwdEncryptSalt string
	Jar            *cookiejar.Jar
	Clinet         *http.Client
	encryptedPwd   string
}

func setProxyUrl() (proxyUrl *url.URL) {
	proxyUrl, _ = url.Parse("http://127.0.0.1:9090")
	return proxyUrl
}

func (U *User) GetJar() *cookiejar.Jar {
	return U.Jar
}

func (U *User) newHttpClinet() {
	U.Jar, _ = cookiejar.New(nil)

	U.Clinet = &http.Client{
		Transport: &http.Transport{
			// set proxyman proxy
			//Proxy: http.ProxyURL(setProxyUrl()),
		},
		Jar: U.Jar,
	}
}

func (U *User) GetInstance() *User {
	return U
}

func (U *User) getEncryInfo() {
	resp, _ := U.Clinet.Get("https://authserver.szpt.edu.cn/authserver/login")
	U.lt, U.pwdEncryptSalt = utils.GetEncry(resp)
	U.encryptedPwd = utils.EncryPasswd(U.passwrod, U.pwdEncryptSalt)

}

func (U *User) login() error {
	requestForm := strings.NewReader(url.Values{"username": {U.account}, "password": {U.encryptedPwd}, "lt": {U.lt}, "dllt": {"userNamePasswordLogin"}, "execution": {"e1s1"}, "_eventId": {"submit"}, "rmShown": {"1"}}.Encode())
	req, _ := http.NewRequest("POST", "https://authserver.szpt.edu.cn/authserver/login", requestForm)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	loginResp, _ := U.Clinet.Do(req)
	if loginResp.Request.URL.Path == "/authserver/index.do" {
		U.Clinet.Get("https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/*default/index.do#/")
		menuinfoForm := "data=%7B%22APPID%22%3A%225812981499622390%22%2C%22APPNAME%22%3A%22szptpubxsjkxxbs%22%7D"
		menuinfoForms := strings.NewReader(menuinfoForm)
		U.Clinet.Post("https://ehall.szpt.edu.cn/publicappinternet/sys/itpub/MobileCommon/getMenuInfo.do", "application/x-www-form-urlencoded", menuinfoForms)
		return nil
	}
	return errors.New("????????????????????????????????????")
}
