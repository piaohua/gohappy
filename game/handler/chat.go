package handler

import "gohappy/pb"

//ChatTextMsg 文本聊天消息
func ChatTextMsg(seat uint32, userid string, msg string) *pb.SChatText {
	return &pb.SChatText{
		Seat:    seat,
		Userid:  userid,
		Content: msg,
	}
}

//ChatVoiceMsg 语音聊天消息
func ChatVoiceMsg(seat uint32, userid string, msg string) *pb.SChatVoice {
	return &pb.SChatVoice{
		Seat:    seat,
		Userid:  userid,
		Content: msg,
	}
}
