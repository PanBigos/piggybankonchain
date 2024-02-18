// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/fuse/TimeLockedPiggyBankFactory.sol";
import "../src/fuse/TimeLockedPiggyBank.sol";
import "../src/interfaces/IERC20.sol";

// contract MockERC20 is IERC20 {
//     // Implement a simple ERC20 mock for testing
// }

contract TimeLockedPiggyBankTest is Test {
    TimeLockedPiggyBankFactory factory;
    // MockERC20 token;
    address owner;
    uint256 unlockDate;
    address testAddress;

    function setUp() public {
        // Setup code here
        owner = address(this);
        factory = new TimeLockedPiggyBankFactory();
        // token = new MockERC20();
        unlockDate = block.timestamp + 1 days; // Set unlock date to 1 day in the future
        testAddress = address(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266);
        payable(testAddress).transfer(10000000 ether);
        vm.prank(testAddress);

        // Mint some tokens to the piggy bank (if necessary for tests)
        // token.mint(address(factory), 1000 * 1e18);
    }

    function testCreatePiggyBank() public {
        // Test creating a piggy bank
        address piggyBankAddress = factory.createPiggyBank(testAddress, unlockDate);
        assertTrue(piggyBankAddress != address(0), "Piggy bank creation failed");

        // Check that the piggy bank is registered for the owner
        address[] memory banks = factory.getPiggyBanks(testAddress);
        assertTrue(banks.length > 0, "Piggy bank not registered");
        assertEq(banks[0], piggyBankAddress, "Registered piggy bank address mismatch");
    }

    function testWithdrawRestrictions() public {
        address piggyBankAddress = factory.createPiggyBank(testAddress, 0);
        TimeLockedPiggyBank piggyBank = TimeLockedPiggyBank(payable(piggyBankAddress));
        payable(piggyBankAddress).transfer(10000000 ether);
        vm.prank(testAddress);
        piggyBank.withdraw();
        vm.expectRevert("Lock period not yet expired");
    }

