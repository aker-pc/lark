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

// CopyDriveFile 将文件复制到用户云空间的其他文件夹中。不支持复制文件夹。
//
// 如果目标文件夹是我的空间, 则复制的文件会在「我的空间」的「归我所有」列表里。
// 该接口不支持并发拷贝多个文件, 且调用频率上限为 5QPS 且 10000次/天
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/copy
// new doc: https://open.feishu.cn/document/server-docs/docs/drive-v1/file/copy
func (r *DriveService) CopyDriveFile(ctx context.Context, request *CopyDriveFileReq, options ...MethodOptionFunc) (*CopyDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveCopyDriveFile != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CopyDriveFile mock enable")
		return r.cli.mock.mockDriveCopyDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CopyDriveFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/copy",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(copyDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCopyDriveFile mock DriveCopyDriveFile method
func (r *Mock) MockDriveCopyDriveFile(f func(ctx context.Context, request *CopyDriveFileReq, options ...MethodOptionFunc) (*CopyDriveFileResp, *Response, error)) {
	r.mockDriveCopyDriveFile = f
}

// UnMockDriveCopyDriveFile un-mock DriveCopyDriveFile method
func (r *Mock) UnMockDriveCopyDriveFile() {
	r.mockDriveCopyDriveFile = nil
}

// CopyDriveFileReq ...
type CopyDriveFileReq struct {
	FileToken   string                   `path:"file_token" json:"-"`    // 被复制的文件token, 示例值: "doccngpahSdXrFPIBD4XdIabcef"
	UserIDType  *IDType                  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	Name        string                   `json:"name,omitempty"`         // 被复制文件的新名称, 示例值: "test.txt"
	Type        *string                  `json:"type,omitempty"`         // 被复制文件的类型, 如果该值为空或者与文件实际类型不匹配, 接口会返回失败, 示例值: "doc", 可选值有: file: 文件类型, doc: 文档类型, sheet: 电子表格类型, bitable: 多维表格类型, docx: 新版文档类型, mindnote: 思维笔记类型, slides: 幻灯片类型
	FolderToken string                   `json:"folder_token,omitempty"` // 文件被复制到的目标文件夹token, 示例值: "fldbcO1UuPz8VwnpPx5a92abcef"
	Extra       []*CopyDriveFileReqExtra `json:"extra,omitempty"`        // 用户自定义请求附加参数, 用于实现特殊的复制语义
}

// CopyDriveFileReqExtra ...
type CopyDriveFileReqExtra struct {
	Key   string `json:"key,omitempty"`   // 自定义属性键对象, 示例值: "target_type"
	Value string `json:"value,omitempty"` // 自定义属性值对象, 示例值: "docx"
}

// CopyDriveFileResp ...
type CopyDriveFileResp struct {
	File *CopyDriveFileRespFile `json:"file,omitempty"` // 复制后的文件资源
}

// CopyDriveFileRespFile ...
type CopyDriveFileRespFile struct {
	Token        string                             `json:"token,omitempty"`         // 文件标识
	Name         string                             `json:"name,omitempty"`          // 文件名
	Type         string                             `json:"type,omitempty"`          // 文件类型
	ParentToken  string                             `json:"parent_token,omitempty"`  // 父文件夹标识
	URL          string                             `json:"url,omitempty"`           // 在浏览器中查看的链接
	ShortcutInfo *CopyDriveFileRespFileShortcutInfo `json:"shortcut_info,omitempty"` // 快捷方式文件信息
	CreatedTime  string                             `json:"created_time,omitempty"`  // 文件创建时间
	ModifiedTime string                             `json:"modified_time,omitempty"` // 文件最近修改时间
	OwnerID      string                             `json:"owner_id,omitempty"`      // 文件所有者
}

// CopyDriveFileRespFileShortcutInfo ...
type CopyDriveFileRespFileShortcutInfo struct {
	TargetType  string `json:"target_type,omitempty"`  // 快捷方式指向的原文件类型
	TargetToken string `json:"target_token,omitempty"` // 快捷方式指向的原文件token
}

// copyDriveFileResp ...
type copyDriveFileResp struct {
	Code  int64              `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string             `json:"msg,omitempty"`  // 错误描述
	Data  *CopyDriveFileResp `json:"data,omitempty"`
	Error *ErrorDetail       `json:"error,omitempty"`
}
