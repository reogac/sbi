/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	Dnn                        string             `json:"dnn"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
}
