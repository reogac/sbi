/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	Dnn                        string             `json:"dnn"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
}
