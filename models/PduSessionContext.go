/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	Dnn                        string             `json:"dnn"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
}
