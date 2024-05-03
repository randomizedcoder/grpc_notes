package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"context"

	"gitlab.com/sidenio/api.git/pkg/generated/grpc/content_access"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"crypto/tls"

	"google.golang.org/grpc/credentials"
)

const (
	apiHostnameAndPort    = "api.dev.siden.io:443"
	dialTimeout           = 10 * time.Second
	keepaliveDuration     = 30 * time.Second
	GRPCWriteBufferSize   = 2 * 1024 * 1024 // 2MB
	GRPCTimeout           = 15 * time.Second
	GRPCServicePolicyFile = "./grpc_service_policy.yml"

	// nasty debug level allows changing of the logging
	debugLevel = 11
)

func main() {

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
		//InsecureSkipVerify: true,
	}

	keep := &keepalive.ClientParameters{
		Time:                keepaliveDuration,
		PermitWithoutStream: true, // send pings even without active streams
	}

	// see https://github.com/grpc/grpc/blob/master/doc/service_config.md to know more about service config
	// https://github.com/grpc/grpc-go/blob/11feb0a9afd8/examples/features/retry/client/main.go#L36
	// https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	// https://gitlab.com/sidenio/ml/episode-scheduler/-/blob/develop/es/main.py#L104
	servicePolicyBytes, err := os.ReadFile(GRPCServicePolicyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create the GRPC connection
	// https://pkg.go.dev/google.golang.org/grpc@v1.49.0#WithReadBufferSize
	dialCtx, dCancel := context.WithTimeout(context.Background(), dialTimeout)
	defer dCancel()

	// analyticsConn, dErr := grpc.DialContext(
	cloudConnection, dErr := grpc.DialContext(
		dialCtx,
		apiHostnameAndPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(*keep),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		//grpc.WithReadBufferSize(),
		grpc.WithWriteBufferSize(GRPCWriteBufferSize),
		//grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		//grpc.WithPerRPCCredentials(jwtCreds),
		grpc.WithDefaultServiceConfig(string(servicePolicyBytes)),
	)
	if dErr != nil {
		if debugLevel > 10 {
			fmt.Println("batcher gprc dial")
		}
		panic(dErr)
	}

	contentAccessClient := content_access.NewContentAccessClient(cloudConnection)
	var batch content_access.LogRequestBatch

	gCtx := context.Background()

	if !true {
		_, gErr := contentAccessClient.LogRequest(gCtx, &batch)
		if gErr != nil {
			log.Printf("GRPC contentAccessClient.LogRequest(gCtx, &batch) error:%v", err)
		}
	}

}
