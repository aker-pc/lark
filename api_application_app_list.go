// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetApplicationAppList 该接口用于查询企业安装的应用列表, 只能被企业自建应用调用。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYDN3UjL2QzN14iN0cTN
// new doc: https://open.feishu.cn/document/server-docs/application-v6/admin/obtain-the-apps-installed-by-an-organization
func (r *ApplicationService) GetApplicationAppList(ctx context.Context, request *GetApplicationAppListReq, options ...MethodOptionFunc) (*GetApplicationAppListResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationAppList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationAppList mock enable")
		return r.cli.mock.mockApplicationGetApplicationAppList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationAppList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v3/app/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationAppListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationAppList mock ApplicationGetApplicationAppList method
func (r *Mock) MockApplicationGetApplicationAppList(f func(ctx context.Context, request *GetApplicationAppListReq, options ...MethodOptionFunc) (*GetApplicationAppListResp, *Response, error)) {
	r.mockApplicationGetApplicationAppList = f
}

// UnMockApplicationGetApplicationAppList un-mock ApplicationGetApplicationAppList method
func (r *Mock) UnMockApplicationGetApplicationAppList() {
	r.mockApplicationGetApplicationAppList = nil
}

// GetApplicationAppListReq ...
type GetApplicationAppListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页起始位置标示, 不填表示从头开始（不保证 page_token 一定为数字, 请填入上一次请求返回的 page_token）
	PageSize  *int64  `query:"page_size" json:"-"`  // 单页需求最大个数（最大 100）, 0 自动最大个数
	Lang      *string `query:"lang" json:"-"`       // 优先展示的应用信息的语言版本（zh_cn: 中文, en_us: 英文, ja_jp: 日文）
	Status    *int64  `query:"status" json:"-"`     // 要返回的应用的状态, 0:停用；1:启用；-1:全部, 默认为 -1
}

// GetApplicationAppListResp ...
type GetApplicationAppListResp struct {
	PageToken  string                            `json:"page_token,omitempty"`  // 下一个请求页应当给的起始位置
	PageSize   int64                             `json:"page_size,omitempty"`   // 本次请求实际返回的页大小
	TotalCount int64                             `json:"total_count,omitempty"` // 可用的应用总数
	HasMore    bool                              `json:"has_more,omitempty"`    // 是否还有更多应用
	Lang       string                            `json:"lang,omitempty"`        // 当前选择的版本语言
	AppList    *GetApplicationAppListRespAppList `json:"app_list,omitempty"`    // 应用列表
}

// GetApplicationAppListRespAppList ...
type GetApplicationAppListRespAppList struct {
	AppID                string `json:"app_id,omitempty"`                 // 应用 ID
	PrimaryLanguage      string `json:"primary_language,omitempty"`       // 应用首选语言
	AppName              string `json:"app_name,omitempty"`               // 应用名称
	Description          string `json:"description,omitempty"`            // 应用描述
	AvatarURL            string `json:"avatar_url,omitempty"`             // 应用 icon
	AppSceneType         int64  `json:"app_scene_type,omitempty"`         // 应用类型, 0: 企业自建应用；1: 应用商店应用
	Status               int64  `json:"status,omitempty"`                 // 启停状态, 0: 停用；1: 启用
	MobileDefaultAbility int64  `json:"mobile_default_ability,omitempty"` // 移动端默认的应用功能, 0: 未开启；1: 小程序；2: H5；8: 机器人
	PcDefaultAbility     int64  `json:"pc_default_ability,omitempty"`     // PC客户端默认的应用功能, 0: 未开启；1: 小程序；2: H5；8: 机器人
}

// getApplicationAppListResp ...
type getApplicationAppListResp struct {
	Code  int64                      `json:"code,omitempty"` // 返回码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 返回码的描述
	Data  *GetApplicationAppListResp `json:"data,omitempty"` // 返回的业务信息, 仅 code = 0 时有效
	Error *ErrorDetail               `json:"error,omitempty"`
}
