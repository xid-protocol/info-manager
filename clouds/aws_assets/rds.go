package aws_assets

// import (
// 	"log"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/service/ec2"
// 	"github.com/aws/aws-sdk-go/service/rds"
// )

// type IpPermission struct {
// 	PortRange string
// 	Protocol  string
// 	CidrIps   []string
// }

// type Sg struct {
// 	GroupID   string
// 	GroupName string

// 	IpPermissions []IpPermission
// }

// type RdsAssets struct {
// 	RdsId      string
// 	SecGoupIds [][]Sg
// }

// type RdsAsset struct {
// 	DBName string
// 	Sgs    []SecGroup
// }

// func GetRds() {
// 	sess := NewSession()

// 	//创建EC2实例
// 	ec2Cli := ec2.New(sess)

// 	// 创建 RDS 服务客户端
// 	rdsCli := rds.New(sess)

// 	// 创建描述 DB 实例的请求
// 	input := &rds.DescribeDBInstancesInput{}
// 	result, err := rdsCli.DescribeDBInstances(input)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	rdsAssets := dords(result, ec2Cli)
// 	var data [][]string
// 	for _, rdsAsset := range rdsAssets {
// 		for _, secgroup := range rdsAsset.Sgs {
// 			for _, rule := range secgroup.Rules {
// 				//
// 				for _, iprange := range rule.Ipranges {
// 					log.Println(rdsAsset.DBName, rule.Protocol, rule.PortRange, iprange.CidrIp, iprange.Description)
// 					var ips string
// 					for _, ip := range iprange.CidrIp {
// 						ips = ips + " " + ip
// 					}
// 					data = append(data, []string{rdsAsset.DBName, rule.Protocol, rule.PortRange, ips, iprange.Description})
// 				}

// 			}

// 		}
// 	}
// 	WriteToCSV(data, "./rds.csv")
// }

// func dords(result *rds.DescribeDBInstancesOutput, ec2Cli *ec2.EC2) []RdsAsset {
// 	var rdsAssets []RdsAsset
// 	for _, db := range result.DBInstances {
// 		var rdsAsset RdsAsset
// 		log.Println(aws.StringValue(db.DBInstanceIdentifier))
// 		//var secgroupids []string
// 		rdsAsset.DBName = aws.StringValue(db.DBInstanceIdentifier)
// 		// log.Println(db)
// 		// if aws.StringValue(db.DBInstanceIdentifier) != "ltp-security" {
// 		// 	continue
// 		// }

// 		// var sgs [][]Sg
// 		var secgroups []SecGroup
// 		for _, group := range db.VpcSecurityGroups {
// 			//获取安全组ID
// 			secgroupID := group.VpcSecurityGroupId
// 			//通过安全组ID获取安全组规则
// 			secgroupRule := GetSgByID(secgroupID, ec2Cli)
// 			//log.Println(secgroupRule)
// 			secgroups = append(secgroups, secgroupRule)
// 			// sgs = append(sgs, secgroups)
// 		}
// 		rdsAsset.Sgs = secgroups

// 		rdsAssets = append(rdsAssets, rdsAsset)
// 	}

// 	return rdsAssets
// }
