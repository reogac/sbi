/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "EapAuthMethod",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/:authCtxId/eap-session",
		Handler: OnEapAuthMethod,
	},
	{
		Label:   "UeAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications",
		Handler: OnUeAuthenticationsPost,
	},
	{
		Label:   "UeAuthentications5gAkaConfirmationPut",
		Method:  http.MethodPut,
		Path:    "/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnUeAuthentications5gAkaConfirmationPut,
	},
	{
		Label:   "Delete5gAkaAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnDelete5gAkaAuthenticationResult,
	},
	{
		Label:   "ProseAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/prose-authentications",
		Handler: OnProseAuthenticationsPost,
	},
	{
		Label:   "ProseAuth",
		Method:  http.MethodPost,
		Path:    "/prose-authentications/:authCtxId/prose-auth",
		Handler: OnProseAuth,
	},
	{
		Label:   "DeleteProSeAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/prose-authentications/:authCtxId/prose-auth",
		Handler: OnDeleteProSeAuthenticationResult,
	},
	{
		Label:   "UeAuthenticationsDeregisterPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/deregister",
		Handler: OnUeAuthenticationsDeregisterPost,
	},
	{
		Label:   "DeleteEapAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/ue-authentications/:authCtxId/eap-session",
		Handler: OnDeleteEapAuthenticationResult,
	},
	{
		Label:   "RgAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/rg-authentications",
		Handler: OnRgAuthenticationsPost,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
