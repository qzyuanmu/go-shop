

  

var easyconfig ={ 
	rulesFlg : "rules",
	msgFlg: "msg",
	validateFlg :"jeValidate",
	 rules : validateRules,
	 msg :validatePrompt
};



var easyValidate = function() {
	var rulesFlg=easyconfig.rulesFlg;
	var msgFlg = easyconfig.msgFlg; 
	
	this.init = function() {
		this.CreateErrorHtml(); 
		$("." + easyconfig.validateFlg).each(function(index, dom) {
			var e = new easyDom(dom);
			e.setFocus();
		    e.setBlur();
		    

		});

	};

///提交的时候验证所传递 的全部控件
this.Check=function(doms){
	var b = true;
	$(doms).each(function(index,d) {
	        
	        var id = $(d).attr("id") + "_error";
	        var e = new easyDom(d);
			if(!e.domCheck(d))
			 b=false;
			e.setDomClass(d,b);  
	});
	return b; 
};

this.CreateErrorHtml = function() {

	$("." + easyconfig.validateFlg).each(function(index, d) {
		var e = document.createElement("span");
		var id = $(d).attr("id") + "_error";
		e.setAttribute("id", id);
		e.setAttribute("class", "jeMsg");
		$(d).after($(e));
	});
};





};

var easyDom = function(d) {

	var dom = d;
	var rules = $(dom).attr(easyconfig.rulesFlg);
	var msg = $(dom).attr(easyconfig.msgFlg);
    var errorID=$(dom).attr("id") + "_error";
    ///设置获取焦点时候方法
	this.setFocus = function() {

		switch($(dom).attr("type")) {
		default:
			$(dom).bind('focus', function(event) { 
				$("#" + errorID).html(easyconfig.msg[msg].onFocus);
			});
			break;
		};

	};

///设置移开时的方法
	this.setBlur = function() {
		switch($(dom).attr("type")) {
		default:
		     var check =this.domCheck;
		     var setClass = this.setDomClass;
			$(dom).bind('blur', function(event) {
				 
				var b =check(dom);
				setClass(dom, b);
			});

			break;
		};
	};



///对dom进行验证
	this.domCheck = function(dom) {

		switch($(dom).attr("type")) {
		default:
			var value = $(dom).val();
			return easyconfig.rules[rules](value);
			break;
		};
	};

///设置css
	this.setDomClass = function(dom,b) {
		var errorID =$(dom).attr("id") + "_error";
		if (!b) {
			$("#" + errorID).html(easyconfig.msg[msg].error).removeClass('succeed').addClass("error");

		} else {
			$("#" + errorID).html(easyconfig.msg[msg].succeed).removeClass('error').addClass('succeed');
		}
	};
};










