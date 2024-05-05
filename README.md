# TabbyKat

### Installation Instructions
Clone the repository 
```bash
git clone git@github.com:DennisTheMenace780/tabbyKat.git /usr/local/bin/
```
### How to run the tests

Because Tabbykat operates within a git repository, one needs to be created in
order to ensure the program behaves as expected. 

1. Run `./bin/init_test_repo.sh` to create the `TestRepo` sub-module.
2. Run `go test -v ./... -update` to create `.golden` files used by [Bubble Tea
   for testing](https://charm.sh/blog/teatest/)
3. Run `go test -v ./..` to execute the tests recursively

### Uninstall Instructions
```bash
cd /usr/local/bin/
sudo rm -rf TabbyKat
```
## Building Binary
```go
GOOD=darwin GOARCH=amd644 go build -o mac-tabbykat
GOOD=linux GOARCH=amd644 go build -o mac-linux
```
