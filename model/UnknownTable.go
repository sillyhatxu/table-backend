package model

import "time"

const (
	StatusEnable  = 1
	StatusDisable = 0
	/**
	L：超龄、逾期申请表
	姓名,性别,年龄,学历,身份证号,手机号,邮箱号,身份证地址,工作单位,单位地址,单位电话,申请原因,银行名称,收卡地址,申请时间,推荐人姓名,推荐人手机号码,备     注
	*/
	TableTypeOne = 1
)

type UnknownTable struct {
	Id               string    `json:"id" mapstructure:"id"`
	TableType        int       `json:"table_type" mapstructure:"table_type"`
	Status           int       `json:"status" mapstructure:"status"`
	Content          string    `json:"content" mapstructure:"content"`
	LastModifiedTime time.Time `json:"last_modified_time" mapstructure:"last_modified_time"`
	CreatedTime      time.Time `json:"created_time" mapstructure:"created_time"`
}
