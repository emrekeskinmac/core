package task

import (
	"github.com/mesg-foundation/application/pubsub"
	"github.com/mesg-foundation/application/service"
	"github.com/mesg-foundation/application/types"
	"golang.org/x/net/context"
)

// Execute a task
func (s *Server) Execute(ctx context.Context, request *types.ExecuteTaskRequest) (reply *types.TaskReply, err error) {
	service := service.New(request.Service)

	reply = &types.TaskReply{
		Task: request.Task,
		Data: request.Data,
	}

	go pubsub.Publish(service.TaskSubscriptionChannel(), reply)

	return
}