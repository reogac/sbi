/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProseServiceAuth struct {
	ProseL3RemoteAuth            UeAuth `json:"proseL3RemoteAuth,omitempty"`
	ProseDirectDiscoveryAuth     UeAuth `json:"proseDirectDiscoveryAuth,omitempty"`
	ProseDirectCommunicationAuth UeAuth `json:"proseDirectCommunicationAuth,omitempty"`
	ProseL2RelayAuth             UeAuth `json:"proseL2RelayAuth,omitempty"`
	ProseL3RelayAuth             UeAuth `json:"proseL3RelayAuth,omitempty"`
	ProseL2RemoteAuth            UeAuth `json:"proseL2RemoteAuth,omitempty"`
}
