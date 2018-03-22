//验证规则
var validateRules = {
    isNull: function (str) {
        return(!(str == "" || typeof str != "string"));
    },
    betweenLength: function (str, _min, _max) {
        return (str.length >= _min && str.length <= _max);
    },
    isUid: function (str) {
        return new RegExp(validateRegExp.username).test(str);
    },
    isPwd: function (str) {
        return new RegExp(validateRegExp.password).test(str);
    },
    isPwd2: function (str1, str2) {
        return (str1 == str2);
    },
    isEmail: function (str) {
        return new RegExp(validateRegExp.email).test(str);
    },
    isTel: function (str) {
        return new RegExp(validateRegExp.tel).test(str);
    },
    isMobile: function (str) {
        return new RegExp(validateRegExp.mobile).test(str);
    },
    checkType: function (element) {
        return (element.attr("type") == "checkbox" || element.attr("type") == "radio" || element.attr("rel") == "select");
    },
    isChinese: function (str) {
        return new RegExp(validateRegExp.chinese).test(str);
    },
    isRealName: function (str) {
        return new RegExp(validateRegExp.realname).test(str);
    },
    isDeptname: function (str) {
        return new RegExp(validateRegExp.deptname).test(str);
    },
    isCompanyname: function (str) {
        return new RegExp(validateRegExp.companyname).test(str);
    },
    isCompanyaddr: function (str) {
        return new RegExp(validateRegExp.companyaddr).test(str);
    },
    isCompanysite: function (str) {
        return new RegExp(validateRegExp.companysite).test(str);
    }
    ,
    isMoney: function (str) {
        return new RegExp(validateRegExp.num).test(str);
    }  ,
    isQty: function (str) {
        return new RegExp(validateRegExp.intege1).test(str);
    }





};