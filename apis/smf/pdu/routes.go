/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package pdu

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "ReleaseSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/release",
		Handler: OnReleaseSmContext,
	},
	{
		Label:   "SendMoData",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/send-mo-data",
		Handler: OnSendMoData,
	},
	{
		Label:   "ReleasePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/release",
		Handler: OnReleasePduSession,
	},
	{
		Label:   "RetrievePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/retrieve",
		Handler: OnRetrievePduSession,
	},
	{
		Label:   "TransferMoData",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/transfer-mo-data",
		Handler: OnTransferMoData,
	},
	{
		Label:   "PostSmContexts",
		Method:  http.MethodPost,
		Path:    "/sm-contexts",
		Handler: OnPostSmContexts,
	},
	{
		Label:   "RetrieveSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/retrieve",
		Handler: OnRetrieveSmContext,
	},
	{
		Label:   "UpdateSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/modify",
		Handler: OnUpdateSmContext,
	},
	{
		Label:   "PostPduSessions",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions",
		Handler: OnPostPduSessions,
	},
	{
		Label:   "UpdatePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/modify",
		Handler: OnUpdatePduSession,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
