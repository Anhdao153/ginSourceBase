package healthCheck

import (
	context "context"
	sync "sync"
)

type HealthCheckService struct {
	UnimplementedHealthCheckServer
	mu sync.Mutex
}

func (hc *HealthCheckService) HealthCheck(cxt context.Context, req *HealthCheckReq) (*HealthCheckReply, error) {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	return &HealthCheckReply{
		Message: "",
	}, nil
}
