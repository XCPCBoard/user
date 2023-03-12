package rank

import (
	"context"
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/utils/keys"
	"github.com/go-redis/redis/v8"
)

//-----------------------------------------------------------------
//			日后要做迁移，解代码的耦合度，这里迁移到XCPCBoard/utils/keys中
//-----------------------------------------------------------------

// GetAllRedisData 获取用户全部网站的数据（目前先包含cf和nowcoder
func GetAllRedisData(id string, data map[string]string) *errors.MyError {
	type Func func(string) (string, string)
	funcList := []Func{
		NowcoderRatingRedis, NowcoderPassAmountRedis, CodeforcesRatingRedis,
		CodeforcesMaxRankingRedis, CodeforcesPassAmountRedis,
	}
	for _, i := range funcList {
		k, v := i(id)
		redisValue, err := dao.RedisClient.Get(context.Background(), v).Result()
		if err == redis.Nil {
			redisValue = "nil" //redis里没有数据
		} else if err != nil {
			return errors.CreateError(errors.INNER_ERROR.Code, "get redis数据错误", v)
		}
		data[k] = redisValue
	}
	return nil
}

func NowcoderRatingRedis(id string) (string, string) {
	return keys.BuildKeyWithSiteKind(keys.NowcoderKey, keys.RatingKey), keys.NowcoderRatingKey(id)
}
func NowcoderPassAmountRedis(id string) (string, string) {
	return keys.BuildKeyWithSiteKind(keys.NowcoderKey, keys.PassAmountKey), keys.NowcoderPassAmountKey(id)
}

func CodeforcesRatingRedis(id string) (string, string) {
	return keys.BuildKeyWithSiteKind(keys.CodeforcesKey, keys.RatingKey), keys.CodeforcesRatingKey(id)
}
func CodeforcesMaxRankingRedis(id string) (string, string) {
	return keys.BuildKeyWithSiteKind(keys.CodeforcesKey, keys.MaxRankingKey), keys.CodeforcesMaxRankingKey(id)
}
func CodeforcesPassAmountRedis(id string) (string, string) {
	return keys.BuildKeyWithSiteKind(keys.CodeforcesKey, keys.PassAmountKey), keys.CodeforcesPassAmountKey(id)
}
