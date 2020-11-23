pragma solidity ^0.5.0;

import "../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";
import "../node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

/**
 * @title TokenTemplate
 * @dev Very simple ERC20 Token that can be minted.
 * It is meant to be used in all crowdsale contracts.
 */
contract TokenTemplate is ERC20Detailed, ERC20Mintable {
    event Debug(string _message);

    string private _logoURL;
    bytes32 private _logoHash;
    bytes32 private _contractHash;

    constructor(
        string memory name,
        string memory symbol,
        uint8 decimals,
        string memory logoURL,
        bytes32 logoHash,
        uint256 totalSupply,
        address owner,
        bytes32 contractHash
    ) ERC20Detailed(name, symbol, decimals) ERC20Mintable() public {
        require(owner != address(0), "Owner must be defined");
        _logoURL = logoURL;
        _addMinter(owner);
        if (totalSupply > 0) {
            _mint(owner, totalSupply);
            _removeMinter(owner);
        }

        _contractHash = contractHash;
        _logoHash = logoHash;
    }

    /**
     * @return The Logo of the URL of the Token.
     */
    function logoURL() public view returns(string memory) {
        return _logoURL;
    }

    /**
     * @return The hash of the logo of the token.
     */
    function logoHash() public view returns(bytes32) {
        return _logoHash;
    }

    /**
     * @return The hash of the PDF contract of the Token.
     */
    function contractHash() public view returns(bytes32) {
        return _contractHash;
    }
}