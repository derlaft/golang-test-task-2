package db

import (
	"configstore/models"
	"math/rand"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// use copies to make everything stable after models change
type DevelopMrRobotConfig_201712081842 struct {
	models.Model

	Host     string
	Port     int
	Database string
	User     string
	Password string
	Schema   string
}

func (DevelopMrRobotConfig_201712081842) TableName() string {
	return "develop_mr_robot_configs"
}

type TestVPNConfig_201712081848 struct {
	models.Model

	Host        string
	Port        int
	Virtualhost string
	User        string
	Password    string
}

func (TestVPNConfig_201712081848) TableName() string {
	return "test_vpn_configs"
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

var Migrations = []*gormigrate.Migration{

	// create tables (use copies to avoid shit after changing fields)
	{
		ID: "201712081842",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&DevelopMrRobotConfig_201712081842{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable(&DevelopMrRobotConfig_201712081842{}).Error
		},
	},

	{
		ID: "201712081848",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&TestVPNConfig_201712081848{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable(&TestVPNConfig_201712081848{}).Error
		},
	},

	{ // push test data
		ID: "201712081849",
		Migrate: func(tx *gorm.DB) error {

			// now push test data
			var (
				vpn   TestVPNConfig_201712081848
				robot DevelopMrRobotConfig_201712081842
			)

			for i := 0; i < 1000; i++ {

				vpn.Data = RandStringBytes(10)
				vpn.Host = RandStringBytes(4)
				vpn.Port = rand.Int() % 65535
				vpn.Virtualhost = RandStringBytes(5)
				vpn.User = RandStringBytes(5)
				vpn.Password = RandStringBytes(8)

				err := tx.Save(&vpn).Error
				if err != nil {
					return err
				}

				robot.Data = RandStringBytes(10)
				robot.Host = RandStringBytes(4)
				robot.Port = rand.Int() % 65535
				robot.Database = RandStringBytes(6)
				robot.User = RandStringBytes(5)
				robot.Password = RandStringBytes(8)
				robot.Schema = RandStringBytes(8)

				err = tx.Save(&robot).Error
				if err != nil {
					return err
				}

			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error { // delete everything
			err := tx.Delete(&TestVPNConfig_201712081848{}).Error
			if err != nil {
				return err
			}

			err = tx.Delete(&DevelopMrRobotConfig_201712081842{}).Error
			return err
		},
	},
}
