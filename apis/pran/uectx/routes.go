/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uectx

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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

func Routes() []sbi.Route[Producer] {
	return _routes
}
