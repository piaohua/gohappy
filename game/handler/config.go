package handler

import (
	"fmt"

	"gohappy/data"
	"gohappy/game/config"
	"gohappy/glog"
	"gohappy/pb"

	jsoniter "github.com/json-iterator/go"
)

//单条配置修改同步都按map格式

//SyncConfig 同步配置
func SyncConfig(arg *pb.SyncConfig) (err error) {
	switch arg.Type {
	case pb.CONFIG_ENV: //变量
		b := make(map[string]int32)
		err = jsoniter.Unmarshal(arg.Data, &b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		for k, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
				config.DelEnv(k)
			case pb.CONFIG_UPSERT:
				config.SetEnv(k, v)
			}
		}
	case pb.CONFIG_NOTICE: //公告
		b := make(map[string]data.Notice)
		err = jsoniter.Unmarshal(arg.Data, &b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		for k, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
				config.DelNotice(k)
			case pb.CONFIG_UPSERT:
				config.SetNotice(v)
			}
		}
	case pb.CONFIG_SHOP: //商城
		b := make(map[string]data.Shop)
		err = jsoniter.Unmarshal(arg.Data, &b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		for k, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
				config.DelShop(k)
			case pb.CONFIG_UPSERT:
				config.SetShop(v)
			}
		}
	case pb.CONFIG_GAMES: //游戏
		b := make(map[string]data.Game)
		err = jsoniter.Unmarshal(arg.Data, &b)
		glog.Debugf("Sync Games %#v", b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		for k, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
				config.DelGame(k)
			case pb.CONFIG_UPSERT:
				config.SetGame(v)
			}
		}
	case pb.CONFIG_VIP: //游戏
		b := make(map[string]data.Vip)
		err = jsoniter.Unmarshal(arg.Data, &b)
		if err != nil {
			glog.Errorf("syncConfig Unmarshal err %v, data %#v", err, arg.Data)
			return
		}
		for k, v := range b {
			switch arg.Atype {
			case pb.CONFIG_DELETE:
				config.DelVip(k)
			case pb.CONFIG_UPSERT:
				config.SetVip(v)
			}
		}
	default:
		glog.Errorf("syncConfig unknown type %s", arg.Type)
		err = fmt.Errorf("type not exist %d", arg.Type)
	}
	return
}

//打包配置
func syncConfigMsg(d interface{}) ([]byte, error) {
	result, err := jsoniter.Marshal(d)
	if err != nil {
		glog.Errorf("syncConfig Marshal err %v", err)
	}
	return result, err
}

//SyncConfig2 打包消息
func SyncConfig2(ctype pb.ConfigType, atype pb.ConfigAtype,
	data []byte) (msg *pb.SyncConfig) {
	msg = new(pb.SyncConfig)
	msg.Type = ctype
	msg.Atype = atype
	msg.Data = data
	return
}

//GetSyncConfig2 同步配置
func GetSyncConfig2(ctype pb.ConfigType) (msg *pb.SyncConfig, err error) {
	msg = new(pb.SyncConfig)
	msg.Type = ctype
	msg.Atype = pb.CONFIG_UPSERT
	switch ctype {
	case pb.CONFIG_ENV: //变量
		msg.Data, err = syncConfigMsg(config.GetEnvs())
	case pb.CONFIG_NOTICE: //公告
		msg.Data, err = syncConfigMsg(config.GetNotices2())
	case pb.CONFIG_SHOP: //商城
		msg.Data, err = syncConfigMsg(config.GetShops2())
	case pb.CONFIG_GAMES: //游戏列表
		msg.Data, err = syncConfigMsg(config.GetGames2())
	case pb.CONFIG_VIP: //vip列表
		msg.Data, err = syncConfigMsg(config.GetVips2())
	default:
		err = fmt.Errorf("type not exist %d", ctype)
	}
	return
}

//GetSyncConfig 获取
func GetSyncConfig(ctype pb.ConfigType) (msg *pb.SyncConfig) {
	msg, err := GetSyncConfig2(ctype)
	if err != nil {
		glog.Errorf("err %s", err)
	}
	return
}
