// SPDX-License-Identifier: MIT
pragma solidity 0.8.14;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract AirDrop is Ownable, ReentrancyGuard {
    function batchTransfer(
        address from,
        address nft,
        address[] memory tos,
        uint256[] memory tokenIds
    ) external onlyOwner nonReentrant {
        for (uint256 i = 0; i < tos.length; ++i) {
            IERC721(nft).safeTransferFrom(from, tos[i], tokenIds[i]);
        }
    }
}
