package core

import "fmt"

type TxPoolImpl struct {
	TxPoolRaw
	txMap map[string]Tx
}

func NewTxPool() TxPool {
	txp := TxPoolImpl{
		txMap: make(map[string]Tx),
		TxPoolRaw: TxPoolRaw{
			Txs:    make([]Tx, 0),
			TxHash: make([][]byte, 0),
		},
	}
	return &txp
}

func (tp *TxPoolImpl) Add(tx Tx) error {
	tp.txMap[common.Base58Encode(tx.Hash())] = tx
	return nil
}

func (tp *TxPoolImpl) Del(tx Tx) error {
	delete(tp.txMap, common.Base58Encode(tx.Hash()))
	return nil
}

func (tp *TxPoolImpl) Find(txHash []byte) (Tx, error) {
	tx, ok := tp.txMap[common.Base58Encode(txHash)]
	if !ok {
		return tx, fmt.Errorf("not found")
	}
	return tx, nil
}
