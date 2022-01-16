package model

type Goods struct {
	Id            int    `orm:"id" json:"id"`                           // 商品ID
	CatId         int    `orm:"cat_id" json:"cat_id"`                   // 分类ID
	Name          string `orm:"name" json:"name"`                       // 商品名称
	Detail        string `orm:"detail" json:"detail"`                   // 商品描述
	CreateTime    int64  `orm:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int64  `orm:"update_time" json:"update_time"`         // 更新时间
	CreateTimeStr string `orm:"create_time_str" json:"create_time_str"` // 创建时间格式化
	Status        int    `orm:"status" json:"status"`                   // 状态
}

func (*Goods) TableName() string {
	return "cms_goods"
}
