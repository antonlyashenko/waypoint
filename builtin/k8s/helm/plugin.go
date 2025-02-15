package helm

import (
	"github.com/hashicorp/waypoint-plugin-sdk"
)

//go:generate protoc -I ../../../.. --go_out=../../../.. --go-grpc_out=../../../.. waypoint/builtin/k8s/helm/plugin.proto

// Options are the SDK options to use for instantiation for the plugin.
var Options = []sdk.Option{
	sdk.WithComponents(&Platform{}),
}
