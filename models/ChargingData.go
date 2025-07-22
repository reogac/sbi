/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ChargingData struct {
	MeteringMethod       MeteringMethod `json:"meteringMethod,omitempty"`
	SdfHandl             *bool          `json:"sdfHandl,omitempty"`
	SponsorId            string         `json:"sponsorId,omitempty"`
	AfChargingIdentifier *int           `json:"afChargingIdentifier,omitempty"`
	AfChargId            string         `json:"afChargId,omitempty"`
	ChgId                string         `json:"chgId"`
	Online               *bool          `json:"online,omitempty"`
	RatingGroup          *int           `json:"ratingGroup,omitempty"`
	ReportingLevel       ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId            *int           `json:"serviceId,omitempty"`
	AppSvcProvId         string         `json:"appSvcProvId,omitempty"`
	Offline              *bool          `json:"offline,omitempty"`
}
