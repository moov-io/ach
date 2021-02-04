---
layout: page
title: AWS

hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Build a Go file for use in AWS Lambda

Lambdas require a Linux executable packaged in a zip folder. The following commands can be used to build a compatible executable (see [Lambda Packaging Docs](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html) for more info):

### Mac/Linux:
```
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```

### Windows:
You must create a binary that is executable on Linux. Amazon makes a handy tool for this, you can get it with:
```
set GO111MODULE=on
go.exe get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip
```

then build your file: <br/>

PowerShell:
```
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main
```

Windows (cmd):
```
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0
go build -o main main.go
%USERPROFILE%\Go\bin\build-lambda-zip.exe -o main.zip main
```

## Next Steps

- [Convert JSON to NACHA format](./json_to_nacha/lambda_json_to_nacha.md)
- [Convert NACHA To JSON format](./nacha_to_json/lambda_nacha_to_json.md)
