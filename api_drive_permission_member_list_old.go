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

// GetDriveMemberPermissionListOld 该接口用于根据 filetoken 查询协作者, 目前包括人("user")和群("chat") 。
//
// - 该接口为旧版接口。推荐你使用新版接口接入业务, 详情参见[获取协作者列表（新版本）](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/list)。
// - 你能获取到协作者列表的前提是你对该文档有分享权限。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATN3UjLwUzN14CM1cTN
// new doc: https://open.feishu.cn/document/server-docs/docs/permission/permission-member/obtain-a-collaborator-list
func (r *DriveService) GetDriveMemberPermissionListOld(ctx context.Context, request *GetDriveMemberPermissionListOldReq, options ...MethodOptionFunc) (*GetDriveMemberPermissionListOldResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveMemberPermissionListOld != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#GetDriveMemberPermissionListOld mock enable")
		return r.cli.mock.mockDriveGetDriveMemberPermissionListOld(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveMemberPermissionListOld",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/permission/member/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveMemberPermissionListOldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveMemberPermissionListOld mock DriveGetDriveMemberPermissionListOld method
func (r *Mock) MockDriveGetDriveMemberPermissionListOld(f func(ctx context.Context, request *GetDriveMemberPermissionListOldReq, options ...MethodOptionFunc) (*GetDriveMemberPermissionListOldResp, *Response, error)) {
	r.mockDriveGetDriveMemberPermissionListOld = f
}

// UnMockDriveGetDriveMemberPermissionListOld un-mock DriveGetDriveMemberPermissionListOld method
func (r *Mock) UnMockDriveGetDriveMemberPermissionListOld() {
	r.mockDriveGetDriveMemberPermissionListOld = nil
}

// GetDriveMemberPermissionListOldReq ...
type GetDriveMemberPermissionListOldReq struct {
	Token string `json:"token,omitempty"` // 文件的 token, 获取方式见 [如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)
	Type  string `json:"type,omitempty"`  // 文档类型, 可选 doc、docx、sheet、bitable、file、minutes
}

// GetDriveMemberPermissionListOldResp ...
type GetDriveMemberPermissionListOldResp struct {
	Members []*GetDriveMemberPermissionListOldRespMember `json:"members,omitempty"` // 协作者列表
}

// GetDriveMemberPermissionListOldRespMember ...
type GetDriveMemberPermissionListOldRespMember struct {
	MemberType   string `json:"member_type,omitempty"`    // 协作者类型 "user" or "chat"
	MemberOpenID string `json:"member_open_id,omitempty"` // 协作者openid
	MemberUserID string `json:"member_user_id,omitempty"` // 协作者userid(仅当member_type="user"时有效)
	Perm         string `json:"perm,omitempty"`           // 协作者权限 (注意: 有"edit"权限的协作者一定有"view"权限)
}

// getDriveMemberPermissionListOldResp ...
type getDriveMemberPermissionListOldResp struct {
	Code  int64                                `json:"code,omitempty"`
	Msg   string                               `json:"msg,omitempty"`
	Data  *GetDriveMemberPermissionListOldResp `json:"data,omitempty"`
	Error *ErrorDetail                         `json:"error,omitempty"`
}
