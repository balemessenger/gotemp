package db

import (
	"fmt"
	"github.com/scylladb/gocqlx/qb"
	"{{ProjectName}}/pkg"
	"strings"
	"time"
	"io/ioutil"
	"github.com/gocql/gocql"
)

type CassandraDB struct {
	session *gocql.Session
}

type CassConfig struct {
	Hosts         []string
	Port          int
	Username      string
	Password      string
	KeySpace      string
	Consistency   string
	PageSize      int
	Timeout       time.Duration
	DataCenter    string
	PartitionSize int32
}

func NewCassandra(option CassConfig) *CassandraDB {
	db := CassandraDB{}
	var err error
	consistency := gocql.LocalOne
	err = consistency.UnmarshalText([]byte(option.Consistency))
	if err != nil {
		pkg.Logger.Errorf("Error in unmarshaling consistency level. Set default consistency LocalOne. Error: %s", err.Error())
	} else {
		pkg.Logger.Infof("Consistency level set to %s", consistency)
	}

	pkg.Logger.Debug("Cassandra configs : ", option)

	cluster := gocql.NewCluster(option.Hosts...)
	cluster.Keyspace = option.KeySpace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: option.Username,
		Password: option.Password}
	cluster.PageSize = option.PageSize
	cluster.Port = option.Port
	cluster.Consistency = consistency
	cluster.Timeout = option.Timeout
	if option.DataCenter != "" {
		pkg.Logger.Infof("Set datacenter to %s", option.DataCenter)
		cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(option.DataCenter)
		cluster.HostFilter = gocql.DataCentreHostFilter(option.DataCenter)
	}

	db.session, err = cluster.CreateSession()

	if err != nil {
		pkg.Logger.Panic(err)
	}
	return &db
}

func (c *CassandraDB) PrintQuery(stmt string, mp qb.M) {
	var result = stmt

	for key, value := range mp {
		t1 := key + "=?"
		t2 := key + "<=?"
		t3 := key + ">=?"
		v := fmt.Sprintf("%v", value)
		result = strings.Replace(result, t1, key+"="+v, 1)
		result = strings.Replace(result, t2, key+"<="+v, 1)
		result = strings.Replace(result, t3, key+">="+v, 1)
	}
	pkg.Logger.Debug(result)
}

//CreateTables(os.Getenv("PWD") + "/../assets/cassandra/example.cql")
func (c *CassandraDB) CreateTables(paths []string) {
	for _, path := range paths {
		result, err := ioutil.ReadFile(path)
		if err != nil {
			pkg.Logger.Error(err)
		}
		query := string(result)

		err = c.session.Query(query).Exec()

		if err != nil {
			pkg.Logger.Error(err)
		}
	}
}
