//微信登录
package data

import (
	"testing"

	"utils"

	"github.com/globalsign/mgo/bson"
	//"github.com/globalsign/mgo"
	"gohappy/pb"
	"log"
)

func TestRecord(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	list1, list2, list3, err := GetRecords("10026", 1)
	t.Log(err)
	t.Logf("list1 %#v, list2 %#v, list3 %#v", list1, list2, list3)
	for k, v := range list1 {
		t.Log("k ", k, v.Roomid, v.Ctime)
	}
}

func TestGroup(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	var types []string
	ListByQ(UserRecords, bson.M{"$group": bson.M{"userid": "$userid"}}, &types)
	t.Logf("%#v", types)
	m := bson.M{
		"$match": bson.M{
		//"diamond": bson.M{"$ne": 0},
		},
	}
	o := bson.M{
		"$project": bson.M{
			"_id":     1,
			"diamond": 1,
		},
	}
	n := bson.M{
		"$group": bson.M{
			"_id": "$_id",
		},
	}
	operations := []bson.M{m, o, n}
	result := []bson.M{}
	pipe := PlayerUsers.Pipe(operations)
	err2 := pipe.All(&result)
	t.Logf("%#v", err2)
	t.Logf("%#v", result)
}

func TestOr(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	or := []bson.M{bson.M{"userid": "10030"}}
	m := bson.M{
		"$match": bson.M{
			"$or": or,
		},
	}
	operations := []bson.M{m}
	result := []bson.M{}
	pipe := TradeRecords.Pipe(operations)
	err2 := pipe.All(&result)
	t.Logf("%#v", err2)
	t.Logf("%#v", result)
}

func TestID(t *testing.T) {
	//dayStamp := utils.Stamp2Time(utils.TimestampYesterday())
	dayStamp := utils.Stamp2Time(utils.TimestampToday())
	id := bson.NewObjectIdWithTime(dayStamp).Hex()
	t.Logf("%s", id)
	id2 := bson.NewObjectIdWithTime(dayStamp).String()
	t.Logf("%s", id2)
	t.Logf("%#v", dayStamp)
}

func TestProfit(t *testing.T) {
	InitMgo("127.0.0.1", "27017", "", "", "test")
	arg := &pb.CAgentDayProfit{
		Selfid: "108377",
		//StartTime: "2018-07-20 00:00:00",
		//EndTime: "2018-08-02 00:00:00",
	}
	list, err := GetAgentDayProfit(arg)
	log.Printf("list %#v, err %v\n", list, err)
	num, err2 := GetAgentDayProfitCount(arg)
	log.Printf("num %d, err2 %v\n", num, err2)

	/*
	arg2 := &pb.LogProfit{
		Userid    : "104143",
		Gtype     : 1,
		Level     : 1,
		Rate      : 0,
		Profit    : 100,
		Agentid   : "100105",
		Type      : 52,
		Agentnote : "test",
		Nickname  : "aaaa",
	}
	for i := 0; i <= 2; i++ {
		go data.DayProfitRecord(arg2)
	}
	*/
}