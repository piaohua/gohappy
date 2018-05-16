package main

import (
	"gohappy/glog"
	"gohappy/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家桌子常用共有操作请求处理
func (rs *RoleActor) handlerDesk(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.EnteredDesk:
		arg := msg.(*pb.EnteredDesk)
		glog.Debugf("EnteredDesk %#v", arg)
		rs.enterdDesk(arg, ctx)
	case *pb.MatchedDesk:
		arg := msg.(*pb.MatchedDesk)
		glog.Debugf("MatchedDesk %#v", arg)
		rs.matchedDesk(arg, ctx)
	case *pb.CreatedDesk:
		arg := msg.(*pb.CreatedDesk)
		glog.Debugf("CreatedDesk %#v", arg)
		rs.createdDesk(arg, ctx)
	case *pb.LeftDesk:
		arg := msg.(*pb.LeftDesk)
		glog.Debugf("LeftDesk %#v", arg)
		if arg.Error == pb.OK {
			rs.gamePid = nil
		}
	case *pb.SetRecord:
		arg := msg.(*pb.SetRecord)
		glog.Debugf("SetRecord %#v", arg)
		if rs.User != nil {
			rs.status = true
			rs.User.SetRecord(arg.Rtype)
		}
	default:
		//glog.Errorf("unknown message %v", msg)
		rs.handlerNiu(msg, ctx)
	}
}

//加入游戏结果
func (rs *RoleActor) enterdDesk(arg *pb.EnteredDesk, ctx actor.Context) {
	if arg.Error != pb.OK {
		//失败消息
		rs.enterdDeskErr(arg, ctx)
		return
	}
	if arg.Desk == nil {
		//失败消息
		arg.Error = pb.EnterFail
		rs.enterdDeskErr(arg, ctx)
		return
	}
	rs.gamePid = arg.Desk
	//加入成功后获取房间数据
	if rs.enterDeskMsg(arg, ctx) {
		return
	}
}

//进入房间数据,返回房间数据
func (rs *RoleActor) enterDeskMsg(msg *pb.EnteredDesk, ctx actor.Context) bool {
	if rs.gamePid == nil {
		return false
	}
	//区分类型消息
	switch msg.Gtype {
	case int32(pb.NIU):
		switch msg.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			msg2 := new(pb.CNNCoinEnterRoom) //加入消息
			rs.gamePid.Request(msg2, ctx.Self())
		case int32(pb.ROOM_TYPE1): //私人
			msg2 := new(pb.CNNEnterRoom) //加入消息
			rs.gamePid.Request(msg2, ctx.Self())
		case int32(pb.ROOM_TYPE2): //百人
			msg2 := new(pb.CNNFreeEnterRoom) //加入消息
			rs.gamePid.Request(msg2, ctx.Self())
		default:
			glog.Errorf("enterDesk match fail %#v", msg)
		}
	default:
		glog.Errorf("enterDesk match fail %#v", msg)
	}
	return true
}

//查找不同类型房间远程节点
func (rs *RoleActor) selectDesk(msg *pb.MatchDesk, ctx actor.Context) {
	msg.Sender = ctx.Self()
	switch msg.Gtype {
	case int32(pb.NIU):
		switch msg.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			//存在房间id直接查找
			if msg.Roomid != "" {
				rs.roomPid.Request(msg, ctx.Self())
			} else {
				//或者不存在时去节点中匹配
				rs.dbmsPid.Request(msg, ctx.Self())
			}
		case int32(pb.ROOM_TYPE1): //私人
			rs.roomPid.Request(msg, ctx.Self())
		case int32(pb.ROOM_TYPE2): //百人
			rs.dbmsPid.Request(msg, ctx.Self())
		default:
			glog.Errorf("selectDesk match fail %#v", msg)
		}
	default:
		glog.Errorf("selectDesk match fail %#v", msg)
	}
}

//配置远程节点结果,然后加入游戏
func (rs *RoleActor) matchedDesk(arg *pb.MatchedDesk, ctx actor.Context) {
	if arg.Error != pb.OK {
		//失败消息
		rs.matchedDeskErr(arg, ctx)
		return
	}
	if arg.Desk == nil {
		//失败消息
		arg.Error = pb.MatchFail
		rs.matchedDeskErr(arg, ctx)
		return
	}
	msg := new(pb.EnterDesk)
	msg.Gtype = arg.Gtype
	msg.Rtype = arg.Rtype
	msg.Dtype = arg.Dtype
	msg.Ltype = arg.Ltype
	if !rs.enterDesk(msg, ctx) {
		//失败消息
		rs.matchedDeskErr(arg, ctx)
		return
	}
	//请求消息
	arg.Desk.Request(msg, ctx.Self())
}

//加入房间消息
func (rs *RoleActor) enterDesk(msg *pb.EnterDesk, ctx actor.Context) bool {
	result4, err4 := json.Marshal(rs.User)
	if err4 != nil {
		glog.Errorf("user Marshal err %v", err4)
		return false
	}
	//玩家数据
	msg.Sender = ctx.Self()
	msg.Data = result4
	return true
}

//响应加入游戏失败消息
func (rs *RoleActor) enterdDeskErr(msg *pb.EnteredDesk, ctx actor.Context) {
	switch msg.Gtype {
	case int32(pb.NIU):
		switch msg.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			msg2 := new(pb.SNNCoinEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		case int32(pb.ROOM_TYPE1): //私人
			msg2 := new(pb.SNNEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		case int32(pb.ROOM_TYPE2): //百人
			msg2 := new(pb.SNNFreeEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		default:
			glog.Errorf("enter Desk match fail %#v", msg)
		}
	default:
		glog.Errorf("enter Desk match fail %#v", msg)
	}
}

//响应加入游戏失败消息
func (rs *RoleActor) matchedDeskErr(msg *pb.MatchedDesk, ctx actor.Context) {
	switch msg.Gtype {
	case int32(pb.NIU):
		switch msg.Rtype {
		case int32(pb.ROOM_TYPE0): //自由
			msg2 := new(pb.SNNCoinEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		case int32(pb.ROOM_TYPE1): //私人
			msg2 := new(pb.SNNEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		case int32(pb.ROOM_TYPE2): //百人
			msg2 := new(pb.SNNFreeEnterRoom) //加入消息
			msg2.Error = msg.Error
			rs.Send(msg2)
		default:
			glog.Errorf("matched DeskErr fail %#v", msg)
		}
	default:
		glog.Errorf("matched DeskErr fail %#v", msg)
	}
}

//加入游戏结果
func (rs *RoleActor) createdDesk(arg *pb.CreatedDesk, ctx actor.Context) {
	if arg.Error != pb.OK {
		rsp := new(pb.SNNCreateRoom)
		rsp.Error = arg.Error
		rs.Send(rsp)
		return
	}
	//TODO 扣除
	msg := new(pb.EnterDesk)
	msg.Gtype = arg.Gtype
	msg.Rtype = arg.Rtype
	if !rs.enterDesk(msg, ctx) || arg.Desk == nil {
		//失败消息
		rsp := new(pb.SNNCreateRoom)
		rsp.Error = pb.CreateRoomFail
		rs.Send(rsp)
		return
	}
	arg.Desk.Request(msg, ctx.Self())
}
