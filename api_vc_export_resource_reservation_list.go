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

// ExportVCResourceReservationList 导出会议室预定数据, 具体权限要求请参考「资源介绍」。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/resource_reservation_list
// new doc: https://open.feishu.cn/document/server-docs/vc-v1/export/resource_reservation_list
func (r *VCService) ExportVCResourceReservationList(ctx context.Context, request *ExportVCResourceReservationListReq, options ...MethodOptionFunc) (*ExportVCResourceReservationListResp, *Response, error) {
	if r.cli.mock.mockVCExportVCResourceReservationList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#ExportVCResourceReservationList mock enable")
		return r.cli.mock.mockVCExportVCResourceReservationList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "ExportVCResourceReservationList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/exports/resource_reservation_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(exportVCResourceReservationListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCExportVCResourceReservationList mock VCExportVCResourceReservationList method
func (r *Mock) MockVCExportVCResourceReservationList(f func(ctx context.Context, request *ExportVCResourceReservationListReq, options ...MethodOptionFunc) (*ExportVCResourceReservationListResp, *Response, error)) {
	r.mockVCExportVCResourceReservationList = f
}

// UnMockVCExportVCResourceReservationList un-mock VCExportVCResourceReservationList method
func (r *Mock) UnMockVCExportVCResourceReservationList() {
	r.mockVCExportVCResourceReservationList = nil
}

// ExportVCResourceReservationListReq ...
type ExportVCResourceReservationListReq struct {
	RoomLevelID string   `json:"room_level_id,omitempty"` // 会议室层级id, 示例值: "omb_57c9cc7d9a81e27e54c8fabfd02759e7"
	NeedTopic   *bool    `json:"need_topic,omitempty"`    // 是否展示会议主题, 示例值: true
	StartTime   string   `json:"start_time,omitempty"`    // 查询开始时间（unix时间, 单位sec）, 示例值: "1655276858"
	EndTime     string   `json:"end_time,omitempty"`      // 查询结束时间（unix时间, 单位sec）, 示例值: "1655276858"
	RoomIDs     []string `json:"room_ids,omitempty"`      // 待筛选的会议室id列表, 示例值: ["omm_eada1d61a550955240c28757e7dec3af"]
	IsExclude   *bool    `json:"is_exclude,omitempty"`    // 若为true表示导出room_ids范围外的会议室, 默认为false, 示例值: false
}

// ExportVCResourceReservationListResp ...
type ExportVCResourceReservationListResp struct {
	TaskID string `json:"task_id,omitempty"` // 任务id
}

// exportVCResourceReservationListResp ...
type exportVCResourceReservationListResp struct {
	Code int64                                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                               `json:"msg,omitempty"`  // 错误描述
	Data *ExportVCResourceReservationListResp `json:"data,omitempty"`
}
