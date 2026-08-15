# zkWF

Zero-knowledge workflow engine: Go CLI compiling BPMN models to gnark zkSNARK circuits, Solidity/Hardhat contracts, and a bpmn-js editor. An agent here is usually touching the Go CLI/circuits (`cmds/zkwf`, `pkg/`), the contracts, or the editor.

## Always on

1. **Never push to master.** This is the upstream ftsrg org repo (`git@github.com:ftsrg/zkWF.git`) and SSH push is configured, so a bad push lands on the org's default branch. The local checkout sits on master — branch before any commit; land everything via PR.
2. **Never run `generate-key` unprompted** — it writes an EdDSA private key to `keys.json` in cwd, silently overwriting any existing one.
3. **Never run `deploy-ecdh`, `scripts/deploy.js`, or `scripts/sendFunds.js` unprompted** — chain- and money-touching. Propose to Martin first, every time.
4. **Never run `setup` unprompted** — trusted-setup phase, CPU/RAM-heavy enough to stall the machine.
5. **CLI artifacts clobber cwd files.** `compile`/`setup`/`fill-inputs`/`prove` drop `circuit.r1cs`, `pk.bin`, `full.wtns`, `keys.json` into the working directory over whatever is there. Run CLI experiments from a scratch directory, not the repo root.
6. **`hardhat.config.js` has no `networks` block — that is deliberate (local-only).** Do not add one; adding networks is what makes rule 3 dangerous.

## Commands

- Build the CLI: `go build -o bin/zkWF ./cmds/zkwf`
  README line 69 says `./cmds/zkWF` — wrong case; the actual directory is `cmds/zkwf` and the README form fails on Linux. Flagged here on purpose: fix the README only via PR.
- Go tests: `go test ./pkg/...` (the only Go tests live in `pkg/circuits/gmimc` and `pkg/crypto/{gmimc,hkdf}`)
- Contract tests: `npx hardhat test`
- Editor: `cd editor && npm install && npm start` — long-running webpack-dev-server; global PID-save-and-kill rule applies.

## Where things are

- `cmds/zkwf/` — CLI entry (one file per subcommand); `pkg/` — circuits, crypto, BPMN parsing, web3.
- `contracts/` — Solidity sources. README's "Project Structure" calls this `solidity/`; that section is stale, trust the tree.
- `models/` — BPMN models + JSON test cases the CLI examples use; `editor/` — bpmn-js app.
- `engine/` — a single undocumented script (`generatePaths.js`); ask Martin before building on it.
