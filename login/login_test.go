package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"gohappy/data"
	"gohappy/pb"

	"github.com/globalsign/mgo/bson"
	jsoniter "github.com/json-iterator/go"
)

func TestRun(t *testing.T) {
	doc := aesEn("127.0.0.1:8080")
	t.Logf("doc decode %s", string(doc))
	t.Logf("doc encode %s", aesDe(doc))
}

func TestWebJson(t *testing.T) {
	sendCoin("101418", 1000)
	sendNotice()
	//sendBuild("101418", "101133")
	//sendState("101418", 1, 1)
	//sendRate("101418", 23)
}

func sendCoin(userid string, coin int64) {
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
	if webRequest(pb.WebGive, b) {
		log.Printf("userid %s, coin %d send successfully.", userid, coin)
	} else {
		log.Printf("userid %s, coin %d send failed.", userid, coin)
	}
}

func sendNotice() {
	msg := make(map[string]data.Notice) //key: Notice.Id
	notice := data.Notice{
		Id:      bson.NewObjectId().Hex(),
		Userid:  "",
		Rtype:   data.NOTICE_TYPE3,
		Content: "恭喜成功购买100金豆",
		Ctime:   bson.Now(),
		Etime:   bson.Now().AddDate(0, 0, 7),
	}
	msg[notice.Id] = notice
	b, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b %s\n", string(b))
	if webRequest(pb.WebNotice, b) {
		log.Printf("msg %#v send successfully.", msg)
	} else {
		log.Printf("msg %#v send failed.", msg)
	}
}

func sendRate(userid string, rate uint32) {
	log.Printf("userid %s, rate %d\n", userid, rate)
	msg := &pb.SetAgentProfitRate{
		Userid: userid,
		Rate:   rate,
	}
	b, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b %s\n", string(b))
	if webRequest(pb.WebRate, b) {
		log.Printf("userid %s, rate %d send successfully.", userid, rate)
	} else {
		log.Printf("userid %s, rate %d send failed.", userid, rate)
	}
}

func sendBuild(userid, agent string) {
	log.Printf("userid %s, agent %s\n", userid, agent)
	msg := &pb.SetAgentBuild{
		Userid: userid,
		Agent:  agent,
	}
	b, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b %s\n", string(b))
	if webRequest(pb.WebBuild, b) {
		log.Printf("userid %s, agent %s send successfully.", userid, agent)
	} else {
		log.Printf("userid %s, agent %s send failed.", userid, agent)
	}
}

func sendState(userid string, state, level uint32) {
	log.Printf("userid %s, state %d, level %d\n", userid, state, level)
	msg := &pb.SetAgentState{
		Userid: userid,
		State:  state,
		Level:  level,
	}
	b, err := jsoniter.Marshal(msg)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("b %s\n", string(b))
	if webRequest(pb.WebState, b) {
		log.Printf("userid %s, state %d, level %d send successfully.", userid, state, level)
	} else {
		log.Printf("userid %s, state %d, level %d send failed.", userid, state, level)
	}
}

func webRequest(code pb.WebCode, b []byte) bool {
	msg2 := &pb.WebRequest2{
		Code:  code,
		Atype: pb.CONFIG_UPSERT,
		Data:  string(b),
	}
	b, err := jsoniter.Marshal(msg2)
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
	return msg3.Code == msg2.Code
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
