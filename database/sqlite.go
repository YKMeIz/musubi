package database

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

// initialize sqlite3 database
func init() {
	orm.RegisterModel(new(Device))
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "/home/neil/Projects/device.db")
	orm.RunSyncdb("default", false, true)
}

// IsName checks if given name appears in database.
//
// It returns true if name appears in the database; returns false
// if name does not appear in the database
func IsName(nam string) bool {
	o := orm.NewOrm()
	devi := Device{Name: nam}
	err := o.Read(&devi, "Name")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

// IsName checks if given token appears in database.
//
// It returns true if token appears in the database; returns false
// if token does not appear in the database
func IsToken(toke string) bool {
	o := orm.NewOrm()
	devi := Device{Token: toke}
	err := o.Read(&devi, "Token")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

// CompareToken compares given token with the token stored in database.
//
// It returns true if both tokens are same; returns false
// if tokens are different.
func CompareToken(nam, toke string) bool {
	o := orm.NewOrm()
	devi := Device{Name: nam}
	err := o.Read(&devi, "Name")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return false
	}
	if devi.Token != toke {
		return false
	}
	return true
}

func GetName(toke string) string {
	o := orm.NewOrm()
	devi := Device{Token: toke}
	err := o.Read(&devi, "Token")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return ""
	} else {
		return devi.Name
	}
}
