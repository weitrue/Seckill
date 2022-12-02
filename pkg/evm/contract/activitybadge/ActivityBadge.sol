// SPDX-License-Identifier: MIT
pragma solidity 0.8.14;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "@openzeppelin/contracts/interfaces/IERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./ActivityChecker.sol";

import "erc721a/contracts/ERC721A.sol";

contract ActivityBadge is ERC721A, IERC2981, Ownable, ReentrancyGuard {
    string private metaURI;

    mapping(uint => ActivityTime) public config; // 活动时间配置
    mapping(address => bool)[] public addressMinted;
    mapping(uint => uint256) public activityTotalMinted;
    mapping(uint256 => uint) public tokenMap;    // 记录token是哪个活动创建的

    address public checker;

    event Received(address indexed, uint256);

    struct ActivityTime {
        uint beginTime;
        uint endTime;
    }

    constructor() ERC721A("Nswap badge NFT", "NSBN") {
        metaURI = "https://s3-nswap-base.nswap.com/nswapbadgenft/";
        checker = address(new Checker());
    }

    function mint(Checker.SignInfo calldata info) public {
        uint256 activity = info.nftType;
        address _address = _msgSender();
        require(info.collectionAddr == address(this), "Contract address Incorrect");
        require(Checker(checker).verifySignature(_address, info), "Signature correct");
        require(addressMinted.length > activity, "Activity is not enabled");
        require(!addressMinted[activity][_address], "Only mint one activity");
        ActivityTime memory _config = config[activity];
        uint time = block.timestamp;
        require(_config.beginTime <= time && _config.endTime >= time, "Out of action time");

        uint256 tokenId = _nextTokenId();
        addressMinted[activity][_address] = true;
        activityTotalMinted[activity]++;
        tokenMap[tokenId] = activity;
        _claim(_address, 1);
    }

    function _claim(address _address, uint64 quantity) internal {
        require(quantity > 0, "invalid number of tokens");
        _safeMint(_address, quantity);
    }

    function setCheckerAddress(address _checker) public onlyOwner {
        checker = _checker;
    }

    function setConfig(uint activity, ActivityTime calldata _config) public onlyOwner {
        require(addressMinted.length >= activity, "Activity config incorrect");
        if (addressMinted.length == activity) {
            addressMinted.push();
        }
        config[activity] = _config;
    }

    function getConfig(uint activity) public view returns (ActivityTime memory) {
        return config[activity];
    }

    function isOpen(uint activity) public view returns (bool) {
        ActivityTime memory _config = config[activity];
        uint time = block.timestamp;
        return _config.beginTime <= time && _config.endTime >= time;
    }

    function totalActivity() public view returns (uint) {
        return addressMinted.length;
    }

    function burn(uint256 tokenId) external {
        _burn(tokenId, true);
    }

    function totalMinted() public view returns (uint256) {
        return _totalMinted();
    }

    /***************Royalty***************/
    function supportsInterface(bytes4 interfaceId)
    public
    view
    virtual
    override(ERC721A, IERC165)
    returns (bool)
    {
        return
        interfaceId == type(IERC2981).interfaceId ||
        super.supportsInterface(interfaceId);
    }

    function royaltyInfo(uint256 tokenId, uint256 salePrice)
    external
    view
    override
    returns (address receiver, uint256 royaltyAmount)
    {
        require(_exists(tokenId), "query for nonexistent token");
        return (address(this), (salePrice * 250) / 10000);
    }

    function withdraw() external onlyOwner nonReentrant {
        (bool success, ) = _msgSender().call{value: address(this).balance}("");
        require(success, "withdraw failed");
    }

    function withdrawTokens(IERC20 token) external onlyOwner nonReentrant {
        uint256 balance = token.balanceOf(address(this));
        token.transfer(_msgSender(), balance);
    }

    receive() external payable {
        emit Received(_msgSender(), msg.value);
    }

    /***************TokenURI***************/
    function setTokenURI(string calldata tokenURI_) external onlyOwner {
        metaURI = tokenURI_;
    }

    function tokenURI(uint256 tokenId)
    public
    view
    virtual
    override
    returns (string memory)
    {
        require(_exists(tokenId), "query for nonexistent token");
        return string(abi.encodePacked(metaURI, Strings.toString(tokenMap[tokenId])));
    }

}