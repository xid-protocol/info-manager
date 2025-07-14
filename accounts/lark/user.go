package lark

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/colin-404/logx"
	"github.com/tidwall/gjson"
	"github.com/xid-protocol/info-manager/common"
)

const LarkURL = "https://open.larksuite.com"

func buildPaginatedURL(baseURL string, pageSize int, pageToken string) string {
	u, _ := url.Parse(baseURL)
	q := u.Query()

	q.Add("fetch_child", "true")
	q.Add("page_size", strconv.Itoa(pageSize))

	if pageToken != "" {
		q.Add("page_token", pageToken)
	}

	u.RawQuery = q.Encode()
	return u.String()
}
func getLarkDepartments(token string) (map[string]struct{}, error) {

	pageToken := ""
	pageSize := 50
	depts := make(map[string]struct{})
	for {
		reqURL := buildPaginatedURL(
			LarkURL+"/open-apis/contact/v3/departments/0/children",
			pageSize,
			pageToken,
		)

		headers := map[string]string{"Authorization": "Bearer " + token}
		resp, err := common.DoHttp("GET", reqURL, nil, headers)
		if err != nil {
			logx.Errorf("获取部门信息失败: %v", err)
			return nil, err
		}
		logx.Infof("resp: %v", resp.String())
		data := gjson.Get(resp.String(), "data")
		hasMore := data.Get("has_more").Bool()
		items := data.Get("items").Array()
		//set

		for _, item := range items {
			deptID := item.Get("open_department_id").String()
			depts[deptID] = struct{}{}
		}

		// 解析部门信息
		// for _, item := range items {
		// 	dept := item.(map[string]interface{})
		// 	logx.Infof("dept: %v", dept)
		// }

		// 检查是否还有下一页
		if !hasMore {
			break
		}

		// 获取下一页的token
		if nextToken := data.Get("page_token").String(); nextToken != "" {
			pageToken = nextToken
		} else {
			break
		}

		logx.Infof("已获取 %d 个部门，继续获取下一页...", len(items))
	}

	logx.Infof("总共获取到 %d 个部门", len(depts))
	logx.Infof("depts: %v", depts)
	return depts, nil
}

func getLarkUsers(token string) (*map[string]gjson.Result, error) {
	depts, err := getLarkDepartments(token)
	if err != nil {
		return nil, err
	}
	users := make(map[string]gjson.Result)
	for deptID := range depts {
		userURL := fmt.Sprintf("%s/open-apis/contact/v3/users/find_by_department/?department_id=%s&page_size=10&department_id_type=open_department_id", LarkURL, deptID)
		headers := map[string]string{"Authorization": "Bearer " + token}
		resp, err := common.DoHttp("GET", userURL, nil, headers)
		if err != nil {
			return nil, err
		}

		data := gjson.Get(resp.String(), "data")
		if !data.Exists() {
			continue
		}

		//提取用户列表
		items := data.Get("items").Array()
		for _, item := range items {
			email := item.Get("email").String()
			if email != "" {
				users[email] = item
			}
		}
	}
	return &users, nil
}
