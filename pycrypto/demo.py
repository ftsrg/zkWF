import hashlib
#import js2py
#import subprocess
#import json
import sys

from zokrates_pycrypto.eddsa import PrivateKey, PublicKey
from zokrates_pycrypto.field import FQ
from zokrates_pycrypto.utils import write_signature_for_zokrates_cli

if __name__ == "__main__":

    testKeys = [1997011358982923168928344992199991480689546837621580239342656433234255379027 , 1997011358982923168928344992199991480689546837621580239342656433234255379026, 1997011358982923168928344992199991480689546837621580239342656433234255379025, 1997011358982923168928344992199991480689546837621580239342656433234255379024]

    h_curr = sys.argv[1]
    h_next = sys.argv[2]

    msg = h_curr+h_next
    
    msg = bytes.fromhex(msg)
    # sk = PrivateKey.from_rand()
    # Seeded for debug purpose
    key = FQ(testKeys[int(sys.argv[3])])
    sk = PrivateKey(key)
    sig = sk.sign(msg)

    pk = PublicKey.from_private(sk)
    is_verified = pk.verify(sig, msg)
    assert(is_verified)
    path = 'zokrates_inputs.txt'
    sig_R, sig_S = sig
    args =[sig_R.x, sig_R.y, sig_S, pk.p.x.n, pk.p.y.n,sk.fe]
    
    args = " ".join(map(str, args))


    with open(path, "w+") as file:
        for l in args:
            file.write(l)
            print(l, end='')
