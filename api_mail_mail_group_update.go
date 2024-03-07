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

// UpdateMailGroup 更新邮件组所有信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/mailgroup/update
// new doc: https://open.feishu.cn/document/server-docs/mail-v1/mail-group/mailgroup/update
func (r *MailService) UpdateMailGroup(ctx context.Context, request *UpdateMailGroupReq, options ...MethodOptionFunc) (*UpdateMailGroupResp, *Response, error) {
	if r.cli.mock.mockMailUpdateMailGroup != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Mail#UpdateMailGroup mock enable")
		return r.cli.mock.mockMailUpdateMailGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Mail",
		API:                   "UpdateMailGroup",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/mail/v1/mailgroups/:mailgroup_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateMailGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMailUpdateMailGroup mock MailUpdateMailGroup method
func (r *Mock) MockMailUpdateMailGroup(f func(ctx context.Context, request *UpdateMailGroupReq, options ...MethodOptionFunc) (*UpdateMailGroupResp, *Response, error)) {
	r.mockMailUpdateMailGroup = f
}

// UnMockMailUpdateMailGroup un-mock MailUpdateMailGroup method
func (r *Mock) UnMockMailUpdateMailGroup() {
	r.mockMailUpdateMailGroup = nil
}

// UpdateMailGroupReq ...
type UpdateMailGroupReq struct {
	MailGroupID    string  `path:"mailgroup_id" json:"-"`       // 邮件组ID或者邮件组地址, 示例值: "xxxxxxxxxxxxxxx 或 test_mail_group@xxx.xx"
	Email          *string `json:"email,omitempty"`             // 邮件组地址, 示例值: "test_mail_group@xxx.xx"
	Name           *string `json:"name,omitempty"`              // 邮件组名称, 示例值: "test mail group"
	Description    *string `json:"description,omitempty"`       // 邮件组描述, 示例值: "mail group for testing"
	WhoCanSendMail *string `json:"who_can_send_mail,omitempty"` // 谁可发送邮件到此邮件组, 示例值: "ALL_INTERNAL_USERS", 可选值有: ANYONE: 任何人, ALL_INTERNAL_USERS: 仅组织内部成员, ALL_GROUP_MEMBERS: 仅邮件组成员, CUSTOM_MEMBERS: 自定义成员
}

// UpdateMailGroupResp ...
type UpdateMailGroupResp struct {
	MailGroupID             string `json:"mailgroup_id,omitempty"`               // 邮件组ID
	Email                   string `json:"email,omitempty"`                      // 邮件组地址
	Name                    string `json:"name,omitempty"`                       // 邮件组名称
	Description             string `json:"description,omitempty"`                // 邮件组描述
	DirectMembersCount      string `json:"direct_members_count,omitempty"`       // 邮件组成员数量
	IncludeExternalMember   bool   `json:"include_external_member,omitempty"`    // 是否包含外部成员
	IncludeAllCompanyMember bool   `json:"include_all_company_member,omitempty"` // 是否是全员邮件组
	WhoCanSendMail          string `json:"who_can_send_mail,omitempty"`          // 谁可发送邮件到此邮件组, 可选值有: ANYONE: 任何人, ALL_INTERNAL_USERS: 仅组织内部成员, ALL_GROUP_MEMBERS: 仅邮件组成员, CUSTOM_MEMBERS: 自定义成员
}

// updateMailGroupResp ...
type updateMailGroupResp struct {
	Code  int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string               `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateMailGroupResp `json:"data,omitempty"`
	Error *ErrorDetail         `json:"error,omitempty"`
}
