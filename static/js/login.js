$(document).ready(function () {

    var loginViewModel = function () {
        this.userName       = ko.observable("");
        this.userNameGet    = ko.computed(function () {
            if (getCookie("userName") != "") {
                return this.userName(getCookie("userName"));
            }
        }, this);
        this.password       = ko.observable("");
		/*this.passwordGet = function(){
			if( getCookie("userPass") != ""  ){
				this.password( getCookie("userPass") );
			}
		};*/
        this.passwordGet    = ko.computed(function () {
            if (getCookie("userPass") != "") {
                return this.password(getCookie("userPass"));
            }
        }, this);
        this.rememberPass   = ko.observable(true);
        this.passkey        = ko.observable();
    }
    var newModel = new loginViewModel();
    ko.applyBindings(newModel, document.getElementById("loginContainer"));

    function checkTime(i) {
        if (i < 10) { i = "0" + i }
        return i;
    }

    function getNowFormatDate() {
        var date = new Date();
        var month = date.getMonth() + 1;
        var strDate = date.getDate();
        var currentdate = date.getFullYear() + checkTime(month) + checkTime(strDate) + checkTime(date.getHours()) + checkTime(date.getMinutes());
        return currentdate;
    }

    var loginForm = $("#loginForm").Validform({
        btnSubmit: "#loginSubmit",
        tiptype: 3,
        showAllError: true,
        callback: function (form) {
            if (newModel.rememberPass() == true) {
                setCookie("userName", newModel.userName(), 1095);
                setCookie("userPass", newModel.password(), 1095);
            };
            var name = $.trim(newModel.userName());
            //var pssd = $.trim(newModel.password());
            $.ajax({
                type: "get",
                url: "",
                async: true,
                cache: false,
                data: { name: name },
                dataType: "json",
                error: function () {
                    layer.alert("获取密钥失败，请稍后再试！", {
                        skin: 'layui-layer-lan',//样式类名
                        closeBtn: 0
                    });
                },
                success: function (data) {
                    console.log(data)
                    if (data.code == 200) {
                        newModel.passkey(data.info);
                        login();
                    } else {
                        layer.alert(data.message + '!', {
                            skin: 'layui-layer-lan',//样式类名
                            closeBtn: 0
                        });
                    }
                }
            });

            return false;
        }
    });

    function login() {
        var name = $.trim(newModel.userName());
        var pssd = $.trim(newModel.password());
        $.ajax({
            type: "post",
            url: "",
            async: true,
            cache: false,
            data: { name: name, password: hex_md5(hex_md5(newModel.passkey() + hex_md5(pssd)) + getNowFormatDate()) },
            dataType: "json",
            error: function () {
                layer.alert("登录失败，请稍后再试！", {
                    skin: 'layui-layer-lan',//样式类名
                    closeBtn: 0
                });
            },
            success: function (result) {
                console.log(result)
                if (result.code == 200) {
                    window.location.href = 'index.html';
                } else {
                    layer.alert(result.message + '!', {
                        skin: 'layui-layer-lan',//样式类名
                        closeBtn: 0
                    });
                }
            }
        });
    }

    loginForm.addRule([
        {
            ele: "#username",
            datatype: "*1-15",
            nullmsg: "请输入用户名！",
            errormsg: "请输入1-15字符长度，支持汉字、字母、数字及_ !",
            sucmsg: ""
        },
        {
            ele: "#password",
            datatype: "*5-30",
            nullmsg: "请输入密码！",
            errormsg: "请输入5-30位密码，支持字母、数字及_ !",
            sucmsg: ""
        }
    ]);
})






