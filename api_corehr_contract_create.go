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

// CreateCoreHRContract 创建合同。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/contract/create
// new doc: https://open.feishu.cn/document/server-docs/corehr-v1/contract/create
func (r *CoreHRService) CreateCoreHRContract(ctx context.Context, request *CreateCoreHRContractReq, options ...MethodOptionFunc) (*CreateCoreHRContractResp, *Response, error) {
	if r.cli.mock.mockCoreHRCreateCoreHRContract != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] CoreHR#CreateCoreHRContract mock enable")
		return r.cli.mock.mockCoreHRCreateCoreHRContract(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "CreateCoreHRContract",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/contracts",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCoreHRContractResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRCreateCoreHRContract mock CoreHRCreateCoreHRContract method
func (r *Mock) MockCoreHRCreateCoreHRContract(f func(ctx context.Context, request *CreateCoreHRContractReq, options ...MethodOptionFunc) (*CreateCoreHRContractResp, *Response, error)) {
	r.mockCoreHRCreateCoreHRContract = f
}

// UnMockCoreHRCreateCoreHRContract un-mock CoreHRCreateCoreHRContract method
func (r *Mock) UnMockCoreHRCreateCoreHRContract() {
	r.mockCoreHRCreateCoreHRContract = nil
}

// CreateCoreHRContractReq ...
type CreateCoreHRContractReq struct {
	ClientToken         *string                               `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求, 示例值: 12454646
	EffectiveTime       string                                `json:"effective_time,omitempty"`         // 合同开始日期, 示例值: "2050-01-01 00:00:00"
	ExpirationTime      *string                               `json:"expiration_time,omitempty"`        // 实际结束日期, 示例值: "9999-12-31 23:59:59"
	EmploymentID        string                                `json:"employment_id,omitempty"`          // 雇佣 ID, 详细信息可通过[【搜索员工信息】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employee/search)接口查询获得, 示例值: "6893013238632416776"
	ContractType        *CreateCoreHRContractReqContractType  `json:"contract_type,omitempty"`          // 合同类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "contract_type"
	FirstPartyCompanyID string                                `json:"first_party_company_id,omitempty"` // 甲方, 引用Company的ID, 详细信息可通过[【查询单个公司】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/get)接口查询获得, 示例值: "6892686614112241165"
	PersonID            *string                               `json:"person_id,omitempty"`              // Person ID, 详细信息可通过[【查询单个个人信息】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/get)接口查询获得, 示例值: "151515151"
	CustomFields        []*CreateCoreHRContractReqCustomField `json:"custom_fields,omitempty"`          // 自定义字段
	DurationType        *CreateCoreHRContractReqDurationType  `json:"duration_type,omitempty"`          // 期限类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "duration_type", 示例值: fixed_term
	ContractEndDate     *string                               `json:"contract_end_date,omitempty"`      // 合同结束日期, 示例值: "2006-01-02"
	ContractNumber      *string                               `json:"contract_number,omitempty"`        // 合同编号, 示例值: "6919737965274990093"
	SigningType         *CreateCoreHRContractReqSigningType   `json:"signing_type,omitempty"`           // 签订类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "signing_type"
}

// CreateCoreHRContractReqContractType ...
type CreateCoreHRContractReqContractType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRContractReqCustomField ...
type CreateCoreHRContractReqCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名, 示例值: "name"
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05"), 示例值: "\"Sandy\""
}

// CreateCoreHRContractReqDurationType ...
type CreateCoreHRContractReqDurationType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRContractReqSigningType ...
type CreateCoreHRContractReqSigningType struct {
	EnumName string `json:"enum_name,omitempty"` // 枚举值, 示例值: "type_1"
}

// CreateCoreHRContractResp ...
type CreateCoreHRContractResp struct {
	Contract *CreateCoreHRContractRespContract `json:"contract,omitempty"` // 创建成功的合同信息
}

// CreateCoreHRContractRespContract ...
type CreateCoreHRContractRespContract struct {
	ID                  string                                         `json:"id,omitempty"`                     // 合同ID
	EffectiveTime       string                                         `json:"effective_time,omitempty"`         // 合同开始日期
	ExpirationTime      string                                         `json:"expiration_time,omitempty"`        // 实际结束日期
	EmploymentID        string                                         `json:"employment_id,omitempty"`          // 雇佣 ID, 详细信息可通过[【搜索员工信息】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/employee/search)接口查询获得
	ContractType        *CreateCoreHRContractRespContractContractType  `json:"contract_type,omitempty"`          // 合同类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "contract_type"
	FirstPartyCompanyID string                                         `json:"first_party_company_id,omitempty"` // 甲方, 引用Company的ID, 详细信息可通过[【查询单个公司】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/company/get)接口查询获得
	PersonID            string                                         `json:"person_id,omitempty"`              // Person ID, 详细信息可通过[【查询单个个人信息】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/get)接口查询获得
	CustomFields        []*CreateCoreHRContractRespContractCustomField `json:"custom_fields,omitempty"`          // 自定义字段
	DurationType        *CreateCoreHRContractRespContractDurationType  `json:"duration_type,omitempty"`          // 期限类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "duration_type"
	ContractEndDate     string                                         `json:"contract_end_date,omitempty"`      // 合同结束日期
	ContractNumber      string                                         `json:"contract_number,omitempty"`        // 合同编号
	SigningType         *CreateCoreHRContractRespContractSigningType   `json:"signing_type,omitempty"`           // 签订类型, 枚举值 api_name 可通过[【获取字段详情】](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/custom_field/get_by_param)接口查询, 查询参数如下: object_api_name = "contract", custom_api_name = "signing_type"
}

// CreateCoreHRContractRespContractContractType ...
type CreateCoreHRContractRespContractContractType struct {
	EnumName string                                                 `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRContractRespContractContractTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRContractRespContractContractTypeDisplay ...
type CreateCoreHRContractRespContractContractTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRContractRespContractCustomField ...
type CreateCoreHRContractRespContractCustomField struct {
	FieldName string `json:"field_name,omitempty"` // 字段名
	Value     string `json:"value,omitempty"`      // 字段值, 是json转义后的字符串, 根据元数据定义不同, 字段格式不同(如123, 123.23, "true", [\"id1\", \"id2\"], "2006-01-02 15:04:05")
}

// CreateCoreHRContractRespContractDurationType ...
type CreateCoreHRContractRespContractDurationType struct {
	EnumName string                                                 `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRContractRespContractDurationTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRContractRespContractDurationTypeDisplay ...
type CreateCoreHRContractRespContractDurationTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// CreateCoreHRContractRespContractSigningType ...
type CreateCoreHRContractRespContractSigningType struct {
	EnumName string                                                `json:"enum_name,omitempty"` // 枚举值
	Display  []*CreateCoreHRContractRespContractSigningTypeDisplay `json:"display,omitempty"`   // 枚举多语展示
}

// CreateCoreHRContractRespContractSigningTypeDisplay ...
type CreateCoreHRContractRespContractSigningTypeDisplay struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// createCoreHRContractResp ...
type createCoreHRContractResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCoreHRContractResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
