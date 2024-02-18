// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "./TimeLockedPiggyBankPrice.sol";

contract TimeLockedPiggyBankPriceFactory {
    mapping(address => address[]) piggyBanks;
    mapping(address => address[]) createdPiggyBanks;

    function getPiggyBanks(
        address _user
    ) public view returns (address[] memory) {
        return piggyBanks[_user];
    }
    
    function getCreatedPiggyBanks(
        address _user
    ) public view returns (address[] memory) {
        return createdPiggyBanks[_user];
    }
    function createPiggyBankPriceLimit(
        address _owner,
        uint256 _unlockDate,
        uint256 _price,
        uint256 _pairIndex,
        address _supraOracle
    ) public returns (address piggyBank) {
        TimeLockedPiggyBankPrice contractPiggyBank = new TimeLockedPiggyBankPrice(
            msg.sender,
            _owner,
            _unlockDate,
            _price,
            _pairIndex,
            _supraOracle
        );
        piggyBank = address(contractPiggyBank);
        piggyBanks[_owner].push(piggyBank);
        createdPiggyBanks[msg.sender].push(piggyBank);

        emit CreatedPiggyBankPrice(
            piggyBank,
            msg.sender,
            _owner,
            block.timestamp,
            _unlockDate,
            _price
        );
    }

    event CreatedPiggyBankPrice(
        address piggyBank,
        address from,
        address to,
        uint256 createdAt,
        uint256 unlockDate,
        uint256 price
    );
}
