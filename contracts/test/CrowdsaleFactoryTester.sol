pragma solidity ^0.5.0;

import "../CrowdsaleFactory.sol";
import "../TokenCrowdsale.sol";
import "./TokenCrowdsaleTester.sol";

/**
 * @title CrowdsaleFactoryTester is a crowdsale with time-travel features.
 *        This is used to test the crowdsale contract.
 */
contract CrowdsaleFactoryTester is CrowdsaleFactory {
    constructor() CrowdsaleFactory () public { }

    /**
     * @dev createCrowdsale creates a new crowdsale.
     * @param _crowdsaleID the crowdsale ID, must be unique.
     * @param _tokenToGive the TokenTemplate instance used to give new tokens.
     * @param _tokenToAccept the TokenTemplate instance used to accept contributions.
     * @param _start the start time of the crowdsale.
     * @param _end the end time of the crowdsale.
     * @param _acceptRatio how many tokens to accept in ratio to _giveRatio.
     * @param _giveRatio how many tokens to give in ration to _acceptRatio.
     * @param _maxCap the threshold the TokenToGive tokens are released.
     * @return the address of the created crowdsale.
     */
    function createCrowdsale(
        string memory _crowdsaleID,
        address _tokenToGive,
        address _tokenToAccept,
        uint256 _start,
        uint256 _end,
        uint8 _acceptRatio,
        uint8 _giveRatio,
        uint256 _maxCap
        ) public returns(address) {
        address crowdsaleAddr = address(new TokenCrowdsaleTester(
            TokenTemplate(_tokenToGive),
            TokenTemplate(_tokenToAccept),
            _start,
            _end,
            _acceptRatio,
            _giveRatio,
            _maxCap,
            msg.sender
        ));

        CrowdsaleData memory tempData = CrowdsaleData({
            crowdsaleContract: crowdsaleAddr,
            tokenToGive: TokenTemplate(_tokenToGive),
            tokenToAccept: TokenTemplate(_tokenToAccept),
            start: _start,
            end: _end,
            acceptRatio: _acceptRatio,
            giveRatio: _giveRatio,
            maxCap: _maxCap,
            owner: msg.sender
        });

        crowdsales[_crowdsaleID] = tempData;

        emit CrowdsaleAdded(
            _crowdsaleID,
            msg.sender,
            now,
            crowdsaleAddr,
            _start,
            _end,
            _acceptRatio,
            _giveRatio,
            _maxCap,
            msg.sender
        );

        return crowdsaleAddr;
    }
}