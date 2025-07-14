package sealsuite

import (
	"context"

	"github.com/colin-404/logx"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"github.com/xid-protocol/info-manager/common"
	"github.com/xid-protocol/info-manager/db/repositories"
	"github.com/xid-protocol/xidp/protocols"
)

type sealsuite struct {
	endpoint string
	token    string
}

func setSealsuiteInfo(usersByEmail map[string]gjson.Result) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	for email := range usersByEmail {

		// info := map[string]interface{}{"email": email, "type": "user_email"}
		xid := protocols.GenerateXid(email)
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/sealsuite")
		if err != nil {
			logx.Errorf("check sealsuite failed: %v", err)
			continue
		}

		if exists {
			logx.Infof("sealsuite %s exists, need to modify", email)
			info := protocols.NewInfo(email, "email", false)
			metadata := protocols.NewMetadata("modify", "/info/sealsuite", "application/json")
			xidRecord := protocols.NewXID(info, metadata, usersByEmail[email].Value())

			logx.Infof("xid: %v", xidRecord)
			// 插入MongoDB
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, xidRecord)
			if err != nil {
				logx.Errorf("modify sealsuite %s failed: %v", email, err)
			} else {
				logx.Infof("modify sealsuite %s success", email)
			}
			continue
		}

		info := protocols.NewInfo(email, "email", false)
		metadata := protocols.NewMetadata("create", "/info/sealsuite", "application/json")
		xidRecord := protocols.NewXID(info, metadata, usersByEmail[email].Value())

		logx.Infof("xid: %v", xidRecord)
		// 插入MongoDB
		err = repo.Insert(ctx, xidRecord)
		if err != nil {
			logx.Errorf("insert sealsuite %s failed: %v", email, err)
		} else {
			logx.Infof("insert sealsuite %s success", email)
		}
	}
}

func RunSealsuite() {
	users := SealsuiteAccount()
	// logx.Infof("users: %v", users)
	//获取所有所有邮箱并去重，struct{}{}

	usersByEmail := make(map[string]gjson.Result)
	for _, user := range *users {
		if user.Get("email").Exists() {
			email := user.Get("email").String()
			if email != "" {
				// 如果email已存在，会自动覆盖（去重）
				usersByEmail[email] = user

			}
		}
	}
	// emailXidInit(usersByEmail)
	setSealsuiteInfo(usersByEmail)
}

func SealsuiteAccount() *[]gjson.Result {
	endpoint := viper.GetString("sealsuite.endpoint")
	accessKeyId := viper.GetString("sealsuite.access_key_id")
	accessKeySecret := viper.GetString("sealsuite.access_key_secret")

	ss := &sealsuite{
		endpoint: endpoint,
		token:    getToken(endpoint, accessKeyId, accessKeySecret),
	}
	logx.Infof("ss: %v", ss)
	departments := ss.getDeparment()
	return ss.getAllUsersForDepartments(departments)

}

func getToken(endpoint string, accessKeyId string, accessKeySecret string) string {
	url := endpoint + "/api/open/v1/token"
	body := map[string]string{
		"access_key_id":     accessKeyId,
		"access_key_secret": accessKeySecret,
	}
	resp, err := common.DoHttp("POST", url, body, nil)
	if err != nil {
		logx.Errorf("get token failed: %v", err)
		return ""
	}
	token := gjson.Get(resp.String(), "data.access_token").String()

	return token

}
