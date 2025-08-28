build:
	go build 

clean:
	rm -rf ./generated/fyd/*

pump_curve:build
	./anchor-go -name pump_curve  --idl ./idl/fyd/pump_curve.json --output ./generated/fyd/pump/pump_curve --program-id 6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P --no-go-mod 

pump_amm:build
	./anchor-go -name pump_amm  --idl ./idl/fyd/pump_amm.json --output ./generated/fyd/pump/pump_amm --program-id pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA --no-go-mod 

raydium_amm:build
#	anchor idl convert ./idl/fyd/raydium_amm.json >./idl/fyd/raydium_amm_new.json
	./anchor-go -name raydium_amm --type-id uint8   --idl ./idl/fyd/raydium_amm_new.json --output ./generated/fyd/raydium/raydium_amm --program-id 675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8 --no-go-mod

raydium_clmm:build
	./anchor-go -name raydium_clmm  --idl ./idl/fyd/raydium_clmm.json --output ./generated/fyd/raydium/raydium_clmm --program-id CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK --no-go-mod

raydium_cpmm:build
	./anchor-go -name raydium_cpmm  --idl ./idl/fyd/raydium_cpmm.json --output ./generated/fyd/raydium/raydium_cpmm --program-id CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C --no-go-mod

raydium_launchpad:build
	./anchor-go -name raydium_launchpad  --idl ./idl/fyd/raydium_launchpad.json --output ./generated/fyd/raydium/raydium_launchpad --program-id LanMV9sAd7wArD4vJFi2qDdfnVhFxYSUg6eADduJ3uj --no-go-mod

moonit:build
#	anchor idl convert ./idl/fyd/moonit.json >./idl/fyd/moonit_new.json
	./anchor-go -name moonit  --idl ./idl/fyd/moonit_new.json --output ./generated/fyd/moonit --program-id MoonCVVNZFSYkqNXP6bxHLPL6QQJiMagDL3qcqUQTrG --no-go-mod

whirlpool:build
#	anchor idl convert ./idl/fyd/whirlpool.json >./idl/fyd/whirlpool_new.json
	./anchor-go -name orca_whirlpool  --idl ./idl/fyd/whirlpool_new.json --output ./generated/fyd/orca/orca_whirlpool --program-id whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc --no-go-mod

meteora_damm:build
	./anchor-go -name meteora_damm  --idl ./idl/fyd/meteora_damm.json --output ./generated/fyd/meteora/meteora_damm --program-id cpamdpZCGKUy5JxQXB4dcpGPiikHawvSWAd6mEn1sGG --no-go-mod

meteora_curve:build
	./anchor-go -name meteora_curve  --idl ./idl/fyd/meteora_curve.json --output ./generated/fyd/meteora/meteora_curve --program-id dbcij3LWUppWqq96dh6gJWwBifmcGfLSB5D4DuSMaqN --no-go-mod

meteora_dlmm:build
	./anchor-go -name meteora_dlmm  --idl ./idl/fyd/meteora_dlmm.json --output ./generated/fyd/meteora/meteora_dlmm --program-id LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo --no-go-mod

meteora_pools:build
	./anchor-go -name meteora_pools  --idl ./idl/fyd/meteora_pools.json --output ./generated/fyd/meteora/meteora_pools --program-id Eo7WjKq67rjJQSZxS6z3YkapzY3eMj6Xy8X5EQVn5UaB --no-go-mod

meteora_vault:build
	./anchor-go -name meteora_vault  --idl ./idl/fyd/meteora_vault.json --output ./generated/fyd/meteora/meteora_vault --program-id 24Uqj9JCLxUeoC3hGfh5W3s9FM9uCHDS2SG3LYwBpyTi --no-go-mod

token2022:build
	./anchor-go -name token2022_program --type-id uint8  --idl ./idl/fyd/token2022_program.json --output ./generated/fyd/system/token2022_program --program-id TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb --no-go-mod

token_metadata:build
	./anchor-go -name token_metadata --type-id uint8 --idl ./idl/fyd/token_metadata.json --output ./generated/fyd/system/token_metadata --program-id metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s --no-go-mod

token_program:build
	./anchor-go -name token_program  --type-id uint8 --idl ./idl/fyd/token_program_new.json --output ./generated/fyd/system/token_program --program-id TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA --no-go-mod

associated_token:build
	./anchor-go -name associated_token --type-id uint8  --idl ./idl/fyd/associated_token_new.json --output ./generated/fyd/system/associated_token --program-id ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL --no-go-mod

all:build clean pump_curve pump_amm raydium_amm raydium_clmm raydium_cpmm raydium_launchpad moonit whirlpool meteora_damm meteora_curve meteora_dlmm meteora_pools meteora_vault

upgrade-anchor:
	avm install 0.30.1 && anchor --version 

convert:
#   anchor idl convert ./test/old.json >./test/new.json
#   "discriminator"\s*:\s*\[\s*[\s\d,]*?\s*\]\s*,? 正则表达式替换
	anchor idl convert ./test/raydium.json >./test/raydium_new.json