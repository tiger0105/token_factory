pragma solidity ^0.5.0;

import "./TokenTemplate.sol";
import "../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";

/**
 * @title TokenFactory is a contract following Factory
 *        Pattern to generate tokens.
 */
contract TokenFactory {
    using SafeMath for uint256;

    struct TokenData {
        address contractAddress;
        string name;
        string symbol;
        uint8 decimals;
        string logoURL;
        address owner;
        uint256 hardCap;
        bool mintable;
    }


    mapping(string => TokenData) internal tokens;
    mapping(address => bool) internal isAdmin;

    //TokenData[] public fullData;

    event TokenAdded(
        address indexed _from,
        uint256 _timestamp,
        address _contractAddress,
        string _name,
        string _symbol,
        uint8 _decimals,
        string _logoURL,
        uint256 _hardCap
    );
    event AdminAdded(address indexed _from, address indexed _who);

    event factoryDebug(address indexed _from, string _op, string _msg);

    constructor() public {
        isAdmin[msg.sender] = true;
        emit AdminAdded(address(0), msg.sender);
    }

    /**
     * @dev Callable only by an admin, makeAdmin adds another specified address as admin
     *      of the TokenFactory.
     * @param _address the address to make admin
     */
    function makeAdmin(address _address) public {
        require(isAdmin[msg.sender], "Must be admin to create another admin");
        isAdmin[_address] = true;
        emit AdminAdded(msg.sender, _address);
    }

    /**
     * @dev getToken gets a created Token data by its symbol.
     * @param _symbol the symbol of the Token.
     * @return the Token data in array form.
     */
    function getToken(string memory _symbol) public view returns(address, string memory, string memory, uint8, string memory, address, bool) {
        return (
            tokens[_symbol].contractAddress,
            tokens[_symbol].name,
            tokens[_symbol].symbol,
            tokens[_symbol].decimals,
            tokens[_symbol].logoURL,
            tokens[_symbol].owner,
            tokens[_symbol].mintable
        );
    }

/*
    function getTokenList() public view returns(TokenData[]) {
        return fullData;
    }

    function getTokenCount() public view returns(uint256) {
        return fullData.length;
    }
*/

    /**
     * @dev createToken creates a new TokenTemplate instance.
     * @param _name the name of the Token
     * @param _symbol the symbol of the Token.
     * @param _decimals the decimals of the Token. Must be between 0 and 18 included.
     * @param _logoURL the URL of the image representing the logo of the Token.
     * @param _logoHash the Hash of the image pointed by _logoURL, to ensure it has not been altered.
     * @param _hardCap the total supply of the Token.
     * @param _contractHash the Hash of the PDF contract bound to this Token.
     */
    function createToken(
        string memory _name,
        string memory _symbol,
        uint8 _decimals,
        string memory _logoURL,
        bytes32 _logoHash,
        uint256 _hardCap,
        bytes32 _contractHash
    ) public {
        // symbol must not be already be created for this token, since it's unique,
        // so we check for 0 value struct (with 0 value contract address).
        require(tokens[_symbol].contractAddress == address(0), "symbol must not be already be created for this token, since it's unique");

        // Generates the new contract and stores its address.
        // if hardCap = 0 token is mintable, otherwise it is capped to hardCap and all money is transfered to the creator
        // of the token.
        TokenTemplate tokenContract = new TokenTemplate(_name, _symbol, _decimals, _logoURL, _logoHash, _hardCap, msg.sender, _contractHash);

        TokenData memory tempTokenData = TokenData({
            contractAddress: address(tokenContract),
            name: _name,
            symbol: _symbol,
            decimals: _decimals,
            logoURL:_logoURL,
            owner: msg.sender,
            hardCap: _hardCap,
            mintable: false
        });

        if (_hardCap == 0) {
            // set mintable to true
            tempTokenData.mintable = true;
            // mint permission for all admins tokenContract.addMinter(admin)
        }

        tokens[_symbol] = tempTokenData;

        //fullData.push(tempTokenData);

        emit TokenAdded(
            msg.sender,
            now,
            address(tokenContract),
            _name,
            _symbol,
            _decimals,
            _logoURL,
            _hardCap
        );
    }
}
