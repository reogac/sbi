/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package alloc

import (
	"github.com/reogac/sbi"
	"net/http"
)

var _routes = []sbi.Route[Producer]{
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
		Label:   "ReleaseIpSegments",
		Method:  http.MethodPost,
		Path:    "/ip-segments/release",
		Handler: OnReleaseIpSegments,
	},
}

func Routes() []sbi.Route[Producer] {
	return _routes
}
