package token_interface

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
)

func NewCreateAccountInstruction(
	newAccount solana.PublicKey,
	payer solana.PublicKey,
	tokenProgram solana.PublicKey,
) solana.Instruction {
	var (
		lamports = uint64(2039280)
		space    = uint64(165)
	)

	if tokenProgram != token.ProgramID {
		lamports = 2157600
		space = 182
	}

	return system.NewCreateAccountInstruction(
		// Number of lamports to transfer to the new account
		lamports,
		// Number of bytes of memory to allocate
		space,
		// Address of program that will own the new account
		tokenProgram,
		// [0] = [WRITE, SIGNER] FundingAccount
		// ··········· Funding account
		payer,
		// [1] = [WRITE, SIGNER] NewAccount
		newAccount,
	).Build()
}
