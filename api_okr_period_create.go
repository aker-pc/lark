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

// CreateOKRPeriod 根据周期规则创建一个 OKR 周期。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/okr-v1/period/create
// new doc: https://open.feishu.cn/document/server-docs/okr-v1/period/create
func (r *OKRService) CreateOKRPeriod(ctx context.Context, request *CreateOKRPeriodReq, options ...MethodOptionFunc) (*CreateOKRPeriodResp, *Response, error) {
	if r.cli.mock.mockOKRCreateOKRPeriod != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] OKR#CreateOKRPeriod mock enable")
		return r.cli.mock.mockOKRCreateOKRPeriod(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "OKR",
		API:                   "CreateOKRPeriod",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/okr/v1/periods",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createOKRPeriodResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockOKRCreateOKRPeriod mock OKRCreateOKRPeriod method
func (r *Mock) MockOKRCreateOKRPeriod(f func(ctx context.Context, request *CreateOKRPeriodReq, options ...MethodOptionFunc) (*CreateOKRPeriodResp, *Response, error)) {
	r.mockOKRCreateOKRPeriod = f
}

// UnMockOKRCreateOKRPeriod un-mock OKRCreateOKRPeriod method
func (r *Mock) UnMockOKRCreateOKRPeriod() {
	r.mockOKRCreateOKRPeriod = nil
}

// CreateOKRPeriodReq ...
type CreateOKRPeriodReq struct {
	PeriodRuleID string `json:"period_rule_id,omitempty"` // 周期规则 id, 示例值: "6969864184272078374"
	StartMonth   string `json:"start_month,omitempty"`    // 周期起始年月, 示例值: "2022-01"
}

// CreateOKRPeriodResp ...
type CreateOKRPeriodResp struct {
	PeriodID   string `json:"period_id,omitempty"`   // 周期id
	StartMonth string `json:"start_month,omitempty"` // 周期起始年月
	EndMonth   string `json:"end_month,omitempty"`   // 周期结束年月
}

// createOKRPeriodResp ...
type createOKRPeriodResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *CreateOKRPeriodResp `json:"data,omitempty"`
}
