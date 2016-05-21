package main

import (
	"log"
	"net"

	pb "github.com/liujunandzhou/snowflake/idserver/idserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

import "github.com/liujunandzhou/snowflake"
import "time"
import "strconv"
import "fmt"
import "os"

const (
	port = ":50051"
)

//定义server结构体
type IdServer struct {
	idIns *idgen.IdMaker
}

//创建一个server
func NewIdServer(host string, duration time.Duration) *IdServer {

	zkIns := idgen.NewZkGen(host, duration)

	if zkIns == nil {
		fmt.Println("idgen.NewZkGen failed")
		return nil
	}

	idIns := idgen.NewIdMaker(zkIns)

	ins := new(IdServer)

	ins.idIns = idIns

	return ins
}

func (this *IdServer) GetId(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	var sUniqId string

	iId := this.idIns.GenId()

	sUniqId = strconv.FormatUint(iId, 10)

	return &pb.Response{Uniqid: &sUniqId}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	ins := NewIdServer("localhost:2181", 5e9)

	if ins == nil {
		log.Fatalf("NewIdServer failed")
		os.Exit(1)
	}

	pb.RegisterIdServerServer(s, ins)

	s.Serve(lis)
}
