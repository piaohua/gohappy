package data

import (
	"io/ioutil"

	"github.com/json-iterator/go"
)

var robotmap map[string]RobotData

type RobotData struct {
	Id        uint32 `csv:"id"`
	Kind      uint32 `csv:"kind"`
	Changci   uint32 `csv:"changci"`
	Nickname  string `csv:"nickname"`
	Sex       uint32 `csv:"sex"`
	Level     uint32 `csv:"level"`
	Coin      uint32 `csv:"coin"`
	Diamond   uint32 `csv:"diamond"`
	Cointime  uint32 `csv:"cointime"`
	Coinup    uint32 `csv:"coinup"`
	Vip       uint32 `csv:"vip"`
	Headframe uint32 `csv:"headframe"`
	Phone     string `csv:"phone"`
}

func GetRobotList() []RobotData {
	return robotList
}

var robotList []RobotData

func GetRobot(phone string) *RobotData {
	if v, ok := robotmap[phone]; ok {
		return &v
	} else {
		return nil
	}

}

/*
func init() {
	// robot.csv UTF-8 格式读取的v.Id为0,ANSI格式读取正常
	f, err := os.Open("robot.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = csv.Unmarshal(data, &robotList)
	if err != nil {
		panic(err)
	}
	robotmap = make(map[string]RobotData)
	for _, v := range robotList {
		robotmap[v.Phone] = v
		//glog.Infoln(v.Phone)
	}
}
*/

//RobotInfo robot info
type RobotInfo struct {
	ID int32 `json:"id"`
	Nickname string `json:"nickname"`
	Sex uint32 `json:"sex"`
	Coin int64 `json:"coin"`
	Diamond int64 `json:"diamond"`
	Vip uint32 `json:"vip"`
	Phone string `json:"phone"`
}

//LoadRobotInfo load robot info by robot.json
func LoadRobotInfo(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal(data, v)
}