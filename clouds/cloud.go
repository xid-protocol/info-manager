package clouds

import (
	"log"

	"github.com/xid-protocol/info-manager/clouds/aws_assets"
)

func CloudMonitor() {

	ec2Instances := aws_assets.GetEc2FromAllRegions()
	aws_assets.SetAWSEc2Info(ec2Instances)
	//安全组
	securityGroups := aws_assets.SecGroupMonitor()
	aws_assets.SetAWSSecGroupInfo(securityGroups)
	log.Println(securityGroups)
}
