package token_interface

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"gitlab.f-yd-8168.top/fyd168/private/solana-scraper/program/system/token2022_program"
)

func NewApproveInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	source solana.PublicKey,
	delegate solana.PublicKey,
	owner solana.PublicKey,
	tokenProgram solana.PublicKey,
) (solana.Instruction, error) {
	if tokenProgram == token.ProgramID {
		return token.NewApproveInstruction(
			amount,
			// [0] = [WRITE] source
			// ··········· The source account.
			source,
			// [1] = [] delegate
			// ··········· The delegate.
			delegate,
			// [2] = [] owner
			// ··········· The source account owner.
			owner,
			// [3...] = [SIGNER] signers
			nil,
		).Build(), nil
	}

	return token2022_program.NewApproveInstruction(
		amount,
		// [0] = [WRITE] source
		// ··········· The source account.
		source,
		// [1] = [] delegate
		// ··········· The delegate.
		delegate,
		// [2] = [] owner
		// ··········· The source account owner.
		owner,
		// [3...] = [SIGNER] signers
		// nil,
	)
}
