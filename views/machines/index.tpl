<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head>

<body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      {{template "inc/user-info.tpl" .}}
      <!--search start-->
      <form class="searchform" action="/machine/manage" method="get">
        <input type="text" class="form-control" id="machineid" name="machineid"  value="{{.condArr.machineid}}" placeholder="机器ID">
        <input type="text" class="form-control" id="business" name="business"  value="{{.condArr.business}}" placeholder="业务用途">
        <input type="text" class="form-control" id="disk" name="disk"  value="{{.condArr.disk}}" placeholder="磁盘容量(T)">
        <input type="text" class="form-control" id="cpu" name="cpu"  value="{{.condArr.cpu}}" placeholder="CPU(核数)">
        <input type="text" class="form-control" id="roles" name="roles"  value="{{.condArr.roles}}" placeholder="角色">
        <input type="text" class="form-control" id="remark" name="remark"  value="{{.condArr.remark}}" placeholder="备注">
        <select name="level" class="form-control">
          <option value="">机器年限</option>
          <option value="1" {{if eq "1" .condArr.level}}selected{{end}}>一年</option>
          <option value="2" {{if eq "2" .condArr.level}}selected{{end}}>两年</option>
          <option value="3" {{if eq "3" .condArr.level}}selected{{end}}>三年</option>
          <option value="4" {{if eq "4" .condArr.level}}selected{{end}}>四年</option>
          <option value="5" {{if eq "5" .condArr.level}}selected{{end}}>五年及以上</option>
        </select>
        <select name="status" class="form-control">
          <option value="">状态</option>
          <option value="1" {{if eq "1" .condArr.status}}selected{{end}}>在用</option>
          <option value="0" {{if eq "0" .condArr.status}}selected{{end}}>未用</option>
          <option value="2" {{if eq "2" .condArr.status}}selected{{end}}>故障</option>
        </select>
        <button type="submit" class="btn btn-primary">搜索</button>
      </form>
      <!--search end-->
      </div>
    <!-- header section end-->
    <!-- page heading start-->
    <div class="page-heading">
      <h4> 机器管理</h4>
      <ul class="breadcrumb pull-left">
        <li> <a href="/user/show/{{.LoginUserid}}">AMS</a> </li>
        <li> <a href="/machine/manage">机器管理</a> </li>
        <li class="active"> 机器列表</li>
      </ul>
      <div class="pull-right"> <a href="/machine/manage" class="hidden-xs btn btn-default">全部</a> <a href="/machine/add" class="btn btn-success">+添加机器</a>
        <a href="/machine/addmore" class="btn btn-success">+批量添加</a></div>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading">总数：{{.countMachine}}<span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a> </span> </header>
            <div class="panel-body">
              <table class="table table-hover general-table">
                <thead>
                <tr>
                  <th>机器ID</th>
                  <th>密码</th>
                  <th>业务用途</th>
                  <th>内存(G)</th>
                  <th>容量(T)</th>
                  <th>磁盘个数</th>
                  <th>CPU</th>
                  <th>年限</th>
                  <th>状态</th>
                  <th>联系人</th>
                  <th>角色</th>
                  <th>备注</th>
                  <th class="hidden-phone hidden-xs">更新日期</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>
                {{range $k,$v := .machines}}
                <tr>
                  <td>{{$v.Id}}</td>
                  <td>****</td>
                  <td>{{$v.Business}}</td>
                  <td>{{$v.Memory}}</td>
                  <td>{{$v.Disk}}</td>
                  <td>{{$v.DiskCount}}</td>
                  <td>{{$v.Cpu}}</td>
                  <td>{{$v.Level}}</td>
                  <td> {{if eq $v.Status 1}} <span class="label label-success label-mini">在用</span> {{else if eq $v.Status 0}} <span class="label label-warning label-mini">未用</span>{{else if eq $v.Status 2}} <span class="label label-danger label-mini">故障</span> {{end}} </td>
                  <td>{{$v.Contact}}</td>
                  <td>{{$v.Roles}}</td>
                  <td>{{$v.Remark}}</td>
                  <td class="hidden-phone hidden-xs">{{getDate $v.Changed}}</td>
                  <td>
                  <div class="btn-group">
                    <button type="button" class="btn btn-primary dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"> 操作<span class="caret"></span> </button>
                    <ul class="dropdown-menu">
                      <li><a href="/machine/edit/{{$v.Id}}">编辑</a></li>
                      <li role="separator" class="divider"></li>
                      <li><a href="javascript:;" class="js-machine-delete" data-op="delete" data-id="{{$v.Id}}">删除</a></li>
                      <li role="separator" class="divider"></li>
                      <li><a href="/machine/edit/{{$v.Id}}">查看</a></li>
                    </ul>
                  </div>
                  </td>
                </tr>
                {{else}}
                <tr>
                  <td colspan="8">暂时还没有机器</td>
                </tr>
                {{end}}
                </tbody>
              </table>
              {{template "inc/page.tpl" .}} </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
</body>
</html>
