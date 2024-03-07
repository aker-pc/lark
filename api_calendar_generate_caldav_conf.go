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

// GenerateCaldavConf 用于为当前用户生成一个CalDAV账号密码, 用于将飞书日历信息同步到本地设备日历。
//
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-enable-bot-ability)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/setting/generate_caldav_conf
// new doc: https://open.feishu.cn/document/server-docs/calendar-v4/setting/generate_caldav_conf
func (r *CalendarService) GenerateCaldavConf(ctx context.Context, request *GenerateCaldavConfReq, options ...MethodOptionFunc) (*GenerateCaldavConfResp, *Response, error) {
	if r.cli.mock.mockCalendarGenerateCaldavConf != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Calendar#GenerateCaldavConf mock enable")
		return r.cli.mock.mockCalendarGenerateCaldavConf(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Calendar",
		API:                 "GenerateCaldavConf",
		Method:              "POST",
		URL:                 r.cli.openBaseURL + "/open-apis/calendar/v4/settings/generate_caldav_conf",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(generateCaldavConfResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCalendarGenerateCaldavConf mock CalendarGenerateCaldavConf method
func (r *Mock) MockCalendarGenerateCaldavConf(f func(ctx context.Context, request *GenerateCaldavConfReq, options ...MethodOptionFunc) (*GenerateCaldavConfResp, *Response, error)) {
	r.mockCalendarGenerateCaldavConf = f
}

// UnMockCalendarGenerateCaldavConf un-mock CalendarGenerateCaldavConf method
func (r *Mock) UnMockCalendarGenerateCaldavConf() {
	r.mockCalendarGenerateCaldavConf = nil
}

// GenerateCaldavConfReq ...
type GenerateCaldavConfReq struct {
	DeviceName *string `json:"device_name,omitempty"` // 需要同步日历的设备名, 在日历中展示用来管理密码, 示例值: "iPhone", 最大长度: `100` 字符
}

// GenerateCaldavConfResp ...
type GenerateCaldavConfResp struct {
	Password      string `json:"password,omitempty"`       // caldav密码
	UserName      string `json:"user_name,omitempty"`      // caldav用户名
	ServerAddress string `json:"server_address,omitempty"` // 服务器地址
	DeviceName    string `json:"device_name,omitempty"`    // 设备名
}

// generateCaldavConfResp ...
type generateCaldavConfResp struct {
	Code  int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                  `json:"msg,omitempty"`  // 错误描述
	Data  *GenerateCaldavConfResp `json:"data,omitempty"`
	Error *ErrorDetail            `json:"error,omitempty"`
}
