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

// BatchDeleteBitableAppRoleMember 批量删除自定义角色的协作者
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role-member/batch_delete
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/advanced-permission/app-role-member/batch_delete
func (r *BitableService) BatchDeleteBitableAppRoleMember(ctx context.Context, request *BatchDeleteBitableAppRoleMemberReq, options ...MethodOptionFunc) (*BatchDeleteBitableAppRoleMemberResp, *Response, error) {
	if r.cli.mock.mockBitableBatchDeleteBitableAppRoleMember != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Bitable#BatchDeleteBitableAppRoleMember mock enable")
		return r.cli.mock.mockBitableBatchDeleteBitableAppRoleMember(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "BatchDeleteBitableAppRoleMember",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id/members/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchDeleteBitableAppRoleMemberResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableBatchDeleteBitableAppRoleMember mock BitableBatchDeleteBitableAppRoleMember method
func (r *Mock) MockBitableBatchDeleteBitableAppRoleMember(f func(ctx context.Context, request *BatchDeleteBitableAppRoleMemberReq, options ...MethodOptionFunc) (*BatchDeleteBitableAppRoleMemberResp, *Response, error)) {
	r.mockBitableBatchDeleteBitableAppRoleMember = f
}

// UnMockBitableBatchDeleteBitableAppRoleMember un-mock BitableBatchDeleteBitableAppRoleMember method
func (r *Mock) UnMockBitableBatchDeleteBitableAppRoleMember() {
	r.mockBitableBatchDeleteBitableAppRoleMember = nil
}

// BatchDeleteBitableAppRoleMemberReq ...
type BatchDeleteBitableAppRoleMemberReq struct {
	AppToken   string                                      `path:"app_token" json:"-"`    // 多维表格文档 Token, 示例值: "bascnnKKvcoUblgmmhZkYqabcef"
	RoleID     string                                      `path:"role_id" json:"-"`      // 自定义角色 ID, 示例值: "rolNGhPqks"
	MemberList []*BatchDeleteBitableAppRoleMemberReqMember `json:"member_list,omitempty"` // 协作者列表, 最大长度: `100`
}

// BatchDeleteBitableAppRoleMemberReqMember ...
type BatchDeleteBitableAppRoleMemberReqMember struct {
	Type *string `json:"type,omitempty"` // 协作者 ID 类型, 示例值: "open_id", 可选值有: open_id: 协作者 ID 类型为 open_id, union_id: 协作者 ID 类型为 union_id, user_id: 协作者 ID 类型为 user_id, chat_id: 协作者 ID 类型为 chat_id, department_id: 协作者 ID 类型为 department_id, open_department_id: 协作者 ID 类型为 open_department_id, 默认值: `open_id`
	ID   string  `json:"id,omitempty"`   // 协作者 ID, 示例值: "ou_35990a9d9052051a2fae9b2f1afabcef"
}

// BatchDeleteBitableAppRoleMemberResp ...
type BatchDeleteBitableAppRoleMemberResp struct {
}

// batchDeleteBitableAppRoleMemberResp ...
type batchDeleteBitableAppRoleMemberResp struct {
	Code  int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                               `json:"msg,omitempty"`  // 错误描述
	Data  *BatchDeleteBitableAppRoleMemberResp `json:"data,omitempty"`
	Error *ErrorDetail                         `json:"error,omitempty"`
}
