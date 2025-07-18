/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
