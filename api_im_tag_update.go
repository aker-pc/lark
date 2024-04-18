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

// UpdateIMTag 修改标签在各个语言下的名称。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/tag/patch
func (r *MessageService) UpdateIMTag(ctx context.Context, request *UpdateIMTagReq, options ...MethodOptionFunc) (*UpdateIMTagResp, *Response, error) {
	if r.cli.mock.mockMessageUpdateIMTag != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#UpdateIMTag mock enable")
		return r.cli.mock.mockMessageUpdateIMTag(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "UpdateIMTag",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v2/tags/:tag_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateIMTagResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageUpdateIMTag mock MessageUpdateIMTag method
func (r *Mock) MockMessageUpdateIMTag(f func(ctx context.Context, request *UpdateIMTagReq, options ...MethodOptionFunc) (*UpdateIMTagResp, *Response, error)) {
	r.mockMessageUpdateIMTag = f
}

// UnMockMessageUpdateIMTag un-mock MessageUpdateIMTag method
func (r *Mock) UnMockMessageUpdateIMTag() {
	r.mockMessageUpdateIMTag = nil
}

// UpdateIMTagReq ...
type UpdateIMTagReq struct {
	TagID    string                  `path:"tag_id" json:"-"`     // 标签 ID, 示例值: "716168xxxxx"
	PatchTag *UpdateIMTagReqPatchTag `json:"patch_tag,omitempty"` // 编辑标签
}

// UpdateIMTagReqPatchTag ...
type UpdateIMTagReqPatchTag struct {
	ID        *string                           `json:"id,omitempty"`         // 标签 ID, 示例值: "716168xxxxx"
	Name      *string                           `json:"name,omitempty"`       // 标签名称, 示例值: "标签名称"
	I18nNames []*UpdateIMTagReqPatchTagI18nName `json:"i18n_names,omitempty"` // i18n 多语言名称集合, 长度范围: `0` ～ `40`
}

// UpdateIMTagReqPatchTagI18nName ...
type UpdateIMTagReqPatchTagI18nName struct {
	Locale string  `json:"locale,omitempty"` // 语言, 示例值: "zh_cn"
	Name   *string `json:"name,omitempty"`   // 名称, 示例值: "标签2"
}

// UpdateIMTagResp ...
type UpdateIMTagResp struct {
	TagInfo            *UpdateIMTagRespTagInfo            `json:"tag_info,omitempty"`              // 编辑后的标签信息
	PatchTagFailReason *UpdateIMTagRespPatchTagFailReason `json:"patch_tag_fail_reason,omitempty"` // 修改失败原因
}

// UpdateIMTagRespPatchTagFailReason ...
type UpdateIMTagRespPatchTagFailReason struct {
	DuplicateID string `json:"duplicate_id,omitempty"` // 名称重复的标签id
}

// UpdateIMTagRespTagInfo ...
type UpdateIMTagRespTagInfo struct {
	ID         string                            `json:"id,omitempty"`          // 标签 ID
	TagType    string                            `json:"tag_type,omitempty"`    // 标签类型
	Name       string                            `json:"name,omitempty"`        // 标签名称
	I18nNames  []*UpdateIMTagRespTagInfoI18nName `json:"i18n_names,omitempty"`  // i18n 多语言名称集合
	CreateTime string                            `json:"create_time,omitempty"` // 创建时间
	UpdateTime string                            `json:"update_time,omitempty"` // 更新时间
}

// UpdateIMTagRespTagInfoI18nName ...
type UpdateIMTagRespTagInfoI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言
	Name   string `json:"name,omitempty"`   // 名称
}

// updateIMTagResp ...
type updateIMTagResp struct {
	Code  int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string           `json:"msg,omitempty"`  // 错误描述
	Data  *UpdateIMTagResp `json:"data,omitempty"`
	Error *ErrorDetail     `json:"error,omitempty"`
}
