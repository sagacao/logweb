$(function(){

    indexModel.onlyTimeShow();
        
    var rechargeTotalViewModel = function() {
        this.serverList1 = ko.observableArray();
        this.serverList2 = ko.observableArray();
        this.serverList3 = ko.observableArray();
        this.serverList4 = ko.observableArray();
        this.serverList5 = ko.observableArray();
        this.selectedServer1 = ko.observable();
        this.selectedServer2 = ko.observable();
        this.selectedServer3 = ko.observable();
        this.selectedServer4 = ko.observable();
        this.selectedServer5 = ko.observable();
    }

    var rechargeTotalModel = new rechargeTotalViewModel();
    ko.applyBindings(rechargeTotalModel, $("#hour-recharge").get(0))
    
    var jsonData = {
        startTime: getStartDate(),
        endTime: getStartDate(),
    }
    console.log(jsonData);
    var timeArr = jsonData.startTime.split("-");
    
    $.ajax({
        url:"/costAnalysis/hourrecharge",
        dataType: "json",
        data: jsonData,
        error: function(){},
        success: function(data){
            console.log(data);
            if (data.info.rows) {

                var yA = [];
                var xA = []; 
                var serData = [];

                $.each(data.info.rows, function (index, obj) {
                    yA.push(obj.servername)
                })
                yA = unique(yA);        //去重

                yA.sort(function (a, b) {
                    if (a.match(/[0-9]+/g) && b.match(/[0-9]+/g)) {
                        return a.match(/[0-9]+/g)[0] - b.match(/[0-9]+/g)[0];
                    }
                });
                console.log("yA:", yA);

                // 加载服务器列表
                rechargeTotalModel.serverList1(yA);
                rechargeTotalModel.serverList2(yA);
                rechargeTotalModel.serverList3(yA);
                rechargeTotalModel.serverList4(yA);
                rechargeTotalModel.serverList5(yA);

                var yA_length = yA.length;
                var arryServer = [
                    rechargeTotalModel.selectedServer1,
                    rechargeTotalModel.selectedServer2,
                    rechargeTotalModel.selectedServer3,
                    rechargeTotalModel.selectedServer4,
                    rechargeTotalModel.selectedServer5
                ]

                for (let i = 0; i < arryServer.length; i++) {   //初始默认选择前5个服务器
                    arryServer[i](yA[i % yA_length])
                }

                // 获取chart所需数据格式
                var chartData = {};
                for (var y of yA) {
                    var t = [];
                    t.length = 25;
                    t.fill(0);

                    for(var s of data.info.rows){
                        if (s.servername == y) {
                            if (s.time != "total") {
                                var _index = parseInt(s.time);
                                t[_index] = s.cash;
                            }
                        } 
                    }
                    chartData[y] = t;
                }
                console.log("chartData:", chartData);
                
            } else {
                console.log("hour-recharge info.rows is no data.");
            }

            // chart的配置项
            var options = {
                chart: {
                    type: "spline",
                    zoomType: 'x',                      //x轴可放大 
                },
                title: {
                    text: "每小时充值"
                },
                xAxis: {
                    title: {
                        text: "时间（单位：小时）"
                    },
                    type: "datetime",
                    crosshair: true,
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
                        },
                        pointStart: Date.UTC(timeArr[0], timeArr[1] - 1, timeArr[2], 0, 0, 0),
                        pointInterval: 1000 * 60 * 60  //时间间隔 5 min
                    }
                },
                series: serData
            };
            var chart = Highcharts.chart("chart-content", options);

            // 初始加入选中服务器数据
            $(".form-group select").each(function (i) {
                if ($(this).val()) {
                    chart.addSeries({
                        name: $(this).val(),
                        data: chartData[$(this).val()]
                    })
                }

                $(this).change(function () {           //change时更新数据
                    chart.series[i].update({
                        name: $(this).val(),
                        data: chartData[$(this).val()]
                    })
                })
            })
            
        }
    })
        
})
