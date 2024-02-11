// SPDX-License-Identifier: MIT
//  _____     _ _         _         _
// |_   _|_ _(_) |_____  | |   __ _| |__ ___
//   | |/ _` | | / / _ \ | |__/ _` | '_ (_-<
//   |_|\__,_|_|_\_\___/ |____\__,_|_.__/__/
//
//   Email: security@taiko.xyz
//   Website: https://taiko.xyz
//   GitHub: https://github.com/taikoxyz
//   Discord: https://discord.gg/taikoxyz
//   Twitter: https://twitter.com/taikoxyz
//   Blog: https://mirror.xyz/labs.taiko.eth
//   Youtube: https://www.youtube.com/@taikoxyz

pragma solidity 0.8.24;

import "lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "../../libs/LibMath.sol";
import "./MerkleClaimable.sol";

/// @title ERC20Airdrop2
/// Contract for managing Taiko token airdrop for eligible users but the
/// withdrawal is not immediate and is subject to a withdrawal window.
contract ERC20Airdrop2 is MerkleClaimable {
    using LibMath for uint256;

    address public token;
    address public vault;
    // Represents the token amount for which the user is (by default) eligible
    mapping(address => uint256) public claimedAmount;
    // Represents the already withdrawn amount
    mapping(address => uint256) public withdrawnAmount;
    // Length of the withdrawal window
    uint64 public withdrawalWindow;

    uint256[45] private __gap;

    event Withdrawn(address user, uint256 amount);

    error WITHDRAWALS_NOT_ONGOING();

    modifier ongoingWithdrawals() {
        if (claimEnd > block.timestamp || claimEnd + withdrawalWindow < block.timestamp) {
            revert WITHDRAWALS_NOT_ONGOING();
        }
        _;
    }

    function init(
        uint64 _claimStart,
        uint64 _claimEnd,
        bytes32 _merkleRoot,
        address _token,
        address _vault,
        uint64 _withdrawalWindow
    )
        external
        initializer
    {
        __Essential_init();
        __MerkleClaimable_init(_claimStart, _claimEnd, _merkleRoot);

        token = _token;
        vault = _vault;
        withdrawalWindow = _withdrawalWindow;
    }

    function claim(address user, uint256 amount, bytes32[] calldata proof) external nonReentrant {
        // Check if this can be claimed
        _verifyClaim(abi.encode(user, amount), proof);

        // Assign the tokens
        claimedAmount[user] += amount;
    }

    /// @notice External withdraw function
    /// @param user User address
    function withdraw(address user) external ongoingWithdrawals {
        (, uint256 amount) = getBalance(user);
        withdrawnAmount[user] += amount;
        IERC20(token).transferFrom(vault, user, amount);

        emit Withdrawn(user, amount);
    }

    /// @notice Getter for the balance and withdrawal amount per given user
    /// The 2nd airdrop is subject to an unlock period. User has to claim his
    /// tokens (within claimStart and claimEnd), but not immediately
    /// withdrawable. With a time of X (withdrawalWindow) it becomes fully
    /// withdrawable - and unlocks linearly.
    /// @param user User address
    /// @return balance The balance the user successfully claimed
    /// @return withdrawableAmount The amount available to withdraw
    function getBalance(address user)
        public
        view
        returns (uint256 balance, uint256 withdrawableAmount)
    {
        balance = claimedAmount[user];
        // If balance is 0 then there is no balance and withdrawable amount
        if (balance == 0) return (0, 0);
        // Balance might be positive before end of claiming (claimEnd - if claimed already) but
        // withdrawable is 0.
        if (block.timestamp < claimEnd) return (balance, 0);

        // Hard cap timestamp - so range cannot go over - to get more allocation over time.
        uint256 timeBasedAllowance = balance
            * (block.timestamp.min(claimEnd + withdrawalWindow) - claimEnd) / withdrawalWindow;

        withdrawableAmount = timeBasedAllowance - withdrawnAmount[user];
    }
}
