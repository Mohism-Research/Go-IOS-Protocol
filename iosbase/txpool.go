package iosbase

type TxPool interface {
	Add(tx Tx) error
	Del(tx Tx) error
	Find(txHash []byte) (Tx, error)
	GetSlice() ([]Tx, error)
	Has(txHash []byte) (bool, error)
	Size() int
	Serializable
}

type TxPoolImpl struct {
	txs map[string]Tx
}

func (tp *TxPoolImpl) Add(tx Tx) error {
	return nil
}

func (tp *TxPoolImpl) Del(tx Tx) error {
	return nil
}

func (tp *TxPoolImpl) Find(txHash []byte) (Tx, error) {
	return Tx{}, nil
}

func (tp *TxPoolImpl) Has(txHash []byte) (bool, error) {
	return true, nil
}

func (tp *TxPoolImpl) GetSlice() ([]Tx, error) {
	return nil, nil
}

func (tp *TxPoolImpl) Size() int {
	return 0
}

func FindTxPool(hash []byte) (TxPool, error) {
	return nil, nil
}