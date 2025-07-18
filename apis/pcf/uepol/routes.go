/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package uepol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "CreateIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies",
		Handler: OnCreateIndividualUEPolicyAssociation,
	},
	{
		Label:   "ReadIndividualUEPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: OnReadIndividualUEPolicyAssociation,
	},
	{
		Label:   "DeleteIndividualUEPolicyAssociation",
		Method:  http.MethodDelete,
		Path:    "/policies/:polAssoId",
		Handler: OnDeleteIndividualUEPolicyAssociation,
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: OnReportObservedEventTriggersForIndividualUEPolicyAssociation,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
