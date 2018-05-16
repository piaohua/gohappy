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
		Unique: d.Id,
		Gtype:  d.Gtype,
		Rtype:  d.Rtype,
		Dtype:  d.Dtype,
		Ltype:  d.Ltype,
		Rname:  d.Name,
		Count:  d.Count,
		Ante:   d.Ante,
		Cost:   d.Cost,
		Vip:    d.Vip,
		Chip:   d.Chip,
		Deal:   d.Deal,
		Carry:  d.Carry,
		Down:   d.Down,
		Top:    d.Top,
		Sit:    d.Sit,
	}
}

//NewFreeGameData 百人房间桌子数据
func NewFreeGameData(node string) *data.Game {
	return &data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  int32(pb.NIU),
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
func NewCoinGameData(node string, ltype int32) *data.Game {
	g := &data.Game{
		Id:     bson.NewObjectId().Hex(),
		Gtype:  int32(pb.NIU),
		Rtype:  int32(pb.ROOM_TYPE0),
		Dtype:  int32(pb.DESK_TYPE1),
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
	case int32(pb.ROOM_LEVEL1):
		g.Ante = 2
		g.Chip = 5000
		g.Sit = 5000
	case int32(pb.ROOM_LEVEL2):
		g.Ante = 5
		g.Chip = 20000
		g.Sit = 20000
	case int32(pb.ROOM_LEVEL3):
		g.Ante = 10
		g.Chip = 200000
		g.Sit = 200000
	case int32(pb.ROOM_LEVEL4):
		g.Ante = 20
		g.Chip = 2000000
		g.Sit = 2000000
	}
	return g
}

//NewPrivGameData 私人房间桌子数据
func NewPrivGameData(arg *pb.CreateDesk) *data.DeskData {
	return &data.DeskData{
		Unique:  bson.NewObjectId().Hex(),
		Gtype:   arg.Gtype,
		Rtype:   arg.Rtype,
		Rname:   arg.Rname,
		Ante:    arg.Ante,
		Count:   arg.Count,
		Round:   arg.Round,
		Payment: arg.Payment,
		Cid:     arg.Cid,
		Expire:  utils.Timestamp() + 600,
		Deal:    true,
	}
}
