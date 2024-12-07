/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:23 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package nasdl

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "NasDl",
		Method:  http.MethodPost,
		Path:    "/dl/:ueId",
		Handler: OnNasDl,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
