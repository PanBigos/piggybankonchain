// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract DividendDistributor {
    address public owner;
    mapping(address => uint) public dividends;

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Not the owner");
        _;
    }

    // Function to deposit Ether into the contract and allocate it as dividends
    function allocateDividends(address[] calldata recipients, uint[] calldata amounts) external payable onlyOwner {
        require(recipients.length == amounts.length, "Recipients and amounts do not match");
        
        uint totalAmount = 0;
        for(uint i = 0; i < amounts.length; i++) {
            totalAmount += amounts[i];
        }
        
        require(msg.value >= totalAmount, "Insufficient Ether sent");
        
        for(uint i = 0; i < recipients.length; i++) {
            dividends[recipients[i]] += amounts[i];
        }

        // Refund any excess Ether sent
        uint excessAmount = msg.value - totalAmount;
        if(excessAmount > 0) {
            payable(msg.sender).transfer(excessAmount);
        }
    }

    // Function for users to withdraw their dividends
    function withdrawDividends() external {
        uint amount = dividends[msg.sender];
        require(amount > 0, "No dividends to withdraw");
        
        dividends[msg.sender] = 0; // Reset dividend balance before transfer to prevent re-entrancy attacks
        payable(msg.sender).transfer(amount);
    }

    receive() external payable {}

    // Function to check the shareholder dividend balance
    function checkUserDividendBalance(address shareholder) public view returns (uint256) {
    return dividends[shareholder];
}
}