// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;
import "../interfaces/IERC20.sol";

contract TimeLockedPiggyBank {
    address public creator;
    address public owner;
    uint256 public unlockDate;
    uint256 public createdAt;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    constructor(
        address _creator,
        address _owner,
        uint256 _unlockDate
    ) {
        creator = _creator;
        owner = _owner;
        unlockDate = _unlockDate;
        createdAt = block.timestamp;
    }
    
    receive() external payable {}

    function withdraw() public onlyOwner {
        require(block.timestamp >= unlockDate);

        uint256 balance = address(this).balance;
        (bool sent, bytes memory data) = payable(owner).call{value: balance}(
            ""
        );
        require(sent, "Transfer failed");
        emit Withdrew(msg.sender, balance);
    }

    function withdrawToken(address _tokenContract) public onlyOwner {
        require(block.timestamp >= unlockDate);
        IERC20 token = IERC20(_tokenContract);

        uint256 tokenBalance = token.balanceOf(address(this));
        require(token.transfer(owner, tokenBalance), "Token transfer failed");
        emit WithdrewTokens(_tokenContract, msg.sender, tokenBalance);
    }

    function info()
        public
        view
        returns (address, address, uint256, uint256, uint256)
    {
        return (
            creator,
            owner,
            unlockDate,
            createdAt,
            address(this).balance
        );
    }

    event Withdrew(address to, uint256 amount);
    event WithdrewTokens(address tokenContract, address to, uint256 amount);
}
