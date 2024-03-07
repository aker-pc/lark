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

// GetCoreHRContract 根据 ID 查询单个合同。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/contract/get
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/contract/get
func (r *CoreHRService) GetCoreHRContract(ctx context.Context, request *GetCoreHRContractReq, options ...MethodOptionFunc) (*GetCoreHRContractResp, *Response, error) {
	if r.cli.mock.mockCoreHRGetCoreHRContract != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#GetCoreHRContract mock enable")
		return r.cli.mock.mockCoreHRGetCoreHRContract(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "GetCoreHRContract",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/contracts/:contract_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHRContractResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRGetCoreHRContract mock CoreHRGetCoreHRContract method
func (r *Mock) MockCoreHRGetCoreHRContract(f func(ctx context.Context, request *GetCoreHRContractReq, options ...MethodOptionFunc) (*GetCoreHRContractResp, *Response, error)) {
	r.mockCoreHRGetCoreHRContract = f
}

// UnMockCoreHRGetCoreHRContract un-mock CoreHRGetCoreHRContract method
func (r *Mock) UnMockCoreHRGetCoreHRContract() {
	r.mockCoreHRGetCoreHRContract = nil
}

// GetCoreHRContractReq ...
type GetCoreHRContractReq struct {
	ContractID string `path:"contract_id" json:"-"` // 合同ID, 示例值: "151515"
}

// GetCoreHRContractResp ...
type GetCoreHRContractResp struct {
	Contract *GetCoreHRContractRespContract `json:"contract,omitempty"` // 合同信息
}

// GetCoreHRContractRespContract ...
type GetCoreHRContractRespContract struct {
	ID                  string                                      `json:"id,omitempty"`                     // 合同ID
	EffectiveTime       string                                      `json:"effective_time,omitempty"`         // 合同开始日期
	ExpirationTime      string                                      `json:"expiration_time,omitempty"`        // 实际结束日期
	EmploymentID        string                                      `json:"employment_id,omitempty"`          // 雇员 ID, 枚举值及详细信息可通过【批量查询雇佣信息】接口查询获得
	ContractType        *GetCoreHRContractRespContractContractType  `json:"contract_type,omitempty"`          // 合同类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同类型（contract_type）枚举定义部分获得
	FirstPartyCompanyID string                                      `json:"first_party_company_id,omitempty"` // 甲方, 引用Company的ID, 枚举值及详细信息可通过【批量查询公司】接口查询获得
	PersonID            string                                      `json:"person_id,omitempty"`              // Person ID, 枚举值及详细信息可通过【批量查询个人信息】接口查询获得
	CustomFields        []*GetCoreHRContractRespContractCustomField `json:"custom_fields,omitempty"`          // 自定义字段
	DurationType        *GetCoreHRContractRespContractDurationType  `json:"duration_type,omitempty"`          // 期限类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)合同期限类型（duration_type）枚举定义部分获得
	ContractEndDate     string                                      `json:"contract_end_date,omitempty"`      // 合同结束日期
	ContractNumber      string                                      `json:"contract_number,omitempty"`        // 合同编号
	SigningType         *GetCoreHRContractRespContractSigningType   `json:"signing_type,omitempty"`           // 签订类型, 枚举值可通过文档[【飞书人事枚举常量】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/feishu-people-enum-constant)签订类型（signing_type）枚举定义部分获得
}

// GetCoreHRContractRespContractContractType ...
type GetCoreHRContractRespContractContractType struct {
	EnumName string                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRContractRespContractContractTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRContractRespContractContractTypeDisplay ...
type GetCoreHRContractRespContractContractTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRContractRespContractCustomField ...
type GetCoreHRContractRespContractCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// GetCoreHRContractRespContractDurationType ...
type GetCoreHRContractRespContractDurationType struct {
	EnumName string                                              `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRContractRespContractDurationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRContractRespContractDurationTypeDisplay ...
type GetCoreHRContractRespContractDurationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// GetCoreHRContractRespContractSigningType ...
type GetCoreHRContractRespContractSigningType struct {
	EnumName string                                             `json:"enum_name,omitempty"` // 枚举值
	Display  []*GetCoreHRContractRespContractSigningTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// GetCoreHRContractRespContractSigningTypeDisplay ...
type GetCoreHRContractRespContractSigningTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHRContractResp ...
type getCoreHRContractResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *GetCoreHRContractResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
