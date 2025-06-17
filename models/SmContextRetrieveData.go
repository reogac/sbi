/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextRetrieveData struct {
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
	NotToTransferEbiList []int            `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
	SmContextType        SmContextType    `json:"smContextType,omitempty"`
}
