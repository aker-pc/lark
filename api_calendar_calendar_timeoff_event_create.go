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

// CreateCalendarTimeoffEvent 为指定用户创建一个请假日程, 可以是一个普通请假日程, 也可以是一个全天日程。
//
// 创建请假日程后, 会在相应时间内, 在用户个人签名页展示请假信息。
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/create
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/timeoff_event/create
func (r *CalendarService) CreateCalendarTimeoffEvent(ctx context.Context, request *CreateCalendarTimeoffEventReq, options ...MethodOptionFunc) (*CreateCalendarTimeoffEventResp, *Response, error) {
	if r.cli.mock.mockCalendarCreateCalendarTimeoffEvent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#CreateCalendarTimeoffEvent mock enable")
		return r.cli.mock.mockCalendarCreateCalendarTimeoffEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "CreateCalendarTimeoffEvent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/timeoff_events",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createCalendarTimeoffEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarCreateCalendarTimeoffEvent mock CalendarCreateCalendarTimeoffEvent method
func (r *Mock) MockCalendarCreateCalendarTimeoffEvent(f func(ctx context.Context, request *CreateCalendarTimeoffEventReq, options ...MethodOptionFunc) (*CreateCalendarTimeoffEventResp, *Response, error)) {
	r.mockCalendarCreateCalendarTimeoffEvent = f
}

// UnMockCalendarCreateCalendarTimeoffEvent un-mock CalendarCreateCalendarTimeoffEvent method
func (r *Mock) UnMockCalendarCreateCalendarTimeoffEvent() {
	r.mockCalendarCreateCalendarTimeoffEvent = nil
}

// CreateCalendarTimeoffEventReq ...
type CreateCalendarTimeoffEventReq struct {
	UserIDType  *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserID      string  `json:"user_id,omitempty"`      // 用户id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: "ou_XXXXXXXXXX"
	Timezone    string  `json:"timezone,omitempty"`     // 时区, 示例值: "Asia/Shanghai"
	StartTime   string  `json:"start_time,omitempty"`   // 休假开始时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式, 其它格式无效；, 时间戳格式是按小时休假日程, 日期格式是全天休假日程；, start_time与end_time格式需保持一致, 否则无效, 示例值: "2021-01-01"
	EndTime     string  `json:"end_time,omitempty"`     // 休假结束时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式, 其它格式无效；, 时间戳格式是按小时休假日程, 日期格式是全天休假日程；, start_time与end_time格式需保持一致, 否则无效, 示例值: "2021-01-01"
	Title       *string `json:"title,omitempty"`        // 自定义请假日程标题, 没有设置则为默认日程标题, 示例值: "请假中(全天) / 1-Day Time Off"
	Description *string `json:"description,omitempty"`  // 自定义请假日程描述, 没有设置则为默认日程描述, 示例值: "若删除此日程, 飞书中相应的“请假”标签将自动消失, 而请假系统中的休假申请不会被撤销。"
}

// CreateCalendarTimeoffEventResp ...
type CreateCalendarTimeoffEventResp struct {
	TimeoffEventID string `json:"timeoff_event_id,omitempty"` // 请假日程ID。参见[请假日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/introduction#b6611a02)
	UserID         string `json:"user_id,omitempty"`          // 用户id, 参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction)
	Timezone       string `json:"timezone,omitempty"`         // 时区
	StartTime      string `json:"start_time,omitempty"`       // 休假开始时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式, 其它格式无效；, 时间戳格式是按小时休假日程, 日期格式是全天休假日程；, start_time与end_time格式需保持一致, 否则无效。
	EndTime        string `json:"end_time,omitempty"`         // 休假结束时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式, 其它格式无效；, 时间戳格式是按小时休假日程, 日期格式是全天休假日程；, start_time与end_time格式需保持一致, 否则无效。
	Title          string `json:"title,omitempty"`            // 自定义请假日程标题, 没有设置则为默认日程标题
	Description    string `json:"description,omitempty"`      // 自定义请假日程描述, 没有设置则为默认日程描述
}

// createCalendarTimeoffEventResp ...
type createCalendarTimeoffEventResp struct {
	Code  int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                          `json:"msg,omitempty"`  // 错误描述
	Data  *CreateCalendarTimeoffEventResp `json:"data,omitempty"`
	Error *ErrorDetail                    `json:"error,omitempty"`
}
