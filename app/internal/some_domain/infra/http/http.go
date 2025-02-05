package http

type SomeClient struct {
}

func NewSomeClient() *SomeClient {
	return &SomeClient{}
}

func (s *SomeClient) CanAccess(methodID int) (bool, error) {
	return false, nil
}
