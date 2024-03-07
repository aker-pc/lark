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

// CreateCoreHRWorkingHoursType 创建工时制度。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/working_hours_type/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/working_hours_type/create
func (r *CoreHRService) CreateCoreHRWorkingHoursType(ctx context.Context, request *CreateCoreHRWorkingHoursTypeReq, options ...MethodOptionFunc) (*CreateCoreHRWorkingHoursTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRWorkingHoursType != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRWorkingHoursType mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRWorkingHoursType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRWorkingHoursType",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/working_hours_types",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRWorkingHoursTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRWorkingHoursType mock CoreHRCreateCoreHRWorkingHoursType method
func (r *Mock) MockCoreHRCreateCoreHRWorkingHoursType(f func(ctx context.Context, request *CreateCoreHRWorkingHoursTypeReq, options ...MethodOptionFunc) (*CreateCoreHRWorkingHoursTypeResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRWorkingHoursType = f
}

// UnMockCoreHRCreateCoreHRWorkingHoursType un-mock CoreHRCreateCoreHRWorkingHoursType method
func (r *Mock) UnMockCoreHRCreateCoreHRWorkingHoursType() {
	r.mockCoreHRCreateCoreHRWorkingHoursType = nil
}

// CreateCoreHRWorkingHoursTypeReq ...
type CreateCoreHRWorkingHoursTypeReq struct {
	ClientToken         *string                                       `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	Code                *string                                       `json:"code,omitempty"`                   // 编码, 示例值: "1"
	Name                []*CreateCoreHRWorkingHoursTypeReqName        `json:"name,omitempty"`                   // 名称
	CountryRegionIDList []string                                      `json:"country_region_id_list,omitempty"` // 国家/地区 ID 列表, 示例值: ["6890452208593356295"]
	DefaultForJob       bool                                          `json:"default_for_job,omitempty"`        // 职务默认值, 示例值: true
	Active              bool                                          `json:"active,omitempty"`                 // 是否启用, 示例值: true
	CustomFields        []*CreateCoreHRWorkingHoursTypeReqCustomField `json:"custom_fields,omitempty"`          // 自定义字段
}

// CreateCoreHRWorkingHoursTypeReqCustomField ...
type CreateCoreHRWorkingHoursTypeReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHRWorkingHoursTypeReqName ...
type CreateCoreHRWorkingHoursTypeReqName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言, 示例值: "zh-CN"
	Value string `json:"value,omitempty"` // 名称信息的内容, 示例值: "张三"
}

// CreateCoreHRWorkingHoursTypeResp ...
type CreateCoreHRWorkingHoursTypeResp struct {
	WorkingHoursType *CreateCoreHRWorkingHoursTypeRespWorkingHoursType `json:"working_hours_type,omitempty"` // 工时制度
}

// CreateCoreHRWorkingHoursTypeRespWorkingHoursType ...
type CreateCoreHRWorkingHoursTypeRespWorkingHoursType struct {
	ID                  string                                                         `json:"id,omitempty"`                     // 工时制度 ID
	Code                string                                                         `json:"code,omitempty"`                   // 编码
	Name                []*CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeName        `json:"name,omitempty"`                   // 名称
	CountryRegionIDList []string                                                       `json:"country_region_id_list,omitempty"` // 国家/地区 ID 列表
	DefaultForJob       bool                                                           `json:"default_for_job,omitempty"`        // 职务默认值
	Active              bool                                                           `json:"active,omitempty"`                 // 是否启用
	CustomFields        []*CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField `json:"custom_fields,omitempty"`          // 自定义字段
}

// CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField ...
type CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeName ...
type CreateCoreHRWorkingHoursTypeRespWorkingHoursTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHRWorkingHoursTypeResp ...
type createCoreHRWorkingHoursTypeResp struct {
	Code  int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                            `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHRWorkingHoursTypeResp `json:"data,omitempty"`
	Error *ErrorDetail                      `json:"error,omitempty"`
}
