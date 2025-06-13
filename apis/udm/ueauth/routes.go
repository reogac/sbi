/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ueauth

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "GenerateProseAV",
		Method:  http.MethodPost,
		Path:    "/or/:supiOrSuci/prose-security-information/generate-av",
		Handler: OnGenerateProseAV,
	},
	{
		Label:   "GenerateAuthData",
		Method:  http.MethodPost,
		Path:    "/or/:supiOrSuci/security-information/generate-auth-data",
		Handler: OnGenerateAuthData,
	},
	{
		Label:   "GetRgAuthData",
		Method:  http.MethodGet,
		Path:    "/or/:supiOrSuci/security-information-rg",
		Handler: OnGetRgAuthData,
	},
	{
		Label:   "ConfirmAuth",
		Method:  http.MethodPost,
		Path:    "/:supi/auth-events",
		Handler: OnConfirmAuth,
	},
	{
		Label:   "GenerateAv",
		Method:  http.MethodPost,
		Path:    "/:supi/hss-security-information/:hssAuthType/generate-av",
		Handler: OnGenerateAv,
	},
	{
		Label:   "DeleteAuth",
		Method:  http.MethodPut,
		Path:    "/:supi/auth-events/:authEventId",
		Handler: OnDeleteAuth,
	},
	{
		Label:   "GenerateGbaAv",
		Method:  http.MethodPost,
		Path:    "/:supi/gba-security-information/generate-av",
		Handler: OnGenerateGbaAv,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
