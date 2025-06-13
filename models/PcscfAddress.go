/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PcscfAddress struct {
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
	Fqdn      string   `json:"fqdn,omitempty"`
}
