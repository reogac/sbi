/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	Dnn                        string             `json:"dnn"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
}
