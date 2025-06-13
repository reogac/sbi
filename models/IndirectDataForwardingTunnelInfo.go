/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type IndirectDataForwardingTunnelInfo struct {
	DrbId           *int   `json:"drbId,omitempty"`
	AdditionalTnlNb *int   `json:"additionalTnlNb,omitempty"`
	Ipv4Addr        string `json:"ipv4Addr,omitempty"`
	Ipv6Addr        string `json:"ipv6Addr,omitempty"`
	GtpTeid         string `json:"gtpTeid"`
}
