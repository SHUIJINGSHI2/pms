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
        <li class="active"> 批量添加 </li>
      </ul>
    </div>
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-lg-12">
          <section class="panel">
              <div class="alert alert-block alert-info fade in">
                <form class="form-horizontal adminex-form" id="machine-more-form">
                  <div class="form-group">
                    <label>机器列表(每行依次为机器ID;密码;业务用途;内存G;机器容量T;磁盘个数;CPU核数;联系人;角色;备注;使用年限;状态(1在用0未用2故障)example: tw07a001;Wqt5Fx8uAeUQ;tw-7a-hadoop;128;64;13;24;zhouyaoyong;NodeManager;tw07a-hadoop集群;1;1)</label>
                    <div>
                      <textarea class="form-control" rows="30" name="addText" value=""></textarea>
                    </div>
                  </div>
                  <div class="form-group">
                    <label class="col-lg-2 col-sm-2 control-label"></label>
                    <div>
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
