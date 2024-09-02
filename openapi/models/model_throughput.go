/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * API version: 3.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type Throughput struct {
	GuaranteedThpt float32 `json:"guaranteedThpt,omitempty" yaml:"guaranteedThpt" bson:"guaranteedThpt" mapstructure:"GuaranteedThpt"`
	MaximumThpt    float32 `json:"maximumThpt,omitempty" yaml:"maximumThpt" bson:"maximumThpt" mapstructure:"MaximumThpt"`
}
