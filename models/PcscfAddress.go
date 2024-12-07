/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PcscfAddress struct {
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
	Fqdn      string   `json:"fqdn,omitempty"`
	Ipv4Addrs []string `json:"ipv4Addrs,omitempty"`
}
