package config

import (
	"strconv"
	"time"

	"gitlab.com/sidenio/go/common.git/pkg/env"
	"google.golang.org/grpc/keepalive"
)

const (
	bitSizeFloat64 = 64

	decimalBase  = 10
	bitSizeInt64 = 64

	MegaByte           = 1024 * 1024
	ThirtyTwoKiloBytes = 0.032
)

var (
	// GRPCReadBufferSizeMBs and friends tweaks
	// See also: https://siden.atlassian.net/wiki/spaces/ENG/pages/2389704710/GRPC+Tweaks
	// Query: avg(rate(orderproc_oc_grpc_clients_request_size_sum[5m])) by (instance, response_code)
	GRPCReadBufferSizeMBs  = env.GetEnvVariableOrDefault("GRPC_READ_BUFFER_MBs", "0.032")  // default is 32kB
	GRPCWriteBufferSizeMBs = env.GetEnvVariableOrDefault("GRPC_WRITE_BUFFER_MBs", "0.032") // default is 32kB
	// GRPCMaxRecvMsgSizeMBs allows overwriting default for 4MB
	// https://pkg.go.dev/google.golang.org/grpc#MaxRecvMsgSize
	// https://github.com/grpc/grpc-go/blob/0fe49e823fcd9904afba6cd5e5980da4390d1899/server.go#L58
	GRPCMaxRecvMsgSizeMBs = env.GetEnvVariableOrDefault("GRPC_MAX_RECV_MBs", "4.0") // default is 4MB
	// GRPCMaxConcurrentStreams comes from here
	// https://github.com/grpc/grpc-go/blob/87eb5b7502493f758e76c4d09430c0049a81a557/internal/transport/defaults.go#L26
	GRPCMaxConcurrentStreams = env.GetEnvVariableOrDefault("GRPC_MAX_CONCURRENT", "10000")      // default is 100
	GRPCKeepAliveTimeMins    = env.GetEnvVariableOrDefault("GRPC_KEEPALIVE_MINs", "5")          // default is 2 hours
	GRPCKeepAliveTimeoutSecs = env.GetEnvVariableOrDefault("GRPC_KEEPALIVE_TIMEOUT_SECs", "20") // default is 20 seconds
	GRPCMaxIdleMins          = env.GetEnvVariableOrDefault("GRPC_MAX_IDLE_MINs", "20")          // default is infinity

	GRPCClientKeepAliveTimeSecs = env.GetEnvVariableOrDefault("GRPC_CLIENT_KEEPALIVE_SECs", "60") // default is 2 hours
)

// GRPC Tweaks start
func GetGRPCReadBufferSize() int {
	i64, err := strconv.ParseFloat(GRPCReadBufferSizeMBs, bitSizeFloat64)
	if err != nil {
		i64 = ThirtyTwoKiloBytes // default 32 kb
	}
	return int(i64 * MegaByte)
}

func GetGRPCWriteBufferSize() int {
	i64, err := strconv.ParseFloat(GRPCWriteBufferSizeMBs, bitSizeFloat64)
	if err != nil {
		i64 = ThirtyTwoKiloBytes // default 32 kb
	}
	return int(i64 * MegaByte)
}

func GetGRPCMaxRecvMsgSize() int {
	i64, err := strconv.ParseFloat(GRPCMaxRecvMsgSizeMBs, bitSizeFloat64)
	if err != nil {
		i64 = 4 // default 4 MB
	}
	return int(i64 * MegaByte)
}
func GetGRPCMaxConcurrentStreams() uint32 {
	i64, err := strconv.ParseInt(GRPCMaxConcurrentStreams, decimalBase, bitSizeInt64)
	if err != nil {
		i64 = 100 // default 100, recommend increasing a LOT
	}
	return uint32(i64)
}
func GetGRPCKeepAliveTime() time.Duration {
	i64, err := strconv.ParseInt(GRPCKeepAliveTimeMins, decimalBase, bitSizeInt64)
	if err != nil {
		i64 = 120 // default 2 hours = 120 mins
	}
	return time.Duration(i64) * time.Minute
}
func GetGRPCKeepAliveTimeout() time.Duration {
	i64, err := strconv.ParseInt(GRPCKeepAliveTimeoutSecs, decimalBase, bitSizeInt64)
	if err != nil {
		i64 = 20 // default 20 seconds
	}
	return time.Duration(i64) * time.Second
}
func GetGRPCMaxIdleTime() time.Duration {
	i64, err := strconv.ParseInt(GRPCMaxIdleMins, decimalBase, bitSizeInt64)
	if err != nil {
		i64 = 20 // default 20 mins
	}
	return time.Duration(i64) * time.Minute
}

func GetGRPCKeepAliveConfig() keepalive.ServerParameters {
	return keepalive.ServerParameters{
		Time:              GetGRPCKeepAliveTime(),
		Timeout:           GetGRPCKeepAliveTimeout(),
		MaxConnectionIdle: GetGRPCMaxIdleTime(),
	}
}

func GetGRPCClientKeepAliveTime() time.Duration {
	i64, err := strconv.ParseInt(GRPCClientKeepAliveTimeSecs, decimalBase, bitSizeInt64)
	if err != nil {
		i64 = 30 // default 2 hours = 120 mins
	}
	return time.Duration(i64) * time.Second
}

// GRPC Tweaks end
