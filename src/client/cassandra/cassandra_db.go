package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to a CassandraDB cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
