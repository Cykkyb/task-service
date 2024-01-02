package handler

import (
	taskproto "github.com/cykkyb/proto/gen/go/task"
	"google.golang.org/grpc"
	"task-service/internal/service"
)

type serverAPI struct {
	taskproto.UnimplementedTaskServer
	task service.Task
}

func RegisterServerAPI(gRPC *grpc.Server, task service.Task) {
	taskproto.RegisterTaskServer(gRPC, &serverAPI{task: task})
}
