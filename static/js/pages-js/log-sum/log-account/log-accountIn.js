$(function(){
    
    indexModel.systemChannelHide();

    var logAccountInViewModel = function () {
        this.dataList               = ko.observableArray();
        this.accountId              = ko.observable("");
        this.getServerIdChange      = ko.computed(function () {             //加载列表项第一项
            if (indexModel.zeusidChange() == "" && $(".pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-account/log-accountIn.html") {
                var idCurrent = $('.clothLis li:first').attr("id");
                $('.clothLis li:first').parents("ul").siblings("input").attr("data-id", idCurrent);
                indexModel.zeusidChange($('.clothLis li:first').children("a").text());
                getServerId = idCurrent;
            }
        }, this);
    }

    var newLogModel = new logAccountInViewModel();
    ko.applyBindings(newLogModel, $("#accountLogin").get(0));

    var pageClick = function (pageclickednumber) {
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        var jsonData = { 
            zeusid: getServerId, 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            pageIndex: pageclickednumber, 
            pageSize: 50 
        };
        newLogModel.accountId() ? $.extend(jsonData, { userId: newLogModel.accountId() }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/gameLog/getAccountLoginLogPage",
            data: jsonData,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log("dataList:", data)
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#accountLogin_tbody").empty();
                    newLogModel.dataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            //data.info.rows[i].genderName 	 = genderType[data.info.rows[i].gender];
                            newLogModel.dataList.push(data.info.rows[i]);
                        }
                        $("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
                    }else{
                        console.log("log-accountIn data.info.rows is no data");
                    }
                   
                } else {
                    layer.alert('数据加载失败！', {
                        skin: 'layui-layer-lan',//样式类名
                        closeBtn: 0
                    });
                }
            }
        });
    }

    pageClick(1);

    $("#accountIdMatch").on("click", function() {
        pageClick(1);
    })
    
    /* var roleForm = $("#accountLoginForm").Validform({//初始化参数
        btnSubmit: "#accountIdMatch",
        tiptype: function (msg, o, cssctl) {
            var objtip = $("#tipShow");
            cssctl(objtip, o.type);
            objtip.text(msg);
        },
        callback: function (form) {//如果不是ajax方式提交表单，传入callback，这时data参数是当前表单对象
            pageClick(1);
            return false;
        }
    }); */

});