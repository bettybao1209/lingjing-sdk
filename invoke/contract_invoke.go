package invoke

import (
	"errors"

	"github.com/bettybao1209/lingjing-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func GetName(client *ethclient.Client, contractAddress string) (string, error) {
	card, err := contracts.NewCard(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Warnf("get card contract error: %s", err)
		return "", errors.New("contract access error")
	}
	if name, err := card.Name(nil); err == nil {
		return name, nil
	} else {
		log.Warnf("get name error: %s", err)
		return "", errors.New("contract access error")
	}
}
