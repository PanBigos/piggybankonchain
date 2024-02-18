// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "../interfaces/IERC20.sol";

contract PiggyBankRouter {
    address payable public owner;
    address payable public feeCollector;
    uint public feeRate; // Fee rate in percentage
    uint public maxFee; // Maximum fee in percentage


    // Event to emit when a transfer occurs
    event Transfer(address indexed to, uint amount, uint fee);
    event TokenTransfer(
        address indexed token,
        address indexed to,
        uint amount,
        uint fee
    );

    // Modifier to check if the caller is the owner
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action.");
        _;
    }

    // Constructor to set the owner, the fee rate, and the maximum fee
    constructor(uint _feeRate, uint _maxFee, address payable _feeCollector) {
        require(
            _feeRate <= _maxFee,
            "Initial fee rate must not exceed the maximum fee"
        );
        require(_feeCollector != address(0), "_newFeeCollector address cannot be the zero address");
        
        feeCollector = _feeCollector;
        owner = payable(msg.sender);
        feeRate = _feeRate;
        maxFee = _maxFee;
    }
    receive() external payable {
    }
    // Function to transfer amount minus fee to a specific address
    function transferWithFee(address payable _to) public payable {
        require(msg.value > 0, "Value must be greater than 0");

        // Calculate fee
        uint fee = (msg.value * feeRate) / 10000;

        // Calculate remaining amount
        uint remainingAmount = msg.value - fee;

        // Transfer fee to the owner
        (bool sent, ) = feeCollector.call{value:fee}("");
        require(sent, "Failed to send fee Ether");

        // Transfer the remaining amount to the specified address
        (bool sentTo, ) = _to.call{value: remainingAmount}("");
        require(sentTo, "Failed to send donation Ether");
        emit Transfer(_to, remainingAmount, fee);
    }

    // Function to transfer ERC20 token amount minus fee to a specific address
    function transferTokenWithFee(
        IERC20 token,
        address _to,
        uint _amount
    ) public {
        require(_amount > 0, "Amount must be greater than 0");

        // Calculate fee
        uint fee = (_amount * feeRate) / 10000;

        uint remainingAmount = _amount - fee;

        // Transfer fee to the owner
        require(
            token.transferFrom(msg.sender, feeCollector, fee),
            "Fee transfer failed"
        );

        // Transfer the remaining amount to the specified address
        require(
            token.transferFrom(msg.sender, _to, remainingAmount),
            "Transfer failed"
        );

        emit TokenTransfer(address(token), _to, remainingAmount, fee);
    }

    // Allow owner to change the feeCollector
    function setFeeCollector(address payable _newFeeCollector) public onlyOwner {        
        require(_newFeeCollector != address(0), "_newFeeCollector address cannot be the zero address");
        feeCollector = _newFeeCollector;
    }
    // Allow owner to change the fee rate
    function setFeeRate(uint _newRate) public onlyOwner {
        require(
            _newRate <= maxFee,
            "New fee rate must not exceed the maximum fee"
        );
        feeRate = _newRate;
    }    

    // Function to withdraw funds stored in the contract (for owner)
    function withdraw() public onlyOwner {
        owner.transfer(address(this).balance);
    }

    // Function to withdraw ERC20 funds stored in the contract (for owner)
    function withdrawTokens(IERC20 token) public onlyOwner {
        uint256 balance = token.balanceOf(address(this));
        require(token.transfer(owner, balance), "Transfer failed");
    }
}