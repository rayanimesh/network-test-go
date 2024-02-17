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
	"context"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
)

type mockClientClient struct {
	route53iface.ClientPI
}

func TestNewDefaultAPIService(t *testing.T) {
	tests := []struct {
		name string
		want DefaultAPIServicer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultAPIService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultAPIService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultAPIService_ListRoute53HostedZoneRecords(t *testing.T) {
	type fields struct {
		Route53Client *route53.Client
	}
	type args struct {
		ctx    context.Context
		domain string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ImplResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DefaultAPIService{
				Route53Client: tt.fields.Route53Client,
			}
			got, err := s.ListRoute53HostedZoneRecords(tt.args.ctx, tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultAPIService.ListRoute53HostedZoneRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultAPIService.ListRoute53HostedZoneRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
