// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run github.com/golang/mock/mockgen -destination mock.go github.com/terramate-io/tf/cloudplugin/cloudproto1 CommandServiceClient,CommandService_ExecuteClient

package mock_cloudproto1
