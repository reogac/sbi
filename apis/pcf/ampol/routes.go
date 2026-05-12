/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:43 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "CreateIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies",
		Handler: OnCreateIndividualAMPolicyAssociation,
	},
	{
		Label:   "DeleteIndividualAMPolicyAssociation",
		Method:  http.MethodDelete,
		Path:    "/policies/:polAssoId",
		Handler: OnDeleteIndividualAMPolicyAssociation,
	},
	{
		Label:   "ReadIndividualAMPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: OnReadIndividualAMPolicyAssociation,
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: OnReportObservedEventTriggersForIndividualAMPolicyAssociation,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
