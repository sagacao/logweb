$(function(){
    
    var virtualCoinViewModel = function () {
        var self = this;
        self.goldDataList       = ko.observableArray();
        self.moneyDataList      = ko.observableArray();
        self.endTimeOnly        = ko.observable('');
        self.getData            = function () {
            initsTableMoneyPage(1);
        };
        self.endTimeOnlyFn      = ko.computed(function () {
            switch (indexModel.configUrl["货币排行"].initTab) {
                case 1:
                    indexModel.dataChannelHide();
                    if (self.endTimeOnly()) {
                        initsTableGoldPage({ moneyType: 3, pageclickednumber: 1 });
                    };
                    break;
                case 2:
                    indexModel.dataChannelHide();
                    if (self.endTimeOnly()) {
                        initsTableGoldPage({ moneyType: 1, pageclickednumber: 1 });
                    };
                    break;
                case 3:
                    indexModel.dataChannelHide();
                    if (self.endTimeOnly()) {
                        initsTableGoldPage({ moneyType: 2, pageclickednumber: 1 });
                    };
                    break;
                default:
                    break;
            }

        }, self);
        self.goldRank           = function () {
            indexModel.datetimepickerObj.format = 'yyyy-mm-dd';
            $(".endTimeOnly").datetimepicker(indexModel.datetimepickerObj);
            self.endTimeOnly('');
            indexModel.dataChannelHide();
            
            indexModel.configUrl["货币排行"].initTab = 1;
        };
        self.diaRank            = function () {
            indexModel.datetimepickerObj.format = 'yyyy-mm-dd';
            $(".endTimeOnly").datetimepicker(indexModel.datetimepickerObj);
            self.endTimeOnly('');
            indexModel.dataChannelHide();
           
            indexModel.configUrl["货币排行"].initTab = 2;
        };
        self.magicDiaRank       = function () {
            indexModel.datetimepickerObj.format = 'yyyy-mm-dd';
            $(".endTimeOnly").datetimepicker(indexModel.datetimepickerObj);
            self.endTimeOnly('');
            indexModel.dataChannelHide();
           
            indexModel.configUrl["货币排行"].initTab = 3;
        };
        self.moneyRank          = function () {
            indexModel.channelHide();
           
            indexModel.configUrl["货币排行"].initTab = 4;
        };
        self.exportExcel        = function(){
            if (indexModel.configUrl["货币排行"].initTab == 4) {
                $("#rechargeTable").table2excel({
                    filename: "充值排名"
                });
            }else{
                $("#moneyRankTable").table2excel({
                    filename: "货币排名"
                });
            }
            
        }
    }

    var newVirtualCoinModel = new virtualCoinViewModel();
    ko.applyBindings(newVirtualCoinModel, $("#virtualCoinControl").get(0));

    var initsTableGoldPage = function (conf) {		        //金币、钻石、绑钻
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        var jsonData = { 
            endTime: newVirtualCoinModel.endTimeOnly(), 
            moneyType: conf.moneyType, 
            pageIndex: conf.pageclickednumber, 
            pageSize: 50 
        };

        console.log(jsonData);
        
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            url: "",
            async: true,
            data: jsonData,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(jsonData);
                console.log("moneyTable: ",data);
                layer.close(index);
                if (data.code == 200) {
                    $("#goldRankTbody").empty();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].zeusidName = zeusidsGlobal[data.info.rows[i].zeusid];
                            newVirtualCoinModel.goldDataList.push(data.info.rows[i]);
                        }
                        $("#pageClick_1").pager({
                            pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                                initsTableGoldPage({ pageclickednumber: pageclickednumber, moneyType: conf.moneyType })
                            }
                        });
                    }else{
                        console.log("consumption-virCoin data.info.rows is no data.");
                    }
                   
                }
            }
        });
    };

    var initsTableMoneyPage = function (pageIndex) {        //充值排名
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            pageIndex: pageIndex, 
            pageSize: 10 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            url: "",
            async: true,
            data: jsonData,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(jsonData);
                console.log("rechargeTable: ",data);
                layer.close(index);
                if (data.code == 200) {
                    $("#moneyRankTbody").empty();
                    for (var i = 0; i < data.info.rows.length; i++) {
                        newVirtualCoinModel.moneyDataList.push(data.info.rows[i]);
                    };
                    $("#pageClick_money").pager({ pagenumber: pageIndex, pagecount: data.info.totalPage, buttonClickCallback: initsTableMoneyPage });
                }
            }
        });
    };
    
    var currentTab = indexModel.configUrl["货币排行"].initTab;
    switch (currentTab) {
        case 1:
            $("#goldRank").addClass("active").siblings().removeClass("active");
            $("#goldRankWrap").addClass("in active").siblings().removeClass("in active");
            newVirtualCoinModel.goldRank();
            break;
        case 2:
            $("#diaRank").addClass("active").siblings().removeClass("active");
            $("#goldRankWrap").addClass("in active").siblings().removeClass("in active");
            newVirtualCoinModel.diaRank();
            break;
        case 3:
            $("#magicDiaRank").addClass("active").siblings().removeClass("active");
            $("#goldRankWrap").addClass("in active").siblings().removeClass("in active");
            newVirtualCoinModel.magicDiaRank();
            break;
        case 4:
            $("#moneyRank").addClass("active").siblings().removeClass("active");
            $("#moneyRankWrap").addClass("in active").siblings().removeClass("in active");
            newVirtualCoinModel.moneyRank();
            break;
        default:
            break;
    }
});