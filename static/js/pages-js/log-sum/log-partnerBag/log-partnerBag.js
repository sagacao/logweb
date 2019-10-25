$(function () {

    indexModel.systemChannelHide();
    
    var partnerBagViewModel = function () {
        var self = this;
        self.DataList = ko.observableArray();
        self.partnerequip = function () {
            initTableTotal({ pageclickednumber: 1, type: 1 });
            indexModel.configUrl["侠客背包日志"].initTab = 1;
        };
        self.partnerexp = function () {
            initTableTotal({ pageclickednumber: 1, type: 2 });
            indexModel.configUrl["侠客背包日志"].initTab = 2;
        };
        self.partnerstone = function () {
            initTableTotal({ pageclickednumber: 1, type: 3 });
            indexModel.configUrl["侠客背包日志"].initTab = 3;
        };
        self.partner = function () {
            initTableTotal({ pageclickednumber: 1, type: 4 });
            indexModel.configUrl["侠客背包日志"].initTab = 4;
        };
    }

    var newPartnerBagModel = new partnerBagViewModel();
    ko.applyBindings(newPartnerBagModel, $("#partnerBag").get(0));

    var jsonDataTotal = {
        startTime: getStartDate(),
        endTime: getEndDate(),
    }

    var initTableTotal = function (conf) {
        var index = layer.load(1, {
            shade: [0.3, '#666666'],
            content: "数据获取中......"
        });

        jsonDataTotal.pageIndex = conf.pageclickednumber;
        jsonDataTotal.type = conf.type;
        getServerId ? $.extend(jsonDataTotal, { zeusid: getServerId }) : $.extend(jsonDataTotal, {});
        console.log(jsonDataTotal);
        $.ajax({
            type: "get",
            async: true,
            url: "",
            data: jsonDataTotal,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log(data);
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#virCoinTotal").empty();
                    newPartnerBagModel.DataList.removeAll();
                    newPartnerBagModel.DataList(data.info.rows);                     //

                    $("#pager_1").pager({
                        pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                            initTableTotal({ pageclickednumber: pageclickednumber, type: conf.type });
                        }
                    });
                }
            }
        });
    }

    var currentTab = indexModel.configUrl["侠客背包日志"].initTab;            //获取当前的标签
    switch (currentTab) {
        case 1:
            $("#parequip").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 1 });
            break;
        case 2:
            $("#parexp").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 2 });
            break;
        case 3:
            $("#parstone").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 3 });
            break;
        case 4:
            $("#par").addClass("active").siblings().removeClass("active");
            initTableTotal({ pageclickednumber: 1, type: 4 });
            break;
        default:
            break;
    }

});