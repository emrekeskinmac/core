package ethwallet

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mesg-foundation/core/client/service"
)

type importInputs struct {
	Account    encryptedKeyJSONV3 `json:"account"`
	Passphrase string             `json:"passphrase"`
}

type importOutputSuccess struct {
	Address common.Address `json:"address"`
}

func (s *Ethwallet) importA(execution *service.Execution) (string, interface{}) {
	var inputs importInputs
	if err := execution.Data(&inputs); err != nil {
		return OutputError(err)
	}

	accountJSON, err := json.Marshal(inputs.Account)
	if err != nil {
		return OutputError(err)
	}

	account, err := s.keystore.Import(accountJSON, inputs.Passphrase, inputs.Passphrase)
	if err != nil {
		return OutputError(err)
	}

	return "success", importOutputSuccess{
		Address: account.Address,
	}
}
