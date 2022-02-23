package note

import (
	"context"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	leveldbUtil "github.com/syndtr/goleveldb/leveldb/util"

	"github.com/go-redis/redis/v8"
)

//11.1.1~11.1.3 LevelDB基本使用
func LeveldbBasic() {
	//打开/新建数据库
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	//关闭数据库
	defer db.Close()

	//放入/修改数据
	db.Put([]byte("user-1"), []byte("{\"username\":\"1\"}"), nil)

	//删除数据
	// db.Delete([]byte("user-1"), nil)

	//取出数据,注意不要修改此data
	data, _ := db.Get([]byte("user-1"), nil)
	fmt.Println("data=", string(data))

	//判断是否存在数据
	ok, _ := db.Has([]byte("user-1"), nil)
	fmt.Println("Has \"user-1\" ?", ok)

	//批量写数据
	batch := new(leveldb.Batch)
	batch.Put([]byte("user-2"), []byte("{\"username\":\"2\"}"))
	batch.Put([]byte("user-3"), []byte("{\"username\":\"3\"}"))
	batch.Delete([]byte("user-1"))
	batch.Put([]byte("user-n1"), []byte("{\"username\":\"n1\"}"))
	db.Write(batch, nil)
}

//11.1.4 LevelDB遍历
func LeveldbIterate() {
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	batch := new(leveldb.Batch)
	for i := 1; i < 11; i++ {
		batch.Put(
			[]byte(fmt.Sprintf("user-%v", i)),
			[]byte(fmt.Sprintf("{\"name\":\"u%v\"}", i)))
	}
	db.Write(batch, nil)

	//遍历指定范围的数据//&leveldbUtil.Range处填nil为完整数据库
	iter := db.NewIterator(
		&leveldbUtil.Range{
			Start: []byte("user-3"),
			Limit: []byte("user-8"),
		}, nil)
	for iter.Next() {
		fmt.Printf("%v=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	//遍历指定前缀的数据
	iter = db.NewIterator(leveldbUtil.BytesPrefix([]byte("user-")), nil)
	for iter.Next() {
		fmt.Printf("%v=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}

//11.1.5~11.1.6 LevelDB事务与快照
func LeveldbTransactionAndSnapshot() {
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	ss, err := db.GetSnapshot()
	if err != nil {
		panic(err)
	}
	defer ss.Release()
	t, err := db.OpenTransaction()
	if err != nil {
		panic(err)
	}
	batch := new(leveldb.Batch)
	for i := 1; i < 11; i++ {
		batch.Put(
			[]byte(fmt.Sprintf("cat-%v", i)),
			[]byte(fmt.Sprintf("{\"name\":\"c%v\"}", i)))
	}
	t.Write(batch, nil)
	//t.Discard()
	t.Commit()
	ok, _ := db.Has([]byte("cat-1"), nil)
	fmt.Println("db Has \"cat-1\" ?", ok)
	ok, _ = ss.Has([]byte("cat-1"), nil)
	fmt.Println("ss Has \"cat-1\" ?", ok)
}

//11.2 Redis的基本操作
func RedisBasic() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	db.Do(ctx, "set", "k1", "v1")
	res, err := db.Do(ctx, "get", "k2").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("该key不存在")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("res=", res.(string))
	}
	db.Do(ctx, "set", "b1", true)
	db.Do(ctx, "set", "b2", 0)
	b, err := db.Do(ctx, "mget", "b1", "b2").BoolSlice()
	if err == nil {
		fmt.Println("b=", b)
	}
	db.Set(ctx, "t1", time.Now(), 0)
	t := db.Get(ctx, "t1").Val()
	if err == nil {
		fmt.Println("t=", t)
	}
}

//11.2.6 Redis管道
func RedisPipeline() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	pipe := db.Pipeline()
	t1 := pipe.Get(ctx, "t1")
	fmt.Println("pipe执行前的t1=", t1)
	for i := 0; i < 10; i++ {
		pipe.Set(ctx, fmt.Sprintf("p%v", i), i, 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("pipe执行后的t1=", t1)

	cmds, err := db.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 10; i++ {
			pipe.Get(ctx, fmt.Sprintf("p%v", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i, cmd := range cmds {
		fmt.Printf("p%v=%v\n", i, cmd.(*redis.StringCmd).Val())
	}
}

//11.2.7 Redis事务
func RedisTransaction() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		err := db.Watch(ctx, func(tx *redis.Tx) (err error) {
			pipe := tx.Pipeline()
			err = pipe.IncrBy(ctx, "p1", 100).Err()
			if err != nil {
				return
			}
			err = pipe.DecrBy(ctx, "p0", 100).Err()
			if err != nil {
				return
			}
			_, err = pipe.Exec(ctx)
			return
		}, "p0")
		if err == nil {
			fmt.Println("事务commit成功")
			break
		} else if err == redis.TxFailedErr {
			fmt.Println("事务执行失败,这次是第", i, "次尝试")
			continue
		} else {
			panic(err)
		}
	}
}

//11.2.8 Redis遍历
func RedisIterate() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	iter := db.Scan(ctx, 0, "p*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Printf("key=%v, value=%v\n", iter.Val(), db.Get(ctx, iter.Val()).Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	db.HSet(ctx, "h1", "f1", "v1", "f2", "v2", "f3", "v3")
	iter = db.HScan(ctx, "h1", 0, "*", 0).Iterator()
	for i := 0; iter.Next(ctx); i++ {
		if i%2 == 0 {
			fmt.Printf("field=%v, ", iter.Val())
		} else {
			fmt.Printf("value=%v\n", iter.Val())
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

//11.2.9 将Redis Hash扫描至Go结构体
type RedisHash struct {
	Name   string `redis:"name"`
	Id     int    `redis:"id"`
	Online bool   `redis:"online"`
}

func RedisHashToStruct() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()
	var rh1 = RedisHash{
		Name:   "rhName",
		Id:     123,
		Online: true,
	}
	db.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, "rh1", "name", rh1.Name)
		pipe.HSet(ctx, "rh1", "id", rh1.Id)
		pipe.HSet(ctx, "rh1", "online", rh1.Online)
		return nil
	})
	var rh2 RedisHash
	if err := db.HGetAll(ctx, "rh1").Scan(&rh2); err != nil {
		panic(err)
	}
	fmt.Printf("rh2=%+v\n", rh2)
}
