const hre = require("hardhat");
const elliptic = require("elliptic");

// Using the P256 curve
const ec = new elliptic.ec("p256");

async function main() {
  // Generate key pairs for 3 participants
  const keyPairA = ec.genKeyPair();
  const keyPairB = ec.genKeyPair();
  const keyPairC = ec.genKeyPair();

  // Get the public keys in uncompressed format (64 bytes)
  const pubKeyA = keyPairA.getPublic(false, "hex");
  const pubKeyB = keyPairB.getPublic(false, "hex");
  const pubKeyC = keyPairC.getPublic(false, "hex");

  // Log private keys (for demonstration purposes)
  console.log("Private Key A:", keyPairA.getPrivate("hex"));
  console.log("Private Key B:", keyPairB.getPrivate("hex"));
  console.log("Private Key C:", keyPairC.getPrivate("hex"));

  // Deploy contract with public keys
  const publicKeys = [
    "0x" + pubKeyA,
    "0x" + pubKeyB,
    "0x" + pubKeyC,
  ];

  const MultiPartyECDH = await hre.ethers.getContractFactory("MultiPartyECDH");
  const ecdh = await MultiPartyECDH.deploy(publicKeys);
  await ecdh.deployed();

  console.log("MultiPartyECDH deployed to:", ecdh.address);
}

// Execute the deployment script
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
