from zokrates_pycrypto.eddsa import PrivateKey, PublicKey
from zokrates_pycrypto.field import FQ
from zokrates_pycrypto.utils import write_signature_for_zokrates_cli
import sys


if __name__ == "__main__":
    testKeys = [1997011358982923168928344992199991480689546837621580239342656433234255379027 , 1997011358982923168928344992199991480689546837621580239342656433234255379026, 1997011358982923168928344992199991480689546837621580239342656433234255379025, 1997011358982923168928344992199991480689546837621580239342656433234255379024]

    # sk = PrivateKey.from_rand()
    # Seeded for debug purpose
    key = FQ(testKeys[int(sys.argv[1])])
    sk = PrivateKey(key)

    pk = PublicKey.from_private(sk)
    print(str(pk.p.x.n) +" "+  str(pk.p.y.n) +" "+ str(sk.fe))