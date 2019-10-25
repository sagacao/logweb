$(function(){

    indexModel.systemChannelHide();
    
    var payUserViewModel = function () {
        this.userDataList       = ko.observableArray();
        this.xCatagoryDaily     = ko.observableArray();
        this.yValueDaily        = ko.observableArray();
        this.channelDataList    = ko.observableArray();
        this.allPayUsers        = function () {
            initTableDaily(1);
            initChartDaily();
            indexModel.configUrl["付费用户"].initTab = 1;
        };
        this.channelPayUsers    = function () {
            initTableChannel(1);
            initChartChannel();
            indexModel.configUrl["付费用户"].initTab = 2;
        };
        this.option             = {
            tooltip: {
                trigger: 'axis',
                show: true
            },
            toolbox: {
                show: true,
                x:"80%",
                y:"5%",
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
                    name:"时间",
                    type: 'category',
                    data: this.xCatagoryDaily()
                }
            ],
            yAxis: [
                {
                    name:"付费账号数",
                    type: 'value'
                }
            ],
            series: [
                {
                    type: 'bar',
                    barMaxWidth: 45,
                    data: this.yValueDaily(),
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
        this.exportExcel        = function(){
            switch (indexModel.configUrl["付费用户"].initTab){
                case 1:
                    $("#allPayUsersWrap").table2excel({
                        // 不被导出的表格行的CSS class类
                        exclude: ".noExl",
                        // Excel文件的名称
                        filename: "付费用户"
                    });
                    break;
                case 2:
                    $("#channelPayUsersWrap").table2excel({
                        // 不被导出的表格行的CSS class类
                        exclude: ".noExl",
                        // Excel文件的名称
                        filename: "渠道付费用户"
                    });
                    break;
            }
        }
    };
    var newPayUserModel = new payUserViewModel();
    ko.applyBindings(newPayUserModel, document.getElementById("payUser"));

    var myChart = echarts.init(document.getElementById('chart-content'));
    
    //付费用户
    var initTableDaily = function (pageIndex) {
        var jsonData = { startTime: getStartDate(), endTime: getEndDate(), pageIndex: pageIndex, pageSize: 10 };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        //getChannelId ? $.extend(jsonData,{channelid: getChannelId}) : $.extend(jsonData,{});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paydailydata",
            data: jsonData,
            dataType: "json",
            error: function () {
               
            },
            success: function (data) {
                if (data.code == 200) {
                    console.log("付费用户:",data);
                    var tabDom = $("#payUserTbody").empty();
                    newPayUserModel.userDataList.removeAll();

                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            newPayUserModel.userDataList.push(data.info.rows[i]);
                        };
                        // $("#pager_1").pager({ pagenumber: pageIndex, pagecount: data.info.totalPage, buttonClickCallback: initTableDaily });            
                    }else{
                        console.log("付费用户. data.info.rows is no data");
                    }
                }
            }
        });
    }
    
    var initChartDaily = function () {
        myChart.showLoading({ text: '正在努力的读取数据中...' });
        var jsonData = { startTime: getStartDate(), endTime: getEndDate() };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        //getChannelId ? $.extend(jsonData,{channelid: getChannelId}) : $.extend(jsonData,{});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paydailydata",
            timeout:10000,
            data: jsonData,
            dataType: "json",
            error: function (xhr,status) { 
                if (status=="timeout") {
                    myChart.hideLoading();
                    console.log("error: 超时！");
                }
            },
            success: function (data) {
                if (data.code == 200 && data.info.rows) {
                    console.log("付费用户图表：",data);
                    newPayUserModel.xCatagoryDaily.removeAll();
                    newPayUserModel.yValueDaily.removeAll();
                    for (var i = 0; i < data.info.rows.length; i++) {
                        newPayUserModel.xCatagoryDaily.push(data.info.rows[i].dimDay);
                        newPayUserModel.yValueDaily.push(data.info.rows[i].payerCount);
                    };

                    myChart.hideLoading();
                    myChart.setOption(newPayUserModel.option);
                }else{
                    myChart.hideLoading();
                }
            }
        });
    }
    //渠道付费用户
    var initTableChannel = function (pageIndex) {
        var jsonData = { startTime: getStartDate(), endTime: getEndDate(), pageIndex: pageIndex, pageSize: 10 };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        getChannelId ? $.extend(jsonData, { channelid: getChannelId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paychanneldata",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                if (data.code == 200) {
                    console.log("渠道付费用户:",data)
                    var tabDom = $("#payChannelTbody").empty();
                    newPayUserModel.channelDataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            data.info.rows[i].channelidName = platGlobal[data.info.rows[i].channelid];
                            newPayUserModel.channelDataList.push(data.info.rows[i]);
                        };
                        // $("#pager_2").pager({ pagenumber: pageIndex, pagecount: data.info.totalPage, buttonClickCallback: initTableDaily });
                    }else{
                        console.log("渠道付费用户. data.info.rows is no data");
                    }
                    
                }
            }
        });
    }
    var initChartChannel = function () {
        myChart.showLoading({ text: '正在努力的读取数据中...' });
        var jsonData = { startTime: getStartDate(), endTime: getEndDate() };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        //getChannelId ? $.extend(jsonData,{channelid: getChannelId}) : $.extend(jsonData,{});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paychanneldata",
            data: jsonData,
            dataType: "json",
            error: function () {
                myChart.hideLoading();
             },
            success: function (data) {
                if (data.code == 200 && data.info.rows) {
                    newPayUserModel.xCatagoryDaily.removeAll();
                    newPayUserModel.yValueDaily.removeAll();
                   
                    for (var i = 0; i < data.info.rows.length; i++) {
                        data.info.rows[i].channelidName = platGlobal[data.info.rows[i].channelid];
                        newPayUserModel.xCatagoryDaily.push(data.info.rows[i].channelidName);
                        newPayUserModel.yValueDaily.push(data.info.rows[i].payerCount);
                    };

                    myChart.hideLoading();
                    myChart.setOption(newPayUserModel.option);
                }else{
                    myChart.hideLoading();
                }
            }
        });
    }
    
    
    var currentTab = indexModel.configUrl["付费用户"].initTab;//获取当前的标签
    switch (currentTab) {
        case 1:
            $("#allPayUsers").addClass("active").siblings().removeClass("active");
            $("#allPayUsersWrap").addClass("in active").siblings().removeClass("in active");
            newPayUserModel.allPayUsers();
            break;
        case 2:
            $("#channelPayUsers").addClass("active").siblings().removeClass("active");
            $("#channelPayUsersWrap").addClass("in active").siblings().removeClass("in active");
            newPayUserModel.channelPayUsers();
            break;
        default:
            break;
    }
    
});