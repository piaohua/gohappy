package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"gohappy/pb"

	jsoniter "github.com/json-iterator/go"
)

func TestRun(t *testing.T) {
	doc := aesEn("127.0.0.1:8080")
	t.Logf("doc decode %s", string(doc))
	t.Logf("doc encode %s", aesDe(doc))
}

func TestWebJson(t *testing.T) {
	webJson("101418", 1000)
}

func webJson(userid string, coin int64) {
	//var userid string
	//var coin int64
	//flag.StringVar(&userid, "userid", "", "userid")
	//flag.Int64Var(&coin, "coin", 0, "coin")
	//flag.Parse()
	log.Printf("userid %s, coin %d\n", userid, coin)
	msg := &pb.PayCurrency{
		Type:   int32(pb.LOG_TYPE9),
		Userid: userid,
		Coin:   coin,
	}
	b, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b %s\n", string(b))
	msg2 := &pb.WebRequest2{
		Code: pb.WebGive,
		Data: string(b),
	}
	b, err = jsoniter.Marshal(msg2)
	if err != nil {
		log.Panic(err)
	}
	url := "http://127.0.0.1/happy/webjson"
	b, err = doHTTPPost(url, b)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("result %s", string(b))
	msg3 := new(pb.WebResponse2)
	err = jsoniter.Unmarshal(b, msg3)
	if err != nil {
		log.Panic(err)
	}
	if msg3.Code == msg2.Code {
		log.Printf("userid %s, coin %d send successfully.", userid, coin)
	} else {
		log.Printf("userid %s, coin %d send failed.", userid, coin)
	}
}

func doHTTPPost(targetURL string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(body))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/json;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}
