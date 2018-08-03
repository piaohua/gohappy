//微信登录
package data

import (
	"testing"

	"github.com/globalsign/mgo"
)

func TestWxLogin(t *testing.T) {
	collection := "col_wx_login"
	_DBNAME := "test"
	host := "127.0.0.1:2225"
	var se *mgo.Session
	var err error
	se, err = mgo.Dial(host)
	se.DB(_DBNAME).C(collection)
	t.Log(err)

	openid := "o48K-wCLzOgvD7b2_kbllFcNHDmQ"
	d := &WXLogin{
		OpenId: openid,
		//AccessToken: openid,
	}
	//err = d.Get()
	//err = d.GetByToken()
	//err = d.Save()
	ok := d.Update()
	t.Log(d, ok)

	//u := &User{
	//	Userid: "16007",
	//}
	//err = u.Get()
	//t.Log(u, err)
}