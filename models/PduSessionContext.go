/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionContext struct {
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	Dnn                        string             `json:"dnn"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
}
