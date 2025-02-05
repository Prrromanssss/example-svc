package grpc

import "context"

type SomeClient struct {
}

func NewSomeClient() *SomeClient {
	return &SomeClient{}
}

func (s *SomeClient) GetChatIDByUserID(ctx context.Context, userID int64) (int64, error) {
	return 0, nil
}
