/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
