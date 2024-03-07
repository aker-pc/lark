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

// CreateDriveDoc :::warning
//
// 此接口已废弃。要创建文档, 使用[创建文档](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/document-docx/docx-v1/document/create)接口。
// 该接口用于创建并初始化文档。
// ## 前提条件
// 在使用此接口前, 请仔细阅读[文档概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/docs-doc-overview)和[准备接入文档 API](https://open.feishu.cn/document/ukTMukTMukTM/ugzNzUjL4czM14CO3MTN/guide/getting-start)了解文档调用的规则和约束, 确保你的文档数据不会丢失或出错。
// 文档数据结构定义可参考: [文档数据结构概述](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugDM2YjL4AjN24COwYjN
// new doc: https://open.feishu.cn/document/server-docs/docs/docs/docs/apiRef/create-document
//
// Deprecated
func (r *DriveService) CreateDriveDoc(ctx context.Context, request *CreateDriveDocReq, options ...MethodOptionFunc) (*CreateDriveDocResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveDoc != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveDoc mock enable")
		return r.cli.mock.mockDriveCreateDriveDoc(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveDoc",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/doc/v2/create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveDocResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveDoc mock DriveCreateDriveDoc method
func (r *Mock) MockDriveCreateDriveDoc(f func(ctx context.Context, request *CreateDriveDocReq, options ...MethodOptionFunc) (*CreateDriveDocResp, *Response, error)) {
	r.mockDriveCreateDriveDoc = f
}

// UnMockDriveCreateDriveDoc un-mock DriveCreateDriveDoc method
func (r *Mock) UnMockDriveCreateDriveDoc() {
	r.mockDriveCreateDriveDoc = nil
}

// CreateDriveDocReq ...
type CreateDriveDocReq struct {
	FolderToken *string `json:"FolderToken,omitempty"` // 文件夹 token, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6)；空表示根目录, tenant_access_token应用权限仅允许操作应用创建的目录
	Content     *string `json:"Content,omitempty"`     // 传入符合[文档数据结构](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)的字符串, 若为空表示创建空文档
}

// CreateDriveDocResp ...
type CreateDriveDocResp struct {
	ObjToken string `json:"objToken,omitempty"` // 新建文档的token
	URL      string `json:"url,omitempty"`      // 新建文档的访问链接
}

// createDriveDocResp ...
type createDriveDocResp struct {
	Code  int64               `json:"code,omitempty"`
	Msg   string              `json:"msg,omitempty"`
	Data  *CreateDriveDocResp `json:"data,omitempty"`
	Error *ErrorDetail        `json:"error,omitempty"`
}
