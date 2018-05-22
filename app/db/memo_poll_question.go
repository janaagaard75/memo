package db

import (
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/jchavannes/jgo/jerr"
	"time"
)

type MemoPollQuestion struct {
	Id         uint   `gorm:"primary_key"`
	TxHash     []byte `gorm:"unique;size:50"`
	NumOptions uint
	PollType   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (m MemoPollQuestion) Save() error {
	result := save(&m)
	if result.Error != nil {
		return jerr.Get("error saving memo poll question", result.Error)
	}
	return nil
}

func (m MemoPollQuestion) GetTransactionHashString() string {
	hash, err := chainhash.NewHash(m.TxHash)
	if err != nil {
		jerr.Get("error getting chainhash from memo poll question", err).Print()
		return ""
	}
	return hash.String()
}

func GetMemoPollQuestion(txHash []byte) (*MemoPollQuestion, error) {
	var memoPollQuestion MemoPollQuestion
	err := find(&memoPollQuestion, MemoPollQuestion{
		TxHash: txHash,
	})
	if err != nil {
		return nil, jerr.Get("error getting memo poll question", err)
	}
	return &memoPollQuestion, nil
}
