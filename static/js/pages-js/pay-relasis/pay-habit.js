$(function(){

    indexModel.systemChannelHide();
    
    var payHabitViewModel = function () {
        var self                = this;
        this.priceDataList      = ko.observableArray();
        this.timeDataList       = ko.observableArray();
        this.yCatagory          = ko.observableArray();
        this.xValue             = ko.observableArray();
        this.canvassApShow      = ko.observable(true);
        this.excelApShow        = ko.observable(false);     
        this.pageApShow         = ko.observable(false);
        this.payHabitType       = ko.observable(0);

        this.timeSpread_x       = ko.observableArray();
        this.serData            = ko.observableArray();
        this.amountData         = ko.observableArray();     //
        this.payerCountData     = ko.observableArray();     //
        
        this.chartApShow        = function () { //
            this.canvassApShow(true);
            this.excelApShow(false);
            // this.pageApShow(false);
        };
        this.tableApShow        = function(){
            this.canvassApShow(false);
            this.excelApShow(true);
            // this.pageApShow(true);
            // initTableTab({pageclickednumber:1,habitType:this.payHabitType});
        }
       
        this.priceHabit         = function () {
            indexModel.configUrl["付费习惯"].initTab = 1;
            this.payHabitType = 0; 
            initChart(this.payHabitType);
        };
        this.priceGold          = function () {
            indexModel.configUrl["付费习惯"].initTab = 2;
            this.payHabitType = 1;
            initChart(this.payHabitType);
        };
        this.priceMonth         = function () {
            indexModel.configUrl["付费习惯"].initTab = 3;
            this.payHabitType = 2;
            initChart(this.payHabitType);
        };
        this.priceFund          = function () {
            indexModel.configUrl["付费习惯"].initTab = 4;
            this.payHabitType = 3;
            initChart(this.payHabitType);
        };
        this.priceWeek          = function () {
            indexModel.configUrl["付费习惯"].initTab = 5;
            this.payHabitType = 4;
            initChart(this.payHabitType);
        };
        this.priceDayLimit      = function () {
            indexModel.configUrl["付费习惯"].initTab = 6;
            this.payHabitType = 5;
            initChart(this.payHabitType);
        };
        this.priceSpecial       = function () {
            indexModel.configUrl["付费习惯"].initTab = 7;
            this.payHabitType = 6;
            initChart(this.payHabitType);
        };
        this.timeSpreadMoney    = function() {
            
        }
        this.timeSpreadAccount = function () {

        };

        //图表参数
        this.option             ={
            title:{
                text: ""
            },
            tooltip: {
                trigger: 'axis'
            },
            toolbox: {
                show: true,
                feature:{
                    dataZoom:{
                        yAxisIndex:'none'
                    },
                    dataView: { readOnly: true },
                    // magicType: { type: ['line', 'bar'] },
                    restore: {},
                    saveAsImage: {}
                }
            },
            yAxis: [
                {   
                    name:"充值角色数",
                    type: 'value',
                    boundaryGap: [0, 0.01]
                }
            ],
            xAxis: [
                {   
                    name:"充值类型",
                    type: 'category',
                    data: this.yCatagory()
                }
            ],
            series: [
                {   
                    name:"充值角色数",
                    type: 'bar',
                    data:this.xValue(),
                    itemStyle: {
                        normal: {
                            color: 'lightblue',
                            barStyle: {        // 系列级个性化折线样式
                                width: 5,
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
        this.exportExcel1       = function () {
            $("#payAnalysisTable").table2excel({
                filename: "付费分析"
            });
        }
        this.exportExcel2       = function () {
            $("#timeSpreadTable").table2excel({
                filename: "付费时间分布"
            });
        }
    };
    
    var newHabitModel = new payHabitViewModel();
    ko.applyBindings(newHabitModel, $("#payHabit").get(0));

    /* var initTableTab = function (conf) {            //获取表格数据
        jsonData.pageIndex = conf.pageclickednumber;
        jsonData.pageSize = 10;
        jsonData.habitType = conf.habitType;
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        // getChannelId ? $.extend(jsonData, { channelid: getChannelId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            async: true,
            url: "/pay/payhabitdata",
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log(jsonData)
                console.log("payhabitdata-Table:",data)
                if (data.code == 200) {
                    var tabDom = $("#payAmountTbody").empty();
                    newHabitModel.priceDataList.removeAll();
                    if (data.info.rows) {
                        newHabitModel.priceDataList(data.info.rows);
                    }else{
                        console.log("initTableTab info.rows is no data!");
                    }
                }
            }
        });
    }
     */

    var myChart = echarts.init($("#chart").get(0));
    var initChart = function (habitType){               //获取图表数据
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate()
        };
        
        myChart.showLoading({ text:"正在努力的读取数据中..."});
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        // getChannelId ? $.extend(jsonData, { channelid: getChannelId }) : $.extend(jsonData, {});
        jsonData.habitType = habitType;
        
        $.ajax({
            type:"get",
            async:true,
            url:"/pay/payhabitdata",
            data:jsonData,
            dataType:"json",
            error:function(){
                
            },
            success:function(data){
                console.log(jsonData);
                console.log("payAnalysis-chart:",data);
                if (data.code == 200) {
                    newHabitModel.yCatagory.removeAll();
                    newHabitModel.xValue.removeAll();
                    newHabitModel.priceDataList.removeAll();

                    if (data.info.rows) {
                        newHabitModel.priceDataList(data.info.rows);        //添加表格数据
                        
                        for (var i = 0; i < data.info.rows.length; i++) {   //添加图表数据
                            if (rechargeGlobal[data.info.rows[i].chargeId] !== undefined) {     
                                data.info.rows[i].chargeId = rechargeGlobal[data.info.rows[i].chargeId];    //根据配置表取得对照值
                            }
                            newHabitModel.yCatagory.push(data.info.rows[i].chargeId);
                            newHabitModel.xValue.push(data.info.rows[i].payerCount);
                        };
                    }else{
                        console.log("initChart info.rows is no data!")
                    }
                    myChart.hideLoading();
                    myChart.setOption(newHabitModel.option);
                }
            }
        });
    }
    
    //付费时间列表
    var timeList = getTimeList(getStartDate(), getEndDate(), 1);    //获取时间列表
    $('.habit_datetime').datetimepicker(indexModel.datetimepickerObj); //配置日历组件
    var chartData = {};
    
    var initTimeSpreadCharts = function (time, index) {
        
        var jsonData = {
            startTime: time,
            endTime: time
        };

        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});
        $.ajax({
            type: "get",
            url: "/pay/payhourdata",
            async: true,
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                
                console.log(jsonData);
                console.log("payhourdata:",data);
                if (data.code == 200) {
                    
                    newHabitModel.amountData.removeAll();
                    newHabitModel.payerCountData.removeAll();
                    newHabitModel.amountData().length = 24;
                    newHabitModel.payerCountData().length = 24;
                    newHabitModel.amountData().fill(0);
                    newHabitModel.payerCountData().fill(0);

                    if (data.info.rows) {
                        for(var j = 0; j < data.info.rows.length; j++){
                            var _index = parseInt(data.info.rows[j].hour);
                            newHabitModel.amountData()[_index] = data.info.rows[j].amount;
                            newHabitModel.payerCountData()[_index] = data.info.rows[j].payerCount;
                        }
                    }

                    chart1.series[index].update({
                        name: time,
                        data: newHabitModel.amountData()
                    })
                    chart2.series[index].update({
                        name: time,
                        data: newHabitModel.payerCountData()
                    })

                }
            }
        });
    }
    
    var options = {
        chart: {
            type: "spline",
            zoomType: 'x',                      //x轴可放大 
        },
        title: {
            text: ""
        },
        xAxis: {
            title: {
                text: "时间点"
            },
            crosshair: true,
            tickInterval: 1                     // 刻度间隔为1
        },
        yAxis: {
            title: {
                text: "数值"
            }
        },
        tooltip: {
            shared: true                        // 开启提示框共享
        },
        plotOptions: {
            spline: {
                dataLabels: {
                    enabled: true               // 开启数据标签
                }
            }
        },
        series: newHabitModel.serData()
    };
    var chart1 = Highcharts.chart("chart-content1", options);
    var chart2 = Highcharts.chart("chart-content2", options);
    
    var arr = [];
    arr.length = 24;
    arr.fill(0);
    
    $(".habit_datetime").each(function (i) {
        chart1.addSeries({
            name: timeList[i % timeList.length],
            data: arr
        })
        chart2.addSeries({
            name: timeList[i % timeList.length],
            data: arr
        })
        
        $(this).data("datetimepicker").setDate(new Date(timeList[i % timeList.length]));      //默认加载的时间
        initTimeSpreadCharts(timeList[i % timeList.length], i)

        $(this).change(function () {
            initTimeSpreadCharts($(this).val(), i)
        })
    })
    
    //--------------------
    var currentTab = indexModel.configUrl["付费习惯"].initTab;//获取当前的标签
    switch (currentTab) {
        case 1:
            $("#habit_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceHabit();
            break;
        case 2:
            $("#gold_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceGold()
            break;
        case 3:
            $("#month_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceMonth()
            break;
        case 4:
            $("#fund_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceFund()
            break;
        case 5:
            $("#week_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceWeek()
            break;
        case 6:
            $("#daylimit_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceDayLimit()
            break;
        case 7:
            $("#special_get").addClass("active").siblings().removeClass("active");
            newHabitModel.priceSpecial()
            break;
        default:
            break;
    }

    
});