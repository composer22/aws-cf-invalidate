# aws-cf-invalidate

A simple CLI to refresh files in cloudfront cache written in [Go.](http://golang.org)

## Configuration

Before using, ensure that you've configured credentials. The best way to configure
credentials is to use the ~/.aws/credentials file, which might look like:
```
[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY
```
Alternatively, you can set the following environment variables:
```
AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY
```
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
```
GOOS=windows GOARCH=amd64 go build -o aws-cf-invalidate-amd64.exe -i
GOOS=linux GOARCH=amd64 go build -o aws-cf-invalidate-linux-amd64 -i
```
To compile all versions, run the build.sh script:
```
. build.sh
```
Information on Golang installation, including pre-built binaries, is available at
<http://golang.org/doc/install>.

Run `go version` to see the version of Go which you have installed.

Run `go build` inside the directory to build.

Run `go test ./...` to run the unit regression tests.

A successful build run produces no messages and creates an executable called `aws-cf-invalidate` in this
directory.

Run `go help` for more guidance, and visit <http://golang.org/> for tutorials, presentations, references and more.

## License

(The MIT License)

Copyright (c) 2015 Pyxxel Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
