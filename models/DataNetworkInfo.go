/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Apr 29 09:34:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkInfo struct {
	Pcscf *IpAddr `json:"pcscf,omitempty"`
	Name  string  `json:"name"`
	Cidr  string  `json:"cidr"`
	Dns   *IpAddr `json:"dns,omitempty"`
}
