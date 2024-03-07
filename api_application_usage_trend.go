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

// GetApplicationUsageTrend 查询应用在指定时间段内企业员工的使用趋势信息。
//
// :::warning
// 此接口目前仅支持小程序的使用情况查询, 不支持网页应用和机器人应用的使用情况查询;仅支持查询自建应用, 不支持查询商店应用
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITN0YjLyUDN24iM1QjN
func (r *ApplicationService) GetApplicationUsageTrend(ctx context.Context, request *GetApplicationUsageTrendReq, options ...MethodOptionFunc) (*GetApplicationUsageTrendResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUsageTrend != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUsageTrend mock enable")
		return r.cli.mock.mockApplicationGetApplicationUsageTrend(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUsageTrend",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v1/app_usage_trend",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUsageTrendResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationUsageTrend mock ApplicationGetApplicationUsageTrend method
func (r *Mock) MockApplicationGetApplicationUsageTrend(f func(ctx context.Context, request *GetApplicationUsageTrendReq, options ...MethodOptionFunc) (*GetApplicationUsageTrendResp, *Response, error)) {
	r.mockApplicationGetApplicationUsageTrend = f
}

// UnMockApplicationGetApplicationUsageTrend un-mock ApplicationGetApplicationUsageTrend method
func (r *Mock) UnMockApplicationGetApplicationUsageTrend() {
	r.mockApplicationGetApplicationUsageTrend = nil
}

// GetApplicationUsageTrendReq ...
type GetApplicationUsageTrendReq struct {
	AppID        string                               `json:"app_id,omitempty"`        // 目标应用的 ID, 支持自建应用
	Ability      string                               `json:"ability,omitempty"`       // 应用能力, mp: 小程序
	TimeStart    int64                                `json:"time_start,omitempty"`    // 起始时间戳（秒）, 时间跨度最长支持180天
	TimeEnd      int64                                `json:"time_end,omitempty"`      // 截止时间戳（秒）, 时间跨度最长支持180天
	TimeInterval int64                                `json:"time_interval,omitempty"` // 步长（秒）, 最小步长为60秒, 需满足(${time_end} - ${time_start}) / ${time_interval} <= 2*24*60
	Filters      []*GetApplicationUsageTrendReqFilter `json:"filters,omitempty"`       // 过滤条件
}

// GetApplicationUsageTrendReqFilter ...
type GetApplicationUsageTrendReqFilter struct {
	Key   string `json:"key,omitempty"`   // 过滤字段, 支持`department_id
	Op    string `json:"op,omitempty"`    // 过滤操作, 支持`in`、`=
	Value string `json:"value,omitempty"` // 过滤字段值, 多个使用英文逗号分隔
}

// GetApplicationUsageTrendResp ...
type GetApplicationUsageTrendResp struct {
	Item map[string]*GetApplicationUsageTrendRespItem `json:"item,omitempty"` // 返回项
}

// GetApplicationUsageTrendRespItem ...
type GetApplicationUsageTrendRespItem struct {
	Trends []*GetApplicationUsageTrendRespItemTrend `json:"trends,omitempty"` // 趋势数据
}

// GetApplicationUsageTrendRespItemTrend ...
type GetApplicationUsageTrendRespItemTrend struct {
	Timestamp int64 `json:"timestamp,omitempty"` // 时间戳
	Pv        int64 `json:"pv,omitempty"`        // 应用使用pv
	Uv        int64 `json:"uv,omitempty"`        // 应用使用uv
}

// getApplicationUsageTrendResp ...
type getApplicationUsageTrendResp struct {
	Code  int64                         `json:"code,omitempty"` // 返回码, 非0表示失败
	Msg   string                        `json:"msg,omitempty"`  // 返回码的描述
	Data  *GetApplicationUsageTrendResp `json:"data,omitempty"` // 返回的业务信息, 仅code = 0时有效
	Error *ErrorDetail                  `json:"error,omitempty"`
}
