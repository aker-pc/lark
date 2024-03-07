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

// UpdateApprovalMessage 此接口可以根据审批bot消息id及相应状态, 更新相应的审批bot消息, 只可用于更新待审批模板的bot消息。例如, 给用户推送了审批待办消息, 当用户处理该消息后, 可以将之前推送的Bot消息更新为已审批。
//
// 注意: 该接口只能更新模板为 1008「收到审批待办」的卡片。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjNyYjLwYjM24CM2IjN
// new doc: https://open.feishu.cn/document/server-docs/approval-v4/message/update-bot-messages
func (r *ApprovalService) UpdateApprovalMessage(ctx context.Context, request *UpdateApprovalMessageReq, options ...MethodOptionFunc) (*UpdateApprovalMessageResp, *Response, error) {
	if r.cli.mock.mockApprovalUpdateApprovalMessage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Approval#UpdateApprovalMessage mock enable")
		return r.cli.mock.mockApprovalUpdateApprovalMessage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "UpdateApprovalMessage",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/approval/v1/message/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateApprovalMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalUpdateApprovalMessage mock ApprovalUpdateApprovalMessage method
func (r *Mock) MockApprovalUpdateApprovalMessage(f func(ctx context.Context, request *UpdateApprovalMessageReq, options ...MethodOptionFunc) (*UpdateApprovalMessageResp, *Response, error)) {
	r.mockApprovalUpdateApprovalMessage = f
}

// UnMockApprovalUpdateApprovalMessage un-mock ApprovalUpdateApprovalMessage method
func (r *Mock) UnMockApprovalUpdateApprovalMessage() {
	r.mockApprovalUpdateApprovalMessage = nil
}

// UpdateApprovalMessageReq ...
type UpdateApprovalMessageReq struct {
	MessageID        string  `json:"message_id,omitempty"`         // 卡片 id, 发送卡片时会拿到
	Status           string  `json:"status,omitempty"`             // 状态类型, 用于更新第一个action文字内容, 枚举: APPROVED:-已同意 REJECTED:-已拒绝 CANCELLED:-已撤回 FORWARDED:-已转交 ROLLBACK:-已回退 ADD:-已加签 DELETED:-已删除 PROCESSED:-已处理 CUSTOM:-自定义按钮状态
	StatusName       *string `json:"status_name,omitempty"`        // status=CUSTOM时可以自定义审批同意/拒绝后title状态
	DetailActionName *string `json:"detail_action_name,omitempty"` // status=CUSTOM时可以自定义审批同意/拒绝后“查看详情按钮名称”
	I18nResources    *string `json:"i18n_resources,omitempty"`     // i18n国际化文案
}

// UpdateApprovalMessageResp ...
type UpdateApprovalMessageResp struct {
	MessageID string `json:"message_id,omitempty"` // 消息 id, 用于卡片更新、撤回
}

// updateApprovalMessageResp ...
type updateApprovalMessageResp struct {
	Code  int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                     `json:"msg,omitempty"`  // 返回码的描述
	Data  *UpdateApprovalMessageResp `json:"data,omitempty"` // 返回业务信息
	Error *ErrorDetail               `json:"error,omitempty"`
}
