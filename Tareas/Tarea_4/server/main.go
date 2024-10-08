package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"tarea4/server/db"
	pb "tarea4/server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Name  string
	Album string
	Year  string
	Rank  string
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	data := Data{
		Year:  in.GetYear(),
		Album: in.GetAlbum(),
		Name:  in.GetName(),
		Rank:  in.GetRank(),
	}
	fmt.Println("data -> ", data)

	db := db.ConnectToDB()

	go db.InsertData(data.Name, data.Album, data.Year, data.Rank)

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el album"}, nil
}

func main() {

	_ = db.ConnectToDB()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
