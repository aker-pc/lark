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

// SearchCoreHRBankBranch 根据银行 ID、支行 ID 、支行名称查询银行信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-bank_branch/search
func (r *CoreHRService) SearchCoreHRBankBranch(ctx context.Context, request *SearchCoreHRBankBranchReq, options ...MethodOptionFunc) (*SearchCoreHRBankBranchResp, *Response, error) {
	if r.cli.mock.mockCoreHRSearchCoreHRBankBranch != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHR#SearchCoreHRBankBranch mock enable")
		return r.cli.mock.mockCoreHRSearchCoreHRBankBranch(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHR",
		API:                   "SearchCoreHRBankBranch",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v2/basic_info/bank_branchs/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(searchCoreHRBankBranchResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHRSearchCoreHRBankBranch mock CoreHRSearchCoreHRBankBranch method
func (r *Mock) MockCoreHRSearchCoreHRBankBranch(f func(ctx context.Context, request *SearchCoreHRBankBranchReq, options ...MethodOptionFunc) (*SearchCoreHRBankBranchResp, *Response, error)) {
	r.mockCoreHRSearchCoreHRBankBranch = f
}

// UnMockCoreHRSearchCoreHRBankBranch un-mock CoreHRSearchCoreHRBankBranch method
func (r *Mock) UnMockCoreHRSearchCoreHRBankBranch() {
	r.mockCoreHRSearchCoreHRBankBranch = nil
}

// SearchCoreHRBankBranchReq ...
type SearchCoreHRBankBranchReq struct {
	PageSize           int64    `query:"page_size" json:"-"`             // 分页大小, 最大 100, 示例值: 100, 取值范围: `1` ～ `100`
	PageToken          *string  `query:"page_token" json:"-"`            // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 6891251722631890445
	BankIDList         []string `json:"bank_id_list,omitempty"`          // 银行 ID 列表, 与「支行 ID 列表」、「支行名称列表」至少填写一项, 示例值: ["6891251722631891445"], 最大长度: `100`
	BankBranchIDList   []string `json:"bank_branch_id_list,omitempty"`   // 支行 ID 列表, 示例值: ["6891251722631891415"], 最大长度: `100`
	BankBranchNameList []string `json:"bank_branch_name_list,omitempty"` // 支行名称列表, 支持对支行名称精确搜索, 示例值: ["招商银行杭州未科支行"], 最大长度: `100`
	StatusList         []int64  `json:"status_list,omitempty"`           // 状态列表, 示例值: [1], 可选值有: 1: 生效, 0: 失效, 默认值: `[1]`, 最大长度: `2`
}

// SearchCoreHRBankBranchResp ...
type SearchCoreHRBankBranchResp struct {
	Items     []*SearchCoreHRBankBranchRespItem `json:"items,omitempty"`      // 查询的支行信息
	PageToken string                            `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                              `json:"has_more,omitempty"`   // 是否还有更多项
}

// SearchCoreHRBankBranchRespItem ...
type SearchCoreHRBankBranchRespItem struct {
	BankBranchID   string                                          `json:"bank_branch_id,omitempty"`   // 支行 ID
	BankBranchName []*SearchCoreHRBankBranchRespItemBankBranchName `json:"bank_branch_name,omitempty"` // 支行名称
	BankID         string                                          `json:"bank_id,omitempty"`          // 所属银行 ID, 可通过[查询银行信息](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/corehr-v2/basic_info-bank/search)接口查询
	Code           string                                          `json:"code,omitempty"`             // 金融分支机构编码
	SwiftCode      string                                          `json:"swift_code,omitempty"`       // 银行代码
	Status         int64                                           `json:"status,omitempty"`           // 状态, 可选值有: 1: 生效, 0: 失效
}

// SearchCoreHRBankBranchRespItemBankBranchName ...
type SearchCoreHRBankBranchRespItemBankBranchName struct {
	Lang  string `json:"lang,omitempty"`  // 语言
	Value string `json:"value,omitempty"` // 内容
}

// searchCoreHRBankBranchResp ...
type searchCoreHRBankBranchResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *SearchCoreHRBankBranchResp `json:"data,omitempty"`
}
