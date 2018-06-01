package handler

import (
	"fmt"
	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"
	"utils"
)

////GetNotice 公告列表
//func GetNotice(atype int32) (stoc *pb.SNotice) {
//	stoc = new(pb.SNotice)
//	list := config.GetNotices(atype)
//	for _, v := range list {
//		body := &pb.Notice{
//			Rtype:   int32(v.Rtype),
//			Acttype: int32(v.Acttype),
//			Content: v.Content,
//		}
//		stoc.List = append(stoc.List, body)
//	}
//	return
//}

//PackNotice 打包公告消息
func PackNotice(msg *pb.SNotice) {
	list := config.GetNotices(data.NOTICE_TYPE1)
	for _, v := range list {
		body := &pb.Notice{
			Rtype:   int32(v.Rtype),
			Acttype: int32(v.Acttype),
			Content: v.Content,
		}
		msg.List = append(msg.List, body)
	}
}

//PackUserNotice 打包玩家消息
func PackUserNotice(arg *pb.CNotice) (msg *pb.SNotice) {
	msg = new(pb.SNotice)
	list, err := data.GetLogNotices(arg.Userid, int(arg.Page))
	if err != nil {
		glog.Errorf("get notice err : %v, arg %#v", err, arg)
		return
	}
	for _, v := range list {
		body := &pb.Notice{
			Rtype:   int32(v.Rtype),
			Acttype: int32(v.Acttype),
			Content: v.Content,
			Time:    utils.Time2Str(v.Ctime),
		}
		msg.List = append(msg.List, body)
	}
	return
}

//SaveNotice 保存消息记录
func SaveNotice(arg *pb.LogNotice) {
	record := &data.Notice{
		Userid:  arg.Userid,
		Rtype:   int(arg.Rtype),
		Acttype: int(arg.Acttype),
		Content: arg.Content,
	}
	record.Save()
}

//NewNotice 新消息
func NewNotice(rtype, atype int32, userid,
	content string) (record *pb.LogNotice, msg *pb.SPushNotice) {
	record = &pb.LogNotice{
		Userid:  userid,
		Rtype:   rtype,
		Acttype: atype,
		Content: content,
	}
	msg = new(pb.SPushNotice)
	msg.Info = &pb.Notice{
		Time:    utils.Time2Str(utils.LocalTime()),
		Rtype:   rtype,
		Acttype: atype,
		Content: content,
	}
	return
}

//BuyNotice 兑换金币消息
func BuyNotice(coin int64, userid string) (record *pb.LogNotice,
	msg *pb.SPushNotice) {
	if coin <= 0 {
		return
	}
	content := fmt.Sprintf("恭喜你成功兑换%d金币", coin)
	return NewNotice(0, 0, userid, content)
}
