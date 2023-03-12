package rank

import (
	"github.com/XCPCBoard/common/dao"
	"github.com/XCPCBoard/common/errors"
	"github.com/XCPCBoard/common/logger"
	"github.com/XCPCBoard/user/entity"
	"github.com/XCPCBoard/user/service"
	"gorm.io/gorm/utils"
)

// UpdateRankById 更新用户rank
func UpdateRankById(id int, rank int) *errors.MyError {
	res := dao.DBClient.Model(&entity.Accounts{}).Where("id=?", id).Update("rank", rank)
	return service.CreatError(res, "更新用户rank数", "id")
}

// UpdateTotalById 更新用户total 过题总数
func UpdateTotalById(id int, total int) *errors.MyError {
	res := dao.DBClient.Model(&entity.Accounts{}).Where("id=?", id).Update("total", total)
	return service.CreatError(res, "更新用户rank数", "id")
}

// SelectAccountPagingOrderByTotal 分页查询，order by 总分数
//
//	@param pageNum 页数
//	@param pageNum 页的大小
func SelectAccountPagingOrderByTotal(pageNum int, pageSize int) (map[string]interface{}, *errors.MyError) {
	offset := (pageNum - 1) * pageSize
	var users []entity.Accounts
	res := dao.DBClient.Model(&entity.Accounts{}).Offset(offset).Limit(pageSize).Find(&users)
	if res.Error != nil {
		return nil, errors.CreateError(errors.INNER_ERROR.Code, "分页查询，order by 总分数失败", res.Error)
	}
	data := make(map[string]interface{})
	for idx, i := range users {
		temp := map[string]string{
			"id":    utils.ToString(i.Id),
			"total": utils.ToString(i.Total),
			"rank":  utils.ToString(i.Rank),
		}
		if err := GetAllRedisData(utils.ToString(i.Id), temp); err != nil {
			logger.L.Err(err, 0)
			return nil, err
		}
		tempIdx := idx + offset
		data[utils.ToString(tempIdx)] = temp
	}

	return data, nil
}
