# aws-cf-invalidate

A simple CLI to refresh files in cloudfront cache written in [Go.](http://golang.org)

## Configuration

Before using, ensure that you've configured credentials. The best way to configure
credentials is to use the ~/.aws/credentials file, which might look like:

[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY

Alternatively, you can set the following environment variables:

AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY

## Usage

```
Description: Send a batch of url paths to AWS Cloudfront to invalidate caches.

Usage: aws-cf-invalidate [options...] [assetpaths...]

Application options:
    -d, --distribution_id            ID of the distribution (mandatory)

Common options:
    -h, --help                       Show this message
    -V, --version                    Show version

Examples:
	aws-cf-invalidate -d ABC123 /path/img1.jpg
    aws-cf-invalidate --distribution_id ABC123 /path/img1.jpg /path/img2.jpg

	aws-cf-invalidate --version
```

## Building

This code currently requires version 1.42 or higher of Go.  You'll also need to
compile different platforms. For example:

GOOS=windows GOARCH=386 go build -o aws-cf-invalidate_i386.exe aws-cf-invalidate.go
GOOS=linux GOARCH=386 go build -o aws-cf-invalidate_linuxi386 aws-cf-invalidate.go

Information on Golang installation, including pre-built binaries, is available at
<http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `aws-cf-invalidate` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.
