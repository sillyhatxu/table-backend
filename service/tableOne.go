package service

import (
	"encoding/json"
	"fmt"
	"github.com/sillyhatxu/convenient-utils/uuid"
	"github.com/sillyhatxu/table-backend/dao"
	"github.com/sillyhatxu/table-backend/dto"
	"github.com/sillyhatxu/table-backend/model"
	"github.com/sillyhatxu/table-backend/utils"
	"strings"
)

func TableOneList() ([]dto.TableOne, error) {
	array, err := dao.FindByTableType(model.TableTypeOne)
	if err != nil {
		return nil, err
	}
	var resultArray []dto.TableOne
	for _, ut := range array {
		var table dto.TableOne
		err := json.Unmarshal([]byte(ut.Content), &table)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, table)
	}
	return resultArray, nil
}

func TableOneAdd(addDTO dto.AddDTO) error {
	if addDTO.Content == "" {
		return fmt.Errorf("add table one error. content is null")
	}
	content := utils.FormatContext(addDTO.Content)
	contentArray := strings.Split(content, "\n")
	data := make(map[string]string)
	for i, c := range contentArray {
		if i == 0 {
			continue
		}
		rows := strings.Split(c, ":")
		if len(rows) > 2 {
			continue
		}
		if len(rows) == 1 {
			data[utils.FormatString(rows[0])] = ""
		} else {
			data[utils.FormatString(rows[0])] = utils.FormatString(rows[1])
		}
	}
	table := dto.TableOne{
		UserName:            data["姓名"],
		Gender:              data["性别"],
		Age:                 data["年龄"],
		Education:           data["学历"],
		Identity:            data["身份证"],
		PhoneNumber:         data["手机号"],
		Email:               data["邮箱"],
		IdentityAddress:     data["身份证地址"],
		Work:                data["工作单位"],
		WorkAddress:         data["单位地址"],
		WorkPhone:           data["单位电话"],
		ReasonForApplying:   data["申请原因"],
		BankName:            data["银行名称"],
		ReceiveAddress:      data["收卡地址"],
		ApplyTime:           data["申请时间"],
		ReferrerName:        data["推荐人姓名"],
		ReferrerPhoneNumber: data["推荐人手机号"],
		Remark:              data["备注"],
	}
	contentJSON, err := json.Marshal(table)
	if err != nil {
		return err
	}
	return dao.Insert(&model.UnknownTable{
		Id:        uuid.V4Upper32(),
		TableType: model.TableTypeOne,
		Status:    model.StatusEnable,
		Content:   string(contentJSON),
	})
}
