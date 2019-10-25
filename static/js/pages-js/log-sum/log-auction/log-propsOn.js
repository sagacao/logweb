$(function(){

    indexModel.systemChannelHide();

    var logPropsOnViewModel = function () {
        this.auctionTypes               = ko.observableArray(indexModel.auctionTypeArr());
        this.selectedAuctionType        = ko.observable();
        this.dataList                   = ko.observableArray();
        this.sellerId                   = ko.observable("");
        this.itemId                     = ko.observable("");
        this.getServerIdChange          = ko.computed(function () {
            if (indexModel.zeusidChange() == "" && $(".pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-auction/log-propsOn.html") {
                var idCurrent = $('.clothLis li:first').attr("id");
                $('.clothLis li:first').parents("ul").siblings("input").attr("data-id", idCurrent);
                indexModel.zeusidChange($('.clothLis li:first').children("a").text());
                getServerId = idCurrent;
            }
        }, this);
        this.changeValue                = function () {
            pageClick(1);
        };
    }

    var newLogModel = new logPropsOnViewModel();
    ko.applyBindings(newLogModel, $("#logPropsOn").get(0));

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
        newLogModel.sellerId() ? $.extend(jsonData, { sellerId: newLogModel.sellerId() }) : $.extend(jsonData, {});
        newLogModel.itemId() ? $.extend(jsonData, { itemId: newLogModel.itemId() }) : $.extend(jsonData, {});
        newLogModel.selectedAuctionType() ? $.extend(jsonData, { type: newLogModel.selectedAuctionType() }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/gameLog/getAddAuctionLogPage",
            data: jsonData,
            dataType: "json",
            error: function () {
                layer.close(index);
            },
            success: function (data) {
                console.log("道具上架：",data)
                layer.close(index);
                if (data.code == 200) {
                    var tabDom = $("#logPropsOn_tbody").empty();
                    newLogModel.dataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            //data.info.rows[i].auctionType = indexModel.auctionTypeMap()[data.info.rows[i].type];
                            newLogModel.dataList.push(data.info.rows[i]);
                        }
                        $("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
                    }else{
                        console.log("log-propsOn data.info.rows is no data.");
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

    var roleForm = $("#logPropsOnForm").Validform({
        btnSubmit: "#accountIdMatch",
        tiptype: function (msg, o, cssctl) {
            var objtip = $("#tipShow");
            cssctl(objtip, o.type);
            objtip.text(msg);
        },
        callback: function (form) {
            pageClick(1);
            return false;
        }
    });
    
});