package aws_assets

import (
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// 通过security group ID 获取安全组规则
func GetSgByID(securityGroupID *string, ec2Cli *ec2.EC2) *ec2.DescribeSecurityGroupsOutput {

	// //实例化安全组
	// var secGroup SecGroup

	securityGroupInput := &ec2.DescribeSecurityGroupsInput{
		GroupIds: []*string{securityGroupID},
	}
	securityGroupOutput, err := ec2Cli.DescribeSecurityGroups(securityGroupInput)
	// log.Println(securityGroupOutput)
	if err != nil {
		log.Println("Failed to describe security groups:", err)
	}

	return securityGroupOutput
}

func SecGroupMonitor() *map[string][]*ec2.DescribeSecurityGroupsOutput {
	// 获取所有区域的客户端
	clients := GetAllRegionEc2Clients()
	// 遍历每个区域获取实例信息
	securityGroups := make(map[string][]*ec2.DescribeSecurityGroupsOutput)
	for region, client := range clients {
		result, err := client.DescribeInstances(nil)
		if err != nil {
			log.Printf("no instances in this region %s: %v", region, err)
			continue
		}

		// 处理该区域的实例
		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				instanceSecGroups := instance.SecurityGroups
				for _, securityGroup := range instanceSecGroups {
					//获取groupid
					groupID := *securityGroup.GroupId
					secGroup := GetSgByID(&groupID, client)
					securityGroups[groupID] = append(securityGroups[groupID], secGroup)
					// log.Println(secGroup)
				}
				// instances[*instance.InstanceId] = instance
			}
		}
	}
	return &securityGroups
}

// type Iprange struct {
// 	CidrIp      []string
// 	Description string
// }

// type SecGroupRule struct {
// 	RuleId    string
// 	IpVersion string
// 	Protocol  string
// 	PortRange string
// 	Ipranges  []Iprange
// }

// type SecGroup struct {
// 	Id    string
// 	Name  string
// 	Rules []SecGroupRule
// }

// // 通过security group ID 获取安全组规则
// func GetSgByID(securityGroupID *string, ec2Cli *ec2.EC2) SecGroup {

// 	//实例化安全组
// 	var secGroup SecGroup

// 	securityGroupInput := &ec2.DescribeSecurityGroupsInput{
// 		GroupIds: []*string{securityGroupID},
// 	}
// 	securityGroupOutput, err := ec2Cli.DescribeSecurityGroups(securityGroupInput)
// 	// log.Println(securityGroupOutput)
// 	if err != nil {
// 		log.Println("Failed to describe security groups:", err)
// 	}

// 	for _, securityGroup := range securityGroupOutput.SecurityGroups {
// 		//fmt.Println("Security Group:", *securityGroup.GroupName, *securityGroup.GroupId)

// 		secGroup.Id = *securityGroup.GroupId
// 		secGroup.Name = *securityGroup.GroupName
// 		//获取安全组规则
// 		secGroup.Rules = getSecGroupRule(securityGroup.IpPermissions)

// 	}
// 	return secGroup
// }

// func getSecGroupRule(permissions []*ec2.IpPermission) []SecGroupRule {
// 	var rules []SecGroupRule
// 	for _, permission := range permissions {
// 		var rule SecGroupRule
// 		rule.Protocol = getProtocal(permission)
// 		rule.PortRange = getPortRange(permission)
// 		rangep := getIprange(permission)
// 		rule.Ipranges = append(rule.Ipranges, rangep...)
// 		rules = append(rules, rule)
// 	}
// 	// log.Println(permissions)
// 	// log.Println(rules)
// 	return rules
// }

// func getIprange(permission *ec2.IpPermission) []Iprange {
// 	// log.Println(permission)
// 	var ipranges []Iprange
// 	for _, ipr := range permission.IpRanges {
// 		// log.Println(ipr)
// 		var iprange Iprange
// 		if ipr.CidrIp != nil {
// 			iprange.CidrIp = append(iprange.CidrIp, *ipr.CidrIp)
// 		}
// 		if ipr.Description != nil {
// 			iprange.Description = *ipr.Description
// 		}

// 		ipranges = append(ipranges, iprange)
// 	}

// 	//log.Println(permission)
// 	//如果UserIdGroupPairs不为空，也就是一个安全组引用了其他安全组的情况
// 	if permission.UserIdGroupPairs != nil {
// 		// log.Println(permission)
// 		for _, pair := range permission.UserIdGroupPairs {
// 			var iprange Iprange
// 			if pair.Description != nil {
// 				iprange.Description = *pair.Description
// 			}
// 			iprange.CidrIp = findInstancesUsingSecurityGroup(*pair.GroupId)
// 			ipranges = append(ipranges, iprange)
// 		}
// 	}

// 	return ipranges
// }

// func findInstancesUsingSecurityGroup(securityGroupID string) []string {
// 	sess := NewSession()
// 	//创建EC2实例
// 	ec2Cli := ec2.New(sess)
// 	input := &ec2.DescribeInstancesInput{
// 		Filters: []*ec2.Filter{
// 			{
// 				Name:   aws.String("instance.group-id"),
// 				Values: []*string{aws.String(securityGroupID)},
// 			},
// 		},
// 	}

// 	var instances []string
// 	err := ec2Cli.DescribeInstancesPages(input,
// 		func(page *ec2.DescribeInstancesOutput, lastPage bool) bool {
// 			for _, reservation := range page.Reservations {
// 				//instances = append(instances, reservation.)
// 				for _, instance := range reservation.Instances {
// 					instanceid := *instance.InstanceId
// 					instances = append(instances, instanceid)
// 				}
// 			}
// 			return !lastPage
// 		})

// 	if err != nil {
// 		log.Println("failed to describe instances: %w", err)
// 	}

// 	return instances
// }

// func getGroupPairs(permission *ec2.IpPermission) []string {
// 	var pairranges []string
// 	for _, pair := range permission.UserIdGroupPairs {

// 		groupID := *pair.GroupId
// 		instances := findInstancesUsingSecurityGroup(groupID)
// 		log.Println(len(instances))

// 		//外部引入的安全组，这个安全组没有被任何实例使用
// 		if len(instances) == 0 {
// 			continue
// 		} else {
// 			for _, instance := range instances {
// 				// var iprange Iprange
// 				// iprange.InstanceId = instance
// 				// iprange.GroupId = groupID
// 				// if pair.Description != nil {
// 				// 	iprange.Description = *pair.Description
// 				// }
// 				pairranges = append(pairranges, instance)
// 			}
// 		}

// 		log.Println(pairranges)

// 	}

// 	return pairranges
// }

// func getProtocal(permission *ec2.IpPermission) string {
// 	var protocol string
// 	protocol = fmt.Sprintf(*permission.IpProtocol)
// 	if protocol == "-1" {
// 		protocol = "all"
// 	}
// 	return protocol
// }

// func getPortRange(permission *ec2.IpPermission) string {
// 	var portrange string

// 	var fromPort string
// 	if permission.FromPort == nil {
// 		fromPort = "-1"
// 	} else {
// 		fromPort = fmt.Sprint(*permission.FromPort)
// 	}

// 	if fromPort == "-1" {
// 		fromPort = "all"
// 	}
// 	var toPort string
// 	if permission.ToPort == nil {
// 		toPort = "-1"
// 	} else {
// 		toPort = fmt.Sprint(*permission.ToPort)
// 	}
// 	if toPort == "-1" {
// 		toPort = "all"
// 	}

// 	if fromPort == toPort {
// 		portrange = fromPort
// 	} else {
// 		portrange = fmt.Sprintf("[%s-%s]", fromPort, toPort)
// 	}
// 	return portrange
// }
