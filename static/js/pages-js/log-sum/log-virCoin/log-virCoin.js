$(function () {

    indexModel.systemChannelHide();
    
    var virCoinViewModel = function () {
        var self = this;
        self.DataList = ko.observableArray();
        self.virExp = function () {
            initTableTotal({ pageclickednumber: 1, type: 2 });
            indexModel.configUrl["虚拟货币日志"].initTab = 1;
        };
        self.virCon = function () {
            initTableTotal({ pageclickednumber: 1, type: 7 });
            indexModel.configUrl["虚拟货币日志"].initTab = 2;
        };
        self.virSkillexp = function () {
            initTableTotal({ pageclickednumber: 1, type: 9 });
            indexModel.configUrl["虚拟货币日志"].initTab = 3;
        };
        self.virCamp = function () {
            initTableTotal({ pageclickednumber: 1, type: 13 });
            indexModel.configUrl["虚拟货币日志"].initTab = 4;
        };
        self.virMush1 = function () {
            initTableTotal({ pageclickednumber: 1, type: 27 });
            indexModel.configUrl["虚拟货币日志"].initTab = 5;
        };
        self.virMush2 = function () {
            initTableTotal({ pageclickednumber: 1, type: 28 });
            indexModel.configUrl["虚拟货币日志"].initTab = 6;
        };
        self.virMush3 = function () {
            initTableTotal({ pageclickednumber: 1, type: 29 });
            indexModel.configUrl["虚拟货币日志"].initTab = 7;
        };

    }

    var newVirCoinModel = new virCoinViewModel();
    ko.applyBindings(newVirCoinModel, $("#virCoin").get(0));

    var initTableTotal = function (conf) {
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });

        var jsonDataTotal = {
            zeusid: getServerId,
            startTime: getStartDate(),
            endTime: getEndDate(),
            pageSize: 50
        }

        jsonDataTotal.pageIndex = conf.pageclickednumber;
        jsonDataTotal.type = conf.type;
        getServerId ? $.extend(jsonDataTotal, { zeusid: getServerId }) : $.extend(jsonDataTotal, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/gameLog/getVirtualMoneyLogPage",
            data: jsonDataTotal,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(jsonDataTotal);
                console.log("virtualmoney:",data);
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#virCoinTotal").empty();
                    newVirCoinModel.DataList.removeAll();
                    if (data.info.rows) {
                        newVirCoinModel.DataList(data.info.rows);                     //
                        
                        $("#pager_1").pager({
                            pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                                initTableTotal({ pageclickednumber: pageclickednumber, type: conf.type });
                            }
                        });
                    } else {
                        console.log("log-virCoin data.info.rows is no data.")
                    }

                }
            }
        });
    }

    var currentTab = indexModel.configUrl["虚拟货币日志"].initTab;//获取当前的标签
    switch (currentTab) {
        case 1:
            $("#exp").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 2 });
            break;
        case 2:
            $("#con").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 7 });
            break;
        case 3:
            $("#skill-exp").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 9 });
            break;
        case 4:
            $("#camp-pre").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 13 });
            break;
        case 5:
            $("#mushroom1").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 27 });
            break;
        case 6:
            $("#mushroom2").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 28 });
            break;
        case 7:
            $("#mushroom3").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 29 });
            break;
        default:
            break;
    }

});