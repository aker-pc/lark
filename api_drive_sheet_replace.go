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

// ReplaceSheet 按照指定的条件查找子表的某个范围内的数据符合条件的单元格并替换值, 返回替换成功的单元格位置。一次请求最多允许替换5000个单元格, 如果超过请将range缩小范围再操作。请求体中的 range、find、replacement 字段必填。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/replace
// new doc: https://open.feishu.cn/document/server-docs/docs/sheets-v3/data-operation/replace
func (r *DriveService) ReplaceSheet(ctx context.Context, request *ReplaceSheetReq, options ...MethodOptionFunc) (*ReplaceSheetResp, *Response, error) {
	if r.cli.mock.mockDriveReplaceSheet != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#ReplaceSheet mock enable")
		return r.cli.mock.mockDriveReplaceSheet(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "ReplaceSheet",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(replaceSheetResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveReplaceSheet mock DriveReplaceSheet method
func (r *Mock) MockDriveReplaceSheet(f func(ctx context.Context, request *ReplaceSheetReq, options ...MethodOptionFunc) (*ReplaceSheetResp, *Response, error)) {
	r.mockDriveReplaceSheet = f
}

// UnMockDriveReplaceSheet un-mock DriveReplaceSheet method
func (r *Mock) UnMockDriveReplaceSheet() {
	r.mockDriveReplaceSheet = nil
}

// ReplaceSheetReq ...
type ReplaceSheetReq struct {
	SpreadSheetToken string                        `path:"spreadsheet_token" json:"-"` // Spreadsheet token, 示例值: "shtcnmBA*yGehy8"
	SheetID          string                        `path:"sheet_id" json:"-"`          // Sheet id, 示例值: "0b**12"
	FindCondition    *ReplaceSheetReqFindCondition `json:"find_condition,omitempty"`   // 查找条件
	Find             string                        `json:"find,omitempty"`             // 查找的字符串, 示例值: "hello"
	Replacement      string                        `json:"replacement,omitempty"`      // 替换的字符串, 示例值: "world"
}

// ReplaceSheetReqFindCondition ...
type ReplaceSheetReqFindCondition struct {
	Range           string `json:"range,omitempty"`             // 查找范围, 参考 [名词解释 Range](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview), 示例值: "PNIfrm!A1:C5"
	MatchCase       *bool  `json:"match_case,omitempty"`        // 是否忽略大小写, 默认为 false, `true`: 表示忽略字符串中字母大小写差异, `false`: 表示区分字符串中字母大小写, 示例值: true
	MatchEntireCell *bool  `json:"match_entire_cell,omitempty"` // 是否完全匹配整个单元格, 默认值为 false, `true`: 表示完全匹配单元格, 比如 find 取值为 "hello", 则单元格中的内容必须为 "hello", `false`: 表示允许部分匹配单元格, 比如 find 取值为 "hello", 则单元格中的内容包含 "hello" 即可, 示例值: false
	SearchByRegex   *bool  `json:"search_by_regex,omitempty"`   // 是否为正则匹配, 默认值为 false, `true`: 表示使用正则匹配, `false`: 表示不使用正则匹配, 示例值: false
	IncludeFormulas *bool  `json:"include_formulas,omitempty"`  // 是否仅搜索单元格公式, 默认值为 false, `true`: 表示仅搜索单元格公式, `false`: 表示仅搜索单元格内容, 示例值: false
}

// ReplaceSheetResp ...
type ReplaceSheetResp struct {
	ReplaceResult *ReplaceSheetRespReplaceResult `json:"replace_result,omitempty"` // 符合查找条件并替换的单元格信息
}

// ReplaceSheetRespReplaceResult ...
type ReplaceSheetRespReplaceResult struct {
	MatchedCells        []string `json:"matched_cells,omitempty"`         // 符合查找条件的单元格数组, 不包含公式, 例如["A1", "A2"...]
	MatchedFormulaCells []string `json:"matched_formula_cells,omitempty"` // 符合查找条件的含有公式的单元格数组, 例如["B3", "H7"...]
	RowsCount           int64    `json:"rows_count,omitempty"`            // 符合查找条件的总行数
}

// replaceSheetResp ...
type replaceSheetResp struct {
	Code  int64             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string            `json:"msg,omitempty"`  // 错误描述
	Data  *ReplaceSheetResp `json:"data,omitempty"`
	Error *ErrorDetail      `json:"error,omitempty"`
}
