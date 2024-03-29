// Copyright 2022 CYBERCRYPT
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by copy-client.sh. DO NOT EDIT.
// version: v2.0.0
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 05a3b04a1916d3e6c1bfea65041b7536cf93778e

package client

import (
	pbauthn "github.com/cybercryptio/d1-client-go/v2/d1-generic/protobuf/authn"
	pbauthz "github.com/cybercryptio/d1-client-go/v2/d1-generic/protobuf/authz"
	pbindex "github.com/cybercryptio/d1-client-go/v2/d1-generic/protobuf/index"
	pbversion "github.com/cybercryptio/d1-client-go/v2/d1-generic/protobuf/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// BaseClient represents the shared functionality between various D1 services.
type BaseClient struct {
	Version    pbversion.VersionClient
	Authn      pbauthn.AuthnClient
	Authz      pbauthz.AuthzClient
	Health     grpc_health_v1.HealthClient
	Index      pbindex.IndexClient
	Connection *grpc.ClientConn
}

// Option is used configure optional settings on the client.
type Option func(*BaseClient) grpc.DialOption

// WithGrpcOption returns an Option which configures the underlying gRPC connection.
func WithGrpcOption(option grpc.DialOption) Option {
	return func(*BaseClient) grpc.DialOption {
		return option
	}
}

// NewBaseClient creates a new client for the given endpoint, configured with the provided options.
func NewBaseClient(endpoint string, opts ...Option) (BaseClient, error) {
	var err error
	baseClient := BaseClient{}

	grpcOpts := []grpc.DialOption{}
	for _, opt := range opts {
		grpcOpts = append(grpcOpts, opt(&baseClient))
	}

	// Initialize connection with the service
	baseClient.Connection, err = grpc.Dial(endpoint, grpcOpts...)
	if err != nil {
		return BaseClient{}, err
	}

	baseClient.Version = pbversion.NewVersionClient(baseClient.Connection)
	baseClient.Authn = pbauthn.NewAuthnClient(baseClient.Connection)
	baseClient.Authz = pbauthz.NewAuthzClient(baseClient.Connection)
	baseClient.Health = grpc_health_v1.NewHealthClient(baseClient.Connection)
	baseClient.Index = pbindex.NewIndexClient(baseClient.Connection)

	return baseClient, nil
}

// Close closes all connections to the server.
func (b *BaseClient) Close() error {
	return b.Connection.Close()
}
