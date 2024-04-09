package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"proyecto2/servergRPC/model"
	pb "proyecto2/servergRPC/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	data := model.Data{
		Year:  in.GetYear(),
		Album: in.GetAlbum(),
		Name:  in.GetName(),
		Rank:  in.GetRank(),
	}
	fmt.Println("data -> ", data)

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el album"}, nil
}

func main() {

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("listening in tpc port -> ", port)
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
