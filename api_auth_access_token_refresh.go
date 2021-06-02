// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// RefreshAccessToken 该接口用于在 access_token 过期时用 refresh_token 重新获取 access_token。此时会返回新的 refresh_token，再次刷新 access_token 时需要使用新的 refresh_token。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN
func (r *AuthService) RefreshAccessToken(ctx context.Context, request *RefreshAccessTokenReq, options ...MethodOptionFunc) (*RefreshAccessTokenResp, *Response, error) {
	if r.cli.mock.mockAuthRefreshAccessToken != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Auth#RefreshAccessToken mock enable")
		return r.cli.mock.mockAuthRefreshAccessToken(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Auth",
		API:                 "RefreshAccessToken",
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedAppAccessToken:  true,
		NeedUserAccessToken: true,
	}
	resp := new(refreshAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAuthRefreshAccessToken(f func(ctx context.Context, request *RefreshAccessTokenReq, options ...MethodOptionFunc) (*RefreshAccessTokenResp, *Response, error)) {
	r.mockAuthRefreshAccessToken = f
}

func (r *Mock) UnMockAuthRefreshAccessToken() {
	r.mockAuthRefreshAccessToken = nil
}

type RefreshAccessTokenReq struct {
	GrantType    string `json:"grant_type,omitempty"`    // 本流程中，此值为 refresh_token
	RefreshToken string `json:"refresh_token,omitempty"` // 来自[获取登录用户身份(新)](https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN) 或 本接口返回值
}

type refreshAccessTokenResp struct {
	Code int64                   `json:"code,omitempty"` // 返回码，0表示请求成功，其他表示请求失败
	Msg  string                  `json:"msg,omitempty"`  // 返回信息
	Data *RefreshAccessTokenResp `json:"data,omitempty"` // 返回业务数据
}

type RefreshAccessTokenResp struct {
	AccessToken      string   `json:"access_token,omitempty"`       // user_access_token，用于获取用户资源
	AvatarURL        string   `json:"avatar_url,omitempty"`         // 用户头像
	AvatarThumb      string   `json:"avatar_thumb,omitempty"`       // 用户头像 72x72
	AvatarMiddle     string   `json:"avatar_middle,omitempty"`      // 用户头像 240x240
	AvatarBig        string   `json:"avatar_big,omitempty"`         // 用户头像 640x640
	ExpiresIn        int64    `json:"expires_in,omitempty"`         // access_token 的有效期，单位: 秒
	Name             string   `json:"name,omitempty"`               // 用户姓名
	EnName           string   `json:"en_name,omitempty"`            // 用户英文名称
	OpenID           string   `json:"open_id,omitempty"`            // 用户在应用内的唯一标识
	UnionID          string   `json:"union_id,omitempty"`           // 用户统一ID
	Email            []string `json:"email,omitempty"`              // 申请了"获取用户邮箱"权限的应用返回该字段
	UserID           string   `json:"user_id,omitempty"`            // 申请了"获取用户 user_id"权限的应用返回该字段
	Mobile           []string `json:"mobile,omitempty"`             // 申请了"获取用户手机号"权限的应用返回该字段
	TenantKey        string   `json:"tenant_key,omitempty"`         // 当前企业标识
	RefreshExpiresIn int64    `json:"refresh_expires_in,omitempty"` // refresh_token 的有效期，单位: 秒
	RefreshToken     string   `json:"refresh_token,omitempty"`      // 刷新用户 access_token 时使用的 token
	TokenType        string   `json:"token_type,omitempty"`         // 此处为 Bearer
}
