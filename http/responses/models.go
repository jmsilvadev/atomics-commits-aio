package responses

import "github.com/jmsilvadev/atomics-commits-aio/internal/entities"

type Transaction struct {
	Block string  `json:"block"`
	Hash  float64 `json:"hash"`
}

func (t *Transaction) ToEntity() *entities.Transaction {
	return &entities.Transaction{
		Block: t.Block,
		Hash:  t.Hash,
	}
}

func (t *Transaction) FromEntity(e *entities.Transaction) *Transaction {
	return &Transaction{
		Block: e.Block,
		Hash:  e.Hash,
	}
}
