package utils

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func GetInstances() ([]*ec2.Instance, error) {
	svc := ec2.New(nil)
	instances := make([]*ec2.Instance, 0)

	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		return nil, err
	}

	for idx, _ := range resp.Reservations {
		for _, instance := range resp.Reservations[idx].Instances {
			instances = append(instances, instance)
		}
	}
	return instances, nil
}

func Filter(instances []*ec2.Instance, f func(*ec2.Instance) bool) []*ec2.Instance {
	filtered := make([]*ec2.Instance, 0)

	for _, instance := range instances {
		if f(instance) {
			filtered = append(filtered, instance)
		}
	}
	return filtered
}

func exist(tag *ec2.Tag, tags []*ec2.Tag) bool {
	for _, t := range tags {
		if *t.Key == *tag.Key && *t.Value == *tag.Value {
			return true
		}
	}
	return false
}

func CreateFilter(c []string) func(inst *ec2.Instance) bool {
	tags := make(map[string]string)

	for _, tag := range c {
		s := strings.Split(tag, ":")
		if len(s) == 2 {
			tags[s[0]] = s[1]
		}
	}
	return createFilterMap(tags)
}

func createFilterMap(m map[string]string) func(inst *ec2.Instance) bool {
	return func(inst *ec2.Instance) bool {
		for key, value := range m {
			if !exist(&ec2.Tag{Key: &key, Value: &value}, inst.Tags) {
				return false
			}
		}
		return true
	}
}
