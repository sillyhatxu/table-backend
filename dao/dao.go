package dao

import (
	"github.com/sillyhatxu/convenient-utils/mysqlclient"
	"github.com/sillyhatxu/table-backend/model"
	"github.com/sirupsen/logrus"
)

const insertSQL = `
insert into unknown_table (id, identification, content, status, table_type)
values (?, ?, ?, ?, ?)
`

func Insert(ut *model.UnknownTable) error {
	logrus.Infof("Insert UnknownTable:%+v", ut)
	_, err := mysqlclient.Client.Insert(insertSQL, ut.Id, ut.Identification, ut.Content, ut.Status, ut.TableType)
	return err
}

const updateContentSQL = `
update unknown_table set content = ? where id = ?
`

func UpdateContent(ut *model.UnknownTable) error {
	logrus.Infof("UpdateContent UnknownTable:%+v", ut)
	_, err := mysqlclient.Client.Update(updateContentSQL, ut.Content, ut.Id)
	return err
}

const updateStatusSQL = `
update unknown_table set status = ? where id = ?
`

func UpdateStatus(ut *model.UnknownTable) error {
	logrus.Infof("UpdateStatus UnknownTable:%+v", ut)
	_, err := mysqlclient.Client.Update(updateStatusSQL, ut.Status, ut.Id)
	return err
}

const clearAllSQL = `
update unknown_table set status = ?
`

func ClearAll(status int) error {
	logrus.Infof("ClearAll")
	_, err := mysqlclient.Client.Update(clearAllSQL, status)
	return err
}

const findByIdSQL = `
select * from unknown_table where id = ?
`

func FindById(id string) (*model.UnknownTable, error) {
	logrus.Infof("FindById:%v", id)
	var ut model.UnknownTable
	err := mysqlclient.Client.FindOneRecord(findByIdSQL, &ut, id)
	if err != nil {
		return nil, err
	}
	return &ut, nil
}

const findByTableTypeSQL = `
select * from unknown_table where table_type = ? and status = 1 order by last_modified_time desc
`

func FindByTableType(tableType int) ([]model.UnknownTable, error) {
	logrus.Infof("FindByTableType:%v", tableType)
	var array []model.UnknownTable
	err := mysqlclient.Client.FindList(findByTableTypeSQL, &array, tableType)
	if err != nil {
		return nil, err
	}
	return array, nil
}

const findByIdentificationAndTableTypeSQL = `
select * from unknown_table where identification = ? and table_type = ? and status = 1 order by last_modified_time desc
`

func FindByIdentificationTableType(identification string, tableType int) ([]model.UnknownTable, error) {
	logrus.Infof("FindByIdentificationTableType; identification: %s; tableType: %d", identification, tableType)
	var array []model.UnknownTable
	err := mysqlclient.Client.FindList(findByIdentificationAndTableTypeSQL, &array, identification, tableType)
	if err != nil {
		return nil, err
	}
	return array, nil
}
