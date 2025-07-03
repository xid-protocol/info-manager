package jumpserver

import (
	"context"

	"github.com/colin-404/logx"
	"github.com/tidwall/gjson"
	"github.com/xid-protocol/xidp/common"
	"github.com/xid-protocol/xidp/db/models"
	"github.com/xid-protocol/xidp/db/repositories"
)

// type jumpserver struct {
// 	accessKey string
// 	secret    string
// }

func setJumpserverInfo(userInfo gjson.Result) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	if email := userInfo.Get("email").String(); email != "" {
		xid := common.GenerateXid(email)
		logx.Infof("xid: %v", xid)
		//检查是否存在
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/jumpserver")
		if err != nil {
			logx.Errorf("check jumpserver %s failed: %v", email, err)
			return
		}
		if exists {
			xidRecord := models.XID{
				Name:    "xid-protocol",
				Xid:     xid,
				Payload: userInfo.Value(),
				Version: "0.1.1",
				Metadata: models.Metadata{
					CardId:      common.GenerateCardId(),
					CreatedAt:   common.GetTimestamp(),
					Operation:   "modify",
					Path:        "/info/jumpserver",
					ContentType: "application/json",
				},
			}
			logx.Infof("xid: %v", xidRecord)
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, &xidRecord)
			if err != nil {
				logx.Errorf("modify jumpserver %s failed: %v", email, err)
			} else {
				logx.Infof("modify jumpserver %s success", email)
			}
			return
		}

		xidRecord := models.XID{
			Name:    "xid-protocol",
			Xid:     xid,
			Payload: userInfo.Value(),
			Version: "0.1.1",
			Metadata: models.Metadata{
				CardId:      common.GenerateCardId(),
				CreatedAt:   common.GetTimestamp(),
				Operation:   "create",
				Path:        "/info/jumpserver",
				ContentType: "application/json",
			},
		}
		err = repo.Insert(ctx, &xidRecord)
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
