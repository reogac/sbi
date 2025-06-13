/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package smpol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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

func Routes() []sbi.Route[Producer] {
	return _routes
}
