/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.SbiRoute{
	{
		Label:   "DeleteSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies/:smPolicyId/delete",
		Handler: OnDeleteSMPolicy,
	},
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
}

func Service(p Producer) sbi.SbiService {
	return sbi.SbiService{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
