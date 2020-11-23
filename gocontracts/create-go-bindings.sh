#!/bin/sh

DIR=$(dirname $0)

# Compile the contract
truffle compile

# Create ABI files
echo "Creating ABI Files..."
truffle exec $DIR/abi-scripts/TokenFactory.js --network ganache
truffle exec $DIR/abi-scripts/TokenTemplate.js --network ganache
echo "DONE"

# Use abigen to create go bindings
echo "Binding contracts to newly created Go structures in .go files..."
abigen --abi TokenFactory.abi --pkg contracts --type TokenFactory --out $DIR/TokenFactory.go
abigen --abi TokenTemplate.abi --pkg contracts --type TokenTemplate --out $DIR/TokenTemplate.go
echo "DONE"
# Clean everything left
echo "Cleaning garbage..."
rm *.abi
echo "DONE"
