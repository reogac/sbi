/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package ampol

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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
	{
		Label:   "ReadIndividualAMPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: OnReadIndividualAMPolicyAssociation,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
