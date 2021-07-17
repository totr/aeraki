package model

import (
	"encoding/json"

	istiomodel "istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/config/host"
)

// TrafficDirection defines whether traffic exists a service instance or enters a service instance
type TrafficDirection string

const (
	// TrafficDirectionInbound indicates inbound traffic
	TrafficDirectionInbound TrafficDirection = "inbound"
	// TrafficDirectionOutbound indicates outbound traffic
	TrafficDirectionOutbound TrafficDirection = "outbound"
)

// BuildClusterName the cluster name referencing service instances for a given service name, a subset and a port, which is the same as the cluster name generated by Istio.
func BuildClusterName(direction TrafficDirection, subsetName string, hostname string, port int) string {
	if direction == TrafficDirectionInbound {
		// On 1.8+ Proxies, Istio uses format inbound|port||. Telemetry no longer requires the hostname
		return istiomodel.BuildSubsetKey(istiomodel.TrafficDirection(direction), subsetName, "", port)
	}
	return istiomodel.BuildSubsetKey(istiomodel.TrafficDirection(direction), subsetName, host.Name(hostname), port)
}

// Struct2JSON convert a go struct to a json object
func Struct2JSON(ojb interface{}) interface{} {
	b, err := json.Marshal(ojb)
	if err != nil {
		return ojb
	}
	return string(b)
}
