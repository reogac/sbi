/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "CreateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies",
		Handler: OnCreateSMPolicy,
	},
	{
		Label:   "GetSMPolicy",
		Method:  http.MethodGet,
		Path:    "/sm-policies/:smPolicyId",
		Handler: OnGetSMPolicy,
	},
	{
		Label:   "UpdateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies/:smPolicyId/update",
		Handler: OnUpdateSMPolicy,
	},
	{
		Label:   "DeleteSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies/:smPolicyId/delete",
		Handler: OnDeleteSMPolicy,
	},
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
