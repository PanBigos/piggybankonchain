// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "./TimeLockedPiggyBank.sol";

contract TimeLockedPiggyBankFactory {
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

    function createPiggyBank(
        address _owner,
        uint256 _unlockDate
    ) public returns (address piggyBank) {
        TimeLockedPiggyBank contractPiggyBank = new TimeLockedPiggyBank(
            msg.sender,
            _owner,
            _unlockDate
        );
        piggyBank = address(contractPiggyBank);
        piggyBanks[_owner].push(piggyBank);
        createdPiggyBanks[msg.sender].push(piggyBank);

        emit CreatedPiggyBank(
            piggyBank,
            msg.sender,
            _owner,
            block.timestamp,
            _unlockDate
        );
    }

    event CreatedPiggyBank(
        address piggyBank,
        address creator,
        address owner,
        uint256 createdAt,
        uint256 unlockDate
    );
}