package format

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func elasticache_replication_group_import(r *InstanceDiff, cond map[string]string) string {
	name, ok := cond["replication_group_id"]

	if !ok {
		return ""
	}

	svc := elasticache.New(session.New())

	input := &elasticache.DescribeReplicationGroupsInput{
		ReplicationGroupId: aws.String(name),
	}

	result, err := svc.DescribeReplicationGroups(input)

	var buffer bytes.Buffer

	if err != nil {
		buffer.WriteString("No Import: There is no redis with replication_group_id " + name)
		buffer.WriteString("\n")
		return buffer.String()
		//log.Fatal(err.Error())
	}

	if len(result.ReplicationGroups) == 1 {
		buffer.WriteString("terraform import  ")
		buffer.WriteString(r.Addr.String() + "  ")
		buffer.WriteString(*(result.ReplicationGroups[0].ReplicationGroupId) + "\n\n")
		return buffer.String()
	}

	return buffer.String()
}
