$(function(){

    
    var payRankViewModel = function () {
        var self = this;
        self.totalDataList          = ko.observableArray();
        self.intervalDataList       = ko.observableArray();
        self.goldLeft               = function (data, event) {
            /* if ($(event.currentTarget).find("span").hasClass("icon-arrow-down")) {
                $(event.currentTarget).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
                jsonDataTotal.order = "gold";
                jsonDataTotal.sort = "asc";
                initTableTotal(1)
            } else {
                $(event.currentTarget).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
                jsonDataTotal.order = "gold";
                jsonDataTotal.sort = "desc";
                initTableTotal(1)
            } */
        };
        self.curRank                 = function (data, event) {
            /* if ($(event.currentTarget).find("span").hasClass("icon-arrow-down")) {
                $(event.currentTarget).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
                jsonDataTotal.order = "level";
                jsonDataTotal.sort = "asc";
                initTableTotal(1)
            } else {
                $(event.currentTarget).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
                jsonDataTotal.order = "level";
                jsonDataTotal.sort = "desc";
                initTableTotal(1)
            } */
        };
        self.loginTime               = function (data, event) {
            /* if ($(event.currentTarget).find("span").hasClass("icon-arrow-down")) {
                $(event.currentTarget).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
                jsonDataTotal.order = "createTime";
                jsonDataTotal.sort = "asc";
                initTableTotal(1)
            } else {
                $(event.currentTarget).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
                jsonDataTotal.order = "createTime";
                jsonDataTotal.sort = "desc";
                initTableTotal(1)
            } */
        };
        self.lastLoginTime          = function (data, event) {
            /* if ($(event.currentTarget).find("span").hasClass("icon-arrow-down")) {
                $(event.currentTarget).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
                jsonDataTotal.order = "lastLoginTime";
                jsonDataTotal.sort = "asc";
                initTableTotal(1)
            } else {
                $(event.currentTarget).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
                jsonDataTotal.order = "lastLoginTime";
                jsonDataTotal.sort = "desc";
                initTableTotal(1)
            } */
        };
        self.payRankTotal           = function () {
            $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            initTableTotal(1);

            indexModel.configUrl["付费排行"].initTab = 1;
            indexModel.onlyZeuidShow();
        };
        self.payRankCur             = function () {
            initTableInterval(1);
            
            indexModel.configUrl["付费排行"].initTab = 2;
            indexModel.systemChannelHide();
            indexModel.onlyTime(false);
        };
        this.exportExcel            = function () {
            switch (indexModel.configUrl["付费排行"].initTab){
                case 1:
                    $("#payRankTotalWrap").table2excel({
                        // 不被导出的表格行的CSS class类
                        exclude: ".noExl",
                        // Excel文件的名称
                        filename: "累计付费排行"
                    });
                    break;
                case 2:
                    $("#payRankCurWrap").table2excel({
                        // 不被导出的表格行的CSS class类
                        exclude: ".noExl",
                        // Excel文件的名称
                        filename: "时段付费排行"
                    });
                    break;
            }

        }
    };

    var newPayRankModel = new payRankViewModel();
    ko.applyBindings(newPayRankModel, $("#payRank").get(0));

    var jsonDataTotal = {
        startTime: getStartDate(),
        endTime:   getEndDate(),
        pageSize:  10
    };

    var initTableTotal = function (pageclickednumber) {     //累计付费排行
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        jsonDataTotal.pageIndex = pageclickednumber;
        jsonDataTotal.rankType  = 0;
        getServerId ? $.extend(jsonDataTotal, { zeusid: getServerId }) : $.extend(jsonDataTotal, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/payrankdata",
            data: jsonDataTotal,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(jsonDataTotal);
                console.log("payRank:",data)
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#pay_rankTotal").empty();
                    newPayRankModel.totalDataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].zeusidName = zeusidsGlobal[data.info.rows[i].zeusid];
                            data.info.rows[i].zeusid < 500 ? data.info.rows[i].systemName = "安卓" : data.info.rows[i].systemName = "IOS";

                            newPayRankModel.totalDataList.push(data.info.rows[i]);
                            
                            $("#pager_1").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: initTableTotal });
                        };
                    }else{
                        console.log("累计付费排行. data.info.rows is no data");
                    }
                   
                }
            }
        });
    }

    var initTableInterval = function (pageclickednumber) {  //时段付费排行
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getStartDate(), 
            pageIndex: pageclickednumber, 
            pageSize: 10, 
            // rankType : 1
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/payrankdaydata",
            data: jsonData,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(jsonData)
                console.log("timePayRank:",data);
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#pay_rankTotal").empty();
                    newPayRankModel.intervalDataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].zeusidName = zeusidsGlobal[data.info.rows[i].zeusid];
                            data.info.rows[i].channelidName = platGlobal[data.info.rows[i].channelid];
                            newPayRankModel.intervalDataList.push(data.info.rows[i]);
                        };
                        $("#pager_2").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: initTableInterval });
                    }else{
                        console.log("timePayRank info.rows is no data");
                    }
                   
                }
            }
        });
    }
    
    var currentTab = indexModel.configUrl["付费排行"].initTab;  //获取当前的标签
    switch (currentTab) {
        case 1:
            $("#payRankTotal").addClass("active").siblings().removeClass("active");
            $("#payRankTotalWrap").addClass("in active").siblings().removeClass("in active");
            newPayRankModel.payRankTotal();
            break;
        case 2:
            $("#payRankCur").addClass("active").siblings().removeClass("active");
            $("#payRankCurWrap").addClass("in active").siblings().removeClass("in active");
            newPayRankModel.payRankCur();
            break;
        default:
            break;
    }
});