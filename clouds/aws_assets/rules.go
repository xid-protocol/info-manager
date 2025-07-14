package aws_assets

// type LarkContent struct {
// 	Region       string
// 	InstanceID   string
// 	PublicIP     []string
// 	PortRange    string
// 	Protocol     string
// 	InstanceName string
// }

// // 获取到0.0.0.0/0的资产
// func Get000000Asset() ([]LarkContent, [][]string) {
// 	secAssets := GetSecAssets()
// 	var Assets0000 [][]string
// 	var larkContent []LarkContent
// 	for _, secAsset := range secAssets {
// 		for _, secGroup := range secAsset.SecGroups {
// 			for _, rule := range secGroup.Rules {
// 				// 获取ip范围
// 				for _, ipRange := range rule.Ipranges {
// 					//如果cidrIp为0.0.0.0/0
// 					for _, cidrIp := range ipRange.CidrIp {
// 						if cidrIp == "0.0.0.0/0" {
// 							larkContent = append(larkContent, LarkContent{
// 								Region:       secAsset.Ec2Asset.Region,
// 								InstanceID:   secAsset.Ec2Asset.InstanceID,
// 								PublicIP:     secAsset.Ec2Asset.PublicIPv4,
// 								PortRange:    rule.PortRange,
// 								Protocol:     rule.Protocol,
// 								InstanceName: secAsset.Ec2Asset.InstanceName,
// 							})
// 							log.Println(larkContent)
// 							publicIP := fmt.Sprintf("%v", secAsset.Ec2Asset.PublicIPv4)
// 							privateIP := fmt.Sprintf("%v", secAsset.Ec2Asset.PrivateIPv4)
// 							//log.Printf("InstanceID: %s, Region: %s, Port: %s, CidrIp: %s", secAsset.Ec2Asset.InstanceID, secAsset.Ec2Asset.Region, rule.PortRange, cidrIp)
// 							Assets0000 = append(Assets0000, []string{secAsset.Ec2Asset.Region, secAsset.Ec2Asset.InstanceID, secAsset.Ec2Asset.InstanceName, publicIP, privateIP, rule.PortRange, rule.Protocol, cidrIp, secGroup.Name, ipRange.Description})
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return larkContent, Assets0000
// }

// func Get000000AssetToCSV() {
// 	_, Assets0000 := Get000000Asset()
// 	WriteToCSV(Assets0000, "./000000.csv")
// }

// func AWSEc2PortAlert() {
// 	larkContent, _ := Get000000Asset()
// 	//告警到lark
// 	msg := ""
// 	instanceWhitelist := InstanceWhitelist()
// 	for _, asset := range larkContent {
// 		//如果publicIP为空，则不告警
// 		if len(asset.PublicIP) == 0 {
// 			continue
// 		}
// 		//排除白名单
// 		whitelist := false
// 		for _, instance := range instanceWhitelist {
// 			if instance.Serveraddress == asset.InstanceID {

// 				log.Println(instance)
// 				log.Println(asset)

// 				if instance.Serverport == asset.PortRange {

// 					if instance.Serverprotocol == asset.Protocol {
// 						whitelist = true
// 					}
// 				}
// 			}

// 		}
// 		if !whitelist {
// 			msg = msg + fmt.Sprintf("%s, %s, %s, %s, %s\\n", asset.Region, asset.InstanceID, asset.InstanceName, asset.Protocol, asset.PortRange)
// 		}

// 	}
// 	content := lark.SgTeamp(msg)
// 	lark.SendToLark(content)

// }
