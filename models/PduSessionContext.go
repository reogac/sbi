/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	Dnn                        string             `json:"dnn"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
}
