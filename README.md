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
### To Do
- [ ] Fetch the upstream branch if it exists to see if it is gone. This is
similar to what you see in Lazygit but just a little more compact.
- [ ] Add a feature that allows you to type in the parent/sub-task number and
the name of the branch and have it auto create
`JOB-100827/JOB-18887/some-sub-task-name` for you. 
- [ ] Add last commit in days to get a sense of how long it has been since the
branch has been in use.
- [ ] If the branch is a feature branch, write it as `<sub-task-number> ->
JOB-xxxxxx/JOB-zzzzz/branch-name`
