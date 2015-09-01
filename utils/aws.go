package utils

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetInstances(c []string) ([]*ec2.Instance, error) {
	tags := parseTags(c)
	filters := make([]*ec2.Filter, len(c))

	for key, value := range tags {
		filters = append(filters, &ec2.Filter{
			Name: aws.String(fmt.Sprintf("tag:%v", key)),
			Values: []*string{
				aws.String(value),
			},
		})
	}

	instances, err := getInstances(filters)

	if err != nil {
		return nil, err
	}
	return instances, nil
}

func getInstances(filters []*ec2.Filter) ([]*ec2.Instance, error) {
	svc := ec2.New(nil)
	instances := make([]*ec2.Instance, 0)
	result, token, err := paginateResult(svc, filters, nil)

	if err != nil {
		return nil, err
	}

	instances = append(instances, result...)

	for token != nil {
		instances, token, err = paginateResult(svc, filters, token)
		if err != nil {
			return nil, err
		}

		instances = append(instances, result...)
	}
	return instances, nil
}

func paginateResult(conn *ec2.EC2, filters []*ec2.Filter, token *string) ([]*ec2.Instance, *string, error) {
	instances := make([]*ec2.Instance, 0)
	params := &ec2.DescribeInstancesInput{
		Filters:   filters,
		NextToken: token,
	}

	result, err := conn.DescribeInstances(params)
	if err != nil {
		return nil, nil, err
	}

	for idx, _ := range result.Reservations {
		for _, instance := range result.Reservations[idx].Instances {
			instances = append(instances, instance)
		}
	}
	return instances, result.NextToken, nil
}

func parseTags(c []string) map[string]string {
	tags := make(map[string]string)

	for _, tag := range c {
		s := strings.Split(tag, ":")
		if len(s) == 2 {
			tags[s[0]] = s[1]
		}
	}
	return tags
}
