package sig

import (
	"context"
	"errors"
	"sync"
	"time"

	"gitlab.chainedfinance.com/chaincore/r2/blockchain"

	"github.com/eddyzhou/log"
)

var (
	keys        = make(map[string]string)
	lock        = new(sync.RWMutex)
	ErrNotFound = errors.New("key not found")
)

func FetchKey(cli blockchain.Client, nodeId string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	v, err := cli.NodesCallerSession(ctx).GetNodeKey(nodeId)
	if err != nil {
		log.Error(err)
		return "", err
	}

	lock.Lock()
	keys[nodeId] = v
	lock.Unlock()
	return v, nil
}

func getKey(nodeId string) (string, error) {
	lock.RLock()
	defer lock.RUnlock()

	if v, ok := keys[nodeId]; ok {
		return v, nil
	} else {
		return "", ErrNotFound
	}
}

func GetKey(cli blockchain.Client, nodeId string) (string, error) {
	if v, err := getKey(nodeId); err != nil {
		return FetchKey(cli, nodeId)
	} else {
		return v, nil
	}
}
