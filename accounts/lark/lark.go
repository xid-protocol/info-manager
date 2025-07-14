package lark

import (
	"context"

	"github.com/colin-404/logx"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/xid-protocol/info-manager/common"
	"github.com/xid-protocol/info-manager/db/repositories"
	"github.com/xid-protocol/xidp/protocols"
)

func getToken() string {
	tokenUrl := "https://open.larksuite.com/open-apis/auth/v3/tenant_access_token/internal"
	body := map[string]string{
		"app_id":     viper.GetString("lark.app_id"),
		"app_secret": viper.GetString("lark.app_secret"),
	}
	tokenResp, err := common.DoHttp("POST", tokenUrl, body, nil)
	if err != nil {
		logx.Errorf("lark token error: %v", err)
	}
	token := gjson.Get(tokenResp.String(), "tenant_access_token").String()
	return token
}

func RunLark() {
	token := getToken()
	users, err := getLarkUsers(token)
	if err != nil {
		logx.Errorf("lark users error: %v", err)
	}
	// logx.Infof("users: %v", users)
	setLarkInfo(users)

}

func setLarkInfo(usersByEmail *map[string]gjson.Result) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	for email := range *usersByEmail {

		// info := map[string]interface{}{"email": email, "type": "user_email"}
		xid := protocols.GenerateXid(email)
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/larksuite")
		if err != nil {
			logx.Errorf("check larksuite failed: %v", err)
			continue
		}

		if exists {
			logx.Infof("larksuite %s exists, need to modify", email)
			info := protocols.NewInfo(email, "email")
			metadata := protocols.NewMetadata("modify", "/info/larksuite", "application/json")
			xidRecord := protocols.NewXID(&info, &metadata, (*usersByEmail)[email].Value())
			logx.Infof("xid: %v", xidRecord)
			// 插入MongoDB
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, xidRecord)
			if err != nil {
				logx.Errorf("modify larksuite %s failed: %v", email, err)
			} else {
				logx.Infof("modify larksuite %s success", email)
			}
			continue
		}

		info := protocols.NewInfo(email, "email")
		metadata := protocols.NewMetadata("create", "/info/larksuite", "application/json")
		xidRecord := protocols.NewXID(&info, &metadata, (*usersByEmail)[email].Value())
		logx.Infof("xid: %v", xidRecord)
		// 插入MongoDB
		err = repo.Insert(ctx, xidRecord)
		if err != nil {
			logx.Errorf("insert larksuite %s failed: %v", email, err)
		} else {
			logx.Infof("insert larksuite %s success", email)
		}
	}
}
