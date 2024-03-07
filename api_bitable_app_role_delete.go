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

// DeleteBitableAppRole 删除自定义角色
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/delete
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/advanced-permission/app-role/delete
func (r *BitableService) DeleteBitableAppRole(ctx context.Context, request *DeleteBitableAppRoleReq, options ...MethodOptionFunc) (*DeleteBitableAppRoleResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableAppRole != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableAppRole mock enable")
		return r.cli.mock.mockBitableDeleteBitableAppRole(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "DeleteBitableAppRole",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles/:role_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteBitableAppRoleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableDeleteBitableAppRole mock BitableDeleteBitableAppRole method
func (r *Mock) MockBitableDeleteBitableAppRole(f func(ctx context.Context, request *DeleteBitableAppRoleReq, options ...MethodOptionFunc) (*DeleteBitableAppRoleResp, *Response, error)) {
	r.mockBitableDeleteBitableAppRole = f
}

// UnMockBitableDeleteBitableAppRole un-mock BitableDeleteBitableAppRole method
func (r *Mock) UnMockBitableDeleteBitableAppRole() {
	r.mockBitableDeleteBitableAppRole = nil
}

// DeleteBitableAppRoleReq ...
type DeleteBitableAppRoleReq struct {
	AppToken string `path:"app_token" json:"-"` // 多维表格的唯一标识符 [app_token 参数说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "appbcbWCzen6D8dezhoCH2RpMAh"
	RoleID   string `path:"role_id" json:"-"`   // 自定义角色的id, 示例值: "roljRpwIUt"
}

// DeleteBitableAppRoleResp ...
type DeleteBitableAppRoleResp struct {
}

// deleteBitableAppRoleResp ...
type deleteBitableAppRoleResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteBitableAppRoleResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
