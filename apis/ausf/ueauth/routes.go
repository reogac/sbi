/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:29 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "ProseAuth",
		Method:  http.MethodPost,
		Path:    "/prose-authentications/:authCtxId/prose-auth",
		Handler: OnProseAuth,
	},
	{
		Label:   "UeAuthenticationsDeregisterPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/deregister",
		Handler: OnUeAuthenticationsDeregisterPost,
	},
	{
		Label:   "UeAuthentications5gAkaConfirmationPut",
		Method:  http.MethodPut,
		Path:    "/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnUeAuthentications5gAkaConfirmationPut,
	},
	{
		Label:   "EapAuthMethod",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/:authCtxId/eap-session",
		Handler: OnEapAuthMethod,
	},
	{
		Label:   "DeleteEapAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/ue-authentications/:authCtxId/eap-session",
		Handler: OnDeleteEapAuthenticationResult,
	},
	{
		Label:   "DeleteProSeAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/prose-authentications/:authCtxId/prose-auth",
		Handler: OnDeleteProSeAuthenticationResult,
	},
	{
		Label:   "UeAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications",
		Handler: OnUeAuthenticationsPost,
	},
	{
		Label:   "Delete5gAkaAuthenticationResult",
		Method:  http.MethodDelete,
		Path:    "/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: OnDelete5gAkaAuthenticationResult,
	},
	{
		Label:   "RgAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/rg-authentications",
		Handler: OnRgAuthenticationsPost,
	},
	{
		Label:   "ProseAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/prose-authentications",
		Handler: OnProseAuthenticationsPost,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
