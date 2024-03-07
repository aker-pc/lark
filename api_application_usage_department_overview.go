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

// GetApplicationUsageDepartmentOverview 查看应用在某一天/某一周/某一个月的使用数据, 可以根据部门做多层子部门的筛选
//
// 1. 仅支持企业版/旗舰版租户使用
// 2. 一般每天早上10点产出前一天的数据
// 3. 已经支持的指标包括: 应用的活跃用户数、累计用户数、新增用户数、访问页面数、打开次数
// 4. 按照部门查看数据时, 可以分别展示当前部门以及其子部门的使用情况
// 5. 如果查询的部门在查询日期没有使用过应用, 只返回指标: 应用的活跃用户数指标
// 6. 数据从飞书4.10版本开始统计, 使用飞书版本4.10及以下版本的用户数据不会被统计到
// 7. 调用频控为100次/分
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/application-v6/application-app_usage/department_overview
// new doc: https://open.feishu.cn/document/server-docs/application-v6/app-usage/department_overview
func (r *ApplicationService) GetApplicationUsageDepartmentOverview(ctx context.Context, request *GetApplicationUsageDepartmentOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageDepartmentOverviewResp, *Response, error) {
	if r.cli.mock.mockApplicationGetApplicationUsageDepartmentOverview != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Application#GetApplicationUsageDepartmentOverview mock enable")
		return r.cli.mock.mockApplicationGetApplicationUsageDepartmentOverview(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Application",
		API:                   "GetApplicationUsageDepartmentOverview",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/application/v6/applications/:app_id/app_usage/department_overview",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApplicationUsageDepartmentOverviewResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApplicationGetApplicationUsageDepartmentOverview mock ApplicationGetApplicationUsageDepartmentOverview method
func (r *Mock) MockApplicationGetApplicationUsageDepartmentOverview(f func(ctx context.Context, request *GetApplicationUsageDepartmentOverviewReq, options ...MethodOptionFunc) (*GetApplicationUsageDepartmentOverviewResp, *Response, error)) {
	r.mockApplicationGetApplicationUsageDepartmentOverview = f
}

// UnMockApplicationGetApplicationUsageDepartmentOverview un-mock ApplicationGetApplicationUsageDepartmentOverview method
func (r *Mock) UnMockApplicationGetApplicationUsageDepartmentOverview() {
	r.mockApplicationGetApplicationUsageDepartmentOverview = nil
}

// GetApplicationUsageDepartmentOverviewReq ...
type GetApplicationUsageDepartmentOverviewReq struct {
	AppID            string            `path:"app_id" json:"-"`              // 目标应用 ID, 示例值: "cli_9f115af860f7901b"
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 调用中使用的部门ID的类型, 示例值: open_department_id, 可选值有: department_id: 以自定义department_id来标识部门, open_department_id: 以open_department_id来标识部门, 默认值: `open_department_id`
	Date             string            `json:"date,omitempty"`               // 查询日期, 格式为yyyy-mm-dd, 若cycle_type为1, date可以为任何自然日；若cycle_type为2, 则输入的date必须为周一； 若cycle_type为3, 则输入的date必须为每月1号, 示例值: "2021-07-08"
	CycleType        int64             `json:"cycle_type,omitempty"`         // 活跃周期的统计类型, 示例值: 1, 可选值有: 1: 日活, 指自然日, 返回当前日期所在日的数据, 2: 周活, 指自然周, 返回当前日期所在周的数据。若到查询时当周还没结束, 则返回周一到当前日期的数值。例如在2021/7/15 查询2021/7/5 这一周的数据, 则代表的是2021/7/5 ~ 2021/7/11。但若是在2021/7/8 查询2021/7/5 这一周的数据, 则返回的是2021/7/5 ~ 2021/7/7 的数据, 3: 月活, 指自然月, 返回当前日期所在月的数据。若不满一个月则返回当月1日到截止日期前的数据。例如在2021/8/15 查询 7月的数据, 则代表2021/7/1~2021/7/31。 若在2021/8/15 查询8月的数据, 则代表2021/8/1~2021/8/14的数据
	DepartmentID     *string           `json:"department_id,omitempty"`      // 查询的部门id, 获取方法可参考[部门ID概述](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/field-overview), 若部门id为空, 则返回当前租户的使用数据；若填写部门id, 则返回当前部门的使用数据（包含子部门的用户） 以及多级子部门的使用数据, 若路径参数中department_id_type为空或者为open_department_id, 则此处应该填写部门的 open_department_id；若路径参数中department_id_type为department_id, 则此处应该填写部门的 department_id, 若不填写则返回整个租户的数据, 示例值: "od-4e6ac4d14bcd5071a37a39de902c7141"
	Recursion        *int64            `json:"recursion,omitempty"`          // 是否需要查询部门下多层子部门的数据。未设置或为0时, 仅查询department_id对应的部门。设置为n时, 查询department_id及其n级子部门的数据。仅在department_id参数传递时有效, 最大值为4, 示例值: 0, 默认值: `0`, 取值范围: `0` ～ `4`
	PageSize         *int64            `json:"page_size,omitempty"`          // 分页大小, 取值范围 1~20, 示例值: 10, 默认值: `10`, 取值范围: `1` ～ `20`
	PageToken        *string           `json:"page_token,omitempty"`         // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "new-1a8f509162ca3c95405838d05ccded09"
}

// GetApplicationUsageDepartmentOverviewResp ...
type GetApplicationUsageDepartmentOverviewResp struct {
	HasMore   bool                                             `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                                           `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	Items     []*GetApplicationUsageDepartmentOverviewRespItem `json:"items,omitempty"`      // 部门内员工使用应用的概览数据
}

// GetApplicationUsageDepartmentOverviewRespItem ...
type GetApplicationUsageDepartmentOverviewRespItem struct {
	DepartmentID string                                                 `json:"department_id,omitempty"` // 租户内部门的唯一标识, ID值与查询参数中的department_id_type 对应。
	App          []*GetApplicationUsageDepartmentOverviewRespItemApp    `json:"app,omitempty"`           // 应用整体使用情况, 指标值包括: uv: 活跃用户数, total_users: 累计用户数, new_users: 新增用户数, pv: 在应用（小程序或网页）中访问的页面数, lifecycle: 打开应用（小程序或网页）的次数
	Gadget       []*GetApplicationUsageDepartmentOverviewRespItemGadget `json:"gadget,omitempty"`        // 小程序使用情况, 没有小程序形态时为null, 指标值包括: uv（小程序活跃用户数）、pv（用户在小程序中的访问页面数）、lifecycle（小程序的打开次数）
	Webapp       []*GetApplicationUsageDepartmentOverviewRespItemWebapp `json:"webapp,omitempty"`        // 网页应用使用情况, 没有网页应用形态时为null, 指标值包括: uv（网页应用活跃用户数）、pv（用户在网页应用中的访问页面数）、lifecycle（网页应用的打开次数）
	Bot          []*GetApplicationUsageDepartmentOverviewRespItemBot    `json:"bot,omitempty"`           // 机器人使用情况, 没有机器人形态时为null, 指标值包括: uv（机器人的活跃用户数）
}

// GetApplicationUsageDepartmentOverviewRespItemApp ...
type GetApplicationUsageDepartmentOverviewRespItemApp struct {
	MetricName  string `json:"metric_name,omitempty"`  // 指标名称
	MetricValue int64  `json:"metric_value,omitempty"` // 指标值
}

// GetApplicationUsageDepartmentOverviewRespItemBot ...
type GetApplicationUsageDepartmentOverviewRespItemBot struct {
	MetricName  string `json:"metric_name,omitempty"`  // 指标名称
	MetricValue int64  `json:"metric_value,omitempty"` // 指标值
}

// GetApplicationUsageDepartmentOverviewRespItemGadget ...
type GetApplicationUsageDepartmentOverviewRespItemGadget struct {
	MetricName  string `json:"metric_name,omitempty"`  // 指标名称
	MetricValue int64  `json:"metric_value,omitempty"` // 指标值
}

// GetApplicationUsageDepartmentOverviewRespItemWebapp ...
type GetApplicationUsageDepartmentOverviewRespItemWebapp struct {
	MetricName  string `json:"metric_name,omitempty"`  // 指标名称
	MetricValue int64  `json:"metric_value,omitempty"` // 指标值
}

// getApplicationUsageDepartmentOverviewResp ...
type getApplicationUsageDepartmentOverviewResp struct {
	Code  int64                                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                                     `json:"msg,omitempty"`  // 错误描述
	Data  *GetApplicationUsageDepartmentOverviewResp `json:"data,omitempty"`
	Error *ErrorDetail                               `json:"error,omitempty"`
}
