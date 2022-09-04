package service

import (
	"context"
	"flag"
	"fmt"
	szpt_login "github.com/Asuka999/szpt-login/login"
	pb "github.com/Asuka999/szpt-login/rpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("Received: Login : %s ---- Password: %s", in.Account, in.Password)

	cookies := szpt_login.Login(in.Account, in.Password).GetCookiesMap()

	var Cookies []*pb.Cookies
	for _, cookie := range cookies {
		Cookies = append(Cookies, &pb.Cookies{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Path:   cookie.Path,
			Domain: cookie.Domain,
		})
	}
	//log.Printf("Received: %v", in.Account())
	return &pb.LoginReply{
		Cookies: Cookies,
	}, nil
}

func Main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
