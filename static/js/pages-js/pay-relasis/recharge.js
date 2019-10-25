$(function(){
    
    indexModel.onlyTimeShow();
    
    var rechargeViewModel = function() {        //创建视图模型
        this.dataList        = ko.observableArray();
        this.tableData       = ko.observableArray();
        this.tableTotal      = ko.observableArray();
        this.serverList1     = ko.observableArray();
        this.serverList2     = ko.observableArray();
        this.serverList3     = ko.observableArray();
        this.serverList4     = ko.observableArray();
        this.serverList5     = ko.observableArray();
        this.selectedServer1 = ko.observable();
        this.selectedServer2 = ko.observable();
        this.selectedServer3 = ko.observable();
        this.selectedServer4 = ko.observable();
        this.selectedServer5 = ko.observable();

        this.rechargeFn      = function() {
            initRecharge("/costAnalysis/recharge", 'all');

            indexModel.configUrl["充值"].initTab = 1;
        };
        this.newRechargeFn        = function() {
            initRecharge("/costAnalysis/newlyrecharge", 'new');

            indexModel.configUrl["充值"].initTab = 1;
        };

    }
    var rechargeModel = new rechargeViewModel();
    ko.applyBindings(rechargeModel, $("#recharge").get(0));
    
    var old_all_sel_server = []
    var old_new_sel_server = []
    
    // recharge
    var initRecharge = function(_url, _type){
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate(),
        }

        $.ajax({
            url: _url,
            data: jsonData,
            dataType: "json",
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log(data);
                if (data.info.rows) {

                    // 移除表格
                    $("#timeCount").empty();
                    $("#serverCount").empty();
                    
                    //
                    var yA = [];
                    var xA = [];
                    var serData = [];

                    $.each(data.info.rows, function (index, obj) {
                        xA.push(obj.time);
                        yA.push(obj.servername);
                    })
                    xA = unique(xA);    //去重
                    yA = unique(yA);
                    
                    xA.sort();          //排序
                    yA.sort(function (a, b) {
                        if (a.match(/[0-9]+/g) && b.match(/[0-9]+/g)) {
                            return a.match(/[0-9]+/g)[0] - b.match(/[0-9]+/g)[0];
                        }
                    });
                    
                    //加载服务器列表
                    rechargeModel.serverList1(yA);
                    rechargeModel.serverList2(yA);
                    rechargeModel.serverList3(yA);
                    rechargeModel.serverList4(yA);
                    rechargeModel.serverList5(yA);

                    var yA_length = yA.length;
                    var arryServer = [
                        rechargeModel.selectedServer1,
                        rechargeModel.selectedServer2,
                        rechargeModel.selectedServer3,
                        rechargeModel.selectedServer4,
                        rechargeModel.selectedServer5
                    ]

                    switch (_type) {
                        case "all":
                            if (old_all_sel_server.length > 0) {
                                for (let i = 0; i < arryServer.length; i++) {   
                                    arryServer[i](old_all_sel_server[i])
                                }
                            } else {
                                for (let i = 0; i < arryServer.length; i++) {   
                                    arryServer[i](yA[i % yA_length])            //初始默认选择前5个服务器
                                    old_all_sel_server[i] = yA[i % yA_length]
                                }
                            }
                            break;
                        case "new":
                            if (old_new_sel_server.length > 0) {
                                for (let i = 0; i < arryServer.length; i++) {
                                    arryServer[i](old_new_sel_server[i])
                                }
                            } else {
                                for (let i = 0; i < arryServer.length; i++) {   
                                    arryServer[i](yA[i % yA_length])            //初始默认选择前5个服务器
                                    old_new_sel_server[i] = yA[i % yA_length]
                                }
                            }
                            break;
                        default:
                            break;
                    }

                    //获取chart所需数据格式

                    console.log("yA:", yA)
                    console.log("xA:", xA)

                    var chartData = {};
                    for (var y of yA) {
                        var t = [];
                        t.length = xA.length - 1; // -1 是为了不显示total
                        t.fill(0)                 //填0

                        for (var s of data.info.rows) {
                            if (s.servername == y && s.time != "total") {
                                var _index = xA.indexOf(s.time);
                                t[_index] = s.cash
                            }
                        }
                        chartData[y] = t;
                    }
                    console.log("chartData", chartData);

                    //获取汇总表格数据
                    var ArrTable = [];
                    var ArrTotal = [];

                    for (i of xA) {
                        var ob = {};
                        ob.time = i;
                        ob.sum = 0;
                        ArrTable.push(ob);
                    }
                    $.each(data.info.rows, function (i, obj) {
                        var objtime = obj.time;
                        var objcash = obj.cash;

                        for (j of ArrTable) {
                            if (j.time == objtime) {
                                j.sum += objcash;
                            }
                        }

                        if (objtime == "total") {       //按服务器统计数据
                            ArrTotal.push(obj)
                        }

                        ArrTotal.sort(function (a, b) {   //排序
                            if (a.servername.match(/[0-9]+/g) && b.servername.match(/[0-9]+/g)) {
                                return a.servername.match(/[0-9]+/g)[0] - b.servername.match(/[0-9]+/g)[0];
                            }
                        })

                    })
                    console.log("ArrTable:", ArrTable);
                    console.log("ArrTotal:", ArrTotal);
                    rechargeModel.tableData(ArrTable);
                    rechargeModel.tableTotal(ArrTotal);

                } else {
                    console.log("recharge info.rows is no data.")
                }

                // chart配置
                var options = {
                    chart: {
                        type: "spline",
                        zoomType: 'x',                      //x轴可放大 
                    },
                    title: {
                        text: "充值"
                    },
                    xAxis: {
                        title: {
                            text: "时间"
                        },
                        crosshair: true,
                        categories: xA
                    },
                    yAxis: {
                        title: {
                            text: "付费数值"
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
                    series: serData
                };
                var chart = Highcharts.chart("chart-content", options);

                $(".form-group select").each(function (i) {
                    
                    if ($(this).val()) {
                        chart.addSeries({
                            name: $(this).val(),
                            data: chartData[$(this).val()]
                        })
                    }

                    $(this).change(function () {           //change时更新数据
                        
                        if (chart.series) {                // 事件多次绑定了
                            
                            switch (_type) {
                                case "all":
                                    old_all_sel_server[i] = $(this).val()
                                    break;
                                case "new":
                                    old_new_sel_server[i] = $(this).val()
                                    break;
                                default:
                                    break;
                            }
                            
                            chart.series[i].update({
                                name: $(this).val(),
                                data: chartData[$(this).val()]
                            })
                        }
                    })
                })

            }
        })
    }

    var currentTab = indexModel.configUrl["充值"].initTab;  //获取当前的标签
    switch (currentTab) {
        case 1:
            $("#rechargeTag").addClass("active").siblings().removeClass("active");
            $("#rechargeWrap").addClass("in active").siblings().removeClass("in active");
            rechargeModel.rechargeFn();
            break;
        case 2:
            $("#incomeTag").addClass("active").siblings().removeClass("active");
            $("#incomeWrap").addClass("in active").siblings().removeClass("in active");
            rechargeModel.newRechargeFn();
            break;
        default:
            break;
    }
    
})

