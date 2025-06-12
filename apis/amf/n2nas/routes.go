/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n2nas

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
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

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
