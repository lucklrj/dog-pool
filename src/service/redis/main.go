package redis

import (
	"github.com/fatih/color"
	"github.com/garyburd/redigo/redis"
	"github.com/lucklrj/pool"
	"lib/config"
	"os"
	"strconv"
)

type Redis string

type SetArgs struct {
	Key      string
	Value    string
	LeftTime int32
}
type GetArgs struct {
	Key string
}
type EmptyArgs string

type ResultError struct {
	Error string
}
type ResultRead struct {
	Data interface{}
}
type ResultWrite struct {
	Success bool
}
type ResourceConn struct {
	redis.Conn
}

func (t *Redis) Set(args *SetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyRedisPool.Put(conn)
	client := conn.Client.(ResourceConn)

	_, err = client.Do("Set", args.Key, args.Value)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	if args.LeftTime > 0 {
		_, err = client.Do("expire", args.Key, args.LeftTime)
		if err != nil {
			*reply = ResultError{Error: err.Error()}
			return nil
		}
	}
	*reply = ResultWrite{Success: true}
	return nil
}

func (t *Redis) Get(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyRedisPool.Put(conn)
	client := conn.Client.(ResourceConn)

	value, err := client.Do("Get", args.Key)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	if value == nil {
		*reply = ResultRead{Data: ""}
	} else {
		*reply = ResultRead{Data: string(value.([]uint8))}
	}
	return nil
}

func (t *Redis) Delete(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyRedisPool.Put(conn)
	client := conn.Client.(ResourceConn)
	
	_, err = client.Do("Del", args.Key)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultWrite{Success: true}
	return nil
}


var MyRedisPool *pool.Pool

func init() {
	MyRedisPool = new(pool.Pool)
	MyRedisPool.MaxOpenConns = config.MemcachedConfig.MaxOpenConns
	MyRedisPool.ConnMaxLifeTime = config.MemcachedConfig.ConnTimeOut
	MyRedisPool.ConnTimeOut = config.MemcachedConfig.ConnTimeOut

	MyRedisPool.CreateClient = func() interface{} {

		client, _ := redis.Dial("tcp", config.RedisConfig.Host+":"+strconv.Itoa(config.RedisConfig.Port))

		//todo 如果设置了密码，需要验证密码是否正确
		return ResourceConn{client}
	}
	MyRedisPool.DestroyClient = func(c interface{}) {
		c.(*pool.Coon).Client.(ResourceConn).Close()
	}

	err := MyRedisPool.Init()

	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
}
func getconn() (conn *pool.Coon, err error) {
	conn, err = MyRedisPool.Get()
	if err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}
