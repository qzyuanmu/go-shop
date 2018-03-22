//验证文本
var validatePrompt = {
    userName: {
        onFocus: "4-20位字符，由中文、英文、数字及“_”、“-”组成",
        succeed: "OK",
        isNull: "请输入用户名",
        error: "不可为空"
    },
    pwd: {
        onFocus: "6-16位字符，可由英文、数字及“_”、“-”组成",
        succeed: "OK",
        isNull: "请输入密码",
        error:  "密码长度只能在6-16位字符之间" 
    },
    pwd2: {
        onFocus: "请再次输入密码",
        succeed: "OK",
        isNull: "请输入密码",
        error:   "两次输入密码不一致"  
    },
    mail: {
        onFocus: "请输入常用的邮箱，将用来找回密码、接收订单通知等",
        succeed: "",
        isNull: "请输入邮箱",
        error: {
            beUsed: "该邮箱已被使用，请更换其它邮箱，或使用该邮箱<a href='http://passport.360buy.com/retrievepassword.aspx' class='flk13'>找回密码</a>",
            badFormat: "邮箱格式不正确",
            badLength: "您填写的邮箱过长，邮件地址只能在50个字符以内"
        }
    },
    authcode: {
        onFocus: "请输入图片中的字符，不区分大小写",
        succeed: "",
        isNull: "请输入验证码",
        error: "验证码错误"
    },
    protocol: {
        onFocus: "",
        succeed: "",
        isNull: "请先阅读并同意《京东商城用户协议》",
        error: ""
    },
    referrer: {
        onFocus: "如您注册并完成订单，推荐人有机会获得积分",
        succeed: "",
        isNull: "",
        error: ""
    },
    schoolinput: {
        onFocus: "您可以用简拼、全拼、中文进行校名模糊查找",
        succeed: "",
        isNull: "请填选学校名称",
        error: "请填选学校名称"
    },
    empty: {
        onFocus: "",
        succeed: "",
        isNull: "",
        error: ""
    },
    
   title: {
    onFocus: "",
    succeed: "",
    isNull: "不可为空",
    error: ""
   },
   money:
       {
           onFocus: "请输入具体金额,保留2位小数",
           succeed: "OK",
           isNull: "请输入具体金额,保留2位小数",
           error: "输入格式错误!" 
       }
    


};