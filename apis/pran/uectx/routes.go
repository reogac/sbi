/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:01:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "UpdateAmfUeContextInfo",
		Method:  http.MethodPut,
		Path:    "/amfinfo/:ueId",
		Handler: OnUpdateAmfUeContextInfo,
	},
	{
		Label:   "UeContextModify",
		Method:  http.MethodPut,
		Path:    "/uectx/modify/:ueId",
		Handler: OnUeContextModify,
	},
	{
		Label:   "UeContextRelease",
		Method:  http.MethodPut,
		Path:    "/uectx/release/:ueId",
		Handler: OnUeContextRelease,
	},
	{
		Label:   "UeContextSetup",
		Method:  http.MethodPost,
		Path:    "/uectx/setup/:ueId",
		Handler: OnUeContextSetup,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
