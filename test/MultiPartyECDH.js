const { expect } = require("chai");
const { ethers } = require("hardhat");
const elliptic = require("elliptic");

// Using the P256 curve
const ec = new elliptic.ec("curve25519");

describe("MultiPartyECDH", function () {
  let contract;
  let keyPairA, keyPairB, keyPairC;
  let publicKeys;

  beforeEach(async function () {
    // Generate key pairs for 3 participants
    keyPairA = ec.genKeyPair();
    keyPairB = ec.genKeyPair();
    keyPairC = ec.genKeyPair();

    // Get public keys in uncompressed format
    const pubKeyA = keyPairA.getPublic(false, "hex");
    const pubKeyB = keyPairB.getPublic(false, "hex");
    const pubKeyC = keyPairC.getPublic(false, "hex");

    publicKeys = [
      "0x" + pubKeyA,
      "0x" + pubKeyB,
      "0x" + pubKeyC,
    ];

    // Deploy contract with generated public keys
    const MultiPartyECDH = await ethers.getContractFactory("MultiPartyECDH");
    contract = await MultiPartyECDH.deploy(publicKeys);
    await contract.deployed();
  });

  it("should store the correct number of participants", async function () {
    expect(await contract.numParticipants()).to.equal(3);
  });

  it("should store and retrieve public keys correctly", async function () {
    const storedKeys = await contract.getPublicKeys();
    expect(storedKeys[0]).to.equal(publicKeys[0]);
    expect(storedKeys[1]).to.equal(publicKeys[1]);
    expect(storedKeys[2]).to.equal(publicKeys[2]);
  });

  it("should allow uploading and retrieving an intermediate value", async function () {
    const combo = "ABC";
    const sharedValue = ethers.utils.formatBytes32String("shared_value_abc");
    
    // Upload the intermediate value
    await contract.uploadIntermediateValue(combo, sharedValue);

    // Retrieve and check the value
    const retrievedValue = await contract.getIntermediateValue(combo);
    expect(retrievedValue).to.equal(sharedValue);
  });

  it("should not allow overwriting an existing intermediate value", async function () {
    const combo = "ABC";
    const sharedValue = ethers.utils.formatBytes32String("shared_value_abc");
    await contract.uploadIntermediateValue(combo, sharedValue);

    await expect(
      contract.uploadIntermediateValue(combo, sharedValue)
    ).to.be.revertedWith("Value already set for this combination");
  });

  it("should revert if trying to access a non-existent value", async function () {
    await expect(
      contract.getIntermediateValue("XYZ")
    ).to.be.revertedWith("No value set for this combination");
  });

  it("should perform full key exchange using the contract for intermediate values and verify the final shared secret", async function () {
    // Calculate intermediate shared secrets using scalar multiplications
    const sharedAB = keyPairB.getPublic().mul(keyPairA.getPrivate());
    const sharedAC = keyPairC.getPublic().mul(keyPairA.getPrivate());
    const sharedBC = keyPairC.getPublic().mul(keyPairB.getPrivate());

    // Convert points to byte arrays for storage in the smart contract
    const sharedABBytes = sharedAB.encodeCompressed();
    const sharedACBytes = sharedAC.encodeCompressed();
    const sharedBCBytes = sharedBC.encodeCompressed();
    
    // Store these intermediate values in the contract
    await contract.uploadIntermediateValue("AB", ethers.utils.hexlify(sharedABBytes));
    await contract.uploadIntermediateValue("AC", ethers.utils.hexlify(sharedACBytes));
    await contract.uploadIntermediateValue("BC", ethers.utils.hexlify(sharedBCBytes));
    
    // Retrieve intermediate values from the contract
    const sharedABFromContract = await contract.getIntermediateValue("AB");
    const sharedACFromContract = await contract.getIntermediateValue("AC");
    const sharedBCFromContract = await contract.getIntermediateValue("BC");

    // Decode the retrieved values back to elliptic points
    const decodedAB = ec.keyFromPublic(ethers.utils.arrayify(sharedABFromContract), "hex").getPublic();
    const decodedAC = ec.keyFromPublic(ethers.utils.arrayify(sharedACFromContract), "hex").getPublic();
    const decodedBC = ec.keyFromPublic(ethers.utils.arrayify(sharedBCFromContract), "hex").getPublic();

    // Calculate the final secret for each participant using the decoded intermediate values
    const finalSecretA = decodedAB.mul(keyPairC.getPrivate());
    const finalSecretB = decodedAC.mul(keyPairB.getPrivate());
    const finalSecretC = decodedBC.mul(keyPairA.getPrivate());

    // Ensure that all final secrets match
    expect(finalSecretA.getX().toString(16)).to.equal(finalSecretB.getX().toString(16));
    expect(finalSecretB.getX().toString(16)).to.equal(finalSecretC.getX().toString(16));

    console.log("Final shared secret (A):", finalSecretA.getX().toString(16));
    console.log("Final shared secret (B):", finalSecretB.getX().toString(16));
    console.log("Final shared secret (C):", finalSecretC.getX().toString(16));
  });

});
