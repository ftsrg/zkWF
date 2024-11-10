// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2023 Contributors of the zkWF project
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

pragma solidity ^0.8.0;

import "../contract.sol";

contract Model is PlonkVerifier {

    uint256 encrypted_state_len = 4;
    uint256 hash;
    uint256[] state;

    constructor(uint256 initial_hash, uint256[] memory initial_state) {
        encrypted_state_len = initial_state.length;
        hash = initial_hash;
        state = initial_state;
    }

    function getCurrentHash() public view returns (uint256) {
        return hash;
    }

    function getCurrentState() public view returns (uint256[] memory) {
        return state;
    }

    function update(bytes calldata proof, uint256[] calldata public_inputs) public payable {
       require(public_inputs[0] == hash);
       require(public_inputs[public_inputs.length-2] == msg.value);
       require(7 + encrypted_state_len+2 <= public_inputs.length );

       bool verified = Verify(proof, public_inputs);
       
       assert(verified);
       hash = public_inputs[1];
       // Get the encrypted state from the public inputs [7:7 + encrypted_state_len]
       for (uint256 i = 7; i < 7 + encrypted_state_len; i++) {
          state[i - 7] = public_inputs[i];
       }

       // Handle withdraw
       uint256 amount = public_inputs[public_inputs.length-1];
       if (amount > 0) {
          payable(msg.sender).transfer(amount);
       }
    }
}

