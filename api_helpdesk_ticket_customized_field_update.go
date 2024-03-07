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

// UpdateHelpdeskTicketCustomizedField 该接口用于更新自定义字段。
//
// 注意事项:
// user_access_token 访问, 需要操作者是当前服务台的管理员或所有者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/update-ticket-customized-field
// new doc: https://open.feishu.cn/document/server-docs/helpdesk-v1/ticket-management/ticket_customized_field/update-ticket-customized-field
func (r *HelpdeskService) UpdateHelpdeskTicketCustomizedField(ctx context.Context, request *UpdateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskTicketCustomizedField != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskTicketCustomizedField",
		Method:              "PATCH",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskUpdateHelpdeskTicketCustomizedField mock HelpdeskUpdateHelpdeskTicketCustomizedField method
func (r *Mock) MockHelpdeskUpdateHelpdeskTicketCustomizedField(f func(ctx context.Context, request *UpdateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskTicketCustomizedField = f
}

// UnMockHelpdeskUpdateHelpdeskTicketCustomizedField un-mock HelpdeskUpdateHelpdeskTicketCustomizedField method
func (r *Mock) UnMockHelpdeskUpdateHelpdeskTicketCustomizedField() {
	r.mockHelpdeskUpdateHelpdeskTicketCustomizedField = nil
}

// UpdateHelpdeskTicketCustomizedFieldReq ...
type UpdateHelpdeskTicketCustomizedFieldReq struct {
	TicketCustomizedFieldID string                  `path:"ticket_customized_field_id" json:"-"` // 工单自定义字段ID, 示例值: "6948728206392295444"
	DisplayName             *string                 `json:"display_name,omitempty"`              // 名称, 示例值: "test dropdown"
	Position                *string                 `json:"position,omitempty"`                  // 字段在列表后台管理列表中的位置, 示例值: "3"
	Description             *string                 `json:"description,omitempty"`               // 描述, 示例值: "下拉示例"
	Visible                 *bool                   `json:"visible,omitempty"`                   // 是否可见, 示例值: true
	Required                *bool                   `json:"required,omitempty"`                  // 是否必填, 示例值: false
	DropdownOptions         *HelpdeskDropdownOption `json:"dropdown_options,omitempty"`          // 下拉列表选项
}

// UpdateHelpdeskTicketCustomizedFieldResp ...
type UpdateHelpdeskTicketCustomizedFieldResp struct {
}

// updateHelpdeskTicketCustomizedFieldResp ...
type updateHelpdeskTicketCustomizedFieldResp struct {
	Code  int64                                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                   `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateHelpdeskTicketCustomizedFieldResp `json:"data,omitempty"`
	Error *ErrorDetail                             `json:"error,omitempty"`
}
