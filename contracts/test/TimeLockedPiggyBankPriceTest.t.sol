// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/fuse/TimeLockedPiggyBankPriceFactory.sol";
import "../src/fuse/TimeLockedPiggyBank.sol";
import "../src/interfaces/IERC20.sol";

contract TimeLockedPiggyBankPriceTest is Test {
    TimeLockedPiggyBankPriceFactory factory;

    address owner;
    uint256 unlockDate;
    address testAddress;

    function setUp() public {
        // Setup code here
        owner = address(this);
        factory = new TimeLockedPiggyBankPriceFactory();
        unlockDate = block.timestamp + 1 days; // Set unlock date to 1 day in the future
        testAddress = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        payable(testAddress).transfer(10000000 ether);
        vm.prank(testAddress);


    }

    function testCreatePiggyBank() public {
        // Test creating a piggy bank
        address piggyBankAddress = factory.createPiggyBankPriceLimit(testAddress, unlockDate, 0, 88, 0x79E94008986d1635A2471e6d538967EBFE70A296);
        assertTrue(piggyBankAddress != address(0), "Piggy bank creation failed");

        // Check that the piggy bank is registered for the owner
        address[] memory banks = factory.getPiggyBanks(testAddress);
        assertTrue(banks.length > 0, "Piggy bank not registered");
        assertEq(banks[0], piggyBankAddress, "Registered piggy bank address mismatch");
    }

    function testWithdrawRestrictions() public {
        // Test withdrawal restrictions
        address piggyBankAddress = factory.createPiggyBankPriceLimit(testAddress, 0, 56310000000000000, 88, 0x79E94008986d1635A2471e6d538967EBFE70A296);
        TimeLockedPiggyBankPrice piggyBank = TimeLockedPiggyBankPrice(payable(piggyBankAddress));
        payable(piggyBankAddress).transfer(10000000 ether);
        vm.prank(testAddress);
        piggyBank.withdraw();
    }

}
