package data

import (
	"os"

	"gohappy/glog"
	"gohappy/pb"
	"utils"

	jsoniter "github.com/json-iterator/go"
)

//{"_id":"10076","nickname":"Terry","photo":"http://wx.qlogo.cn/mmhead/weWZwghR1ymaX1pscPZEgOywfofhWlhugvjvHQgSFRw/0"}
type RegistRobotInfo struct {
	ID       string `json:"_id"`
	Nickname string `json:"nickname"`
	Photo    string `json:"photo"`
}

func RegistRobots(phone, passwd string, genid *IDGen) {
	var minDiamond int64 = 1000
	var minCoin int64 = 10000000
	var minChip int64 = 10000000
	var minCard int64 = 100

	var jsonPath string = "./robot.json"

	list := make([]RegistRobotInfo, 0)
	f, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := jsoniter.NewDecoder(f)
	err = decoder.Decode(&list)
	if err != nil {
		panic(err)
	}
	glog.Infof("regist robots -> %d", len(list))
	for _, r := range list {
		user := new(User)
		user.Phone = phone
		user.GetByPhone()
		if user.Userid == "" {
			user.Userid = genid.GenID()
			user.Nickname = r.Nickname
			//user.Sex = r.Sex
			user.Photo = r.Photo
			user.Robot = true
			user.Diamond = utils.RandInt64N(minDiamond)
			user.Coin = utils.RandInt64N(minCoin)
			user.Chip = int64(utils.RandInt64N(minChip))
			user.Card = utils.RandInt64N(minCard)
			//user.Vip = 1
			passwd := utils.Md5(passwd)
			auth := string(utils.GetAuth())
			user.Auth = auth
			user.Password = utils.Md5(passwd + auth)
			user.Ctime = utils.LocalTime()
			user.Save()
			glog.Infof("regist robot userid -> %s, phone -> %s", user.Userid, user.Phone)
		} else {
			var update bool
			if user.Diamond < minDiamond {
				user.AddDiamond(minDiamond)
				update = true
			}
			if user.Coin < minCoin {
				user.AddCoin(minCoin)
				update = true
			}
			if user.Chip < int64(minChip) {
				user.AddChip(int64(minChip))
				update = true
			}
			if user.Card < minCard {
				user.AddCard(minCard)
				update = true
			}
			if update {
				user.Save()
				glog.Infof("update robot userid -> %s, phone -> %s", user.Userid, user.Phone)
			}
		}
		phone = utils.StringAdd(phone)
	}
}

//默认头像
type RegistPhoto struct {
	Photo string `json:"photo"`
}

func RegistPhotos() []RegistPhoto {

	var jsonPath string = "./photo.json"

	list := make([]RegistPhoto, 0)
	f, err := os.Open(jsonPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := jsoniter.NewDecoder(f)
	err = decoder.Decode(&list)
	if err != nil {
		panic(err)
	}
	glog.Infof("regist headimag -> %d", len(list))
	return list
}

func RegistRobots2(head, phone, passwd string, genid *IDGen) {
	HeadImagList := RegistPhotos()
	if head == "" {
		glog.Errorf("regist robot head err: %s", head)
		return
	}
	if len(phone) != 11 {
		glog.Errorf("regist robot phone err: %s", phone)
		return
	}
	var minDiamond int64 = 1000
	var minCoin int64 = 80000
	var minChip int64 = 80000
	var minCard int64 = 100
	//注册
	for j := 0; j < 100; j++ {
		user := new(User)
		user.Phone = phone
		user.GetByPhone()
		if user.Userid == "" {
			user.Userid = genid.GenID()
			user.Nickname = "用户" + phone[len(phone)-4:]
			i := utils.RandIntN(len(HeadImagList))
			user.Photo = head + "/" + HeadImagList[i].Photo
			user.Robot = true
			user.Diamond = utils.RandInt64N(minDiamond)
			user.Coin = utils.RandInt64N(minCoin)
			user.Chip = int64(utils.RandInt64N(minChip))
			user.Card = utils.RandInt64N(minCard)
			//user.Vip = 1
			passwd := utils.Md5(passwd)
			auth := string(utils.GetAuth())
			user.Auth = auth
			user.Password = utils.Md5(passwd + auth)
			user.Save()
			glog.Infof("regist robot userid -> %s, phone -> %s", user.Userid, user.Phone)
		} else {
			var update bool
			if user.Diamond < minDiamond {
				user.AddDiamond(minDiamond)
				update = true
			}
			if user.Coin < minCoin {
				user.AddCoin(minCoin)
				update = true
			}
			if user.Chip < int64(minChip) {
				user.AddChip(int64(minChip))
				update = true
			}
			if user.Card < minCard {
				user.AddCard(minCard)
				update = true
			}
			if update {
				user.Save()
				glog.Infof("update robot userid -> %s, phone -> %s", user.Userid, user.Phone)
			}
		}
		phone = utils.StringAdd(phone)
	}
}

func RegistRobots3(head, phone, passwd string, genid *IDGen) {
	HeadImagList := RegistPhotos()
	if head == "" {
		glog.Errorf("regist robot head err: %s", head)
		return
	}
	if len(phone) != 11 {
		glog.Errorf("regist robot phone err: %s", phone)
		return
	}
	var minDiamond int64 = 1000
	var minCoin int64 = 80000
	var minChip int64 = 80000
	var minCard int64 = 100
	//注册
	list := GetRobotList()
	for _, v := range list {
		user := new(User)
		user.Phone = phone
		user.GetByPhone()
		if user.Userid == "" {
			user.Userid = genid.GenID()
			user.Nickname = v.Nickname
			user.Sex = v.Sex
			i := utils.RandIntN(len(HeadImagList))
			user.Photo = head + "/" + HeadImagList[i].Photo
			user.Robot = true
			user.Diamond = utils.RandInt64N(minDiamond)
			user.Coin = utils.RandInt64N(minCoin)
			user.Chip = int64(utils.RandInt64N(minChip))
			user.Card = utils.RandInt64N(minCard)
			//user.Vip = 1
			passwd := utils.Md5(passwd)
			auth := string(utils.GetAuth())
			user.Auth = auth
			user.Password = utils.Md5(passwd + auth)
			user.Save()
			glog.Infof("regist robot userid -> %s, phone -> %s", user.Userid, user.Phone)
		} else {
			var update bool
			if user.Diamond < minDiamond {
				user.AddDiamond(minDiamond)
				update = true
			}
			if user.Coin < minCoin {
				user.AddCoin(minCoin)
				update = true
			}
			if user.Chip < int64(minChip) {
				user.AddChip(int64(minChip))
				update = true
			}
			if user.Card < minCard {
				user.AddCard(minCard)
				update = true
			}
			if update {
				user.Save()
				glog.Infof("update robot userid -> %s, phone -> %s", user.Userid, user.Phone)
			}
		}
		phone = utils.StringAdd(phone)
	}
}

//RegistRobots4 robot regist
func RegistRobots4(photo, passwd string) (ls []*pb.RobotRegist) {
	rs := make([]RobotInfo, 0)
	err := LoadRobotInfo("./robot.json", &rs)
	if err != nil {
		panic(err)
	}
	for _, v := range rs {
		r := &pb.RobotRegist{
			ID: v.ID,
			Nickname: v.Nickname,
			Sex: v.Sex,
			Coin: v.Coin,
			Diamond: v.Diamond,
			Vip: v.Vip,
			Phone: v.Phone,
			Photo: photo,
			Password: passwd,
		}
		//HeadImagList := RegistPhotos()
		//i := utils.RandIntN(len(HeadImagList))
		//r.Photo = photo + "/" + HeadImagList[i].Photo
		ls = append(ls, r)
	}
	return
}