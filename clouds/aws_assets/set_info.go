package aws_assets

import (
	"context"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/colin-404/logx"
	"github.com/xid-protocol/info-manager/db/repositories"
	"github.com/xid-protocol/xidp/protocols"
)

func SetAWSEc2Info(instances *map[string]*ec2.Instance) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	for instanceId := range *instances {

		// info := map[string]interface{}{"email": email, "type": "user_email"}
		xid := protocols.GenerateXid(instanceId)
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/aws/instance")
		if err != nil {
			logx.Errorf("check aws ec2 failed: %v", err)
			continue
		}

		if exists {
			logx.Infof("aws ec2 %s exists, need to modify", instanceId)

			//创建XID
			info := protocols.NewInfo(instanceId, "aws-instanceid")
			metadata := protocols.NewMetadata("modify", "/info/aws/instance", "application/json")
			xidRecord := protocols.NewXID(&info, &metadata, (*instances)[instanceId])
			logx.Infof("xid: %v", xidRecord)
			// 插入MongoDB
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, xidRecord)
			if err != nil {
				logx.Errorf("modify aws ec2 %s failed: %v", instanceId, err)
			} else {
				logx.Infof("modify aws ec2 %s success", instanceId)
			}
			continue
		}

		info := protocols.NewInfo(instanceId, "aws-instanceid")
		metadata := protocols.NewMetadata("create", "/info/aws/instance", "application/json")
		xidRecord := protocols.NewXID(&info, &metadata, (*instances)[instanceId])
		logx.Infof("xid: %v", xidRecord)
		// 插入MongoDB
		err = repo.Insert(ctx, xidRecord)
		if err != nil {
			logx.Errorf("insert aws ec2 %s failed: %v", instanceId, err)
		} else {
			logx.Infof("insert aws ec2 %s success", instanceId)
		}
	}
}

func SetAWSSecGroupInfo(secGroups *map[string][]*ec2.DescribeSecurityGroupsOutput) {
	repo := repositories.NewXidInfoRepository()
	ctx := context.Background()
	for groupID := range *secGroups {
		for _, secGroup := range (*secGroups)[groupID] {
			for _, sg := range secGroup.SecurityGroups {
				for _, ipPermission := range sg.IpPermissions {
					logx.Infof("ipPermission: %v", ipPermission)
					for _, ipRange := range ipPermission.IpRanges {
						logx.Infof("ipRange: %v", ipRange)
					}
				}
			}
		}

		xid := protocols.GenerateXid(groupID)
		exists, err := repo.CheckXidInfoExists(ctx, xid, "/info/aws/secgroup")
		if err != nil {
			logx.Errorf("check aws secgroup failed: %v", err)
			continue
		}

		if exists {
			logx.Infof("aws secgroup %s exists, need to modify", groupID)
			info := protocols.NewInfo(groupID, "aws-secgroup")
			metadata := protocols.NewMetadata("modify", "/info/aws/secgroup", "application/json")
			xidRecord := protocols.NewXID(&info, &metadata, (*secGroups)[groupID])
			logx.Infof("xid: %v", xidRecord)
			err = repo.UpdateXidInfo(ctx, xidRecord.Xid, xidRecord.Metadata.Path, xidRecord)
			if err != nil {
				logx.Errorf("modify aws secgroup %s failed: %v", groupID, err)
			} else {
				logx.Infof("modify aws secgroup %s success", groupID)
			}
			continue
		}

		info := protocols.NewInfo(groupID, "aws-secgroup")
		metadata := protocols.NewMetadata("create", "/info/aws/secgroup", "application/json")
		xidRecord := protocols.NewXID(&info, &metadata, (*secGroups)[groupID])
		logx.Infof("xid: %v", xidRecord)
		err = repo.Insert(ctx, xidRecord)
		if err != nil {
			logx.Errorf("insert aws secgroup %s failed: %v", groupID, err)
		} else {
			logx.Infof("insert aws secgroup %s success", groupID)
		}
	}
}
