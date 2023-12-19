// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

contract PushOracle {
  address public constant DEPOSITOR_ACCOUNT =
    0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001;

  mapping(uint256 => uint256) public data;

  function set(uint256 slot, uint256 value) external {
    require(
      msg.sender == DEPOSITOR_ACCOUNT,
      'PushOracle: only the depositor account'
    );
    data[slot] = value;
  }
}
