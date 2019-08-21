package cassandra

import (
	"{{ProjectName}}/pkg"
	"io/ioutil"
	"os"
	"time"
	"github.com/gocql/gocql"
)

type CassandraDb struct {
	log     *pkg.Logger
	Session *gocql.Session
}

type Option struct {
	Hosts       []string
	Port        int
	Username    string
	Password    string
	KeySpace    string
	Consistency string
	PageSize    int
	Timeout     time.Duration
	DataCenter  string
}

func New(log *pkg.Logger, option Option) CassandraDb {
	db := CassandraDb{log: log}
	var err error
	consistency := gocql.LocalOne
	err = consistency.UnmarshalText([]byte(option.Consistency))
	if err != nil {
		log.Errorf("Error in unmarshaling consistency level. Set default consistency LocalOne. Error: %s", err)
	} else {
		log.Infof("Consistency level set to %s", consistency)
	}

	log.Debug("Cassandra configs : ", option)

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
		log.Infof("Set datacenter to %s", option.DataCenter)
		cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(option.DataCenter)
		cluster.HostFilter = gocql.DataCentreHostFilter(option.DataCenter)
	}

	db.Session, err = cluster.CreateSession()

	if err != nil {
		log.Panic(err)
	}
	return db
}

func (c *CassandraDb) CreateTables() {
	historyMessagePath := os.Getenv("PWD") + "/../assets/cassandra/history_messages.cql"

	paths := []string{historyMessagePath}

	for _, path := range paths {
		result, err := ioutil.ReadFile(path)
		if err != nil {
			c.log.Error(err)
		}
		query := string(result)

		err = c.Session.Query(query).Exec()

		if err != nil {
			c.log.Error(err)
		}
	}
}
