package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/JonatasMSantos/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Jonatas"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(account.Id))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.BankID, bank.Id)

	_, err = model.NewAccount(bank, "", ownerName)
	require.NotNil(t, err)
}