# Code-Preview-Matchmaker
Just another matchmaker to match slack channel's member to peer-peer groups for code preview.

# Usage
## Method 1
1. Download file
    1. Windows : [matchmaker.exe](https://github.com/nguyenvanduocit/Code-Preview-Matchmaker/raw/master/build/matchmaker.exe)
    2. OSX : [matchmaker_darwin](https://github.com/nguyenvanduocit/Code-Preview-Matchmaker/raw/master/build/matchmaker_darwin)
    3. Linux : [matchmaker_linux](https://github.com/nguyenvanduocit/Code-Preview-Matchmaker/raw/master/build/matchmaker_linux)
2. Config .env
2. Run with .env or aguments
```
./matchmaker -token=yourtoken -name=yourbotname -channel=targetchannel -debug=true
```

## Method 2
1. Get repo
```
go get github.com/nguyenvanduocit/Code-Preview-Matchmaker
```
2. Run
```
Code-Preview-Matchmaker -token=yourtoken -name=yourbotname -channel=targetchannel -debug=true
```
Make sure $GOPATH/bin added to the PATH

## Method 3 : Build from source
1. Clone this repo
```
git clone https://github.com/nguyenvanduocit/Code-Preview-Matchmaker.git
```
2. Rename .env.example to .env and config it
3. Build :
```
cd Code-Preview-Matchmaker
go build matchmaker.go`
```
4. Run :
```
./matchmaker
```
