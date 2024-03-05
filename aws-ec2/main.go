package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {

	ctx := context.Background()
	instanceId, err := createEc2(ctx, "eu-central-1")
	if err != nil {
		fmt.Printf("error while creating ec2 instance: %s", err)
		os.Exit(1)
	}

	fmt.Printf("created instance id: %s", instanceId)
}

func createEc2(ctx context.Context, region string) (string, error) {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}

	ec2Client := ec2.NewFromConfig(cfg)
	_, err = ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
		KeyName: aws.String("go-aws-demo"),
	})
	if err != nil {
		return "", fmt.Errorf("error while creating key pair: %s", err)
	}

	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("error while describing image: %s", err)
	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("no images received: %s", err)
	}

	instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		KeyName:      aws.String("go-aws-demo"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		return "", fmt.Errorf("error while running instances: %s", err)
	}
	if len(instance.Instances) == 0 {
		return "", fmt.Errorf("no instances received: %s", err)
	}

	return *instance.Instances[0].InstanceId, nil
}
