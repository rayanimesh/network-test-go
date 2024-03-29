/*
 * DNS Service
 *
 * Sevice for managing DNS with route53
 *
 * API version: 0.0.1
 * Contact: iitr.animesh@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package route53dns


import (
	"errors"
)



type DnsRecord struct {

	Name string `json:"name,omitempty"`

	Type DnsRecordType `json:"type,omitempty"`

	Value []string `json:"value,omitempty"`

	Ttl int32 `json:"ttl,omitempty"`
}

// AssertDnsRecordRequired checks if the required fields are not zero-ed
func AssertDnsRecordRequired(obj DnsRecord) error {
	return nil
}

// AssertDnsRecordConstraints checks if the values respects the defined constraints
func AssertDnsRecordConstraints(obj DnsRecord) error {
	if obj.Ttl < 60 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Ttl > 86400 {
		return &ParsingError{Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
