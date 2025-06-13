/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	Online               *bool          `json:"online,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
}
