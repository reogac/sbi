/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	ChgId                string         `json:"chgId"`
	Offline              *bool          `json:"offline,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
}
