# grpc_service_policy_check

This was a tiny little go program to allow testing of the grpc_service_policy.yml configuration is valid.

## Useful targets

```bash
make
```

## Running the test

```bash
./grpc_service_policy_check
```

If the .yml doesn't parse, it will get an error, for example:

```bash
das@t:~/Downloads/siden-edge-node/cmd/grpc_service_policy_check$ ./grpc_service_policy_check
batcher gprc dial
panic: grpc: the provided default service config is invalid: invalid character '"' after object key:value pair

goroutine 1 [running]:
main.main()

```

or
```bash
das@t:~/Downloads/siden-edge-node/cmd/grpc_service_policy_check$ ./grpc_service_policy_check
batcher gprc dial
panic: grpc: the provided default service config is invalid: json: cannot unmarshal string into Go struct field jsonSC.LoadBalancingConfig of type serviceconfig.intermediateBalancerConfig

goroutine 1 [running]:
main.main()

```