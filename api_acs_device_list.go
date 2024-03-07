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

// GetACSDeviceList 使用该接口获取租户内所有门禁设备。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/acs-v1/device/list
// new doc: https://open.feishu.cn/document/server-docs/acs-v1/device/list
func (r *ACSService) GetACSDeviceList(ctx context.Context, request *GetACSDeviceListReq, options ...MethodOptionFunc) (*GetACSDeviceListResp, *Response, error) {
	if r.cli.mock.mockACSGetACSDeviceList != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] ACS#GetACSDeviceList mock enable")
		return r.cli.mock.mockACSGetACSDeviceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "ACS",
		API:                   "GetACSDeviceList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/acs/v1/devices",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getACSDeviceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockACSGetACSDeviceList mock ACSGetACSDeviceList method
func (r *Mock) MockACSGetACSDeviceList(f func(ctx context.Context, request *GetACSDeviceListReq, options ...MethodOptionFunc) (*GetACSDeviceListResp, *Response, error)) {
	r.mockACSGetACSDeviceList = f
}

// UnMockACSGetACSDeviceList un-mock ACSGetACSDeviceList method
func (r *Mock) UnMockACSGetACSDeviceList() {
	r.mockACSGetACSDeviceList = nil
}

// GetACSDeviceListReq ...
type GetACSDeviceListReq struct {
}

// GetACSDeviceListResp ...
type GetACSDeviceListResp struct {
	Items []*GetACSDeviceListRespItem `json:"items,omitempty"`
}

// GetACSDeviceListRespItem ...
type GetACSDeviceListRespItem struct {
	DeviceID   string `json:"device_id,omitempty"`   // 门禁设备 ID
	DeviceName string `json:"device_name,omitempty"` // 设备名称
	DeviceSn   string `json:"device_sn,omitempty"`   // 设备 SN 码
}

// getACSDeviceListResp ...
type getACSDeviceListResp struct {
	Code  int64                 `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                `json:"msg,omitempty"`  // 错误描述
	Data  *GetACSDeviceListResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
