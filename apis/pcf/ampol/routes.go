/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "ReadIndividualAMPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: OnReadIndividualAMPolicyAssociation,
	},
	{
		Label:   "DeleteIndividualAMPolicyAssociation",
		Method:  http.MethodDelete,
		Path:    "/policies/:polAssoId",
		Handler: OnDeleteIndividualAMPolicyAssociation,
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: OnReportObservedEventTriggersForIndividualAMPolicyAssociation,
	},
	{
		Label:   "CreateIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies",
		Handler: OnCreateIndividualAMPolicyAssociation,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
