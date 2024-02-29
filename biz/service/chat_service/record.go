package chat_service

import (
	"bibi/biz/dal/db"
	"bibi/biz/model/chat"
	"bibi/pkg/errno"
	"time"
)

func (s *MessageService) MessageRecord(req *chat.MessageRecordReq, uid int64) ([]db.Message, int64, error) {
	ft, _ := time.Parse(time.DateOnly, req.FromTime)
	tt, _ := time.Parse(time.DateOnly, req.ToTime)
	t := tt.Add(time.Hour * 24)
	fts := ft.Unix()
	ts := t.Unix()
	//day:86400
	if fts <= 0 || ts <= 0 || fts >= ts {
		return nil, 0, errno.ParamError
	}
	return db.GetRecordMessagesByTime(uid, req.TargetID, ft, t)
}
