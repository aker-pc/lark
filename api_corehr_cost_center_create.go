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

// CreateCoreHRCostCenter 创建成本中心
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/cost_center/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/organization-management/cost_center/create
func (r *CoreHRService) CreateCoreHRCostCenter(ctx context.Context, request *CreateCoreHRCostCenterReq, options ...MethodOptionFunc) (*CreateCoreHRCostCenterResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRCostCenter != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRCostCenter mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRCostCenter(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRCostCenter",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/cost_centers",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRCostCenterResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRCostCenter mock CoreHRCreateCoreHRCostCenter method
func (r *Mock) MockCoreHRCreateCoreHRCostCenter(f func(ctx context.Context, request *CreateCoreHRCostCenterReq, options ...MethodOptionFunc) (*CreateCoreHRCostCenterResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRCostCenter = f
}

// UnMockCoreHRCreateCoreHRCostCenter un-mock CoreHRCreateCoreHRCostCenter method
func (r *Mock) UnMockCoreHRCreateCoreHRCostCenter() {
	r.mockCoreHRCreateCoreHRCostCenter = nil
}

// CreateCoreHRCostCenterReq ...
type CreateCoreHRCostCenterReq struct {
	UserIDType         *IDType                                 `query:"user_id_type" json:"-"`          // 用户 ID 类型, 示例值: people_corehr_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_corehr_id: 以飞书人事的 ID 来识别用户, 默认值: `people_corehr_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Name               []*CreateCoreHRCostCenterReqName        `json:"name,omitempty"`                  // 成本中心名称。名称不能包含「/」「；」「;」符号, 名称长度不能超过 255 个字符
	Code               *string                                 `json:"code,omitempty"`                  // 编码, 示例值: "MDPD00000023"
	ParentCostCenterID *string                                 `json:"parent_cost_center_id,omitempty"` // 上级成本中心ID, 示例值: "6862995757234914824"
	Managers           []string                                `json:"managers,omitempty"`              // 成本中心负责人ID 列表, 示例值: ["6862995757234914824"]
	Description        []*CreateCoreHRCostCenterReqDescription `json:"description,omitempty"`           // 成本中心描述
	EffectiveTime      string                                  `json:"effective_time,omitempty"`        // 生效时间, 示例值: "2020-01-01"
}

// CreateCoreHRCostCenterReqDescription ...
type CreateCoreHRCostCenterReqDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "张三"
}

// CreateCoreHRCostCenterReqName ...
type CreateCoreHRCostCenterReqName struct {
	Lang  string `json:"lang,omitempty"`  // 语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 内容, 示例值: "张三"
}

// CreateCoreHRCostCenterResp ...
type CreateCoreHRCostCenterResp struct {
	CostCenter *CreateCoreHRCostCenterRespCostCenter `json:"cost_center,omitempty"` // 成本中心信息
}

// CreateCoreHRCostCenterRespCostCenter ...
type CreateCoreHRCostCenterRespCostCenter struct {
	CostCenterID       string                                             `json:"cost_center_id,omitempty"`        // 成本中心ID
	Name               []*CreateCoreHRCostCenterRespCostCenterName        `json:"name,omitempty"`                  // 成本中心名称
	Code               string                                             `json:"code,omitempty"`                  // 编码
	ParentCostCenterID string                                             `json:"parent_cost_center_id,omitempty"` // 上级成本中心ID
	Managers           []string                                           `json:"managers,omitempty"`              // 成本中心负责人ID 列表
	Description        []*CreateCoreHRCostCenterRespCostCenterDescription `json:"description,omitempty"`           // 成本中心描述
	EffectiveTime      string                                             `json:"effective_time,omitempty"`        // 生效时间
	ExpirationTime     string                                             `json:"expiration_time,omitempty"`       // 过期时间
	Active             bool                                               `json:"active,omitempty"`                // 当前实体是否启用
}

// CreateCoreHRCostCenterRespCostCenterDescription ...
type CreateCoreHRCostCenterRespCostCenterDescription struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// CreateCoreHRCostCenterRespCostCenterName ...
type CreateCoreHRCostCenterRespCostCenterName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// createCoreHRCostCenterResp ...
type createCoreHRCostCenterResp struct {
	Code  int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                      `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHRCostCenterResp `json:"data,omitempty"`
	Error *ErrorDetail                `json:"error,omitempty"`
}
