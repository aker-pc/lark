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

// CreateDriveExportTask 创建导出任务, 将云文档导出为指定格式的本地文件, 目前支持新版文档、电子表格、多维表格和旧版文档。该接口为异步接口, 任务创建完成即刻返回, 并不会阻塞等待到任务执行成功, 因此需要结合[查询导出任务结果](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/get)接口获取导出结果。
//
// ::: note
// 请求用户需要对该文档有导出权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/export_task/create
func (r *DriveService) CreateDriveExportTask(ctx context.Context, request *CreateDriveExportTaskReq, options ...MethodOptionFunc) (*CreateDriveExportTaskResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveExportTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveExportTask mock enable")
		return r.cli.mock.mockDriveCreateDriveExportTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveExportTask",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/export_tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveExportTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveExportTask mock DriveCreateDriveExportTask method
func (r *Mock) MockDriveCreateDriveExportTask(f func(ctx context.Context, request *CreateDriveExportTaskReq, options ...MethodOptionFunc) (*CreateDriveExportTaskResp, *Response, error)) {
	r.mockDriveCreateDriveExportTask = f
}

// UnMockDriveCreateDriveExportTask un-mock DriveCreateDriveExportTask method
func (r *Mock) UnMockDriveCreateDriveExportTask() {
	r.mockDriveCreateDriveExportTask = nil
}

// CreateDriveExportTaskReq ...
type CreateDriveExportTaskReq struct {
	FileExtension string  `json:"file_extension,omitempty"` // 导出文件扩展名, 示例值: "csv", 可选值有: docx: Microsoft Word (DOCX) 格式, pdf: pdf 格式, xlsx: Microsoft Excel (XLSX) 格式, csv: csv 格式
	Token         string  `json:"token,omitempty"`          // 导出文档 token [获取文档 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "shtcnxe5OxxxxxxxSNdsJviENsk"
	Type          string  `json:"type,omitempty"`           // 导出文档类型 [文档类型说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#560bf735), 示例值: "sheet", 可选值有: doc: 旧版飞书云文档类型, 支持导出为 docx、pdf 格式, sheet: 飞书电子表格类型, 支持导出为 xlsx、csv 格式, bitable: 飞书多维表格类型, 支持导出为 xlsx、csv 格式, docx: 新版飞书云文档类型, 支持导出为 docx、pdf 格式
	SubID         *string `json:"sub_id,omitempty"`         // 导出子表ID, 仅当将电子表格/多维表格导出为 csv 时使用, [获取电子表格子表ID](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/query) 文档中的 sheet_id, [获取多维表格子表ID](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list) 文档中的 table_id, 示例值: "tblKz5D60T4JlfcT"
}

// CreateDriveExportTaskResp ...
type CreateDriveExportTaskResp struct {
	Ticket string `json:"ticket,omitempty"` // 导出任务ID
}

// createDriveExportTaskResp ...
type createDriveExportTaskResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveExportTaskResp `json:"data,omitempty"`
}
