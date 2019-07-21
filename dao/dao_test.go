package dao

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sillyhatxu/convenient-utils/envconfig"
	"github.com/sillyhatxu/convenient-utils/mysqlclient"
	"github.com/sillyhatxu/convenient-utils/uuid"
	"github.com/sillyhatxu/table-backend/config"
	"github.com/sillyhatxu/table-backend/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	configFile := "/Users/shikuanxu/go/src/github.com/sillyhatxu/table-backend/config.conf"
	envconfig.ParseConfig(configFile, func(content []byte) {
		err := toml.Unmarshal(content, &config.Conf)
		if err != nil {
			logrus.Panicf("unmarshal toml object error. %v", err)
		}
	})
	mysqlclient.InitialDBClient(config.Conf.MysqlDB.DataSource, config.Conf.MysqlDB.MaxIdleConns, config.Conf.MysqlDB.MaxOpenConns)
}

func TestInsert(t *testing.T) {
	ut := &model.UnknownTable{
		Id:        uuid.V4Upper32(),
		Identity:  "Identity",
		TableType: model.TableTypeOne,
		Status:    model.StatusEnable,
		Content:   "Content",
	}
	err := Insert(ut)
	assert.Nil(t, err)
	fmt.Println(ut.Id)
}

func TestUpdateContent(t *testing.T) {
	ut := &model.UnknownTable{
		Id:        "01DG9C4FKS81705Q6DXRQ2MFB3",
		Identity:  "Identity",
		TableType: model.TableTypeOne,
		Status:    model.StatusEnable,
		Content:   "Content-update",
	}
	err := UpdateContent(ut)
	assert.Nil(t, err)
}

func TestUpdateStatus(t *testing.T) {
	ut := &model.UnknownTable{
		Id:        "01DG9C4FKS81705Q6DXRQ2MFB3",
		Identity:  "Identity",
		TableType: model.TableTypeOne,
		Status:    model.StatusDisable,
		Content:   "Content-update",
	}
	err := UpdateStatus(ut)
	assert.Nil(t, err)
}

func TestFindById(t *testing.T) {
	ut, err := FindById("01DG9C4FKS81705Q6DXRQ2MFB3")
	assert.Nil(t, err)
	assert.EqualValues(t, ut.Id, "01DG9C4FKS81705Q6DXRQ2MFB3")
	assert.EqualValues(t, ut.Identity, "Identity")
	assert.EqualValues(t, ut.Status, model.StatusDisable)
	assert.EqualValues(t, ut.TableType, model.TableTypeOne)
	assert.EqualValues(t, ut.Content, "Content-update")
}
