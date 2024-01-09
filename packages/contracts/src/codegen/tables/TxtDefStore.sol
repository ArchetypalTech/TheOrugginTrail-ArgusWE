// SPDX-License-Identifier: MIT
pragma solidity >=0.8.21;

/* Autogenerated file. Do not edit manually. */

// Import schema type
import { SchemaType } from "@latticexyz/schema-type/src/solidity/SchemaType.sol";

// Import store internals
import { IStore } from "@latticexyz/store/src/IStore.sol";
import { StoreSwitch } from "@latticexyz/store/src/StoreSwitch.sol";
import { StoreCore } from "@latticexyz/store/src/StoreCore.sol";
import { Bytes } from "@latticexyz/store/src/Bytes.sol";
import { Memory } from "@latticexyz/store/src/Memory.sol";
import { SliceLib } from "@latticexyz/store/src/Slice.sol";
import { EncodeArray } from "@latticexyz/store/src/tightcoder/EncodeArray.sol";
import { FieldLayout, FieldLayoutLib } from "@latticexyz/store/src/FieldLayout.sol";
import { Schema, SchemaLib } from "@latticexyz/store/src/Schema.sol";
import { PackedCounter, PackedCounterLib } from "@latticexyz/store/src/PackedCounter.sol";
import { ResourceId } from "@latticexyz/store/src/ResourceId.sol";
import { RESOURCE_TABLE, RESOURCE_OFFCHAIN_TABLE } from "@latticexyz/store/src/storeResourceTypes.sol";

// Import user types
import { TxtDefType } from "./../common.sol";

ResourceId constant _tableId = ResourceId.wrap(
  bytes32(abi.encodePacked(RESOURCE_TABLE, bytes14("meat"), bytes16("TxtDefStore")))
);
ResourceId constant TxtDefStoreTableId = _tableId;

FieldLayout constant _fieldLayout = FieldLayout.wrap(
  0x0005020104010000000000000000000000000000000000000000000000000000
);

struct TxtDefStoreData {
  uint32 owner;
  TxtDefType txtDefType;
  string value;
}

library TxtDefStore {
  /**
   * @notice Get the table values' field layout.
   * @return _fieldLayout The field layout for the table.
   */
  function getFieldLayout() internal pure returns (FieldLayout) {
    return _fieldLayout;
  }

  /**
   * @notice Get the table's key schema.
   * @return _keySchema The key schema for the table.
   */
  function getKeySchema() internal pure returns (Schema) {
    SchemaType[] memory _keySchema = new SchemaType[](1);
    _keySchema[0] = SchemaType.BYTES32;

    return SchemaLib.encode(_keySchema);
  }

  /**
   * @notice Get the table's value schema.
   * @return _valueSchema The value schema for the table.
   */
  function getValueSchema() internal pure returns (Schema) {
    SchemaType[] memory _valueSchema = new SchemaType[](3);
    _valueSchema[0] = SchemaType.UINT32;
    _valueSchema[1] = SchemaType.UINT8;
    _valueSchema[2] = SchemaType.STRING;

    return SchemaLib.encode(_valueSchema);
  }

  /**
   * @notice Get the table's key field names.
   * @return keyNames An array of strings with the names of key fields.
   */
  function getKeyNames() internal pure returns (string[] memory keyNames) {
    keyNames = new string[](1);
    keyNames[0] = "txtDefId";
  }

  /**
   * @notice Get the table's value field names.
   * @return fieldNames An array of strings with the names of value fields.
   */
  function getFieldNames() internal pure returns (string[] memory fieldNames) {
    fieldNames = new string[](3);
    fieldNames[0] = "owner";
    fieldNames[1] = "txtDefType";
    fieldNames[2] = "value";
  }

  /**
   * @notice Register the table with its config.
   */
  function register() internal {
    StoreSwitch.registerTable(_tableId, _fieldLayout, getKeySchema(), getValueSchema(), getKeyNames(), getFieldNames());
  }

  /**
   * @notice Register the table with its config.
   */
  function _register() internal {
    StoreCore.registerTable(_tableId, _fieldLayout, getKeySchema(), getValueSchema(), getKeyNames(), getFieldNames());
  }

  /**
   * @notice Get owner.
   */
  function getOwner(bytes32 txtDefId) internal view returns (uint32 owner) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes32 _blob = StoreSwitch.getStaticField(_tableId, _keyTuple, 0, _fieldLayout);
    return (uint32(bytes4(_blob)));
  }

  /**
   * @notice Get owner.
   */
  function _getOwner(bytes32 txtDefId) internal view returns (uint32 owner) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes32 _blob = StoreCore.getStaticField(_tableId, _keyTuple, 0, _fieldLayout);
    return (uint32(bytes4(_blob)));
  }

  /**
   * @notice Set owner.
   */
  function setOwner(bytes32 txtDefId, uint32 owner) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.setStaticField(_tableId, _keyTuple, 0, abi.encodePacked((owner)), _fieldLayout);
  }

  /**
   * @notice Set owner.
   */
  function _setOwner(bytes32 txtDefId, uint32 owner) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.setStaticField(_tableId, _keyTuple, 0, abi.encodePacked((owner)), _fieldLayout);
  }

  /**
   * @notice Get txtDefType.
   */
  function getTxtDefType(bytes32 txtDefId) internal view returns (TxtDefType txtDefType) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes32 _blob = StoreSwitch.getStaticField(_tableId, _keyTuple, 1, _fieldLayout);
    return TxtDefType(uint8(bytes1(_blob)));
  }

  /**
   * @notice Get txtDefType.
   */
  function _getTxtDefType(bytes32 txtDefId) internal view returns (TxtDefType txtDefType) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes32 _blob = StoreCore.getStaticField(_tableId, _keyTuple, 1, _fieldLayout);
    return TxtDefType(uint8(bytes1(_blob)));
  }

  /**
   * @notice Set txtDefType.
   */
  function setTxtDefType(bytes32 txtDefId, TxtDefType txtDefType) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.setStaticField(_tableId, _keyTuple, 1, abi.encodePacked(uint8(txtDefType)), _fieldLayout);
  }

  /**
   * @notice Set txtDefType.
   */
  function _setTxtDefType(bytes32 txtDefId, TxtDefType txtDefType) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.setStaticField(_tableId, _keyTuple, 1, abi.encodePacked(uint8(txtDefType)), _fieldLayout);
  }

  /**
   * @notice Get value.
   */
  function getValue(bytes32 txtDefId) internal view returns (string memory value) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes memory _blob = StoreSwitch.getDynamicField(_tableId, _keyTuple, 0);
    return (string(_blob));
  }

  /**
   * @notice Get value.
   */
  function _getValue(bytes32 txtDefId) internal view returns (string memory value) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    bytes memory _blob = StoreCore.getDynamicField(_tableId, _keyTuple, 0);
    return (string(_blob));
  }

  /**
   * @notice Set value.
   */
  function setValue(bytes32 txtDefId, string memory value) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.setDynamicField(_tableId, _keyTuple, 0, bytes((value)));
  }

  /**
   * @notice Set value.
   */
  function _setValue(bytes32 txtDefId, string memory value) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.setDynamicField(_tableId, _keyTuple, 0, bytes((value)));
  }

  /**
   * @notice Get the length of value.
   */
  function lengthValue(bytes32 txtDefId) internal view returns (uint256) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    uint256 _byteLength = StoreSwitch.getDynamicFieldLength(_tableId, _keyTuple, 0);
    unchecked {
      return _byteLength / 1;
    }
  }

  /**
   * @notice Get the length of value.
   */
  function _lengthValue(bytes32 txtDefId) internal view returns (uint256) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    uint256 _byteLength = StoreCore.getDynamicFieldLength(_tableId, _keyTuple, 0);
    unchecked {
      return _byteLength / 1;
    }
  }

  /**
   * @notice Get an item of value.
   * @dev Reverts with Store_IndexOutOfBounds if `_index` is out of bounds for the array.
   */
  function getItemValue(bytes32 txtDefId, uint256 _index) internal view returns (string memory) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    unchecked {
      bytes memory _blob = StoreSwitch.getDynamicFieldSlice(_tableId, _keyTuple, 0, _index * 1, (_index + 1) * 1);
      return (string(_blob));
    }
  }

  /**
   * @notice Get an item of value.
   * @dev Reverts with Store_IndexOutOfBounds if `_index` is out of bounds for the array.
   */
  function _getItemValue(bytes32 txtDefId, uint256 _index) internal view returns (string memory) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    unchecked {
      bytes memory _blob = StoreCore.getDynamicFieldSlice(_tableId, _keyTuple, 0, _index * 1, (_index + 1) * 1);
      return (string(_blob));
    }
  }

  /**
   * @notice Push a slice to value.
   */
  function pushValue(bytes32 txtDefId, string memory _slice) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.pushToDynamicField(_tableId, _keyTuple, 0, bytes((_slice)));
  }

  /**
   * @notice Push a slice to value.
   */
  function _pushValue(bytes32 txtDefId, string memory _slice) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.pushToDynamicField(_tableId, _keyTuple, 0, bytes((_slice)));
  }

  /**
   * @notice Pop a slice from value.
   */
  function popValue(bytes32 txtDefId) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.popFromDynamicField(_tableId, _keyTuple, 0, 1);
  }

  /**
   * @notice Pop a slice from value.
   */
  function _popValue(bytes32 txtDefId) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.popFromDynamicField(_tableId, _keyTuple, 0, 1);
  }

  /**
   * @notice Update a slice of value at `_index`.
   */
  function updateValue(bytes32 txtDefId, uint256 _index, string memory _slice) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    unchecked {
      bytes memory _encoded = bytes((_slice));
      StoreSwitch.spliceDynamicData(_tableId, _keyTuple, 0, uint40(_index * 1), uint40(_encoded.length), _encoded);
    }
  }

  /**
   * @notice Update a slice of value at `_index`.
   */
  function _updateValue(bytes32 txtDefId, uint256 _index, string memory _slice) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    unchecked {
      bytes memory _encoded = bytes((_slice));
      StoreCore.spliceDynamicData(_tableId, _keyTuple, 0, uint40(_index * 1), uint40(_encoded.length), _encoded);
    }
  }

  /**
   * @notice Get the full data.
   */
  function get(bytes32 txtDefId) internal view returns (TxtDefStoreData memory _table) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    (bytes memory _staticData, PackedCounter _encodedLengths, bytes memory _dynamicData) = StoreSwitch.getRecord(
      _tableId,
      _keyTuple,
      _fieldLayout
    );
    return decode(_staticData, _encodedLengths, _dynamicData);
  }

  /**
   * @notice Get the full data.
   */
  function _get(bytes32 txtDefId) internal view returns (TxtDefStoreData memory _table) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    (bytes memory _staticData, PackedCounter _encodedLengths, bytes memory _dynamicData) = StoreCore.getRecord(
      _tableId,
      _keyTuple,
      _fieldLayout
    );
    return decode(_staticData, _encodedLengths, _dynamicData);
  }

  /**
   * @notice Set the full data using individual values.
   */
  function set(bytes32 txtDefId, uint32 owner, TxtDefType txtDefType, string memory value) internal {
    bytes memory _staticData = encodeStatic(owner, txtDefType);

    PackedCounter _encodedLengths = encodeLengths(value);
    bytes memory _dynamicData = encodeDynamic(value);

    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.setRecord(_tableId, _keyTuple, _staticData, _encodedLengths, _dynamicData);
  }

  /**
   * @notice Set the full data using individual values.
   */
  function _set(bytes32 txtDefId, uint32 owner, TxtDefType txtDefType, string memory value) internal {
    bytes memory _staticData = encodeStatic(owner, txtDefType);

    PackedCounter _encodedLengths = encodeLengths(value);
    bytes memory _dynamicData = encodeDynamic(value);

    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.setRecord(_tableId, _keyTuple, _staticData, _encodedLengths, _dynamicData, _fieldLayout);
  }

  /**
   * @notice Set the full data using the data struct.
   */
  function set(bytes32 txtDefId, TxtDefStoreData memory _table) internal {
    bytes memory _staticData = encodeStatic(_table.owner, _table.txtDefType);

    PackedCounter _encodedLengths = encodeLengths(_table.value);
    bytes memory _dynamicData = encodeDynamic(_table.value);

    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.setRecord(_tableId, _keyTuple, _staticData, _encodedLengths, _dynamicData);
  }

  /**
   * @notice Set the full data using the data struct.
   */
  function _set(bytes32 txtDefId, TxtDefStoreData memory _table) internal {
    bytes memory _staticData = encodeStatic(_table.owner, _table.txtDefType);

    PackedCounter _encodedLengths = encodeLengths(_table.value);
    bytes memory _dynamicData = encodeDynamic(_table.value);

    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.setRecord(_tableId, _keyTuple, _staticData, _encodedLengths, _dynamicData, _fieldLayout);
  }

  /**
   * @notice Decode the tightly packed blob of static data using this table's field layout.
   */
  function decodeStatic(bytes memory _blob) internal pure returns (uint32 owner, TxtDefType txtDefType) {
    owner = (uint32(Bytes.slice4(_blob, 0)));

    txtDefType = TxtDefType(uint8(Bytes.slice1(_blob, 4)));
  }

  /**
   * @notice Decode the tightly packed blob of dynamic data using the encoded lengths.
   */
  function decodeDynamic(
    PackedCounter _encodedLengths,
    bytes memory _blob
  ) internal pure returns (string memory value) {
    uint256 _start;
    uint256 _end;
    unchecked {
      _end = _encodedLengths.atIndex(0);
    }
    value = (string(SliceLib.getSubslice(_blob, _start, _end).toBytes()));
  }

  /**
   * @notice Decode the tightly packed blobs using this table's field layout.
   * @param _staticData Tightly packed static fields.
   * @param _encodedLengths Encoded lengths of dynamic fields.
   * @param _dynamicData Tightly packed dynamic fields.
   */
  function decode(
    bytes memory _staticData,
    PackedCounter _encodedLengths,
    bytes memory _dynamicData
  ) internal pure returns (TxtDefStoreData memory _table) {
    (_table.owner, _table.txtDefType) = decodeStatic(_staticData);

    (_table.value) = decodeDynamic(_encodedLengths, _dynamicData);
  }

  /**
   * @notice Delete all data for given keys.
   */
  function deleteRecord(bytes32 txtDefId) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreSwitch.deleteRecord(_tableId, _keyTuple);
  }

  /**
   * @notice Delete all data for given keys.
   */
  function _deleteRecord(bytes32 txtDefId) internal {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    StoreCore.deleteRecord(_tableId, _keyTuple, _fieldLayout);
  }

  /**
   * @notice Tightly pack static (fixed length) data using this table's schema.
   * @return The static data, encoded into a sequence of bytes.
   */
  function encodeStatic(uint32 owner, TxtDefType txtDefType) internal pure returns (bytes memory) {
    return abi.encodePacked(owner, txtDefType);
  }

  /**
   * @notice Tightly pack dynamic data lengths using this table's schema.
   * @return _encodedLengths The lengths of the dynamic fields (packed into a single bytes32 value).
   */
  function encodeLengths(string memory value) internal pure returns (PackedCounter _encodedLengths) {
    // Lengths are effectively checked during copy by 2**40 bytes exceeding gas limits
    unchecked {
      _encodedLengths = PackedCounterLib.pack(bytes(value).length);
    }
  }

  /**
   * @notice Tightly pack dynamic (variable length) data using this table's schema.
   * @return The dynamic data, encoded into a sequence of bytes.
   */
  function encodeDynamic(string memory value) internal pure returns (bytes memory) {
    return abi.encodePacked(bytes((value)));
  }

  /**
   * @notice Encode all of a record's fields.
   * @return The static (fixed length) data, encoded into a sequence of bytes.
   * @return The lengths of the dynamic fields (packed into a single bytes32 value).
   * @return The dyanmic (variable length) data, encoded into a sequence of bytes.
   */
  function encode(
    uint32 owner,
    TxtDefType txtDefType,
    string memory value
  ) internal pure returns (bytes memory, PackedCounter, bytes memory) {
    bytes memory _staticData = encodeStatic(owner, txtDefType);

    PackedCounter _encodedLengths = encodeLengths(value);
    bytes memory _dynamicData = encodeDynamic(value);

    return (_staticData, _encodedLengths, _dynamicData);
  }

  /**
   * @notice Encode keys as a bytes32 array using this table's field layout.
   */
  function encodeKeyTuple(bytes32 txtDefId) internal pure returns (bytes32[] memory) {
    bytes32[] memory _keyTuple = new bytes32[](1);
    _keyTuple[0] = txtDefId;

    return _keyTuple;
  }
}