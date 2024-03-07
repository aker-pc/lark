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

// GetContactJobFamilyList 该接口用于获取租户序列列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/job_family/list
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/job_family/list
func (r *ContactService) GetContactJobFamilyList(ctx context.Context, request *GetContactJobFamilyListReq, options ...MethodOptionFunc) (*GetContactJobFamilyListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactJobFamilyList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Contact#GetContactJobFamilyList mock enable")
		return r.cli.mock.mockContactGetContactJobFamilyList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactJobFamilyList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/job_families",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactJobFamilyListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactJobFamilyList mock ContactGetContactJobFamilyList method
func (r *Mock) MockContactGetContactJobFamilyList(f func(ctx context.Context, request *GetContactJobFamilyListReq, options ...MethodOptionFunc) (*GetContactJobFamilyListResp, *Response, error)) {
	r.mockContactGetContactJobFamilyList = f
}

// UnMockContactGetContactJobFamilyList un-mock ContactGetContactJobFamilyList method
func (r *Mock) UnMockContactGetContactJobFamilyList() {
	r.mockContactGetContactJobFamilyList = nil
}

// GetContactJobFamilyListReq ...
type GetContactJobFamilyListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `10`, 取值范围: `1` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 3
	Name      *string `query:"name" json:"-"`       // 序列名称, 传入该字段时, 可查询指定序列名称对应的序列信息, 示例值: 2-2, 长度范围: `1` ～ `100` 字符
}

// GetContactJobFamilyListResp ...
type GetContactJobFamilyListResp struct {
	Items     []*GetContactJobFamilyListRespItem `json:"items,omitempty"`      // 序列信息
	PageToken string                             `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                               `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetContactJobFamilyListRespItem ...
type GetContactJobFamilyListRespItem struct {
	Name              string                                            `json:"name,omitempty"`                 // 序列名称。1-100字符, 支持中、英文及符号
	Description       string                                            `json:"description,omitempty"`          // 序列描述, 描述序列详情信息
	ParentJobFamilyID string                                            `json:"parent_job_family_id,omitempty"` // 上级序列ID。需是该租户的序列ID列表中的值, 对应唯一的序列名称。
	Status            bool                                              `json:"status,omitempty"`               // 是否启用
	I18nName          []*GetContactJobFamilyListRespItemI18nName        `json:"i18n_name,omitempty"`            // 多语言序列名称
	I18nDescription   []*GetContactJobFamilyListRespItemI18nDescription `json:"i18n_description,omitempty"`     // 多语言描述
	JobFamilyID       string                                            `json:"job_family_id,omitempty"`        // 职级序列ID
}

// GetContactJobFamilyListRespItemI18nDescription ...
type GetContactJobFamilyListRespItemI18nDescription struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// GetContactJobFamilyListRespItemI18nName ...
type GetContactJobFamilyListRespItemI18nName struct {
	Locale string `json:"locale,omitempty"` // 语言版本
	Value  string `json:"value,omitempty"`  // 字段名
}

// getContactJobFamilyListResp ...
type getContactJobFamilyListResp struct {
	Code  int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                       `json:"msg,omitempty"`  // 错误描述
	Data  *GetContactJobFamilyListResp `json:"data,omitempty"`
	Error *ErrorDetail                 `json:"error,omitempty"`
}
