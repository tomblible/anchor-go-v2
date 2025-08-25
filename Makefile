build:
	go build 

pump_curve:build
# 	./anchor-go  -src=./idl/dex/pump_curve.json -pkg=pump_curve -dst=./generated/prd/pump_curve
	./anchor-go -name pump_curve  --idl ./idl/fyd/pump_curve.json --output ./generated/fyd/pump_curve --program-id 6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P --no-go-mod 

pump_amm:build
	./anchor-go -name pump_amm  --idl ./idl/fyd/pump_amm.json --output ./generated/fyd/pump_amm --program-id pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA --no-go-mod 


convert:
	# anchor idl convert ./test/old.json >./test/new.json
	anchor idl convert ./test/raydium.json >./test/raydium_new.json