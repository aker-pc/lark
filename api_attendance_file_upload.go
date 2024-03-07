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
	"io"
)

// UploadAttendanceFile 上传文件并获取文件 ID, 可用于“修改用户设置”接口中的 face_key 参数。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload
// new doc: https://open.feishu.cn/document/server-docs/attendance-v1/user_setting/upload
func (r *AttendanceService) UploadAttendanceFile(ctx context.Context, request *UploadAttendanceFileReq, options ...MethodOptionFunc) (*UploadAttendanceFileResp, *Response, error) {
	if r.cli.mock.mockAttendanceUploadAttendanceFile != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Attendance#UploadAttendanceFile mock enable")
		return r.cli.mock.mockAttendanceUploadAttendanceFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "UploadAttendanceFile",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/files/upload",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		IsFile:                true,
	}
	resp := new(uploadAttendanceFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceUploadAttendanceFile mock AttendanceUploadAttendanceFile method
func (r *Mock) MockAttendanceUploadAttendanceFile(f func(ctx context.Context, request *UploadAttendanceFileReq, options ...MethodOptionFunc) (*UploadAttendanceFileResp, *Response, error)) {
	r.mockAttendanceUploadAttendanceFile = f
}

// UnMockAttendanceUploadAttendanceFile un-mock AttendanceUploadAttendanceFile method
func (r *Mock) UnMockAttendanceUploadAttendanceFile() {
	r.mockAttendanceUploadAttendanceFile = nil
}

// UploadAttendanceFileReq ...
type UploadAttendanceFileReq struct {
	FileName string    `query:"file_name" json:"-"` // 带后缀的文件名, 示例值: 人脸照片.jpg
	File     io.Reader `json:"file,omitempty"`      // 文件内容, 示例值: 二进制文件
}

// UploadAttendanceFileResp ...
type UploadAttendanceFileResp struct {
	File *UploadAttendanceFileRespFile `json:"file,omitempty"` // 文件
}

// UploadAttendanceFileRespFile ...
type UploadAttendanceFileRespFile struct {
	FileID string `json:"file_id,omitempty"` // 文件 ID
}

// uploadAttendanceFileResp ...
type uploadAttendanceFileResp struct {
	Code  int64                     `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                    `json:"msg,omitempty"`  // 错误描述
	Data  *UploadAttendanceFileResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
