package model

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 不带软删除的定义
type BaseModel struct {
	ID        uint      `gorm:"column:id;type:bigint(20) unsigned;primaryKey;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

// GormModel 带软删除的结构
type GormModel struct {
	BaseModel
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// 幻灯片表
type Slide struct {
	//BaseModel
	ID     uint   `gorm:"column:id;type:bigint(20) unsigned;primaryKey;AUTO_INCREMENT" json:"id"`
	Status uint   `gorm:"column:status;type:tinyint(3) unsigned;default:1;comment:状态,1:显示,0不显示;NOT NULL" json:"status"`
	Name   string `gorm:"column:name;type:varchar(50);comment:幻灯片分类;NOT NULL" json:"name"`
	Remark string `gorm:"column:remark;type:varchar(255);comment:分类备注;NOT NULL" json:"remark"`
}
