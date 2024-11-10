const fs = require('fs');
const path = require('path');
const { ethers } = require('hardhat');

async function main() {
  const deployer = (await ethers.getSigners())[0];
  console.log(`Deployer Address: ${deployer.address}`);

  const files = fs.readdirSync(__dirname+"/..").filter(file => /eth_key\d*\.json$/.test(file));
  
  console.log(`Found key files: ${files}`);

  const amountToSend = ethers.utils.parseEther("0.1"); // Amount to send in ETH

  for (const file of files) {
    const filePath = path.join(__dirname+"/..", file);
    const fileContent = fs.readFileSync(filePath, 'utf8');
    const { PublicKey } = JSON.parse(fileContent);

    if (!ethers.utils.isAddress(PublicKey)) {
      console.error(`Invalid address in ${file}: ${PublicKey}`);
      continue;
    }

    console.log(`Sending ${ethers.utils.formatEther(amountToSend)} ETH to ${PublicKey}`);
    const tx = await deployer.sendTransaction({
      to: PublicKey,
      value: amountToSend
    });
    await tx.wait();
    console.log(`Transaction confirmed: ${tx.hash}`);
  }

  console.log('All transactions completed!');
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
