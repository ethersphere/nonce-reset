# nonce-cleaner

Reset the nonce counter of a bee node. 
Only use this if there is an actual nonce mismatch and there are no pending transactions floating around in the eth network.

Ensure your bee node is shutdown before running and backup your `statestore` directory inside the bee data directory.

```sh
go run ./pkg/main.go path_to_bee_datadir
```

or if using binaries
```sh
./binary path_to_bee_datadir
```
