# TabbyKat

A Git tool developed in Golang that allows you to tab through your open branches and check them out seemlessly.

### Installation Instructions
```bash
git clone git@github.com:DennisTheMenace780/tabbyKat.git /tmp/tabbyKat
sudo mv /tmp/tabbyKat/tabbyKat /usr/local/bin/
rm -rf /tmp/tabbyKat
```
### Uninstall Instructions
```bash
cd /usr/local/bin/
# Darwin (MacOS)
sudo rm /usr/local/bin/mac-tabbyKat
# Linux
sudo rm /usr/local/bin/linux-tabbyKat
```
### How to run the tests

Because tabbyKat is a thin wrapper around git, a git repository needs to be created in
order to ensure the program behaves as expected. 

1. Run `./bin/init_test_repo.sh` to create the `TestRepo` sub-module.
2. Run `go test -v ./... -update` to create `.golden` files used by [Bubble Tea
   for testing](https://charm.sh/blog/teatest/)
3. Run `go test -v ./...` to execute the tests recursively

### Building Binary
```go
GOOD=darwin GOARCH=amd64 go build -o mac-tabbykat
GOOD=linux GOARCH=amd64 go build -o linux-tabbykat
```
