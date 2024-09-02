/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type N2ConnectionChargingInformation struct {
	N2ConnectionMessageType int32                    `json:"n2ConnectionMessageType" yaml:"n2ConnectionMessageType" bson:"n2ConnectionMessageType" mapstructure:"N2ConnectionMessageType"`
	UserInformation         *UserInformation         `json:"userInformation,omitempty" yaml:"userInformation" bson:"userInformation" mapstructure:"UserInformation"`
	UserLocationinfo        *UserLocation            `json:"userLocationinfo,omitempty" yaml:"userLocationinfo" bson:"userLocationinfo" mapstructure:"UserLocationinfo"`
	PSCellInformation       *PsCellInformation       `json:"pSCellInformation,omitempty" yaml:"pSCellInformation" bson:"pSCellInformation" mapstructure:"PSCellInformation"`
	UetimeZone              string                   `json:"uetimeZone,omitempty" yaml:"uetimeZone" bson:"uetimeZone" mapstructure:"UetimeZone"`
	RATType                 RatType                  `json:"rATType,omitempty" yaml:"rATType" bson:"rATType" mapstructure:"RATType"`
	AmfUeNgapId             int32                    `json:"amfUeNgapId,omitempty" yaml:"amfUeNgapId" bson:"amfUeNgapId" mapstructure:"AmfUeNgapId"`
	RanUeNgapId             int32                    `json:"ranUeNgapId,omitempty" yaml:"ranUeNgapId" bson:"ranUeNgapId" mapstructure:"RanUeNgapId"`
	RanNodeId               *GlobalRanNodeId         `json:"ranNodeId,omitempty" yaml:"ranNodeId" bson:"ranNodeId" mapstructure:"RanNodeId"`
	RestrictedRatList       []RatType                `json:"restrictedRatList,omitempty" yaml:"restrictedRatList" bson:"restrictedRatList" mapstructure:"RestrictedRatList"`
	ForbiddenAreaList       []Area                   `json:"forbiddenAreaList,omitempty" yaml:"forbiddenAreaList" bson:"forbiddenAreaList" mapstructure:"ForbiddenAreaList"`
	ServiceAreaRestriction  []ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty" yaml:"serviceAreaRestriction" bson:"serviceAreaRestriction" mapstructure:"ServiceAreaRestriction"`
	RestrictedCnList        []CoreNetworkType        `json:"restrictedCnList,omitempty" yaml:"restrictedCnList" bson:"restrictedCnList" mapstructure:"RestrictedCnList"`
	AllowedNSSAI            []Snssai                 `json:"allowedNSSAI,omitempty" yaml:"allowedNSSAI" bson:"allowedNSSAI" mapstructure:"AllowedNSSAI"`
	RrcEstCause             string                   `json:"rrcEstCause,omitempty" yaml:"rrcEstCause" bson:"rrcEstCause" mapstructure:"RrcEstCause"`
}
