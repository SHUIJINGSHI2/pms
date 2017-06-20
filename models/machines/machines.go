package machines

import (
	"fmt"
	"opms/models"
	"opms/utils"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Machines struct {
	Id          string `orm:"pk;column(machineid);"`
	Password    string
	Business    string
	Memory      float64
	Disk        float64
	DiskCount   int
	Cpu         int
	Roles       string
	Contact     string
	Remark      string
	Level       int
	Status      int
	Created     int64
	Changed     int64
}

func (this *Machines) TableName() string {
	return models.TableName("machines")
}

func init() {
	orm.RegisterModel(new(Machines))
}

func AddMachine(upd Machines) error {
	o := orm.NewOrm()
	machine := new(Machines)

	machine.Id = upd.Id
	machine.Password = upd.Password
	machine.Business = upd.Business
	machine.Memory = upd.Memory
	machine.Disk = upd.Disk
	machine.DiskCount = upd.DiskCount
	machine.Cpu = upd.Cpu
	machine.Contact = upd.Contact
	machine.Roles = upd.Roles
	machine.Remark = upd.Remark
	machine.Level = upd.Level
	machine.Status = upd.Status
	machine.Created = time.Now().Unix()
	machine.Changed = time.Now().Unix()
	_, err := o.Insert(machine)
	return err
}

func UpdateMachine(id string, upd Machines) error {
	var machine Machines
	o := orm.NewOrm()
	machine = Machines{Id: id}

	machine.Password = upd.Password
	machine.Business = upd.Business
	machine.Memory = upd.Memory
	machine.Disk = upd.Disk
	machine.DiskCount = upd.DiskCount
	machine.Cpu = upd.Cpu
	machine.Contact = upd.Contact
	machine.Roles = upd.Roles
	machine.Remark = upd.Remark
	machine.Level = upd.Level
	machine.Status = upd.Status
	machine.Changed = time.Now().Unix()

	var err error
	_, err = o.Update(&machine, "password", "business","memory", "disk","diskcount", "cpu", "contact","roles", "remark", "level", "status", "changed")

	return err
}

func ListMachine(condArr map[string]string, page int, offset int) (num int64, err error, ops []Machines) {
	if page < 1 {
		page = 1
	}
	if offset < 1 {
		offset, _ = beego.AppConfig.Int("pageoffset")
	}
	var machine []Machines
	start := (page - 1) * offset
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("pms_machines ").
		Where("machineid LIKE ?")
	if condArr["business"] != "" {
		qb.And("business LIKE '%" + condArr["business"] +"%'")
	}
	if condArr["roles"] != "" {
		qb.And("roles LIKE '%" + condArr["roles"] +"%'")
	}
	if condArr["remark"] != "" {
		qb.And("remark LIKE '%" + condArr["remark"] +"%'")
	}
	if condArr["disk"] != "" {
		qb.And("disk = " + condArr["disk"])
	}
	if condArr["cpu"] != "" {
		qb.And("cpu = " + condArr["cpu"])
	}
	if condArr["level"] != "" {
		qb.And("level = " + condArr["level"])
	}
	if condArr["status"] != "" {
		qb.And("status = " + condArr["status"])
	}
	qb.OrderBy("machineid").
		Asc().
		Limit(offset).
		Offset(start)

	sql := qb.String()
	o := orm.NewOrm()
	nums, err := o.Raw(sql, "%"+condArr["machineid"]+"%").QueryRows(&machine)
	return nums, err, machine

	/*
		o := orm.NewOrm()
		o.Using("default")
		qs := o.QueryTable(models.TableName("machines"))
		cond := orm.NewCondition()

		if condArr["machineid"] != "" {
			cond = cond.And("machineid", condArr["machineid"])
		}
		if condArr["business"] != "" {
			cond = cond.And("business", condArr["business"])
		}
		if condArr["disk"] != "" {
			cond = cond.And("disk", condArr["disk"])
		}
		if condArr["cpu"] != "" {
			cond = cond.And("cpu", condArr["cpu"])
		}
		if condArr["level"] != "" {
			cond = cond.And("level", condArr["level"])
		}
		if condArr["status"] != "" {
			cond = cond.And("status", condArr["status"])
		}
		qs = qs.SetCond(cond)
		if page < 1 {
			page = 1
		}
		if offset < 1 {
			offset, _ = beego.AppConfig.Int("pageoffset")
		}
		start := (page - 1) * offset
		qs = qs.OrderBy("machineid")
		var machine []Machines
		num, errs := qs.Limit(offset, start).All(&machine)
		return num, errs, machine*/
}

func CountMachine(condArr map[string]string) int64 {
	var machine []Machines
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(" *").From("pms_machines ").
		Where("machineid LIKE ?")
	if condArr["business"] != "" {
		qb.And("business LIKE '%" + condArr["business"] +"%'")
	}
	if condArr["roles"] != "" {
		qb.And("roles LIKE '%" + condArr["roles"] +"%'")
	}
	if condArr["remark"] != "" {
		qb.And("remark LIKE '%" + condArr["remark"] +"%'")
	}
	if condArr["disk"] != "" {
		qb.And("disk = " + condArr["disk"])
	}
	if condArr["cpu"] != "" {
		qb.And("cpu = " + condArr["cpu"])
	}
	if condArr["level"] != "" {
		qb.And("level = " + condArr["level"])
	}
	if condArr["status"] != "" {
		qb.And("status = " + condArr["status"])
	}
	sql := qb.String()
	o := orm.NewOrm()
	nums, _ := o.Raw(sql, "%"+condArr["machineid"]+"%").QueryRows(&machine)
	return nums
	/*o := orm.NewOrm()
	qs := o.QueryTable(models.TableName("machines"))
	cond := orm.NewCondition()

	if condArr["machineid"] != "" {
		cond = cond.And("machineid", condArr["machineid"])
	}
	if condArr["business"] != "" {
		cond = cond.And("business", condArr["business"])
	}
	if condArr["disk"] != "" {
		cond = cond.And("disk", condArr["disk"])
	}
	if condArr["cpu"] != "" {
		cond = cond.And("cpu", condArr["cpu"])
	}
	if condArr["level"] != "" {
		cond = cond.And("level", condArr["level"])
	}
	if condArr["status"] != "" {
		cond = cond.And("status", condArr["status"])
	}
	num, _ := qs.SetCond(cond).Count()
	return num*/
}

func GetMachine(id string) (Machines, error) {
	var machine Machines
	var err error

	err = utils.GetCache("GetMachine.id."+fmt.Sprintf("%d", id), &machine)
	if err != nil {
		cache_expire, _ := beego.AppConfig.Int("cache_expire")
		o := orm.NewOrm()
		machine = Machines{Id: id}
		err = o.Read(&machine)
		utils.SetCache("GetMachine.id."+fmt.Sprintf("%d", id), machine, cache_expire)
	}
	return machine, err
}

func ChangeMachineStatus(id string, status int) error {
	o := orm.NewOrm()

	machine := Machines{Id: id}
	err := o.Read(&machine, "machineid")
	if nil != err {
		return err
	} else {
		machine.Status = status
		_, err := o.Update(&machine)
		return err
	}
}

func DeleteMachine(id string) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Machines{Id: id})

	if err == nil {
		_, err = o.Raw("DELETE FROM "+models.TableName("machines")+" WHERE machineid = ?", id).Exec()
	}
	return err
}
