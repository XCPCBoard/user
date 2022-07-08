package luogu

import (
	"XCPCer_board/dao"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	packageName = "luogu"
)

//SetUserMsgToRedis 将用户信息放入redis
func SetUserMsgToRedis(uid string, ctx context.Context) error {

	//get user msg
	res, err := ScrapeUser(uid)
	if err != nil {
		log.Errorf("%v get uid=%v message err:%v", packageName, uid, err)
	}
	//creat key map
	mapKey := make(map[string]int)
	for key, val := range res {
		mapKey[getUserIDRedisKey(uid, key)] = val
	}
	//set data to redis
	for key, val := range mapKey {
		err := dao.RedisClient.Set(ctx, key, val, 0).Err()
		if err != nil {
			log.Errorf("%v set redis data for uid=%v failed, err:%v\n", packageName, uid, err)
			return err
		}
	}
	return nil
}

//GetUserMsgFromRedis 获取用户某一keyWord的数据
func GetUserMsgFromRedis(uid string, keyWord string, ctx context.Context) (int, error) {

	//get data
	val, err := dao.RedisClient.Get(ctx, getUserIDRedisKey(uid, keyWord)).Result()
	if err != nil {
		log.Errorf("%v get redis data for uid=%v ,keyWord=%v, failed, err:%v\n", packageName, uid, keyWord, err)
		return -1, err
	}
	//str to int
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Errorf("%v get redis data strToInt err:%v\tand the return is %v:", packageName, err, num)
		return -1, err
	}
	return num, nil
}

//GetUserAllMsgFromRedis 获取用户的所有数据
func GetUserAllMsgFromRedis(uid string, ctx context.Context) (map[string]int, error) {

	mp := make(map[string]int)
	//Traverse key word
	for _, keyWord := range KeyWordListOfUser {

		//get data
		val, err := dao.RedisClient.Get(ctx, getUserIDRedisKey(uid, keyWord)).Result()
		if err != nil {
			log.Errorf("%v get redis data for uid=%v ,keyWord=%v, failed, err:%v\n", packageName, uid, keyWord, err)
			return nil, err
		}

		//str to int
		num, err := strconv.Atoi(val)
		if err != nil {
			log.Errorf("%v get redis data strToInt err:%v\tand the return is %v:", packageName, err, num)
			return nil, err
		}
		mp[keyWord] = num
	}

	return mp, nil
}

func getUserIDRedisKey(uid string, key any) string {
	return fmt.Sprintf("%v_id_%v_%v", packageName, uid, key)
}
