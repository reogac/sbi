/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 11:15:54 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package alloc

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
	{
		Label:   "ReleaseIpSegments",
		Method:  http.MethodPost,
		Path:    "/ip-segments/release",
		Handler: OnReleaseIpSegments,
	},
	{
		Label:   "AmfRegister",
		Method:  http.MethodPut,
		Path:    "/amf-pointer",
		Handler: OnAmfRegister,
	},
	{
		Label:   "AllocateIpSegments",
		Method:  http.MethodPut,
		Path:    "/ip-segments",
		Handler: OnAllocateIpSegments,
	},
	{
		Label:   "RenewLeases",
		Method:  http.MethodPut,
		Path:    "/leases",
		Handler: OnRenewLeases,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
