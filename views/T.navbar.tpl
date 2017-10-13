{{define "navbar"}}
<nav class="navbar navbar-default navbar-static-top">
      <div class="container">
        <div class="navbar-header">

          <a class="navbar-brand" href="#">我的博客</a>
        </div>


          <ul class="nav navbar-nav">
            <li{{if .IsHome}} class="active"{{end}}><a href="/">首页</a></li>
            <li{{if .IsCategory}} class="active"{{end}}><a href="/category">分类</a></li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">文章 <span class="caret"></span></a>
              <ul class="dropdown-menu">
                <li><a href="#">Action</a></li>
                <li><a href="#">Another action</a></li>
                <li><a href="#">Something else here</a></li>

              </ul>
            </li>
          </ul>
          <div id="passport" class="nav">
          <form class="navbar-form navbar-right">
            <div class="form-group">
              <input type="text" placeholder="邮箱" class="form-control">
            </div>
            <div class="form-group">
              <input type="password" placeholder="密码" class="form-control">
            </div>
            <button type="submit" class="btn btn-success">登录</button>
            <button  type="button" class="btn btn-success"data-toggle="modal" data-target="#myModal">注册</a>

          </form>
        </div><!--/.navbar-collapse -->
        </div>
        </div>
    </div>
</nav>

  <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">新用户注册</h4>
        </div>
        <div class="modal-body" style="height:400px;">
        
        <div  style="position:absolute; left:20px; top:55px; font-weight:bold;">*姓名</div>
        <div  style="position:absolute; left:20px; top:105px; font-weight:bold;">*邮箱</div>
        <div  style="position:absolute; left:20px; top:160px;font-weight:bold;">*密码</div>
        <div  style="position:absolute; left:20px; top:215px;font-weight:bold;">*确认密码</div>
        <div  style="position:absolute; left:20px; top:266px; font-weight:bold;">*手机</div>
                        
                <form >
                        <input type="email" class="form-control"  placeholder="用户名" style="position:absolute; top:50px; left:100px; width:200px" >                                       
                        <input type="password" class="form-control" placeholder="邮箱" style="position:absolute; top:100px; left:100px; width:200px">
                        <input type="email" class="form-control"  placeholder="密码" style="position:absolute; top:155px; left:100px; width:200px" > 
                        <input type="email" class="form-control"  placeholder="重新输入一次密码" style="position:absolute; top:210px; left:100px; width:200px" > 
                        <input type="email" class="form-control"  placeholder="手机号码" style="position:absolute; top:260px; left:100px; width:200px" > 

                </form>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
          <button type="button" class="btn btn-primary">注册</button>
        </div>
      </div>
    </div>
  </div>

  
{{end}}