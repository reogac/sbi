/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AddrFqdn struct {
	IpAddr *IpAddr `json:"ipAddr,omitempty"`
	Fqdn   string  `json:"fqdn,omitempty"`
}
