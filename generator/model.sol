// SPDX-License-Identifier: GPL-3.0-or-later

pragma solidity ^0.8.0;

import "verifier.sol";


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

    struct Signature {
        uint256[2] R;
        uint256 S;
    }

    // Ezeke a megfelelő módon inicializálandóak (deploy előtt?)
    Hash current_hash;
    Signature sig;
    string current_ciphertext = "";

    constructor(Hash memory start_hash,string memory start_ciphertext) {
        current_hash = start_hash;
        current_ciphertext = start_ciphertext;
    }

    function stepModel(Hash memory hash, string memory ciphertext, Signature memory sig_new, Proof memory p) public {
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

    function getLastSignature() public view returns (Signature memory)  {
        return sig;
    }
    function getCiphertext() public view returns (string memory) {
        return current_ciphertext;
    }
}