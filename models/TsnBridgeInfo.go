/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TsnBridgeInfo struct {
	BridgeId      *int   `json:"bridgeId,omitempty"`
	DsttAddr      string `json:"dsttAddr,omitempty"`
	DsttPortNum   *int   `json:"dsttPortNum,omitempty"`
	DsttResidTime *int   `json:"dsttResidTime,omitempty"`
}
