// SPDX-License-Identifier: MIT
pragma solidity >=0.8.21;

/* Autogenerated file. Do not edit manually. */

/**
 * @title IInventorySystem
 * @dev This interface is automatically generated from the corresponding system contract. Do not edit manually.
 */
interface IInventorySystem {
  function meat_InventorySystem_take(address world, string[] memory tokens, uint32 rId) external returns (uint8 err);

  function meat_InventorySystem_drop(address world, string[] memory tokens, uint32 rId) external returns (uint8 err);
}
