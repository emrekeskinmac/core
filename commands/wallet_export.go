package commands

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	survey "gopkg.in/AlecAivazis/survey.v1"
)

type walletExportCmd struct {
	baseWalletCmd

	address common.Address

	e WalletExecutor
}

func newWalletExportCmd(e WalletExecutor) *walletExportCmd {
	c := &walletExportCmd{e: e}
	c.cmd = newCommand(&cobra.Command{
		Use:     "export",
		Short:   "export an account",
		Long:    "export an existing account in order to backup it and import it in an other wallet",
		Example: "mesg-core wallet export",
		Args:    cobra.ExactArgs(1),
		PreRunE: c.preRunE,
		RunE:    c.runE,
	})
	c.setupFlags()
	return c
}

func (c *walletExportCmd) preRunE(cmd *cobra.Command, args []string) error {
	if !c.noPassphrase && c.passphrase == "" {
		if err := survey.AskOne(&survey.Password{
			Message: "Enter passphrase",
		}, &c.passphrase, survey.MinLength(1)); err != nil {
			return err
		}
	}
	if !common.IsHexAddress(args[0]) {
		return errInvalidAddress
	}
	c.address = common.HexToAddress(args[0])
	return nil
}

func (c *walletExportCmd) runE(cmd *cobra.Command, args []string) error {
	account, err := c.e.Export(c.address, c.passphrase)
	if err != nil {
		return err
	}
	b, err := json.Marshal(account)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}