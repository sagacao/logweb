$(function(){

    indexModel.channelHide();
    
    var consumptionMallViewModel = function () {
        this.dataList       = ko.observableArray();
        this.xCatogary      = ko.observableArray();
        this.yValue         = ko.observableArray();
        this.diaLoss        = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            initsChart(1);
            initsTable({ pageclickednumber: 1, moneyType: 1 });
            
            indexModel.configUrl["商城货币"].initTab = 1;
        };
        this.magicLoss      = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            initsChart(2);
            initsTable({ pageclickednumber: 1, moneyType: 2 });
            
            indexModel.configUrl["商城货币"].initTab = 2;
        };
        this.countOrder     = function (data, event) {
            // order(event.currentTarget, "count", "asc", "desc");
        };
        this.diaOrder       = function (data, event) {
            // order(event.currentTarget, "cost", "asc", "desc");
        };
        this.option         = {
            tooltip: {
                trigger: 'axis',
                show: true
            },
            toolbox: {
                show: true,
                x: "80%",
                y: "5%",
                feature: {
                    dataZoom: {
                        yAxisIndex: 'none'
                    },
                    dataView: { readOnly: true },
                    magicType: { type: ['line', 'bar'] },
                    restore: {},
                    saveAsImage: {}
                }
            },
            calculable: true,
            xAxis: [
                {
                    show: false,
                    type: 'category',
                    data: this.xCatogary()
                }
            ],
            yAxis: [
                {
                    type: 'value'
                }
            ],
            series: [
                {
                    name: '总数量',
                    type: 'bar',
                    barMaxWidth: 40,
                    data: this.yValue(),
                    itemStyle: {
                        normal: {
                            color: 'lightblue',
                            barStyle: {        // 系列级个性化折线样式
                                width: 100,
                                type: 'solid'
                            }
                        },
                        emphasis: {
                            color: 'lightskyblue'
                        }
                    }
                }
            ]
        };
        this.exportExcel    = function(){
            $("#consumMallTable").table2excel({
                filename: "商城货币"
            });
        }
        
    };

    var mallNewModel = new consumptionMallViewModel();
    ko.applyBindings(mallNewModel, $("#mallConsumption").get(0));
    //Table
    var initsTable = function (conf) {
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            moneyType: conf.moneyType, 
            pageIndex: conf.pageclickednumber, 
            pageSize: 10 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        if (conf.order && conf.sort) {
            jsonData.order = conf.order;
            jsonData.sort = conf.sort;
        };

        $.ajax({
            type: "get",
            async: true,
            url: "/costAnalysis/mallcostlist",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log("Table:",data)
                if (data.code == 200) {
                    var tabDom = $("#mallLossTbody").empty();
                    mallNewModel.dataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].toolType = indexModel.itemMallMap()[data.info.rows[i].itemid];
                            mallNewModel.dataList.push(data.info.rows[i]);
                        };
                        $("#pager").pager({
                            pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                                initsTable({ pageclickednumber: pageclickednumber, moneyType: conf.moneyType, order: conf.order, sort: conf.sort })
                            }
                        });
                    }else{
                        console.log("consumption-mall data.info.rows is no data.");
                    }
                    
                }
            }
        });
    }
    //Chart
    var myChart = echarts.init($("#chart-content").get(0));
    var initsChart = function (moneyType) {
        myChart.showLoading({ text: '正在努力的读取数据中...' });
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            moneyType: moneyType 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/costAnalysis/mallcost",
            data: jsonData,
            dataType: "json",
            error: function () { 
                myChart.hideLoading();
                myChart.setOption(mallNewModel.option);
            },
            success: function (data) {

                console.log(jsonData);
                console.log("chart:",data);
                
                if (data.code == 200) {
                    mallNewModel.xCatogary.removeAll();
                    mallNewModel.yValue.removeAll();
                    for (var i = 0; i < data.info.length; i++) {
                        data.info[i].itemName = indexModel.itemMallMap()[data.info.itemid];
                        mallNewModel.xCatogary.push(data.info[i].itemName ? data.info[i].itemName : data.info[i].itemid);
                        mallNewModel.yValue.push(data.info[i].count);
                    };

                    myChart.hideLoading();
                    myChart.setOption(mallNewModel.option);
                }
            }
        });
    }

    var currentTab = indexModel.configUrl["商城货币"].initTab;      //获取当前的标签
    switch (currentTab) {
        case 1:
            $("#dia_loss").addClass("active").siblings().removeClass("active");
            mallNewModel.diaLoss();
            break;
        case 2:
            $("#magic_dia_loss").addClass("active").siblings().removeClass("active");
            mallNewModel.magicLoss();
            break;
        default:
            break;
    };

   /*  //排序
    function order(domElement, orderBy, asc, desc) {
        if ($(domElement).find("span").attr("class") == "icon-arrow-down") {
            $(domElement).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
            switch (indexModel.configUrl["商城货币"].initTab) {
                case 1:
                    initsTable({ pageclickednumber: 1, moneyType: 1, order: orderBy, sort: asc });
                    break;
                case 2:
                    initsTable({ pageclickednumber: 1, moneyType: 2, order: orderBy, sort: asc });
                    break;
            }
        } else {
            $(domElement).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            switch (indexModel.configUrl["商城货币"].initTab) {
                case 1:
                    initsTable({ pageclickednumber: 1, moneyType: 1, order: orderBy, sort: desc });
                    break;
                case 2:
                    initsTable({ pageclickednumber: 1, moneyType: 2, order: orderBy, sort: desc });
                    break;
            }
        };
    } */
});