/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package n4sbi

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "AssociationRequest",
		Method:  http.MethodPost,
		Path:    "/associate",
		Handler: OnAssociationRequest,
	},
	{
		Label:   "DisassociationRequest",
		Method:  http.MethodPut,
		Path:    "/disassociate/:smfId",
		Handler: OnDisassociationRequest,
	},
	{
		Label:   "SessionEstablishment",
		Method:  http.MethodPost,
		Path:    "/session/setup/:smfId",
		Handler: OnSessionEstablishment,
	},
	{
		Label:   "SessionModification",
		Method:  http.MethodPut,
		Path:    "/session/modify/:seid",
		Handler: OnSessionModification,
	},
	{
		Label:   "SessionDeletion",
		Method:  http.MethodPut,
		Path:    "/session/delete/:seid",
		Handler: OnSessionDeletion,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
