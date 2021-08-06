# Smart Contracts repository

This repository contains all smart contracts for the CommonsHood project

### Pattern used
The main pattern here is the contract factory pattern and it works as follows:

There is a **Product** and a **Factory** contracts, based on the *Factory Pattern* model.

The Factory creates product (in the case of the [Token Factory](./contracts/TokenFactory.sol) the Factory creates a series of tokens based on similar behaviors, which are the same of a standard ERC20 model token).

The created products references are maintained inside the factory and there is an **ACL** to control who can create products.

## Products to handle
- Tokens
- Crowdsales
- Coupons

## Tokens
The [Token Template](./contracts/TokenTemplate.sol) contract reprents both Standard and Mintable token.
This is modeled by the constructor, if `totalSupply` is 0 the token is Mintable otherwise the supply is fixed.

## Token Factory

The [Token Factory](./contracts/TokenFactory.sol) allows to create and retrive the tokens with functions and events:
- `createToken` function to create new tokens
- `getToken` function to retrive a token by the ticker
- `mintToken` function to mint tokens 
- `TokenAdded` event to retrive all the tokens created

### Additional notes
This repository follows the truffle model, and it has been created to be used with truffle and ganache on local development, while testnet configuration skeletons can be provided later but require a wallet to be setup and a passphrase to be set and shared among developers.
