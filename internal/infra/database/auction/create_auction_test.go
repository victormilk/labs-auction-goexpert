package auction

import (
	"context"
	"fullcycle-auction_go/internal/internal_error"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

type AuctionRepositoryMock struct {
	mock.Mock
}

func (m *AuctionRepositoryMock) CloseAuction(ctx context.Context, id string) *internal_error.InternalError {
	args := m.Called(ctx, id)
	if args.Error(0) == nil {
		return nil
	}
	return internal_error.NewInternalServerError(args.Error(0).Error())
}

func TestCloseAuctionRoutine(t *testing.T) {
	t.Run("should close auction", func(t *testing.T) {
		repository := &AuctionRepositoryMock{}
		ctx := context.Background()
		repository.On("CloseAuction", ctx, "879caf5f-c4b1-4be9-8184-e15f321130cb").Return(nil)
		closeTime := time.Now().Add(time.Millisecond * 10)
		go closeRoutine(ctx, closeTime, "879caf5f-c4b1-4be9-8184-e15f321130cb", repository)
		time.Sleep(time.Millisecond * 1)
		repository.AssertNumberOfCalls(t, "CloseAuction", 0)
		time.Sleep(time.Millisecond * 15)
		repository.AssertNumberOfCalls(t, "CloseAuction", 1)
		repository.AssertExpectations(t)
	})

	t.Run("should context cancel", func(t *testing.T) {
		repository := &AuctionRepositoryMock{}
		ctx, cancel := context.WithCancel(context.Background())
		repository.On("CloseAuction", ctx, "879caf5f-c4b1-4be9-8184-e15f321130cb").Return(nil)
		closeTime := time.Now().Add(time.Millisecond * 1)
		go closeRoutine(ctx, closeTime, "879caf5f-c4b1-4be9-8184-e15f321130cb", repository)
		cancel()
		time.Sleep(time.Millisecond * 5)
		repository.AssertNotCalled(t, "CloseAuction", ctx, "879caf5f-c4b1-4be9-8184-e15f321130cb")
		repository.AssertNumberOfCalls(t, "CloseAuction", 0)
	})
}
