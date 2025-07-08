/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "InitialUeMessage",
		Method:  http.MethodPost,
		Path:    "/init-ue-msg",
		Handler: OnInitialUeMessage,
	},
	{
		Label:   "NasUl",
		Method:  http.MethodPut,
		Path:    "/nas-ul/:ueId",
		Handler: OnNasUl,
	},
	{
		Label:   "NasErr",
		Method:  http.MethodPut,
		Path:    "/nas-err/:ueId",
		Handler: OnNasErr,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
