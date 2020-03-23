package main
import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	summer "summer/pkg/grpc-schema"
	summerservice "summer/pkg/summer-server"
)
const (
	grpcPort = ":50051"
)
type SummerServer struct{
	SummerService summerservice.SummerService
}

func (s *SummerServer)AddTuple(con context.Context,req *summer.RequestTuple) (*summer.Response, error) {
	//validate input
	sum := s.SummerService.AddTuple(int(req.A), int(req.B))
	return &summer.Response{
		Sum:                  int32(sum),
	}, nil
}

func (s *SummerServer)AddTriple(con context.Context,req *summer.RequestTriple) (*summer.Response, error) {
	//validate input
	sum := s.SummerService.AddTriple(int(req.A), int(req.B), int(req.C))
	return &summer.Response{
		Sum:                  int32(sum),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	grpcServer := grpc.NewServer()
	summer.RegisterSummerServer(grpcServer, &SummerServer{summerservice.SummerService{}})
	reflection.Register(grpcServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
