/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataNetworkConfiguration struct {
	Name             string                         `json:"name"`
	Snssais          []Snssai                       `json:"snssais,omitempty"`
	NetworkInstances []NetworkInstanceConfiguration `json:"networkInstances,omitempty"`
}
