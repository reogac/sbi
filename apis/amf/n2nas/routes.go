/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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
	{
		Label:   "InitialUeMessage",
		Method:  http.MethodPost,
		Path:    "/init-ue-msg",
		Handler: OnInitialUeMessage,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
