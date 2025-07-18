/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PresenceInfoRm struct {
	PresenceState       PresenceState     `json:"presenceState,omitempty"`
	TrackingAreaList    []Tai             `json:"trackingAreaList,omitempty"`
	EcgiList            []Ecgi            `json:"ecgiList,omitempty"`
	NcgiList            []Ncgi            `json:"ncgiList,omitempty"`
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
	GlobaleNbIdList     []GlobalRanNodeId `json:"globaleNbIdList,omitempty"`
	PraId               string            `json:"praId,omitempty"`
	AdditionalPraId     string            `json:"additionalPraId,omitempty"`
}
