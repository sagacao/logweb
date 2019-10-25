$(function(){

    indexModel.systemChannelHide();
    
    var consumptionCoinViewModel = function () {
        this.dataList       = ko.observableArray();
        this.dataListTotal  = ko.observableArray();
        this.xCatagory      = ko.observableArray();
        this.yValue         = ko.observableArray();
        this.Total          = ko.observableArray();
        this.MoneyTotal     = ko.observableArray();
        this.coinGet        = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");        
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 1, type: 1, sort: "desc" });       //不需要分页，数据和initsChart的一致
            initsChart({ moneyType: 1, type: 1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 1;
        };
        this.coinLoss       = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 1, type: -1, sort: "desc" });
            initsChart({ moneyType: 1, type: -1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 2;
        };
        this.diaGet         = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 4, type: 1, sort: "desc" });
            initsChart({ moneyType: 4, type: 1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 3;
        };
        this.diaLoss        = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 4, type: -1, sort: "desc" });
            initsChart({ moneyType: 4, type: -1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 4;
        };
        this.magicDiaGet    = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 3, type: 1, sort: "desc" });
            initsChart({ moneyType: 3, type: 1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 5;
        };
        this.magicDiaLoss   = function () {
            // $(".order-by").find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            $("#partDataPage").show().siblings().hide();
            // initsTable({ pageclickednumber: 1, moneyType: 2, type: -1, sort: "desc" });
            initsChart({ moneyType: 3, type: -1 });
            
            indexModel.configUrl["货币产出/消耗"].initTab = 6;
        };
        this.virCoinTotal   = function () {
            initsTotalTable(1);
            $("#vircoinSummary").show().siblings().hide();
            
            indexModel.configUrl["货币产出/消耗"].initTab = 7;
        };
        this.getTotal       = function (data, event) {
            order(event.currentTarget);
        };
        this.option         = {
            tooltip: {
                trigger: 'axis'
            },
            toolbox: {
                x: "80%",
                y: "5%",
                show: true,
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
                    show: true,
                    type: 'category',
                    data: this.xCatagory()
                }
            ],
            yAxis: [
                {
                    type: 'value',
                }
            ],
            series: [
                {
                    name: '总量',
                    type: 'bar',
                    barMaxWidth: 40,
                    data: this.yValue(),
                    label: {
                        normal:{
                            show:true,
                            color:"rgb(47, 69, 84)"
                        }
                    },
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

        this.exportExcel    = function () {
            var num = indexModel.configUrl["货币产出/消耗"].initTab;
            if (num <= 6) {
                $("#consumTable").table2excel({
                    filename: "货币产出/消耗"
                });
            }else{
                $("#consumTotalTable").table2excel({
                    filename: "货币产出消耗汇总"
                });
            }
        }
    };

    var coinNewModel = new consumptionCoinViewModel();
    ko.applyBindings(coinNewModel, $("#coinConsumption").get(0));
    
    //Table
    /* var initsTable = function (conf) {           //Table的值和Chart的值一致
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            type: conf.type, 
            moneyType: conf.moneyType, 
            // sort: conf.sort, 
            // pageIndex: conf.pageclickednumber, 
            // pageSize: 10 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        $.ajax({
            type: "get",
            async: true,
            url: "/costAnalysis/moneychange",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log("Table:",data);
                if (data.code == 200) {
                    var tabDom = $("#goldCoin").empty();
                    coinNewModel.dataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            if (data.info.rows[i].num == 0) {
                                continue;
                            }
                            data.info.rows[i].reasonType = moneyChangeReason[data.info.rows[i].reason];
                            coinNewModel.dataList.push(data.info.rows[i]);
                        };

                        $("#pageClick_1").pager({
                            pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                                initsTable({ pageclickednumber: pageclickednumber, moneyType: conf.moneyType, type: conf.type, sort: conf.sort })
                            }
                        });
                    }else{
                        console.log("consumption-coin data.info.rows is no data.");
                    }
                }
            }
        });
    } */

    //Chart
    var myChart = echarts.init($("#chart-content").get(0));
    var initsChart = function (conf) {
        myChart.showLoading({ text: '正在努力的读取数据中...' });
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(),
            type: conf.type, 
            moneyType: conf.moneyType 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/costAnalysis/moneychange",
            data: jsonData,
            dataType: "json",
            error: function () { 
                myChart.hideLoading();
                console.log("error!");
            },
            success: function (data) {
                console.log(jsonData);
                console.log("consumption-coin: ", data);                
                
                if (data.code == 200) {    
                    coinNewModel.xCatagory.removeAll();
                    coinNewModel.yValue.removeAll();
                    var tabDom = $("#goldCoin").empty();
                    coinNewModel.dataList.removeAll();
                    var arr = [0];
                    
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].reason = moneyChangeReason[data.info.rows[i].reason];
                            coinNewModel.xCatagory.push(data.info.rows[i].reason);
                            coinNewModel.yValue.push(data.info.rows[i].num);
                            coinNewModel.dataList.push(data.info.rows[i]);              //Table 数据
                        }

                        //MoneyTotal
                        $("#goldCoin").find("tr").each(function (i) {
                            $(this).find("td").each(function (j) {
                                if (j > 0) {
                                    arr[j - 1] += parseInt($(this).text());
                                }
                            })
                        })
                    } else {
                        console.log("consumption-coin is no data!");
                    }
                    coinNewModel.MoneyTotal.removeAll();
                    coinNewModel.MoneyTotal(arr);

                    myChart.hideLoading();
                    myChart.setOption(coinNewModel.option);
                } 
            }
        });
    }
    //消耗汇总表格
    var initsTotalTable = function (pageclickednumber) {
        var jsonData = { 
            startTime: getStartDate(), 
            endTime: getEndDate(), 
            pageIndex: pageclickednumber, 
            pageSize: 10 
        };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/costAnalysis/moneychangelist",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log("消耗汇总表data：",data);
                
                if (data.code == 200) {
                    var tabDom = $("#goldCoin").empty();
                    coinNewModel.dataListTotal.removeAll();
                    var arr = [0, 0, 0, 0, 0, 0, 0, 0, 0];

                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            coinNewModel.dataListTotal.push(data.info.rows[i]);
                        };
                        // $("#pageClick_2").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: initsTotalTable });

                        //Total
                        $("#virCoinLoss").find("tr").each(function (i) {
                            $(this).find("td").each(function (j) {
                                if (j > 0) {
                                    arr[j - 1] += parseInt($(this).text());
                                }
                            })
                        })
                    } else {
                        console.log("TotalTable info.rows is no data!");
                    }
                    coinNewModel.Total.removeAll();
                    coinNewModel.Total(arr);
                }
                
            }
        });
    }
    
    var currentTab = indexModel.configUrl["货币产出/消耗"].initTab;      //获取当前的标签
    switch (currentTab) {
        case 1:
            $("#cion_get").addClass("active").siblings().removeClass("active");
            coinNewModel.coinGet();
            break;
        case 2:
            $("#coin_loss").addClass("active").siblings().removeClass("active");
            coinNewModel.coinLoss();
            break;
        case 3:
            $("#dia_get").addClass("active").siblings().removeClass("active");
            coinNewModel.diaGet();
            break;
        case 4:
            $("#dia_loss").addClass("active").siblings().removeClass("active");
            coinNewModel.diaLoss();
            break;
        case 5:
            $("#magic_dia_get").addClass("active").siblings().removeClass("active");
            coinNewModel.magicDiaGet();
            break;
        case 6:
            $("#magic_dia_loss").addClass("active").siblings().removeClass("active");
            coinNewModel.magicDiaLoss();
            break;
        case 7:
            $("#virCoin_all").addClass("active").siblings().removeClass("active");
            coinNewModel.virCoinTotal();
            break;
        default:
            break;
    }
    
    //升降序
    function order(domElement) {
        /* if ($(domElement).find("span").attr("class") == "icon-arrow-down") {
            $(domElement).find("span").removeClass("icon-arrow-down").addClass("icon-arrow-up");
            switch (indexModel.configUrl["货币产出/消耗"].initTab) {
                case 1:
                    initsTable({ pageclickednumber: 1, moneyType: 3, type: 1, sort: "asc" });
                    break;
                case 2:
                    initsTable({ pageclickednumber: 1, moneyType: 3, type: -1, sort: "asc" });
                    break;
                case 3:
                    initsTable({ pageclickednumber: 1, moneyType: 1, type: 1, sort: "asc" });
                    break;
                case 4:
                    initsTable({ pageclickednumber: 1, moneyType: 1, type: -1, sort: "asc" });
                    break;
                case 5:
                    initsTable({ pageclickednumber: 1, moneyType: 2, type: 1, sort: "asc" });
                    break;
                case 6:
                    initsTable({ pageclickednumber: 1, moneyType: 2, type: -1, sort: "asc" });
                    break;
            }
        } else {
            $(domElement).find("span").removeClass("icon-arrow-up").addClass("icon-arrow-down");
            switch (indexModel.configUrl["货币产出/消耗"].initTab) {
                case 1:
                    initsTable({ pageclickednumber: 1, moneyType: 3, type: 1, sort: "desc" });
                    break;
                case 2:
                    initsTable({ pageclickednumber: 1, moneyType: 3, type: -1, sort: "desc" });
                    break;
                case 3:
                    initsTable({ pageclickednumber: 1, moneyType: 1, type: 1, sort: "desc" });
                    break;
                case 4:
                    initsTable({ pageclickednumber: 1, moneyType: 1, type: -1, sort: "desc" });
                    break;
                case 5:
                    initsTable({ pageclickednumber: 1, moneyType: 2, type: 1, sort: "desc" });
                    break;
                case 6:
                    initsTable({ pageclickednumber: 1, moneyType: 2, type: -1, sort: "desc" });
                    break;
            }
        }; */
    }
});