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

import "./verifier.sol";


contract Model is Verifier{
    struct Hash {
        uint a;
        uint b;
        uint c;
        uint d;
        uint e;
        uint f;
        uint g;
        uint h;
    }

    struct Signiture {
        uint256[2] R;
        uint256 S;
    }

    // Ezeke a megfelelő módon inicializálandóak (deploy előtt?)
    Hash current_hash;
    Signiture sig;
    string current_ciphertext = "";

    constructor(Hash memory start_hash,string memory start_ciphertext) {
        current_hash = start_hash;
        current_ciphertext = start_ciphertext;
    }

    function stepModel(Hash memory hash, string memory ciphertext,Signiture memory sig_new, Proof memory p) public {
        uint[19] memory inputs = [current_hash.a,current_hash.b,current_hash.c,current_hash.d,current_hash.e,current_hash.f,current_hash.g,current_hash.h,sig_new.R[0],sig_new.R[1],sig_new.S,hash.a,hash.b,hash.c,hash.d,hash.e,hash.f,hash.g,hash.h];
        bool verified = verifyTx(p,inputs);
        assert(verified);
        current_ciphertext = ciphertext;
        sig = sig_new;
        current_hash = hash;
    }

    function  getCurrentHash() public view returns (Hash memory) {
        return current_hash;
    }

    function getLastSignature() public view returns (Signiture memory)  {
        return sig;
    }
    function getCiphertext() public view returns (string memory) {
        return current_ciphertext;
    }
}
