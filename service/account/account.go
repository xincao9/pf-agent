package account

import (
	"errors"
	"github.com/jinzhu/gorm"
	"pf-agent/internal/db"
	"pf-agent/internal/logger"
	"pf-agent/model/account"
)

var (
	A *accountService
)

func init() {
	A = new()
}

type accountService struct {
	o *gorm.DB
}

func new() *accountService {
	return &accountService{o: db.O}
}

func (as *accountService) Login(a *account.Account) error {
	logger.L.Infof("验证码登录/手机一键登录 验证")
	return nil
}

func (as *accountService) Save(a *account.Account) error {
	mobile := a.Mobile
	ou, err := as.GetAccountByMobile(mobile)
	if err != nil {
		return err
	}
	if ou == nil { // 第一次登录
		return as.o.Save(a).Error
	}
	a.Id = ou.Id
	err = as.o.Save(a).Error
	return err
}

func (as *accountService) GetAccountByMobile(mobile string) (*account.Account, error) {
	a := &account.Account{}
	err := as.o.Where("`mobile`=?", mobile).First(a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (as *accountService) GetAccountByToken(token string) (*account.Account, error) {
	a := &account.Account{}
	err := as.o.Where("`token`=?", token).First(a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (as *accountService) GetAccountById(id int64) (*account.Account, error) {
	a := &account.Account{}
	err := as.o.Where("`id`=?", id).First(a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (as *accountService) Delete(id int64) error {
	return as.o.Unscoped().Where("`id`=?", id).Delete(account.Account{}).Error
}
