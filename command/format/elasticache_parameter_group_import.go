package format

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func elasticache_parameter_group_import(r *InstanceDiff, cond map[string]string) string {
	name, ok := cond["name"]

	if !ok {
		return ""
	}

	svc := elasticache.New(session.New())

	input := &elasticache.DescribeCacheParameterGroupsInput{
		CacheParameterGroupName: aws.String(name),
	}

	result, err := svc.DescribeCacheParameterGroups(input)

	var buffer bytes.Buffer

	if err != nil {
		buffer.WriteString("No Import: There is no Elastic Cache Cluster Parameter Group named " + name)
		buffer.WriteString("\n")
		return buffer.String()
	}

	buffer.WriteString("terraform import  ")
	buffer.WriteString(r.Addr.String() + "  ")
	buffer.WriteString(*(result.CacheParameterGroups[0].CacheParameterGroupName) + "\n\n")
	return buffer.String()
}
