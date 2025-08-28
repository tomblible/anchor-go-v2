package token_interface

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"gitlab.f-yd-8168.top/fyd168/private/solana-scraper/program/system/token2022_program"
)

func NewInitializeAccount3Instruction(
	// Parameters:
	owner solana.PublicKey,
	// Accounts:
	account solana.PublicKey,
	mint solana.PublicKey,
	tokenProgram solana.PublicKey,
) (solana.Instruction, error) {
	if tokenProgram == token.ProgramID {
		return token.NewInitializeAccount3Instruction(
			// The new account's owner/multisignature.
			owner,
			// [0] = [WRITE] account
			// ··········· The account to initialize.
			account,
			// [1] = [] mint
			// ··········· The mint this account will be associated with.
			mint,
		).Build(), nil
	}

	return token2022_program.NewInitializeAccount3Instruction(
		// The new account's owner/multisignature.
		owner,
		// [0] = [WRITE] account
		// ··········· The account to initialize.
		account,
		// [1] = [] mint
		// ··········· The mint this account will be associated with.
		mint,
	)
}
