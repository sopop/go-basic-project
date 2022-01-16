package service

import (
	"demo/dao"
	"demo/model"
	"time"

	"gorm.io/gorm"
)

type goodsService struct {
	db *gorm.DB
}

func NewGoods() *goodsService {
	return &goodsService{
		db: dao.GormDB,
	}
}

// 根据ID查找商品
func (s *goodsService) GetDetail(id int) (model.Goods, error) {
	var goods model.Goods
	err := s.db.Where("id = ?", id).First(&goods).Error
	return goods, err
}

// 查找所有商品
func (s *goodsService) GetLists(data RequestInfo) ([]model.Goods, error) {
	var goods []model.Goods
	page := data.Page
	pageSize := data.PageSize
	offset := page * pageSize
	err := s.db.Select("*,FROM_UNIXTIME(create_time,'%Y-%m-%d %H:%i:%s') create_time_str").
		Where("status = 1").Limit(pageSize).Offset(offset).Find(&goods).Error

	return goods, err
}

// 添加商品
func (s *goodsService) AddGoods(data map[string]interface{}) error {
	data["status"] = 1
	data["cat_id"] = 1
	data["create_time"] = time.Now().Unix()
	data["update_time"] = time.Now().Unix()
	err := s.db.Model(&model.Goods{}).Create(data).Error
	return err
}

// 更新商品
func (s *goodsService) EditGoods(id int, data map[string]interface{}) error {
	data["update_time"] = time.Now().Unix()
	err := s.db.Model(&model.Goods{}).Where("id = ?", id).Updates(data).Error
	return err
}

// 删除商品
func (s *goodsService) DelGoods(id int) error {
	err := s.db.Where("id = ?", id).Delete(&model.Goods{}).Error
	return err
}
