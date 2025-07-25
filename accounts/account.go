package accounts

import (
	"github.com/xid-protocol/info-manager/accounts/lark"
	"github.com/xid-protocol/info-manager/accounts/sealsuite"
)

func AccountMonitor() {

	sealsuite.RunSealsuite()

	// jumpserver.RunJumpServer()
	lark.RunLark()

}

// func emailXidInit(usersByEmail map[string]gjson.Result) {
// 	repo := repositories.NewXIDRepository()
// 	ctx := context.Background()
// 	for email := range usersByEmail {
// 		// 检查email是否已存在
// 		exists, err := repo.CheckEmailExists(ctx, email)
// 		if err != nil {
// 			logx.Errorf("检查email失败: %v", err)
// 			continue
// 		}

// 		if exists {
// 			logx.Infof("Email %s 已存在，跳过插入", email)
// 			continue
// 		}

// 		userInfo := map[string]interface{}{"email": email, "type": "user_email"}

// 		xidRecord := models.XID{
// 			Name:    "xid-protocol",
// 			Xid:     common.GenerateXid(email),
// 			Version: "0.1.0",
// 			Payload: userInfo,
// 			Metadata: models.Metadata{
// 				CreatedAt:   common.GetTimestamp(),
// 				Operation:   "create",
// 				Path:        "/info",
// 				ContentType: "application/json",
// 			},
// 		}
// 		logx.Infof("xid: %v", xidRecord)
// 		// 插入MongoDB
// 		err = repo.Insert(ctx, &xidRecord)
// 		if err != nil {
// 			logx.Errorf("插入用户 %s 失败: %v", email, err)
// 		} else {
// 			logx.Infof("成功插入用户: %s", email)
// 		}
// 	}
// }
