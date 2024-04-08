package gmail

import (
    "context"
    "testing"
    
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockGmailClient is of Gmail Client interface
type MockGmailClient struct {
    mock.Mock
}

func TestNewGmailClient(t *testing.T) {
    t.Parallel()
    
    ctx := context.TODO()
    
    gs := NewGmailClient(ctx)
    
    assert.NotNil(t, gs, "GmailService should not be nil")
}
