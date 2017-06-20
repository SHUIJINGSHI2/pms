<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{config "String" "globaltitle" ""}}</title>
  {{template "inc/meta.tpl" .}}
  <link href="/static/js/bootstrap-datepicker/css/datepicker-custom.css" rel="stylesheet" />
  <style>
    .form-group .fa{font-size:66px;}
  </style>
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a> {{template "inc/user-info.tpl" .}} </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h4> 机器管理</h4>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">AMS</a> </li>
        <li> <a href="/machine/manage">机器管理</a> </li>
        <li class="active"> 编辑 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
            <header class="panel-heading"> {{.title}} </header>
            <div class="panel-body">
              <div class="alert alert-block alert-info fade in">
                <form class="form-horizontal adminex-form" id="machine-form">
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>机器ID</label>
                    <div class="col-sm-10">
                      <input type="text" name="machineid" value="{{.machine.Id}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>机器密码</label>
                    <div class="col-sm-10">
                      <input type="text" name="password" value="{{.machine.Password}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>业务用途</label>
                    <div class="col-sm-10">
                      <input type="text" name="business" value="{{.machine.Business}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>内存(G)</label>
                    <div class="col-sm-10">
                      <input type="text" name="memory" value="{{.machine.Memory}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>磁盘容量(T)</label>
                    <div class="col-sm-10">
                      <input type="text" name="disk" value="{{.machine.Disk}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>磁盘个数</label>
                    <div class="col-sm-10">
                      <input type="text" name="diskcount" value="{{.machine.DiskCount}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>CPU(核数)</label>
                    <div class="col-sm-10">
                      <input type="text" name="cpu" value="{{.machine.Cpu}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>机器级别</label>
                    <div class="col-sm-10">
                      <input type="text" name="level" value="{{.machine.Level}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>状态</label>
                    <div class="col-sm-10">
                      <!--<input type="text" name="status" value="{{.machine.Status}}" class="form-control">-->
                      <select name="status" class="form-control">
                        <option value="1" {{if eq 1 .machine.Status}}selected{{end}}>在用</option>
                        <option value="0" {{if eq 0 .machine.Status}}selected{{end}}>未用</option>
                        <option value="2" {{if eq 2 .machine.Status}}selected{{end}}>故障</option>
                      </select>
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>联系人</label>
                    <div class="col-sm-10">
                      <input type="text" name="contact" value="{{.machine.Contact}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label"><span>*</span>角色</label>
                    <div class="col-sm-10">
                      <input type="text" name="roles" value="{{.machine.Roles}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-sm-2 col-sm-2 control-label">备注</label>
                    <div class="col-sm-10">
                      <input type="text" name="remark" value="{{.machine.Remark}}" class="form-control">
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-lg-2 col-sm-2 control-label"></label>
                    <div class="col-lg-10">
                      <!--<input type="hidden" name="machineid" id="machineid" value="{{.machine.Id}}">-->
                      <button type="submit" class="btn btn-primary">提交保存</button>
                    </div>
                  </div>
                </form>
              </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/user-dialog.tpl" .}}
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
<script>
  $(function(){

  })
</script>
</body>
</html>
