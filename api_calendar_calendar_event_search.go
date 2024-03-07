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

// SearchCalendarEvent 该接口用于以用户身份搜索某日历下的相关日程。
//
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份必须对日历有reader、writer或owner权限（调用[获取日历](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/get)接口, role字段可查看权限）。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/search
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/calendar-event/search
func (r *CalendarService) SearchCalendarEvent(ctx context.Context, request *SearchCalendarEventReq, options ...MethodOptionFunc) (*SearchCalendarEventResp, *Response, error) {
	if r.cli.mock.mockCalendarSearchCalendarEvent != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#SearchCalendarEvent mock enable")
		return r.cli.mock.mockCalendarSearchCalendarEvent(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Calendar",
		API:                   "SearchCalendarEvent",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/calendar/v4/calendars/:calendar_id/events/search",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(searchCalendarEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarSearchCalendarEvent mock CalendarSearchCalendarEvent method
func (r *Mock) MockCalendarSearchCalendarEvent(f func(ctx context.Context, request *SearchCalendarEventReq, options ...MethodOptionFunc) (*SearchCalendarEventResp, *Response, error)) {
	r.mockCalendarSearchCalendarEvent = f
}

// UnMockCalendarSearchCalendarEvent un-mock CalendarSearchCalendarEvent method
func (r *Mock) UnMockCalendarSearchCalendarEvent() {
	r.mockCalendarSearchCalendarEvent = nil
}

// SearchCalendarEventReq ...
type SearchCalendarEventReq struct {
	CalendarID string                        `path:"calendar_id" json:"-"`   // 日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction), 示例值: "feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	UserIDType *IDType                       `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	PageToken  *string                       `query:"page_token" json:"-"`   // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: xxxxx
	PageSize   *int64                        `query:"page_size" json:"-"`    // 分页大小, 示例值: 10, 默认值: `20`, 最大值: `100`
	Query      string                        `json:"query,omitempty"`        // 搜索关键字, 示例值: "query words", 长度范围: `0` ～ `200` 字符
	Filter     *SearchCalendarEventReqFilter `json:"filter,omitempty"`       // 搜索过滤器
}

// SearchCalendarEventReqFilter ...
type SearchCalendarEventReqFilter struct {
	StartTime *SearchCalendarEventReqFilterStartTime `json:"start_time,omitempty"` // 搜索过滤项, 日程搜索区间的开始时间, 被搜索日程的事件必须与搜索区间有交集
	EndTime   *SearchCalendarEventReqFilterEndTime   `json:"end_time,omitempty"`   // 搜索过滤项, 日程搜索区间的结束时间, 被搜索日程的事件必须与搜索区间有交集
	UserIDs   []string                               `json:"user_ids,omitempty"`   // 搜索过滤项, 参与人的用户ID列表, 被搜索日程中必须包含至少一个其中的参与人。参见[用户相关的 ID 概念](https://open.feishu.cn/document/home/user-identity-introduction/introduction), 示例值: ["ou_e051986ab19f80d16b7b8d74f3f1235"]
	RoomIDs   []string                               `json:"room_ids,omitempty"`   // 搜索过滤项, 会议室ID列表, 被搜索日程中必须包含至少一个其中的会议室, 示例值: ["omm_eada1d61a550955240c28757e7dec3af"]
	ChatIDs   []string                               `json:"chat_ids,omitempty"`   // 搜索过滤项, 群ID列表, 被搜索日程的参与人中必须包含至少一个其中的群。参见[群ID 说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-id-description), 示例值: ["oc_a0553eda9014c201e6969b478895c230"]
}

// SearchCalendarEventReqFilterEndTime ...
type SearchCalendarEventReqFilterEndTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定, 示例值: "2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值: "1602504000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai, 示例值: "Asia/Shanghai"
}

// SearchCalendarEventReqFilterStartTime ...
type SearchCalendarEventReqFilterStartTime struct {
	Date      *string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定, 示例值: "2018-09-01"
	Timestamp *string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区), 示例值: "1602504000"
	Timezone  *string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai, 示例值: "Asia/Shanghai"
}

// SearchCalendarEventResp ...
type SearchCalendarEventResp struct {
	Items     []*SearchCalendarEventRespItem `json:"items,omitempty"`      // 搜索命中的日程列表
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
}

// SearchCalendarEventRespItem ...
type SearchCalendarEventRespItem struct {
	EventID             string                                     `json:"event_id,omitempty"`              // 日程ID。参见[日程ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/introduction)
	OrganizerCalendarID string                                     `json:"organizer_calendar_id,omitempty"` // 日程组织者日历ID。参见[日历ID说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar/introduction)
	Summary             string                                     `json:"summary,omitempty"`               // 日程标题
	Description         string                                     `json:"description,omitempty"`           // 日程描述；目前不支持编辑富文本描述, 如果日程描述通过客户端编辑过, 更新描述会导致富文本格式丢失
	StartTime           *SearchCalendarEventRespItemStartTime      `json:"start_time,omitempty"`            // 日程开始时间
	EndTime             *SearchCalendarEventRespItemEndTime        `json:"end_time,omitempty"`              // 日程结束时间
	Visibility          string                                     `json:"visibility,omitempty"`            // 日程公开范围, 新建日程默认为Default；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: default: 默认权限, 跟随日历权限, 默认仅向他人显示是否“忙碌”, public: 公开, 显示日程详情, private: 私密, 仅自己可见详情
	AttendeeAbility     string                                     `json:"attendee_ability,omitempty"`      // 参与人权限, 可选值有: none: 无法编辑日程、无法邀请其它参与人、无法查看参与人列表, can_see_others: 无法编辑日程、无法邀请其它参与人、可以查看参与人列表, can_invite_others: 无法编辑日程、可以邀请其它参与人、可以查看参与人列表, can_modify_event: 可以编辑日程、可以邀请其它参与人、可以查看参与人列表
	FreeBusyStatus      string                                     `json:"free_busy_status,omitempty"`      // 日程占用的忙闲状态, 新建日程默认为Busy；仅新建日程时对所有参与人生效, 之后修改该属性仅对当前身份生效, 可选值有: busy: 忙碌, free: 空闲
	Location            *SearchCalendarEventRespItemLocation       `json:"location,omitempty"`              // 日程地点
	Color               int64                                      `json:"color,omitempty"`                 // 日程颜色, 颜色RGB值的int32表示。仅对当前身份生效；客户端展示时会映射到色板上最接近的一种颜色；值为0或-1时默认跟随日历颜色。
	Reminders           []*SearchCalendarEventRespItemReminder     `json:"reminders,omitempty"`             // 日程提醒列表
	Recurrence          string                                     `json:"recurrence,omitempty"`            // 重复日程的重复性规则；参考[rfc5545](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10)；, 不支持COUNT和UNTIL同时出现；, 预定会议室重复日程长度不得超过两年。
	Status              string                                     `json:"status,omitempty"`                // 日程状态, 可选值有: tentative: 未回应, confirmed: 已确认, cancelled: 日程已取消
	IsException         bool                                       `json:"is_exception,omitempty"`          // 日程是否是一个重复日程的例外日程
	RecurringEventID    string                                     `json:"recurring_event_id,omitempty"`    // 例外日程的原重复日程的event_id
	EventOrganizer      *SearchCalendarEventRespItemEventOrganizer `json:"event_organizer,omitempty"`       // 日程组织者信息
	AppLink             string                                     `json:"app_link,omitempty"`              // 日程的app_link, 跳转到具体的某个日程
}

// SearchCalendarEventRespItemEndTime ...
type SearchCalendarEventRespItemEndTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// SearchCalendarEventRespItemEventOrganizer ...
type SearchCalendarEventRespItemEventOrganizer struct {
	UserID      string `json:"user_id,omitempty"`      // 日程组织者user ID
	DisplayName string `json:"display_name,omitempty"` // 日程组织者姓名
}

// SearchCalendarEventRespItemLocation ...
type SearchCalendarEventRespItemLocation struct {
	Name      string  `json:"name,omitempty"`      // 地点名称
	Address   string  `json:"address,omitempty"`   // 地点地址
	Latitude  float64 `json:"latitude,omitempty"`  // 地点坐标纬度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
	Longitude float64 `json:"longitude,omitempty"` // 地点坐标经度信息, 对于国内的地点, 采用GCJ-02标准, 海外地点采用WGS84标准
}

// SearchCalendarEventRespItemReminder ...
type SearchCalendarEventRespItemReminder struct {
	Minutes int64 `json:"minutes,omitempty"` // 日程提醒时间的偏移量, 正数时表示在日程开始前X分钟提醒, 负数时表示在日程开始后X分钟提醒, 新建或更新日程时传入该字段, 仅对当前身份生效, 不会对日程其他参与人生效。
}

// SearchCalendarEventRespItemStartTime ...
type SearchCalendarEventRespItemStartTime struct {
	Date      string `json:"date,omitempty"`      // 仅全天日程使用该字段, 如2018-09-01。需满足 RFC3339 格式。不能与 timestamp 同时指定
	Timestamp string `json:"timestamp,omitempty"` // 秒级时间戳, 如1602504000(表示2020/10/12 20:0:00 +8时区)
	Timezone  string `json:"timezone,omitempty"`  // 时区名称, 使用IANA Time Zone Database标准, 如Asia/Shanghai；全天日程时区固定为UTC, 非全天日程时区默认为Asia/Shanghai
}

// searchCalendarEventResp ...
type searchCalendarEventResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *SearchCalendarEventResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
