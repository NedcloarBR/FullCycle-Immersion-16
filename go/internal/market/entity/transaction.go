package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int
	Price        float64
	Total        float64
	DateTime     time.Time
}

func NewTransaction(sellingOrder *Order, buyingOrder *Order, shares int, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		DateTime:     time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int, price float64) {
	t.Total = float64(shares) * price
}

func (t *Transaction) CloseBuyOrderTransaction() {
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellOrderTransaction() {
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) UpdateBuyingInvestor(minShares int) {
	t.BuyingOrder.Investor.UpdateAssetPosition(t.BuyingOrder.Asset.ID, minShares)
	t.BuyingOrder.PendingShares -= minShares
}

func (t *Transaction) UpdateSellingInvestor(minShares int) {
	t.SellingOrder.Investor.UpdateAssetPosition(t.SellingOrder.Asset.ID, -minShares)
	t.SellingOrder.PendingShares -= minShares
}
