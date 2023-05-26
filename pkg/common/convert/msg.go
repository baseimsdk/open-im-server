package convert

import (
	"encoding/json"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/db/table/unrelation"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/proto/sdkws"
)

func MsgPb2DB(msg *sdkws.MsgData) unrelation.MsgDataModel {
	var msgDataModel unrelation.MsgDataModel
	msgDataModel.SendID = msg.SendID
	msgDataModel.RecvID = msg.RecvID
	msgDataModel.GroupID = msg.GroupID
	msgDataModel.ClientMsgID = msg.ClientMsgID
	msgDataModel.ServerMsgID = msg.ServerMsgID
	msgDataModel.SenderPlatformID = msg.SenderPlatformID
	msgDataModel.SenderNickname = msg.SenderNickname
	msgDataModel.SenderFaceURL = msg.SenderFaceURL
	msgDataModel.SessionType = msg.SessionType
	msgDataModel.MsgFrom = msg.MsgFrom
	msgDataModel.ContentType = msg.ContentType
	msgDataModel.Content = string(msg.Content)
	msgDataModel.Seq = msg.Seq
	msgDataModel.SendTime = msg.SendTime
	msgDataModel.CreateTime = msg.CreateTime
	msgDataModel.Status = msg.Status
	msgDataModel.Options = msg.Options
	msgDataModel.OfflinePush = &unrelation.OfflinePushModel{
		Title:         msg.OfflinePushInfo.Title,
		Desc:          msg.OfflinePushInfo.Desc,
		Ex:            msg.OfflinePushInfo.Ex,
		IOSPushSound:  msg.OfflinePushInfo.IOSPushSound,
		IOSBadgeCount: msg.OfflinePushInfo.IOSBadgeCount,
	}
	msgDataModel.AtUserIDList = msg.AtUserIDList
	msgDataModel.AttachedInfo = msg.AttachedInfo
	msgDataModel.Ex = msg.Ex
	return msgDataModel

}

func MsgDB2Pb(msgModel *unrelation.MsgDataModel) *sdkws.MsgData {
	var msg sdkws.MsgData
	msg.SendID = msgModel.SendID
	msg.RecvID = msgModel.RecvID
	msg.GroupID = msgModel.GroupID
	msg.ClientMsgID = msgModel.ClientMsgID
	msg.ServerMsgID = msgModel.ServerMsgID
	msg.SenderPlatformID = msgModel.SenderPlatformID
	msg.SenderNickname = msgModel.SenderNickname
	msg.SenderFaceURL = msgModel.SenderFaceURL
	msg.SessionType = msgModel.SessionType
	msg.MsgFrom = msgModel.MsgFrom
	msg.ContentType = msgModel.ContentType
	b, _ := json.Marshal(msgModel.Content)
	msg.Content = b
	msg.Seq = msgModel.Seq
	msg.SendTime = msgModel.SendTime
	msg.CreateTime = msgModel.CreateTime
	msg.Status = msgModel.Status
	msg.Options = msgModel.Options
	msg.OfflinePushInfo = &sdkws.OfflinePushInfo{
		Title:         msgModel.OfflinePush.Title,
		Desc:          msgModel.OfflinePush.Desc,
		Ex:            msgModel.OfflinePush.Ex,
		IOSPushSound:  msgModel.OfflinePush.IOSPushSound,
		IOSBadgeCount: msgModel.OfflinePush.IOSBadgeCount,
	}
	msg.AtUserIDList = msgModel.AtUserIDList
	msg.AttachedInfo = msgModel.AttachedInfo
	msg.Ex = msgModel.Ex
	return &msg
}
