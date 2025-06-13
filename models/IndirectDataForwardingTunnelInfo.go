/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IndirectDataForwardingTunnelInfo struct {
	Ipv6Addr        string `json:"ipv6Addr,omitempty"`
	GtpTeid         string `json:"gtpTeid"`
	DrbId           *int   `json:"drbId,omitempty"`
	AdditionalTnlNb *int   `json:"additionalTnlNb,omitempty"`
	Ipv4Addr        string `json:"ipv4Addr,omitempty"`
}
