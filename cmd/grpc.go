package cmd

import (
	"ewallet-framework/helpers"
	"log"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func ServeGRPC() {
	listen, err := net.Listen("tcp", ":"+helpers.GetEnv("PORT_GRPC", "7000"))
	if err != nil {
		log.Fatal("failed to listen serve grpc : ", err)
	}

	s := grpc.NewServer()

	logrus.Info("connecting to port : " + helpers.GetEnv("PORT_GRPC", "7000"))
	if err := s.Serve(listen); err != nil {
		log.Fatal("failed to listen serve grpc : ", err)
	}
}
