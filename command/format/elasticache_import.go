package format

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func elasticache_import(r *InstanceDiff, cond map[string]string) string {
	name, ok := cond["cluster_id"]

	if !ok {
		return ""
	}

	svc := elasticache.New(session.New())

	input := &elasticache.DescribeCacheClustersInput{
		CacheClusterId: aws.String(name),
	}

	result, err := svc.DescribeCacheClusters(input)

	var buffer bytes.Buffer

	if err != nil {
		buffer.WriteString("No Import: There is no Elastic Cache Cluster named " + name)
		buffer.WriteString("\n")
		return buffer.String()
	}

	buffer.WriteString("terraform import  ")
	buffer.WriteString(r.Addr.String() + "  ")
	buffer.WriteString(*(result.CacheClusters[0].CacheClusterId) + "\n\n")
	return buffer.String()
}
