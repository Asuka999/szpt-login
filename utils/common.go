package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func GetEncry(resp *http.Response) (Lt string, pwdEncryptSalt string) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	Lt = doc.Find("input[name='lt']").Get(1).Attr[2].Val
	pwdEncryptSalt = doc.Find("input[id='pwdDefaultEncryptSalt']").Get(0).Attr[2].Val
	return Lt, pwdEncryptSalt
}

func EncryPasswd(passwd string, pwdEncryptSalt string) (EncryptedPwd string) {
	var iv string = "TksnjaCpP5ZwfxhN"
	var aes64 string = "MmzS5PzdCrzGDZzN8A4tZ5bZx6Z8BTcrKQTyippksjFYDABmKwGP8TEfRRtmmzrX"
	EncryptedPwd = AesEncrypt(pwdEncryptSalt, iv, aes64+passwd)
	return EncryptedPwd
}
