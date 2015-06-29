// aws-cf-invalidate is a simple CLI for refreshing files in cloudfront cache
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/cloudfront"
)

// PrintVersionAndExit prints the version of the application then exits.
func PrintVersionAndExit() {
	fmt.Printf("aws-cf-invalidate version %s\n", version)
	os.Exit(0)
}

// main entry point to the application.
func main() {
	var distributionID string
	var showVersion bool

	// Get the command line flags and options.
	flag.StringVar(&distributionID, "d", "", "Distribution ID")
	flag.StringVar(&distributionID, "--distribution_id", "", "Distribution ID")
	flag.BoolVar(&showVersion, "V", false, "Show version")
	flag.BoolVar(&showVersion, "--version", false, "Show version")

	flag.Usage = PrintUsageAndExit
	flag.Parse()
	args := flag.Args()
	if showVersion {
		PrintVersionAndExit()
	}

	// Check params are valid.
	if distributionID == "" {
		fmt.Println("Error: Distribution ID is mandatory.")
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println("Error: No asset urls submitted to application.")
		os.Exit(1)
	}

	// Build the batch list of asset URLs to reload.
	assetList := make([]*string, 0)
	for _, arg := range args {
		assetList = append(assetList, aws.String(arg))
	}

	// Connect to AWS CF.
	svc := cloudfront.New(&aws.Config{Region: "us-west-1"})

	// Build the params.
	params := &cloudfront.CreateInvalidationInput{
		DistributionID: aws.String(distributionID),
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(createV4UUID()),
			Paths: &cloudfront.Paths{
				Quantity: aws.Long(int64(len(assetList))),
				Items:    assetList,
			},
		},
	}

	// Send and check for errors.
	resp, err := svc.CreateInvalidation(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	// Print response and return.
	fmt.Println(awsutil.StringValue(resp))
}
