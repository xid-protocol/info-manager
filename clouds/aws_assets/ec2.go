package aws_assets

import (
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

// type Ec2Asset struct {
// 	Ec2Client    *ec2.EC2
// 	InstanceName string
// 	Instance     *ec2.Instance
// 	InstanceID   string
// 	Tags         []string
// 	Region       string
// 	PublicIPv4   []string
// 	PrivateIPv4  []string
// }

// GetEc2FromAllRegions 从所有区域获取 EC2 实例信息
func GetEc2FromAllRegions() *map[string]*ec2.Instance {
	// 获取所有区域的客户端
	clients := GetAllRegionEc2Clients()

	// 创建一个空的 Assets 切片
	// var allAssets []Ec2Asset

	// 遍历每个区域获取实例信息
	instances := make(map[string]*ec2.Instance)
	for region, client := range clients {
		result, err := client.DescribeInstances(nil)
		if err != nil {
			log.Printf("no instances in this region %s: %v", region, err)
			continue
		}

		// 处理该区域的实例
		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				instances[*instance.InstanceId] = instance

				//如果状态不是running，则跳过
				// if *instance.State.Name != "running" {
				// 	continue
				// }
				// var tags []string
				// for _, tag := range instance.Tags {
				// 	tags = append(tags, *tag.Value)
				// }
				// //找到tags中Name的值
				// var InstanceName string
				// for _, tag := range instance.Tags {
				// 	if *tag.Key == "Name" {
				// 		InstanceName = *tag.Value
				// 	}
				// }

				// var publicIPv4, privateIPv4 []string
				// if instance.PublicIpAddress != nil {
				// 	publicIPv4 = append(publicIPv4, *instance.PublicIpAddress)
				// }
				// if instance.PrivateIpAddress != nil {
				// 	privateIPv4 = append(privateIPv4, *instance.PrivateIpAddress)
				// }

				// asset := Ec2Asset{
				// 	InstanceName: InstanceName,
				// 	Instance:     instance,
				// 	InstanceID:   *instance.InstanceId,
				// 	Tags:         tags,
				// 	Region:       region,
				// 	Ec2Client:    client,
				// 	PublicIPv4:   publicIPv4,
				// 	PrivateIPv4:  privateIPv4,
				// }
				// allAssets = append(allAssets, asset)
			}
		}
	}

	return &instances
}

// func GetEc2Instance(client *ec2.EC2) ([]Ec2Asset, error) {
// 	// 创建一个空的 Assets 切片
// 	var allAssets []Ec2Asset
// 	result, err := client.DescribeInstances(nil)
// 	if err != nil {
// 		log.Printf("获取该区域EC2实例失败: %v", err)
// 		return nil, err
// 	}

// 	for _, reservation := range result.Reservations {
// 		for _, instance := range reservation.Instances {
// 			var tags []string
// 			for _, tag := range instance.Tags {
// 				tags = append(tags, *tag.Value)
// 			}

// 			asset := Ec2Asset{
// 				Instance:   instance,
// 				InstanceID: *instance.InstanceId,
// 				Tags:       tags,
// 			}
// 			allAssets = append(allAssets, asset)
// 		}
// 	}
// 	return allAssets, nil
// }

// type Assets []Ec2Asset

// func GetEc2() Assets {

// 	// 创建 EC2 服务客户端
// 	svc := NewEc2ClientWithRegion(config.AwsRegion)
// 	as := newAssets()
// 	as.GetAssets(svc)
// 	return *as
// }

// func (as *Assets) GetAssets(svc *ec2.EC2) {
// 	as.getInstance(svc)
// }

// // 获取Instance、InstanceID、Tags
// func (as *Assets) getInstance(svc *ec2.EC2) {
// 	result, err := svc.DescribeInstances(nil)
// 	if err != nil {
// 		log.Println("Failed to describe instances:", err)
// 	}

// 	for _, reservation := range result.Reservations {
// 		for _, instance := range reservation.Instances {
// 			var tags []string
// 			for _, tag := range instance.Tags {
// 				tags = append(tags, *tag.Value)
// 			}
// 			asset := Asset{
// 				Instance:   instance,
// 				InstanceID: *instance.InstanceId,
// 				Tags:       tags,
// 			}
// 			*as = append(*as, asset)
// 		}

// 	}
// }

// func newAssets() *Assets {
// 	assets := Assets{
// 		{
// 			Instance:   &ec2.Instance{},
// 			InstanceID: "",
// 		},
// 	}

// 	return &assets
// }
