// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

interface ISupraSValueFeedStorage {

    struct dataWithoutHcc {
        uint256 round;
        uint256 decimals;
        uint256 time;
        uint256 price;

    }

function getSvalue(uint256 _pairIndex)
        external
        view
        returns (dataWithoutHcc memory);
}