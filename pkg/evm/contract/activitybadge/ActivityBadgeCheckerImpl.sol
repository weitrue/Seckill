// SPDX-License-Identifier: MIT
pragma solidity 0.8.14;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./ActivityChecker.sol";

contract ActivityBadgeCheckerImpl is Checker, Ownable {

    address private _signer = 0x35253468619BcBB9585CC09aec5e5f7820914Fca;

    function setSigner(address _address) public onlyOwner {
        _signer = _address;
    }

    function verifySignature(address _userAddress, SignInfo calldata signInfo) public override view returns (bool) {
        require(block.chainid == signInfo.chainID, "Incorrect ChainId");
        bytes32 _hash = keccak256(
            abi.encode(
                signInfo.collectionAddr,
                _userAddress,
                signInfo.nftType,
                signInfo.chainID
            )
        );
        address signer = ECDSA.recover(_hash, signInfo.v, signInfo.r, signInfo.s);
        require(signer == _signer, "Incorrect signature");
        return true;
    }

}