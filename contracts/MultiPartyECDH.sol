// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MultiPartyECDH {
    uint8 public numParticipants;
    mapping(uint8 => bytes) public publicKeys; // Maps participant index to their public key
    mapping(string => bytes) private intermediateValues; // Maps participant combo to shared values
    mapping(string => bool) private isValueSet; // Track which values have been uploaded

    constructor(bytes[] memory _publicKeys) {
        numParticipants = uint8(_publicKeys.length);
        require(numParticipants >= 2 && numParticipants <= 128, "Participants must be between 2 and 128");
        
        // Store public keys
        for (uint8 i = 0; i < numParticipants; i++) {
            publicKeys[i] = _publicKeys[i];
        }
    }

    // Upload intermediate shared value
    function uploadIntermediateValue(string memory combo, bytes memory sharedValue) public {
        require(!isValueSet[combo], "Value already set for this combination");
        intermediateValues[combo] = sharedValue;
        isValueSet[combo] = true;
    }

    // Retrieve intermediate value
    function getIntermediateValue(string memory combo) public view returns (bytes memory) {
        require(isValueSet[combo], "No value set for this combination");
        return intermediateValues[combo];
    }

    // Retrieve all public keys
    function getPublicKeys() public view returns (bytes[] memory) {
        bytes[] memory keys = new bytes[](numParticipants);
        for (uint8 i = 0; i < numParticipants; i++) {
            keys[i] = publicKeys[i];
        }
        return keys;
    }
}
