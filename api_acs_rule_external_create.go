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

// CreateACSRuleExternal 创建或更新权限组
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/rule_external/create
func (r *ACSService) CreateACSRuleExternal(ctx context.Context, request *CreateACSRuleExternalReq, options ...MethodOptionFunc) (*CreateACSRuleExternalResp, *Response, error) {
	if r.cli.mock.mockACSCreateACSRuleExternal != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] ACS#CreateACSRuleExternal mock enable")
		return r.cli.mock.mockACSCreateACSRuleExternal(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "ACS",
		API:                 "CreateACSRuleExternal",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/acs/v1/rule_external",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(createACSRuleExternalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSCreateACSRuleExternal mock ACSCreateACSRuleExternal method
func (r *Mock) MockACSCreateACSRuleExternal(f func(ctx context.Context, request *CreateACSRuleExternalReq, options ...MethodOptionFunc) (*CreateACSRuleExternalResp, *Response, error)) {
	r.mockACSCreateACSRuleExternal = f
}

// UnMockACSCreateACSRuleExternal un-mock ACSCreateACSRuleExternal method
func (r *Mock) UnMockACSCreateACSRuleExternal() {
	r.mockACSCreateACSRuleExternal = nil
}

// CreateACSRuleExternalReq ...
type CreateACSRuleExternalReq struct {
	RuleID     *string                       `query:"rule_id" json:"-"`      // 权限组id-为空创建, 不为空则更新, 示例值: 7298933941867135276
	UserIDType *IDType                       `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Rule       *CreateACSRuleExternalReqRule `json:"rule,omitempty"`         // 权限组信息
}

// CreateACSRuleExternalReqRule ...
type CreateACSRuleExternalReqRule struct {
	ID           *string                                  `json:"id,omitempty"`            // 权限组id, 示例值: "34252345234523"
	Name         *string                                  `json:"name,omitempty"`          // 权限组名称, 示例值: "南门"
	Devices      []*CreateACSRuleExternalReqRuleDevice    `json:"devices,omitempty"`       // 权限组包含的设备, 长度范围: `0` ～ `5000`
	UserCount    *string                                  `json:"user_count,omitempty"`    // 权限组包含的员工个数, 示例值: "3"
	Users        []*CreateACSRuleExternalReqRuleUser      `json:"users,omitempty"`         // 权限组包含的员工列表, 长度范围: `0` ～ `10000`
	VisitorCount *string                                  `json:"visitor_count,omitempty"` // 权限组包含的访客个数, 示例值: "3"
	Visitors     []*CreateACSRuleExternalReqRuleVisitor   `json:"visitors,omitempty"`      // 权限组包含的访客列表, 长度范围: `0` ～ `10000`
	RemindFace   *bool                                    `json:"remind_face,omitempty"`   // 是否通知人员录入, 示例值: false
	OpeningTime  *CreateACSRuleExternalReqRuleOpeningTime `json:"opening_time,omitempty"`  // 开门时间段
	IsTemp       *bool                                    `json:"is_temp,omitempty"`       // 是否为临时权限组, 示例值: false
}

// CreateACSRuleExternalReqRuleDevice ...
type CreateACSRuleExternalReqRuleDevice struct {
	ID   *string `json:"id,omitempty"`   // 设备id, 示例值: "534545234523452345"
	Name *string `json:"name,omitempty"` // 设备名称, 示例值: "北门"
}

// CreateACSRuleExternalReqRuleOpeningTime ...
type CreateACSRuleExternalReqRuleOpeningTime struct {
	ValidDay *CreateACSRuleExternalReqRuleOpeningTimeValidDay  `json:"valid_day,omitempty"` // 有效日期
	Weekdays []int64                                           `json:"weekdays,omitempty"`  // 有效星期, 示例值: [1, 2, 3], 长度范围: `0` ～ `7`
	DayTimes []*CreateACSRuleExternalReqRuleOpeningTimeDayTime `json:"day_times,omitempty"` // 有效时间, 长度范围: `0` ～ `1`
}

// CreateACSRuleExternalReqRuleOpeningTimeDayTime ...
type CreateACSRuleExternalReqRuleOpeningTimeDayTime struct {
	StartHhmm int64 `json:"start_hhmm,omitempty"` // 起始时间, 示例值: 1200, 取值范围: `0` ～ `2400`
	EndHhmm   int64 `json:"end_hhmm,omitempty"`   // 结束时间, 示例值: 1400, 取值范围: `0` ～ `2400`
}

// CreateACSRuleExternalReqRuleOpeningTimeValidDay ...
type CreateACSRuleExternalReqRuleOpeningTimeValidDay struct {
	StartDay int64 `json:"start_day,omitempty"` // 权限开始时间, 示例值: 1699031483, 取值范围: `1000000000` ～ `2147483647`
	EndDay   int64 `json:"end_day,omitempty"`   // 权限结束时间, 示例值: 1699931483, 取值范围: `1000000000` ～ `2147483647`
}

// CreateACSRuleExternalReqRuleUser ...
type CreateACSRuleExternalReqRuleUser struct {
	UserType     int64   `json:"user_type,omitempty"`     // 用户类型, 示例值: 1, 可选值有: 1: 员工, 2: 部门, 10: 全体员工, 11: 访客
	UserID       *string `json:"user_id,omitempty"`       // 用户id, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserName     *string `json:"user_name,omitempty"`     // 用户名称, 示例值: "张三"
	PhoneNum     *string `json:"phone_num,omitempty"`     // 电话号码, 示例值: "1357890001"
	DepartmentID *string `json:"department_id,omitempty"` // 部门id, 示例值: "od-f7d44ab733f7602f5cc5194735fd9aaf"
}

// CreateACSRuleExternalReqRuleVisitor ...
type CreateACSRuleExternalReqRuleVisitor struct {
	UserType     int64   `json:"user_type,omitempty"`     // 用户类型, 示例值: 1, 可选值有: 1: 员工, 2: 部门, 10: 全体员工, 11: 访客
	UserID       *string `json:"user_id,omitempty"`       // 用户id, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserName     *string `json:"user_name,omitempty"`     // 用户名称, 示例值: "张三"
	PhoneNum     *string `json:"phone_num,omitempty"`     // 电话号码, 示例值: "1357890001"
	DepartmentID *string `json:"department_id,omitempty"` // 部门id, 示例值: "od-f7d44ab733f7602f5cc5194735fd9aaf"
}

// CreateACSRuleExternalResp ...
type CreateACSRuleExternalResp struct {
	RuleID string `json:"rule_id,omitempty"` // 权限组id
}

// createACSRuleExternalResp ...
type createACSRuleExternalResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateACSRuleExternalResp `json:"data,omitempty"`
}
