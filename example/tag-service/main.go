package main

import (
	"errors"
	"flag"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"strconv"
	"tag-service/global"
	"tag-service/internal/middleware"
	"tag-service/pkg/balance"
	"tag-service/pkg/tracer"
	"tag-service/pkg/weight"
	pb "tag-service/proto"
	"tag-service/server"
)

var grpcPort string
var httpPort string
var w string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "6699", "grpc启动端口号")
	flag.StringVar(&httpPort, "http_port", "0033", "http启动端口号")
	flag.StringVar(&w, "w", "0", "负载均衡权重0~5")
	flag.Parse()
	err := setupTracer()
	if err != nil {
		log.Fatalf("tag-service init.setupTracer err: %v\n", err)
	}
}

const SERVICE_NAME = "tag-service"

func main() {
	g := new(errgroup.Group)
	g.Go(func() error {
		return RunHttpServer()
	})
	g.Go(func() error {
		return RunGrpcServer()
	})
	if err := g.Wait(); err != nil {
		log.Fatalf("Run server err:%v", err)
	}
}

func RunHttpServer() error {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("pong"))
	})
	s := &http.Server{
		Addr:    ":" + httpPort,
		Handler: serveMux,
	}
	return s.ListenAndServe()
}

func RunGrpcServer() error {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.HelloInterceptor,
			middleware.WorldInterceptor,
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
			middleware.ServerTracing,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer()),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	//grpcurl -plaintext -d '{"name":"Go"}' localhost:6699 TagService.GetTagList
	//protoc --go_out=plugins=grpc:. ./proto/*.proto
	reflection.Register(s)
	atoi, err := strconv.Atoi(w)
	if err != nil {
		return err
	}
	if atoi < 0 || atoi > 5 {
		return errors.New(fmt.Sprintf("错误权重w %d", atoi))
	}
	var ser *balance.ServiceRegister
	if atoi == 0 {
		ser, err = balance.NewServiceRegister([]string{"http://localhost:2379"}, SERVICE_NAME, "localhost:"+grpcPort, 5, fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name))
	} else {
		ser, err = balance.NewServiceRegister([]string{"http://localhost:2379"}, SERVICE_NAME, "localhost:"+grpcPort, 5, fmt.Sprintf(`{"LoadBalancingPolicy": "%s","weight":"%s"}`, weight.Name, w))
	}
	if err != nil {
		return err
	}
	defer ser.Close()
	/*target := fmt.Sprintf("/etcdv3://go-programming-tour/grpc/%s", SERVICE_NAME)
	grpcproxy.Register(etcdClient, target, ":"+grpcPort, 60)*/
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		return err
	}
	defer lis.Close()
	return s.Serve(lis)
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("tag-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
