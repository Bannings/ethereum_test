package contract_gen

//go:generate abigen -abi ./build/FuxToken.abi -bin ./build/FuxToken.bin -solc solc -type FuxToken -pkg contract_gen -out ./FuxToken.go
//go:generate abigen -abi ./build/FuxPayBox.abi -bin ./build/FuxPayBox.bin -solc solc -type FuxPayBox -pkg contract_gen -out ./FuxPayBox.go
//go:generate abigen -abi ./build/FuxPayBoxFactory.abi -bin ./build/FuxPayBoxFactory.bin -solc solc -type FuxPayBoxFactory -pkg contract_gen -out ./FuxPayBoxFactory.go
