# Code-Preview-Matchmaker
Just another matchmaker to match slack channel's member to peer-peer groups for code preview.

# Usage
## Method 1 : Clone source
1. Clone this repo
```
git clone https://github.com/nguyenvanduocit/Code-Preview-Matchmaker.git
```
1. Rename .env.example to .env and config it
2. Build :
```
cd Code-Preview-Matchmaker
go build matchmaker.go`
```
3. Run :
```
./matchmaker
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

## Method 3
1. Download file
    1. Windows : matchmaker.exe
    2. OSX : matchmaker_darwin
    3. Linux : matchmaker_linux
2. Run
```
./matchmaker -token=yourtoken -name=yourbotname -channel=targetchannel -debug=true
```
