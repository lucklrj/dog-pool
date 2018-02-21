package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/fatih/color"
	"github.com/lucklrj/pool"
	"lib/config"
	"os"
	"strconv"
	"strings"
)

type Memcache string

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

func (t *Memcache) Set(args *SetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*memcache.Client)

	err = mc.Set(&memcache.Item{Key: args.Key, Value: []byte(args.Value)})
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	err = mc.Touch(args.Key, args.LeftTime)
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
	mc := conn.Client.(*memcache.Client)

	returnValue, err := mc.Get(args.Key)
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	*reply = ResultRead{Data: string(returnValue.Value)}
	return nil
}
func (t *Memcache) GetMulti(args *GetArgs, reply *interface{}) error {
	conn, err := getconn()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}
	defer MyMemcachePool.Put(conn)
	mc := conn.Client.(*memcache.Client)

	KeyList := strings.Split(args.Key, ",")
	returnValue, err := mc.GetMulti(KeyList)

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
	mc := conn.Client.(*memcache.Client)
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
	mc := conn.Client.(*memcache.Client)
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
		return memcache.New(config.MemcachedConfig.Host + ":" + strconv.Itoa(config.MemcachedConfig.Port))
	}
	MyMemcachePool.DestroyClient = func(c interface{}) {

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
