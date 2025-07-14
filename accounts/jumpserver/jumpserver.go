package jumpserver

import (
	"context"

	"github.com/colin-404/logx"
	"github.com/tidwall/gjson"
	"github.com/xid-protocol/info-manager/db/repositories"
	"github.com/xid-protocol/xidp/protocols"
)

// type jumpserver struct {
// 	accessKey string
// 	secret    string
// }

func setJumpserverInfo(userInfo gjson.Result) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	if email := userInfo.Get("email").String(); email != "" {
		xid := protocols.GenerateXid(email)
		logx.Infof("xid: %v", xid)
		//检查是否存在
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/jumpserver")
		if err != nil {
			logx.Errorf("check jumpserver %s failed: %v", email, err)
			return
		}
		if exists {
			info := protocols.NewInfo(email, "email", false)
			metadata := protocols.NewMetadata("modify", "/info/jumpserver", "application/json")
			xidRecord := protocols.NewXID(info, metadata, userInfo.Value())
			logx.Infof("xid: %v", xidRecord)
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, xidRecord)
			if err != nil {
				logx.Errorf("modify jumpserver %s failed: %v", email, err)
			} else {
				logx.Infof("modify jumpserver %s success", email)
			}
			return
		}

		info := protocols.NewInfo(email, "email", false)
		metadata := protocols.NewMetadata("create", "/info/jumpserver", "application/json")
		xidRecord := protocols.NewXID(info, metadata, userInfo.Value())
		err = repo.Insert(ctx, xidRecord)
		if err != nil {
			logx.Errorf("insert jumpserver %s failed: %v", email, err)
		} else {
			logx.Infof("insert jumpserver %s success", email)
		}
	}
}

func RunJumpServer() {

	resp := getUserInfo()
	results := gjson.Parse(resp.String())

	for _, userInfo := range results.Array() {
		setJumpserverInfo(userInfo)
	}

}
