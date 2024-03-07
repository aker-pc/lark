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

// CopyBitableApp 复制一个多维表格, 可以指定复制到某个有权限的文件夹下
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/copy
// new doc: https://open.feishu.cn/document/server-docs/docs/bitable-v1/app/copy
func (r *BitableService) CopyBitableApp(ctx context.Context, request *CopyBitableAppReq, options ...MethodOptionFunc) (*CopyBitableAppResp, *Response, error) {
	if r.cli.mock.mockBitableCopyBitableApp != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Bitable#CopyBitableApp mock enable")
		return r.cli.mock.mockBitableCopyBitableApp(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CopyBitableApp",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/copy",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(copyBitableAppResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCopyBitableApp mock BitableCopyBitableApp method
func (r *Mock) MockBitableCopyBitableApp(f func(ctx context.Context, request *CopyBitableAppReq, options ...MethodOptionFunc) (*CopyBitableAppResp, *Response, error)) {
	r.mockBitableCopyBitableApp = f
}

// UnMockBitableCopyBitableApp un-mock BitableCopyBitableApp method
func (r *Mock) UnMockBitableCopyBitableApp() {
	r.mockBitableCopyBitableApp = nil
}

// CopyBitableAppReq ...
type CopyBitableAppReq struct {
	AppToken       string  `path:"app_token" json:"-"`        // [多维表格 App token](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/bitable/notification#8121eebe), 示例值: "S404b*e9PQsYDWYcNryFn0g"
	Name           *string `json:"name,omitempty"`            // 多维表格 App 名字, 示例值: "一篇新的多维表格"
	FolderToken    *string `json:"folder_token,omitempty"`    // [多维表格 App 归属文件夹 ](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df), 示例值: "fldbco*CIMltVc"
	WithoutContent *bool   `json:"without_content,omitempty"` // 是否复制多维表格内容, 取值: * true: 不复制, * false: 复制, 示例值: false
	TimeZone       *string `json:"time_zone,omitempty"`       // 文档时区, [详见](https://feishu.feishu.cn/docx/YKRndTM7VoyDqpxqqeEcd67MnEf), 示例值: "Asia/Shanghai"
}

// CopyBitableAppResp ...
type CopyBitableAppResp struct {
	App *CopyBitableAppRespApp `json:"app,omitempty"` // 返回响应体
}

// CopyBitableAppRespApp ...
type CopyBitableAppRespApp struct {
	AppToken       string `json:"app_token,omitempty"`        // 多维表格的 app_token
	Name           string `json:"name,omitempty"`             // 多维表格的名字
	FolderToken    string `json:"folder_token,omitempty"`     // 多维表格 App 归属文件夹
	URL            string `json:"url,omitempty"`              // 多维表格 App URL
	DefaultTableID string `json:"default_table_id,omitempty"` // 默认的表格id
	TimeZone       string `json:"time_zone,omitempty"`        // 文档时区
}

// copyBitableAppResp ...
type copyBitableAppResp struct {
	Code  int64               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string              `json:"msg,omitempty"`  // 错误描述
	Data  *CopyBitableAppResp `json:"data,omitempty"`
	Error *ErrorDetail        `json:"error,omitempty"`
}
