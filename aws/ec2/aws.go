package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Ec2Client interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput)
}

type Client struct {
	Config *aws.Config
}

func NewConfig(awsProfile, region string) *Client {
	cfg, _ := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(awsProfile),
		config.WithRegion(region),
	)
	return &Client{
		Config: &cfg,
	}
}

func (c *Client) DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput) {
	ec2Client := ec2.NewFromConfig(*c.Config)
	instanceOut, _ := ec2Client.DescribeInstances(context.TODO(), params)
	for _, reservation := range instanceOut.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Println("Instance ID:", aws.ToString(instance.InstanceId))
			fmt.Println("Instance State:", instance.State.Name)
			fmt.Println("Instance Type:", instance.InstanceType)
			fmt.Println("Private IP Address:", aws.ToString(instance.PrivateIpAddress))
		}
	}
}
