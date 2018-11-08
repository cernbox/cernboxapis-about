package main

import (
	"fmt"
	"log"
	"net"

	"github.com/cernbox/cernboxapis/gen/proto/go/cernbox/about/v1"
	"github.com/cernbox/cernboxapis/gen/proto/go/cernbox/rpc"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	srv := &server{}
	aboutv1pb.RegisterAboutServiceServer(s, srv)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9901))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s.Serve(lis)
}

type server struct{}

func (s *server) ListMembers(ctx context.Context, req *aboutv1pb.ListMembersRequest) (*aboutv1pb.ListMembersResponse, error) {
	status := &rpcpb.Status{
		Code: rpcpb.Code_CODE_OK,
	}

	res := &aboutv1pb.ListMembersResponse{
		Members: getCernboxMembers(),
		Status:  status,
	}

	return res, nil
}

func (s *server) GetDocumentation(ctx context.Context, req *aboutv1pb.GetDocumentationRequest) (*aboutv1pb.GetDocumentationResponse, error) {
	status := &rpcpb.Status{
		Code: rpcpb.Code_CODE_OK,
	}

	res := &aboutv1pb.GetDocumentationResponse{
		Status:        status,
		Documentation: getDoc(),
	}

	return res, nil

}

func getCernboxMembers() []*aboutv1pb.Member {
	people := map[string]string{
		"hugo.gonzalez.labrador@cern.ch": "Hugo Gonzalez Labrador",
	}

	members := []*aboutv1pb.Member{}
	for mail, name := range people {
		m := &aboutv1pb.Member{
			Email:       mail,
			DisplayName: name,
		}
		members = append(members, m)
	}

	return members
}

func getDoc() *aboutv1pb.Documentation {
	return &aboutv1pb.Documentation{
		GithubUrl:  "https://github.com/cernbox",
		ServiceUrl: "https://cernbox.cern.ch",
	}
}
