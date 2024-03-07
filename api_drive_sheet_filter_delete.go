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

// DeleteSheetFilter 删除子表的筛选
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/delete
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/spreadsheet-sheet-filter/delete
func (r *DriveService) DeleteSheetFilter(ctx context.Context, request *DeleteSheetFilterReq, options ...MethodOptionFunc) (*DeleteSheetFilterResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteSheetFilter != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DeleteSheetFilter mock enable")
		return r.cli.mock.mockDriveDeleteSheetFilter(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteSheetFilter",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteSheetFilterResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteSheetFilter mock DriveDeleteSheetFilter method
func (r *Mock) MockDriveDeleteSheetFilter(f func(ctx context.Context, request *DeleteSheetFilterReq, options ...MethodOptionFunc) (*DeleteSheetFilterResp, *Response, error)) {
	r.mockDriveDeleteSheetFilter = f
}

// UnMockDriveDeleteSheetFilter un-mock DriveDeleteSheetFilter method
func (r *Mock) UnMockDriveDeleteSheetFilter() {
	r.mockDriveDeleteSheetFilter = nil
}

// DeleteSheetFilterReq ...
type DeleteSheetFilterReq struct {
	SpreadSheetToken string `path:"spreadsheet_token" json:"-"` // 表格 token, 示例值: "shtcnmBA\*yGehy8"
	SheetID          string `path:"sheet_id" json:"-"`          // 子表 id, 示例值: "0b\**12"
}

// DeleteSheetFilterResp ...
type DeleteSheetFilterResp struct {
}

// deleteSheetFilterResp ...
type deleteSheetFilterResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *DeleteSheetFilterResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
