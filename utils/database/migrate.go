package database

import (
	products "github.com/capstone-kelompok-7/backend-disappear/module/product/domain"
	users "github.com/capstone-kelompok-7/backend-disappear/module/users/domain"
	voucher "github.com/capstone-kelompok-7/backend-disappear/module/voucher/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(users.UserModels{}, users.AddressModels{}, products.Category{}, products.Product{}, products.ProductPhotos{}, products.Review{}, voucher.VoucherModels{})
	if err != nil {
		return
	}
}
