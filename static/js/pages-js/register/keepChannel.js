$(function(){
    indexModel.systemChannelHide();
    
    inits();
    
    function inits() {
        var jsonData = {
            startTime: getStartDate(),
            endTime: getStartDate()
        }    
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        var yA = platGlobalValue;
        var yAkey = platGlobalKey;
        var serData = [];
        
        $.ajax({
            url: "/statistics/remainplatdata",
            dataType: "json",
            data: jsonData,
            error: function () { 
                spinner.spin();
            },
            beforeSend: function () {
                spinner.spin($("#chart-content").get(0));
            },
            success: function (data) {
                console.log(jsonData);
                console.log(data);
                spinner.spin();
                var xA = ["accountDnu", "roleDnu", "macDnu", "accountNdrr1", "roleNdrr1", "macNdrr1", "accountNdrr3", "roleNdrr3", "macNdrr3", "accountNdrr4", "roleNdrr4", "macNdrr4", "accountNdrr5", "roleNdrr5", "macNdrr5", "accountNdrr6", "roleNdrr6", "macNdrr6", "accountNdrr7", "roleNdrr7", "macNdrr7"];

                if (data.info.rows) {
                    serData = heatmapDataformat(xA, yAkey, "dimDay", data.info.rows);       //此处 "dimDay" 指 "plat"
                    console.log(serData);
                }else{
                    console.log("keepChannel info.rows is no data.");
                }

                $("#chart-content").highcharts({
                    chart: {
                        type: "heatmap"
                    },
                    title: {
                        text: "渠道留存"
                    },
                    xAxis: {
                        title: {
                            text: "留存统计"
                        },
                        categories: ["新增账号", "新增角色", "新增设备", "账号次留", "角色次留", "设备次留", "账号3留", "角色3留", "设备3留", "账号4留", "角色4留", "设备4留", "账号5留", "角色5留", "设备5留", "账号6留", "角色6留", "设备6留", "账号7留", "角色7留", "设备7留"]
                    },
                    yAxis: {
                        title: {
                            text: "渠道"
                        },
                        categories: yA
                    },
                    colorAxis: {
                        min: 0,
                        minColor: '#FFFFFF',
                        maxColor: '#439fae'
                    },
                    tooltip: {
                        formatter: function () {
                            return "<b>" + this.series.xAxis.categories[this.point.x] + "</b><br><b>" + this.point.value + "</b><br>" + this.series.yAxis.categories[this.point.y] + "</b>";
                        }
                    },
                    plotOptions: {
                        series: {
                            dataLabels: {
                                enabled: true,
                                formatter: function () {
                                    if (String(this.point.value).indexOf(".") != -1) {
                                        return this.point.value + "%";
                                    }
                                    return this.point.value
                                }
                            }
                        }
                    },
                    series: [{
                        name: "test",
                        borderWidth: 1,
                        color: "#eee",
                        data: serData,
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