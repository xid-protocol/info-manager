package aws_assets

// type Ec2SecAsset struct {
// 	Ec2Asset  Ec2Asset
// 	SecGroups []SecGroup
// }

// func GetSecAssets() []Ec2SecAsset {
// 	var assets []Ec2SecAsset
// 	// regions, err := GetAllRegions()
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// }
// 	// log.Println(regions)

// 	// // 创建所有区域的客户端
// 	// clients := make(map[string]*ec2.EC2)
// 	// for _, region := range regions {
// 	// 	sess, err := session.NewSession(&aws.Config{
// 	// 		Region:      aws.String(region),
// 	// 		Credentials: credentials.NewStaticCredentials(config.AwsApiKey, config.AwsSecretKey, ""),
// 	// 	})
// 	// 	if err != nil {
// 	// 		log.Printf("Failed to create session for region %s: %v", region, err)
// 	// 		continue
// 	// 	}
// 	// 	clients[region] = ec2.New(sess)
// 	// }

// 	ec2Assets, err := GetEc2FromAllRegions()
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	for _, ec2Asset := range ec2Assets {
// 		// if ec2Asset.InstanceID != "i-0620a66a1c1690b01" {
// 		// 	continue
// 		// }
// 		log.Printf("Region: %s, instance: %s", ec2Asset.Region, ec2Asset.InstanceID)
// 		//获取安全组
// 		securityGroups := ec2Asset.Instance.SecurityGroups
// 		for _, securityGroup := range securityGroups {
// 			//获取安全组规则
// 			secGroups := GetSgByID(securityGroup.GroupId, ec2Asset.Ec2Client)
// 			// log.Println(ec2Asset.InstanceID)
// 			log.Println(secGroups)
// 			assets = append(assets, Ec2SecAsset{
// 				Ec2Asset:  ec2Asset,
// 				SecGroups: []SecGroup{secGroups},
// 			})
// 		}
// 	}

// 	return assets
// }
