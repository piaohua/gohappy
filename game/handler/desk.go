package handler

/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-12-17 18:22:30
 * Filename      : desk.go
 * Description   : 桌子数据处理
 * *******************************************************/

import (
	"gohappy/data"
	"gohappy/glog"
	"gohappy/pb"
	"utils"

	"github.com/globalsign/mgo/bson"
	jsoniter "github.com/json-iterator/go"
)

//Desk2Data 打包桌子数据
func Desk2Data(deskData *data.DeskData) []byte {
	result, err := jsoniter.Marshal(deskData)
	if err != nil {
		glog.Errorf("Desk2Data Marshal err %v", err)
		return []byte{}
	}
	return result
}

//Data2Desk 解析桌子数据
func Data2Desk(deskDataStr []byte) *data.DeskData {
	deskData := new(data.DeskData)
	err := jsoniter.Unmarshal(deskDataStr, deskData)
	if err != nil {
		glog.Errorf("Data2Desk Unmarshal err %v", err)
		return nil
	}
	return deskData
}

//NewDeskData 转换为桌子数据
func NewDeskData(d *data.Game) *data.DeskData {
	//TODO 新加字段
	return &data.DeskData{
		Unique:  d.Id,
		Gtype:   d.Gtype,
		Rtype:   d.Rtype,
		Dtype:   d.Dtype,
		Ltype:   d.Ltype,
		Rname:   d.Name,
		Count:   d.Count,
		Ante:    d.Ante,
		Cost:    d.Cost,
		Vip:     d.Vip,
		Chip:    d.Chip,
		Deal:    d.Deal,
		Carry:   d.Carry,
		Down:    d.Down,
		Top:     d.Top,
		Sit:     d.Sit,
		Minimum: d.Minimum,
		Maximum: d.Maximum,
	}
}

//NewFreeGameData 百人房间桌子数据
func NewFreeGameData(node string, gtype int32) *data.Game {
	return &data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  gtype,
		Rtype:  int32(pb.ROOM_TYPE2),
		Status: 1,
		Count:  100,
		Ante:   1,
		Deal:   true,
		Carry:  20000,
		Down:   10000,
		Top:    2000000,
		Sit:    20000,
		Node:   node,
		Ctime:  bson.Now(),
	}
}

//NewCoinGameData 自由房间桌子数据
func NewCoinGameData(node string, gtype, dtype, ltype int32) *data.Game {
	g := &data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  gtype,
		Rtype:  int32(pb.ROOM_TYPE0),
		Dtype:  dtype,
		Ltype:  ltype,
		Status: 1,
		Count:  5,
		Ante:   1,
		Deal:   true,
		Chip:   20000,
		Sit:    20000,
		Node:   node,
		Ctime:  bson.Now(),
	}
	switch ltype {
	case int32(pb.ROOM_LEVEL0):
		g.Ante = 1
		g.Chip = 2000
		g.Sit = 2000
		g.Minimum = 2000
		g.Maximum = 5000
	case int32(pb.ROOM_LEVEL1):
		g.Ante = 2
		g.Chip = 5000
		g.Sit = 5000
		g.Minimum = 5000
		g.Maximum = 50000
	case int32(pb.ROOM_LEVEL2):
		g.Ante = 5
		g.Chip = 20000
		g.Sit = 20000
		g.Minimum = 20000
		g.Maximum = 200000
	case int32(pb.ROOM_LEVEL3):
		g.Ante = 10
		g.Chip = 200000
		g.Sit = 200000
		g.Minimum = 200000
		g.Maximum = 2000000
	case int32(pb.ROOM_LEVEL4):
		g.Ante = 20
		g.Chip = 2000000
		g.Sit = 2000000
		g.Minimum = 2000000
		g.Maximum = 0
	}
	return g
}

//NewPrivGameData 私人房间桌子数据
func NewPrivGameData(arg *pb.CreateDesk) *data.DeskData {
	return &data.DeskData{
		Unique:  bson.NewObjectId().Hex(),
		Gtype:   arg.Gtype,
		Rtype:   arg.Rtype,
		Dtype:   arg.Dtype,
		Rname:   arg.Rname,
		Ante:    arg.Ante,
		Count:   arg.Count,
		Round:   arg.Round,
		Payment: arg.Payment,
		Cost:    arg.Cost,
		Cid:     arg.Cid,
		Ctime:   uint32(utils.Timestamp()),
		Expire:  utils.Timestamp() + 600,
		Deal:    true,
	}
}

/*
//SetGameList 游戏房间配置，测试数据
func SetGameList() {
	g1 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_NIU,
		Rtype:  data.ROOM_TYPE0,
		Ltype:  data.GAME_BJPK10,
		Name:   "牛牛1区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   true,
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    1,
		Ctime:  time.Now(),
	}
	g2 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_SAN,
		Rtype:  data.ROOM_TYPE0,
		Ltype:  data.GAME_BJPK10,
		Name:   "三公1区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   true,
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    1,
		Ctime:  time.Now(),
	}
	g3 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_JIU,
		Rtype:  data.ROOM_TYPE0, //免佣房间
		Ltype:  data.GAME_BJPK10,
		Name:   "北京赛车1区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   false, //无庄
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    0,
		Node:   "game.huiyin1",
		Ctime:  time.Now(),
	}
	g4 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_JIU,
		Rtype:  data.ROOM_TYPE1, //抽佣房间
		Ltype:  data.GAME_BJPK10,
		Name:   "北京赛车2区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   false, //无庄
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    0,
		Node:   "game.huiyin1",
		Ctime:  time.Now(),
	}
	g5 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_JIU,
		Rtype:  data.ROOM_TYPE0, //免佣房间
		Ltype:  data.GAME_MLAFT,
		Name:   "幸运飞艇1区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   false, //有庄
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    0,
		Node:   "game.huiyin2",
		Ctime:  time.Now(),
	}
	g6 := data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  data.GAME_JIU,
		Rtype:  data.ROOM_TYPE1, //抽佣房间
		Ltype:  data.GAME_MLAFT,
		Name:   "幸运飞艇2区",
		Status: 1,
		Count:  100,
		Ante:   1,
		Cost:   5,
		Vip:    0,
		Chip:   0,
		Deal:   false, //有庄
		Carry:  20000,
		Down:   10000,
		Top:    60000,
		Sit:    20000,
		Del:    0,
		Node:   "game.huiyin2",
		Ctime:  time.Now(),
	}
	config.SetGame(g1)
	config.SetGame(g2)
	config.SetGame(g3)
	config.SetGame(g4)
	config.SetGame(g5)
	config.SetGame(g6)
}

//SetShopList 商城配置
func SetShopList() {
	s1 := data.Shop{
		Id:     "111",
		Status: 1,
		Propid: 1,
		Payway: 1,
		Number: 10,
		Price:  1000,
		Name:   "筹码",
		Info:   "筹码",
		Del:    0,
		Etime:  time.Now().AddDate(0, 0, 5),
		Ctime:  time.Now(),
	}
	s2 := data.Shop{
		Id:     "112",
		Status: 1,
		Propid: 1,
		Payway: 1,
		Number: 10,
		Price:  1000,
		Name:   "筹码",
		Info:   "筹码",
		Del:    0,
		Etime:  time.Now().AddDate(0, 0, 5),
		Ctime:  time.Now(),
	}
	s3 := data.Shop{
		Id:     "113",
		Status: 1,
		Propid: 1,
		Payway: 1,
		Number: 10,
		Price:  1000,
		Name:   "筹码",
		Info:   "筹码",
		Del:    0,
		Etime:  time.Now().AddDate(0, 0, 5),
		Ctime:  time.Now(),
	}
	config.SetShop(s1)
	config.SetShop(s2)
	config.SetShop(s3)
}
*/
