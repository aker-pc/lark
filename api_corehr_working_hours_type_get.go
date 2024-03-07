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

// GetCoreHRWorkingHoursType 根据 ID 查询单个工时制度。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/working_hours_type/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/basic-infomation/working_hours_type/get
func (r *CoreHRService) GetCoreHRWorkingHoursType(ctx context.Context, request *GetCoreHRWorkingHoursTypeReq, options ...MethodOptionFunc) (*GetCoreHRWorkingHoursTypeResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRWorkingHoursType != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRWorkingHoursType mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRWorkingHoursType(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRWorkingHoursType",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/working_hours_types/:working_hours_type_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRWorkingHoursTypeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRWorkingHoursType mock CoreHRGetCoreHRWorkingHoursType method
func (r *Mock) MockCoreHRGetCoreHRWorkingHoursType(f func(ctx context.Context, request *GetCoreHRWorkingHoursTypeReq, options ...MethodOptionFunc) (*GetCoreHRWorkingHoursTypeResp, *Response, error)) {
	r.mockCoreHRGetCoreHRWorkingHoursType = f
}

// UnMockCoreHRGetCoreHRWorkingHoursType un-mock CoreHRGetCoreHRWorkingHoursType method
func (r *Mock) UnMockCoreHRGetCoreHRWorkingHoursType() {
	r.mockCoreHRGetCoreHRWorkingHoursType = nil
}

// GetCoreHRWorkingHoursTypeReq ...
type GetCoreHRWorkingHoursTypeReq struct {
	WorkingHoursTypeID string `path:"working_hours_type_id" json:"-"` // 工时制度 ID, 示例值: "1212"
}

// GetCoreHRWorkingHoursTypeResp ...
type GetCoreHRWorkingHoursTypeResp struct {
	WorkingHoursType *GetCoreHRWorkingHoursTypeRespWorkingHoursType `json:"working_hours_type,omitempty"` // 工时制度信息
}

// GetCoreHRWorkingHoursTypeRespWorkingHoursType ...
type GetCoreHRWorkingHoursTypeRespWorkingHoursType struct {
	ID                  string                                                      `json:"id,omitempty"`                     // 工时制度 ID
	Code                string                                                      `json:"code,omitempty"`                   // 编码
	Name                []*GetCoreHRWorkingHoursTypeRespWorkingHoursTypeName        `json:"name,omitempty"`                   // 名称
	CountryRegionIDList []string                                                    `json:"country_region_id_list,omitempty"` // 国家/地区 ID 列表
	DefaultForJob       bool                                                        `json:"default_for_job,omitempty"`        // 职务默认值
	Active              bool                                                        `json:"active,omitempty"`                 // 是否启用
	CustomFields        []*GetCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField `json:"custom_fields,omitempty"`          // 自定义字段
}

// GetCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField ...
type GetCoreHRWorkingHoursTypeRespWorkingHoursTypeCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRWorkingHoursTypeRespWorkingHoursTypeName ...
type GetCoreHRWorkingHoursTypeRespWorkingHoursTypeName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRWorkingHoursTypeResp ...
type getCoreHRWorkingHoursTypeResp struct {
	Code  int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                         `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRWorkingHoursTypeResp `json:"data,omitempty"`
	Error *ErrorDetail                   `json:"error,omitempty"`
}
