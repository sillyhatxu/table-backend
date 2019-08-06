package service

import (
	"encoding/json"
	"fmt"
	"github.com/sillyhatxu/convenient-utils/uuid"
	"github.com/sillyhatxu/table-backend/dao"
	"github.com/sillyhatxu/table-backend/dto"
	"github.com/sillyhatxu/table-backend/model"
	"github.com/sillyhatxu/table-backend/utils"
	"github.com/tealeg/xlsx"
	"strconv"
	"strings"
)

func TableOneList() ([]dto.TableOneDTO, error) {
	array, err := dao.FindByTableType(model.TableTypeOne)
	if err != nil {
		return nil, err
	}
	var resultArray []dto.TableOneDTO
	for _, ut := range array {
		var table dto.TableOne
		err := json.Unmarshal([]byte(ut.Content), &table)
		if err != nil {
			return nil, err
		}
		resultArray = append(resultArray, *&dto.TableOneDTO{Id: ut.Id, TableOne: table})
	}
	return resultArray, nil
}

func TableOneClear() error {
	return dao.ClearAll(model.StatusDisable)
}

func TableOneFindById(id string) (*dto.TableOne, error) {
	ut, err := dao.FindById(id)
	if err != nil {
		return nil, err
	}
	var table dto.TableOne
	err = json.Unmarshal([]byte(ut.Content), &table)
	if err != nil {
		return nil, err
	}
	return &table, nil
}

func TableOneUpdate(id string, table dto.TableOne) error {
	if id == "" {
		return fmt.Errorf("ID不能为空")
	}
	contentJSON, err := json.Marshal(table)
	if err != nil {
		return err
	}
	if table.Identity == "" {
		return fmt.Errorf("身份证号不能为空")
	}
	return dao.UpdateContent(&model.UnknownTable{
		Id:             id,
		Identification: table.Identity,
		TableType:      model.TableTypeOne,
		Status:         model.StatusEnable,
		Content:        string(contentJSON),
	})
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
	if table.Identity == "" {
		return fmt.Errorf("身份证号不能为空")
	}
	checkArray, err := dao.FindByIdentificationTableType(table.Identity, model.TableTypeOne)
	if err != nil {
		return err
	}
	if len(checkArray) > 0 {
		return fmt.Errorf("这条记录已经录入了")
	}
	return dao.Insert(&model.UnknownTable{
		Id:             uuid.V4Upper32(),
		Identification: table.Identity,
		TableType:      model.TableTypeOne,
		Status:         model.StatusEnable,
		Content:        string(contentJSON),
	})
}

func ExportExtra() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	tableArray, err := TableOneList()
	if err != nil {
		return nil, err
	}
	row := sheet.AddRow()
	row.SetHeightCM(0.65) //设置每行的高度
	cell := addCell(row)
	cell.SetString("序号")
	cell = addCell(row)
	cell.SetString("职务")
	cell = addCell(row)
	cell.SetString("姓名")
	cell = addCell(row)
	cell.SetString("性别")
	cell = addCell(row)
	cell.SetString("年龄")
	cell = addCell(row)
	cell.SetString("身份证号")
	cell = addCell(row)
	cell.SetString("手机号码")
	cell = addCell(row)
	cell.SetString("省份")
	cell = addCell(row)
	cell.SetString("推荐人")
	cell = addCell(row)
	cell.SetString("推荐人手机号")
	cell = addCell(row)
	cell.SetString("无智能手机超龄人员")
	cell = addCell(row)
	cell.SetString("备注")
	for i, table := range tableArray {
		row := sheet.AddRow()
		row.SetHeightCM(0.65)
		cell := addCell(row)
		cell.SetString(strconv.Itoa(i + 1))
		cell = addCell(row)
		cell.SetString("会员")
		cell = addCell(row)
		cell.SetString(table.TableOne.UserName)
		cell = addCell(row)
		cell.SetString(table.TableOne.Gender)
		cell = addCell(row)
		cell.SetString(table.TableOne.Age)
		cell = addCell(row)
		cell.SetString(table.TableOne.Identity)
		cell = addCell(row)
		cell.SetString(table.TableOne.PhoneNumber)
		cell = addCell(row)
		cell.SetString(getProvince(table.TableOne.IdentityAddress))
		cell = addCell(row)
		cell.SetString(table.TableOne.ReferrerName)
		cell = addCell(row)
		cell.SetString(table.TableOne.ReferrerPhoneNumber)
		cell = addCell(row)
		cell.SetString("")
		cell = addCell(row)
		cell.SetString("")
	}
	return file, nil
}

func getProvince(identityAddress string) string {
	if strings.Contains(identityAddress, "北京") {
		return "北京"
	} else if strings.Contains(identityAddress, "上海") {
		return "上海"
	} else if strings.Contains(identityAddress, "天津") {
		return "天津"
	} else if strings.Contains(identityAddress, "重庆") {
		return "重庆"
	} else if strings.Contains(identityAddress, "省") {
		identityAddressRunes := []rune(identityAddress)
		endIndex := strings.Index(identityAddress, "省")
		prefix := []byte(identityAddress)[0:endIndex]
		result := []rune(string(prefix))
		identityAddressSubstring := string(identityAddressRunes[0:len(result)])
		return string(identityAddressSubstring)
	}
	return identityAddress
}

func Export() (*xlsx.File, error) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, err
	}
	tableArray, err := TableOneList()
	if err != nil {
		return nil, err
	}
	row := sheet.AddRow()
	row.SetHeightCM(0.65) //设置每行的高度
	cell := addCell(row)
	cell.SetString("序号")
	cell = addCell(row)
	cell.SetString("姓名")
	cell = addCell(row)
	cell.SetString("性别")
	cell = addCell(row)
	cell.SetString("年龄")
	cell = addCell(row)
	cell.SetString("学历")
	cell = addCell(row)
	cell.SetString("身份证号")
	cell = addCell(row)
	cell.SetString("手机号")
	cell = addCell(row)
	cell.SetString("邮箱号")
	cell = addCell(row)
	cell.SetString("身份证地址")
	cell = addCell(row)
	cell.SetString("工作单位")
	cell = addCell(row)
	cell.SetString("单位地址")
	cell = addCell(row)
	cell.SetString("单位电话")
	cell = addCell(row)
	cell.SetString("申请原因")
	cell = addCell(row)
	cell.SetString("银行名称")
	cell = addCell(row)
	cell.SetString("收卡地址")
	cell = addCell(row)
	cell.SetString("申请时间")
	cell = addCell(row)
	cell.SetString("推荐人")
	cell = addCell(row)
	cell.SetString("推荐人手机号")
	cell = addCell(row)
	cell.SetString("备注")
	for i, table := range tableArray {
		row := sheet.AddRow()
		row.SetHeightCM(0.65)
		cell := addCell(row)
		cell.SetString(strconv.Itoa(i + 1))
		cell = addCell(row)
		cell.SetString(table.TableOne.UserName)
		cell = addCell(row)
		cell.SetString(table.TableOne.Gender)
		cell = addCell(row)
		cell.SetString(table.TableOne.Age)
		cell = addCell(row)
		cell.SetString(table.TableOne.Education)
		cell = addCell(row)
		cell.SetString(table.TableOne.Identity)
		cell = addCell(row)
		cell.SetString(table.TableOne.PhoneNumber)
		cell = addCell(row)
		cell.SetString(table.TableOne.Email)
		cell = addCell(row)
		cell.SetString(table.TableOne.IdentityAddress)
		cell = addCell(row)
		cell.SetString(table.TableOne.Work)
		cell = addCell(row)
		cell.SetString(table.TableOne.WorkAddress)
		cell = addCell(row)
		cell.SetString(table.TableOne.WorkPhone)
		cell = addCell(row)
		cell.SetString(table.TableOne.ReasonForApplying)
		cell = addCell(row)
		cell.SetString(table.TableOne.BankName)
		cell = addCell(row)
		cell.SetString(table.TableOne.ReceiveAddress)
		cell = addCell(row)
		cell.SetString(table.TableOne.ApplyTime)
		cell = addCell(row)
		cell.SetString(table.TableOne.ReferrerName)
		cell = addCell(row)
		cell.SetString(table.TableOne.ReferrerPhoneNumber)
		cell = addCell(row)
		cell.SetString(table.TableOne.Remark)
	}
	return file, nil
}

func addCell(row *xlsx.Row) *xlsx.Cell {
	cell := row.AddCell()
	style := xlsx.NewStyle()
	style.Font = *xlsx.NewFont(12, "Calibri")
	style.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	style.ApplyBorder = true
	style.ApplyFill = true
	cell.SetStyle(style)
	return cell
}
