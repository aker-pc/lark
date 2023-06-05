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

// GetCoreHrTransferReasonList 获取异动原因列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/transfer_reason/query
func (r *CoreHrService) GetCoreHrTransferReasonList(ctx context.Context, request *GetCoreHrTransferReasonListReq, options ...MethodOptionFunc) (*GetCoreHrTransferReasonListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrTransferReasonList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrTransferReasonList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrTransferReasonList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrTransferReasonList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/transfer_reasons/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrTransferReasonListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrTransferReasonList mock CoreHrGetCoreHrTransferReasonList method
func (r *Mock) MockCoreHrGetCoreHrTransferReasonList(f func(ctx context.Context, request *GetCoreHrTransferReasonListReq, options ...MethodOptionFunc) (*GetCoreHrTransferReasonListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrTransferReasonList = f
}

// UnMockCoreHrGetCoreHrTransferReasonList un-mock CoreHrGetCoreHrTransferReasonList method
func (r *Mock) UnMockCoreHrGetCoreHrTransferReasonList() {
	r.mockCoreHrGetCoreHrTransferReasonList = nil
}

// GetCoreHrTransferReasonListReq ...
type GetCoreHrTransferReasonListReq struct {
	Active                         *bool    `query:"active" json:"-"`                            // 异动原因状态, 示例值: true
	TransferReasonUniqueIdentifier []string `query:"transfer_reason_unique_identifier" json:"-"` // 异动原因唯一标识, 多条时最多数量为10, 示例值: voluntary_transfer, 最大长度: `10`
}

// GetCoreHrTransferReasonListResp ...
type GetCoreHrTransferReasonListResp struct {
	Items []*GetCoreHrTransferReasonListRespItem `json:"items,omitempty"` // 异动原因列表
}

// GetCoreHrTransferReasonListRespItem ...
type GetCoreHrTransferReasonListRespItem struct {
	TransferReasonUniqueIdentifier       string                                     `json:"transfer_reason_unique_identifier,omitempty"`        // 异动原因唯一标识
	Name                                 []*GetCoreHrTransferReasonListRespItemName `json:"name,omitempty"`                                     // 异动原因的名称信息
	Active                               bool                                       `json:"active,omitempty"`                                   // 异动原因状态
	ParentTransferReasonUniqueIdentifier string                                     `json:"parent_transfer_reason_unique_identifier,omitempty"` // 上级异动原因唯一标识
	CreatedTime                          string                                     `json:"created_time,omitempty"`                             // 创建时间
	UpdatedTime                          string                                     `json:"updated_time,omitempty"`                             // 更新时间
}

// GetCoreHrTransferReasonListRespItemName ...
type GetCoreHrTransferReasonListRespItemName struct {
	Lang  string `json:"lang,omitempty"`  // 名称信息的语言
	Value string `json:"value,omitempty"` // 名称信息的内容
}

// getCoreHrTransferReasonListResp ...
type getCoreHrTransferReasonListResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrTransferReasonListResp `json:"data,omitempty"`
}
