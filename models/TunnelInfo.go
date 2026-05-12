/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TunnelInfo struct {
	Ipv4Addr string     `json:"ipv4Addr,omitempty"`
	Ipv6Addr string     `json:"ipv6Addr,omitempty"`
	GtpTeid  string     `json:"gtpTeid"`
	AnType   AccessType `json:"anType,omitempty"`
}
