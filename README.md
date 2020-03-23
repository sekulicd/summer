# gRPC server/CLI client

- [Simple gRPC user service and its CLI client](#simple-grpc-user-service-and-its-cli-client)
  - [Stack](#stack)
  - [Use](#use)
  - [Install](#install)
  
## Stack

- **gRPC**: Transport
- **Cobra**: CLI
- **CI/CD**: Drone.io
- **OCI**: Kubernetes

## Use

- summer-cli addTuple --num1=2 --num2=2
- summer-cli addTriple --num1=2 --num2=2 --num3=3

## Install
- clone repo
- make ci
- make cd
- kubectl create -f deployment.yaml
