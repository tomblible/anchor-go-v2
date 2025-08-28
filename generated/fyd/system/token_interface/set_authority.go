package token_interface

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"gitlab.f-yd-8168.top/fyd168/private/solana-scraper/program/system/token2022_program"
)

func NewSetAuthorityInstruction(
	newAuthority solana.PublicKey,
	// Accounts:
	subject solana.PublicKey,
	authority solana.PublicKey,
	tokenProgram solana.PublicKey,
) (solana.Instruction, error) {
	if tokenProgram == token.ProgramID {
		return token.NewSetAuthorityInstruction(
			token.AuthorityAccountOwner,
			newAuthority,
			subject,
			authority,
			nil,
		).Build(), nil
	}
	return token2022_program.NewSetAuthorityInstruction(
		token2022_program.AuthorityType(token.AuthorityAccountOwner),
		&newAuthority,
		subject,
		authority,
		// nil,
	)
}
