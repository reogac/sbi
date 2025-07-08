/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IndirectDataForwardingTunnelInfo struct {
	GtpTeid         string `json:"gtpTeid"`
	DrbId           *int   `json:"drbId,omitempty"`
	AdditionalTnlNb *int   `json:"additionalTnlNb,omitempty"`
	Ipv4Addr        string `json:"ipv4Addr,omitempty"`
	Ipv6Addr        string `json:"ipv6Addr,omitempty"`
}
