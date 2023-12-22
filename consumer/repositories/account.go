package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&BankAccount{})
	return accountRepository{db}
}

func (r accountRepository) Save(bankAccount BankAccount) error {
	return r.db.Save(bankAccount).Error
}

func (r accountRepository) Delete(id string) error {
	return r.db.Where("id=?", id).Delete(&BankAccount{}).Error
}

func (r accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = r.db.Find(&bankAccounts).Error
	return bankAccounts, err
}

func (r accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	err = r.db.Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}
