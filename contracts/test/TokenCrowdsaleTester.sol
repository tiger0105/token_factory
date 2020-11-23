pragma solidity ^0.5.0;

import "../TokenCrowdsale.sol";

/**
 * @title TokenCrowdsaleTester is a crowdsale with time-travel features.
 *        This is used to test the crowdsale contract.
 */
contract TokenCrowdsaleTester is TokenCrowdsale {
    constructor(
        TokenTemplate _tokenToGive,
        TokenTemplate _tokenToAccept,
        uint256 _start,
        uint256 _end,
        uint8 _acceptRatio,
        uint8 _giveRatio,
        uint256 _maxCap,
        address _owner
    ) TokenCrowdsale(
        _tokenToGive,
        _tokenToAccept,
        _start,
        _end,
        _acceptRatio,
        _giveRatio,
        _maxCap,
        _owner
    ) public { }

    /**
     * @dev Allows to go back and forth in time, this is used to test start and end
     *      of a crowdsale. Actually moves start and end of the crowdsale, but
     *      it's good for our purposes.
     *      using int256 may have overflow / underflow errors.
     */
    function timeTravel(int256 _timeSpan) public {
        uint256 span;
        if (_timeSpan > 0) {
            span = uint256(_timeSpan);
            start -= span;
            end -= span;
        } else {
            span = uint256(-_timeSpan);
            start += span;
            end += span;
        }
    }
}