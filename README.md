# starfy
> A tool for growing stars rapidly

## Requirements
```
go 1.6.2
```

## Installation
```
glide up -u -s
```

## Usage
```
#Set up tokens adding to secrets.yml
cd ./config && cp ./secrets.yml.copy ./secrets.yml

#Build one binary
go build

#Run
./starfy [OWNER] [REPO]
```
