package machines

import (
	"fmt"
	"opms/controllers"
	. "opms/models/machines"
	. "opms/models/users"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

//我的
type ManageMachineController struct {
	controllers.BaseController
}

func (this *ManageMachineController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-manage") {
		this.Abort("401")
	}
	page, err := this.GetInt("p")
	machineid := this.GetString("machineid")
	business := this.GetString("business")
	//memory := this.GetString("memory")
	disk := this.GetString("disk")
	//diskcount := this.GetString("diskcount")
	cpu := this.GetString("cpu")
	level := this.GetString("level")
	status := this.GetString("status")
	roles := this.GetString("roles")
	remark := this.GetString("remark")

	if err != nil {
		page = 1
	}

	offset, err1 := beego.AppConfig.Int("pageoffset")
	if err1 != nil {
		offset = 15
	}

	condArr := make(map[string]string)
	condArr["machineid"] = machineid
	condArr["business"] = business
	condArr["roles"] = roles
	condArr["remark"] = remark
	condArr["disk"] = disk
	condArr["cpu"] = cpu
	condArr["level"] = level
	condArr["status"] = status

	//condArr["userid"] = fmt.Sprintf("%d", this.BaseController.UserUserId)

	countMachine := CountMachine(condArr)

	paginator := pagination.SetPaginator(this.Ctx, offset, countMachine)
	_, _, machines := ListMachine(condArr, page, offset)
	this.Data["paginator"] = paginator
	this.Data["condArr"] = condArr
	this.Data["machines"] = machines
	this.Data["countMachine"] = countMachine

	this.TplName = "machines/index.tpl"
}

type AddMachineController struct {
	controllers.BaseController
}

func (this *AddMachineController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-add") {
		this.Abort("401")
	}
	var machine Machines
	this.Data["machine"] = machine

	_, _, users := ListUserFind()
	this.Data["users"] = users

	this.TplName = "machines/form.tpl"
}
func (this *AddMachineController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	machineid := this.GetString("machineid")
	if "" == machineid {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器ID"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器密码"}
		this.ServeJSON()
		return
	}

	business := this.GetString("business")
	if "" == business {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器业务用途"}
		this.ServeJSON()
		return
	}

	disk := this.GetString("disk")
	if "" == disk {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器机器硬盘容量（单位T）"}
		this.ServeJSON()
		return
	}

	cpu := this.GetString("cpu")
	if "" == cpu {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写CPU核数"}
		this.ServeJSON()
		return
	}

	contact := this.GetString("contact")
	remark := this.GetString("remark")
	level := this.GetString("level")
	status := this.GetString("status")

	memory := this.GetString("memory")
	diskcount := this.GetString("diskcount")
	roles := this.GetString("roles")

	var machine Machines
	machine.Id = machineid
	machine.Password = password
	machine.Business = business
	machine.Memory,_ = strconv.ParseFloat(memory, 32)
	machine.Disk,_ = strconv.ParseFloat(disk, 32)
	machine.DiskCount,_ = strconv.Atoi(diskcount)
	machine.Cpu,_ = strconv.Atoi(cpu)
	machine.Contact = contact
	machine.Roles = roles
	machine.Remark = remark
	machine.Level,_ = strconv.Atoi(level)
	machine.Status,_ = strconv.Atoi(status)

	err := AddMachine(machine)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功。", "ID": fmt.Sprintf("%d", machineid)}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "添加失败"}
	}
	this.ServeJSON()
}

type AddMoreMachineController struct {
	controllers.BaseController
}

func (this *AddMoreMachineController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-add") {
		this.Abort("401")
	}
	var machine Machines
	this.Data["machine"] = machine

	_, _, users := ListUserFind()
	this.Data["users"] = users

	this.TplName = "machines/more.tpl"
}
func (this *AddMoreMachineController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-add") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}
	addText := this.GetString("addText")

	if "" == addText {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器内容"}
		this.ServeJSON()
		return
	}
	fmt.Println(addText)
	machines := strings.Split(addText, "\n")
	//ss string[] :=addText.split("\n|\r");
	var success int
	var failed int
	for b := range machines {
		tmp := strings.Split(machines[b], ",")
		var machine Machines
		machine.Id = tmp[0]
		machine.Password = tmp[1]
		machine.Business = tmp[2]
		machine.Memory,_ = strconv.ParseFloat(tmp[3], 32)
		machine.Disk, _ = strconv.ParseFloat(tmp[4], 32)
		machine.DiskCount, _ = strconv.Atoi(tmp[5])
		machine.Cpu, _ = strconv.Atoi(tmp[6])
		machine.Contact = tmp[7]
		machine.Roles = tmp[8]
		machine.Remark = tmp[9]
		machine.Level, _ = strconv.Atoi(tmp[10])
		machine.Status, _ = strconv.Atoi(tmp[11])

		err := AddMachine(machine)

		if err == nil {
			success++
		} else {
			failed++
		}
	}
	this.Data["json"] = map[string]interface{}{"code": 1, "message": "添加成功"+ strconv.Itoa(success)+"台，"+"添加失败"+strconv.Itoa(failed)+"台"}
	this.ServeJSON()
}

type EditMachineController struct {
	controllers.BaseController
}

func (this *EditMachineController) Get() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-edit") {
		this.Abort("401")
	}
	//idstrs := this.Ctx.Input.Param(":id")
	idstr := this.Ctx.Input.Param(":id")
	machine, _ := GetMachine(idstr)

	this.Data["machine"] = machine
	this.TplName = "machines/form.tpl"
}
func (this *EditMachineController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-edit") {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权设置"}
		this.ServeJSON()
		return
	}

	password := this.GetString("password")
	if "" == password {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器密码"}
		this.ServeJSON()
		return
	}

	business := this.GetString("business")
	if "" == business {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器业务用途"}
		this.ServeJSON()
		return
	}

	disk := this.GetString("disk")
	if "" == disk {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写机器机器硬盘容量（单位T）"}
		this.ServeJSON()
		return
	}

	cpu := this.GetString("cpu")
	if "" == cpu {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "请填写CPU核数"}
		this.ServeJSON()
		return
	}
	memory := this.GetString("memory")
	diskcount := this.GetString("diskcount")
	contact := this.GetString("contact")
	roles := this.GetString("roles")
	remark := this.GetString("remark")
	level := this.GetString("level")
	status := this.GetString("status")

	var machine Machines
	machine.Password = password
	machine.Business = business
	machine.Memory,_ = strconv.ParseFloat(memory, 32)
	machine.Disk,_ = strconv.ParseFloat(disk, 32)
	machine.DiskCount,_ = strconv.Atoi(diskcount)
	machine.Cpu,_ = strconv.Atoi(cpu)
	machine.Contact = contact
	machine.Roles = roles
	machine.Remark = remark
	machine.Level,_ = strconv.Atoi(level)
	machine.Status,_ = strconv.Atoi(status)

	err := UpdateMachine(this.GetString("machineid"),machine)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功", "id": fmt.Sprintf("%d", this.GetString("machineid"))}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	this.ServeJSON()
}

type AjaxMachineDeleteController struct {
	controllers.BaseController
}

func (this *AjaxMachineDeleteController) Post() {
	//权限检测
	if !strings.Contains(this.GetSession("userPermission").(string), "machine-delete") {
		this.Abort("401")
	}
	id := this.GetString("id")
	machine, _ := GetMachine(id)
	if machine.Contact != this.BaseController.UserUsername {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "无权操作"}
		this.ServeJSON()
		return
	}
	err := DeleteMachine(id)

	if err == nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "删除失败"}
	}
	this.ServeJSON()
}

