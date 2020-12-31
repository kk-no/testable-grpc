package client

import "context"

type MockSampleClient struct{}

func (c *MockSampleClient) Ping(ctx context.Context) {
	println("called mock ping")
}

func NewMockSampleClient() SampleClientManager {
	return &MockSampleClient{}
}
