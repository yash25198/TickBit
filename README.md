# tickbit-relayer

Golang implementation of the TickBit Relayer

## Setup

1. Install Go 1.20 or later
2. Clone the repository
3. Run `go run ./cmd/relay/relay.go`
4. Check env-example for environment variables
5. Run `go run ./cmd/relay/relay.go` to start the relayer


## How it works

1. The relayer listens for new bitcoin blocks being mined.
2. When a new block is mined, the relayer waits until cryptoid.info has the block indexed.
3. The relayer then fetches the block from cryptoid.info and verifies the SXG signature.
4. Then spins a child process (a rust compiled binary) to start the proof generation process.
5. The resultant proof is then submitted to the TickBit contract on chain thus settling the all the bets for that block.

