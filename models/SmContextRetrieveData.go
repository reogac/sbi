/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrieveData struct {
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
	SmContextType        SmContextType    `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
}
