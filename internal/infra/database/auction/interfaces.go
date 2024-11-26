package auction

import (
	"context"
	"fullcycle-auction_go/internal/internal_error"
)

type AuctionRepositoryInterface interface {
	// CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError
	CloseAuction(ctx context.Context, id string) *internal_error.InternalError
	// FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError)
	// FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category string, productName string) ([]auction_entity.Auction, *internal_error.InternalError)
}
