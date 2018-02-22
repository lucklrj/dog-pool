package memcache

import (
	"github.com/kklis/gomemcache"
	"github.com/fatih/color"
	"github.com/lucklrj/pool"
	"lib/config"
	"os"
	"strings"
)

type Memcache string

type SetArgs struct {
	Key      string
	Value    string
	LeftTime int64
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

func (t *Memcache) Set(args *SetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*gomemcache.Memcache)
	
	err = mc.Set(args.Key, []uint8(args.Value), 0, args.LeftTime)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultWrite{Success: true}
	return nil
}

func (t *Memcache) Get(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*gomemcache.Memcache)

	returnValue, _, err := mc.Get(args.Key)
	
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultRead{Data: string(returnValue)}
	return nil
}
func (t *Memcache) GetMulti(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*gomemcache.Memcache)

	KeyList := strings.Split(args.Key, ",")
	returnValue, err := mc.GetMulti(KeyList...)

	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	returnData := make(map[string]string)
	for key, value := range returnValue {
		returnData[key] = string(value.Value)
	}
	*reply = ResultRead{Data: returnData}
	return nil
}

func (t Memcache) Delete(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*gomemcache.Memcache)
	err = mc.Delete(args.Key)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultWrite{Success: true}
	return nil
}

func (t Memcache) FlushAll(args *EmptyArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*gomemcache.Memcache)
	err = mc.FlushAll()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultWrite{Success: true}
	return nil
}

var MyMemcachePool *pool.Pool

func init() {
	MyMemcachePool = new(pool.Pool)
	MyMemcachePool.MaxOpenConns = config.MemcachedConfig.MaxOpenConns
	MyMemcachePool.ConnMaxLifeTime = config.MemcachedConfig.ConnTimeOut
	MyMemcachePool.ConnTimeOut = config.MemcachedConfig.ConnTimeOut

	MyMemcachePool.CreateClient = func() interface{} {
		memc, _ := gomemcache.Connect(config.MemcachedConfig.Host, config.MemcachedConfig.Port)
		return memc
	}
	MyMemcachePool.DestroyClient = func(c interface{}) {
		c.(*gomemcache.Memcache).Close()
	}

	err := MyMemcachePool.Init()

	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
}
func getconn() (conn *pool.Coon, err error) {
	conn, err = MyMemcachePool.Get()
	if err != nil {
		return nil, err
	} else {
		return conn, nil
		//return conn.Client.(*memcache.Client), nil
	}
}
