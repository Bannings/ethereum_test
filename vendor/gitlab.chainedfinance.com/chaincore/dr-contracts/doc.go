package dr_contracts

//go:generate abigen -abi ./build/ar.abi -bin ./build/ar.bin -solc solc -type ARStore -pkg dr_contracts -out ./ARStore.go
//go:generate abigen -abi ./build/node.abi -bin ./build/node.bin -solc solc -type CFNodes -pkg dr_contracts -out ./CFNodes.go
