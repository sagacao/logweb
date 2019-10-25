$(function(){
    indexModel.systemChannelHide();

    var logLevelDisViewModel = function () {
        this.jobType = ko.observableArray([{ name: "所有职业", id: 0 }, { name: "剑客", id: 1 }, { name: "鬼影", id: 2 }, { name: "灵士", id: 3 }, { name: "妙手", id: 4 }]);
        this.selectedJobtype = ko.observable();
        this.math = function () {
            inits();
        };
    }
    var newLogLevelDisModel = new logLevelDisViewModel();
    ko.applyBindings(newLogLevelDisModel, $("#levelDis").get(0));

    inits();
    
    function inits(){
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate(),
            zeusid: getServerId,
            prof: newLogLevelDisModel.selectedJobtype()
        }
        
        var xA = [];
        var yA = platGlobalValue;
        var yAkey = platGlobalKey;
        var serData = [];
        
        $.ajax({
            url: "/level/levellosedata",
            dataType: "json",
            data: jsonData,
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log(data);
                
                if (data.info.rows) {
                    $.each(data.info.rows, function (index, obj) {
                        var t = [];
                        var y, x, z;

                        y = yAkey.indexOf(obj["plat"])      // y轴的坐标
                        if (y == -1) {
                            y = yAkey.length;
                        }
                        x = obj["level"]                     // x轴的坐标
                        z = obj["num"] || 0                  // 值      
                        
                        if (x != -1) {                        //此处无效
                            t.push(x)
                            t.push(y)
                            t.push(z)
                        } else {
                            console.log("xA is no data!")
                        }
                        if (t.length > 0) {
                            serData.push(t)
                        }
                    });
                }
                console.log(serData)

                $("#chart-content").highcharts({
                    chart: {
                        type: "heatmap",
                        zoomType: 'x',      //x轴可放大
                    },
                    title: {
                        text: "等级流失"
                    },
                    xAxis: {
                        title: {
                            text: "（单位：等级）"
                        },
                        crosshair: true,
                        min: 1,
                        categories: xA
                    },
                    yAxis: {
                        title: {
                            text: "渠道"
                        },
                        categories: yA,
                    },
                    colorAxis: {
                        min: 0,
                        minColor: '#FFFFFF',
                        maxColor: "#439fae"
                    },
                    /* plotOptions: {
                        series: {
                            dataLabels: {
                                enabled: true,
                                format: "{point.value:.2f}%"
                            }
                        }
                    }, */
                    tooltip: {
                        formatter: function () {
                            return "<b>等级：" + this.point.x + "</b><br><b>人数：" + this.point.value + "</b><br>" + "<b>渠道：" + this.series.yAxis.categories[this.point.y] + "</b>";
                        }
                    },
                    series: [{
                        name: "等级&流失率",
                        borderWidth: 1,
                        data: serData,
                        color:"#eee",
                        dataLabels: {
                            enabled: true,
                            color: "#000"
                        }
                    }]
                });
            }
        });
        
    }
    
});