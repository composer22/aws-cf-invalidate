package main

import (
	"fmt"
	"os"
)

const usageStr = `
Description: Send a batch of url paths to AWS Cloudfront to invalidate caches.

Requirements:

Before using, ensure that you've configured credentials. The best way to configure
credentials is to use the ~/.aws/credentials file, which might look like:

[default]
aws_access_key_id = AKID1234567890
aws_secret_access_key = MY-SECRET-KEY

Alternatively, you can set the following environment variables:

AWS_ACCESS_KEY_ID=AKID1234567890
AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY

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
`

// end help text

// PrintUsageAndExit is used to print out command line options.
func PrintUsageAndExit() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}
