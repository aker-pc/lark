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

// GetTaskReminderList 返回提醒时间列表, 支持分页, 最大值为50。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task-reminder/list
// new doc: https://open.feishu.cn/document/server-docs/task-v1/task-reminder/list
func (r *TaskService) GetTaskReminderList(ctx context.Context, request *GetTaskReminderListReq, options ...MethodOptionFunc) (*GetTaskReminderListResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskReminderList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#GetTaskReminderList mock enable")
		return r.cli.mock.mockTaskGetTaskReminderList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskReminderList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks/:task_id/reminders",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskReminderListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTaskReminderList mock TaskGetTaskReminderList method
func (r *Mock) MockTaskGetTaskReminderList(f func(ctx context.Context, request *GetTaskReminderListReq, options ...MethodOptionFunc) (*GetTaskReminderListResp, *Response, error)) {
	r.mockTaskGetTaskReminderList = f
}

// UnMockTaskGetTaskReminderList un-mock TaskGetTaskReminderList method
func (r *Mock) UnMockTaskGetTaskReminderList() {
	r.mockTaskGetTaskReminderList = nil
}

// GetTaskReminderListReq ...
type GetTaskReminderListReq struct {
	TaskID    string  `path:"task_id" json:"-"`     // 任务 ID, 示例值: "0d38e26e-190a-49e9-93a2-35067763ed1f"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 50, 默认值: `50`, 取值范围: `0` ～ `50`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: 「填写上次返回的page_token」
}

// GetTaskReminderListResp ...
type GetTaskReminderListResp struct {
	Items     []*GetTaskReminderListRespItem `json:"items,omitempty"`      // 返回提醒时间设置列表
	PageToken string                         `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetTaskReminderListRespItem ...
type GetTaskReminderListRespItem struct {
	ID                 string `json:"id,omitempty"`                   // 提醒时间设置的 ID（在删除时候需要使用这个）
	RelativeFireMinute int64  `json:"relative_fire_minute,omitempty"` // 相对于截止时间的提醒时间（如提前 30 分钟, 截止时间后 30 分钟, 则为 -30） 任务没有截止时间则为全天任务(截止时间为0)
}

// getTaskReminderListResp ...
type getTaskReminderListResp struct {
	Code  int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                   `json:"msg,omitempty"`  // 错误描述
	Data  *GetTaskReminderListResp `json:"data,omitempty"`
	Error *ErrorDetail             `json:"error,omitempty"`
}
