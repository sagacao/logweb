$(function(){

    indexModel.systemChannelHide();
    
    var controlLossViewModel = function () {
        this.dataList = ko.observableArray();
        this.xCatagory = ko.observableArray();
        this.yValue3 = ko.observableArray();
        this.yValue7 = ko.observableArray();
        this.yValue14 = ko.observableArray();
        this.option = {
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
            legend: {
                data: [
                    {
                        name: '3日流失',
                        icon: 'circle',
                        color: 'white',
                        textStyle: { fontWeight: 'bold', color: 'black' }
                    },
                    {
                        name: '7日流失',
                        icon: 'circle',
                        textStyle: { fontWeight: 'bold', color: 'black' }
                    },
                    {
                        name: '14日流失',
                        icon: 'circle',
                        textStyle: { fontWeight: 'bold', color: 'black' }
                    }
                ],
                y: 20
            },
            calculable: true,
            xAxis: [
                {
                    type: 'category',
                    boundaryGap: false,
                    data: this.xCatagory()
                }
            ],
            yAxis: [
                {
                    type: 'value',
                    axisLabel: {
                        formatter: '{value} %'
                    }
                }
            ],
            series: [
                {
                    name: '3日流失',
                    type: 'line',
                    stack: '总量',
                    itemStyle: { normal: { color: "#f79d2f", areaStyle: { color: "#e4e9d7", type: 'default' } } },
                    data: this.yValue3()
                },
                {
                    name: '7日流失',
                    type: 'line',
                    stack: '总量',
                    itemStyle: { normal: { color: "#72ca68", areaStyle: { color: "#e2f3ee", type: 'default' } } },
                    data: this.yValue7()
                },
                {
                    name: '14日流失',
                    type: 'line',
                    stack: '总量',
                    itemStyle: { normal: { color: "#64c0fb", areaStyle: { color: "#f0f9fe", type: 'default' } } },
                    data: this.yValue14()
                }
            ]
        };
        this.payLoss1 = function () {
            initsChart(2);
            initsTable({ pageclickednumber: 1, type: 2 });
			/*if(getServerId){
				$(".export").attr("href","/userAnalysis/exportDailyLossList?"+"startTime="+getStartDate()+"&endTime="+getEndDate()+"&zeusid="+getServerId+"&type="+1);
			}else{
				$(".export").attr("href","/userAnalysis/exportDailyLossList?"+"startTime="+getStartDate()+"&endTime="+getEndDate()+"&type="+1);
			};*/
            indexModel.configUrl["付费流失"].initTab = 1;
        };
        this.payLoss2 = function () {
            initsChart(3);
            initsTable({ pageclickednumber: 1, type: 3 });
			/*if(getServerId){
				$(".export").attr("href","/userAnalysis/exportDailyLossList?"+"startTime="+getStartDate()+"&endTime="+getEndDate()+"&zeusid="+getServerId+"&type="+1);
			}else{
				$(".export").attr("href","/userAnalysis/exportDailyLossList?"+"startTime="+getStartDate()+"&endTime="+getEndDate()+"&type="+1);
			};*/
            indexModel.configUrl["付费流失"].initTab = 2;
        };
        this.test = function(){
            alert("test!");
        }
    };

    var lossNewModel = new controlLossViewModel();
    ko.applyBindings(lossNewModel, $("#controlLoss").get(0));
    
    var initsTable = function (conf) {
        var jsonData = { startTime: getStartDate(), endTime: getEndDate(), type: conf.type, pageIndex: conf.pageclickednumber, pageSize: 10 };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        //var pageIndex = conf.pageclickednumber;
        //var type = conf.type;
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paylossdata",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log("initsTable: ",jsonData);
                console.log("付费流失 intsTable:",data)
                if (data.code == 200) {
                    var tabDom = $("#controlLossTbody").empty();
                    lossNewModel.dataList.removeAll();
                    if (data.info.rows) {
                        for (var i = 0; i < data.info.rows.length; i++) {
                            lossNewModel.dataList.push(data.info.rows[i]);
                        };

                        $("#pager").pager({
                            pagenumber: conf.pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: function (pageclickednumber) {
                                initsTable({ pageclickednumber: pageclickednumber, type: conf.type });
                            }
                        });
                    }else{
                        console.log("data.info.rows is no data");
                    }
                   
                }
            }
        });
    }

    var myChart = echarts.init($("#chart-content").get(0));
    var initsChart = function (type) {
        myChart.showLoading({ text: '正在努力的读取数据中...' });
        var jsonData = { startTime: getStartDate(), endTime: getEndDate(), type: type };
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/paylossdata",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log("initsChart: ",jsonData);
                console.log("付费流失 initsChart:",data)
                if (data.code == 200) {
                    lossNewModel.xCatagory.removeAll();
                    lossNewModel.yValue3.removeAll();
                    lossNewModel.yValue7.removeAll();
                    lossNewModel.yValue14.removeAll();
                    for (var i = 0; i < data.info.length; i++) {
                        lossNewModel.xCatagory.push(data.info[i].dimDay);
                        lossNewModel.yValue3.push(data.info[i].dlr3);
                        lossNewModel.yValue7.push(data.info[i].dlr7);
                        lossNewModel.yValue14.push(data.info[i].dlr14);
                    };

                    myChart.hideLoading();
                    myChart.setOption(lossNewModel.option);
                }
            }
        });
    };

    var currentTab = indexModel.configUrl["付费流失"].initTab;
    switch (currentTab) {
        case 1:
            $("#pay1_loss").addClass("active").siblings().removeClass("active");
            lossNewModel.payLoss1();
            break;
        case 2:
            $("#pay2_loss").addClass("active").siblings().removeClass("active");
            lossNewModel.payLoss2();
            break;
        default:
            break;
    }
    
});