$(function(){
    
    indexModel.systemChannelHide();

    inits();
    
    function inits(){
        
        var jsonData = {
            startTime: getStartDate(),
            endTime: addDays(new Date(getStartDate()), 6).format("yyyy-MM-dd")
        }
        getServerId ? $.extend(jsonData, { zeusid: getServerId }) : $.extend(jsonData, {});

        $.ajax({
            url:"/statistics/remaindata",
            dataType:"json",
            data:jsonData,
            error:function() {},
            beforeSend: function (){
                spinner.spin($("#chart-content").get(0));     
            },
            success:function(data){
                console.log("jsonData:",jsonData);
                console.log(data);
                spinner.spin();

                var yA = getTimeList(jsonData.startTime, addDays(new Date(jsonData.startTime), 6), 1); 
                var xA = ["accountDnu", "roleDnu", "macDnu", "accountNdrr1", "roleNdrr1", "macNdrr1", "accountNdrr3", "roleNdrr3", "macNdrr3", "accountNdrr4", "roleNdrr4", "macNdrr4", "accountNdrr5", "roleNdrr5", "macNdrr5", "accountNdrr6", "roleNdrr6", "macNdrr6", "accountNdrr7", "roleNdrr7", "macNdrr7"];
                var serData = [];

                //spine 
                var accDnu   = [];
                var accNdrr1 = [];
                var accNdrr3 = [];
                var accNdrr4 = [];
                var accNdrr5 = [];
                var accNdrr6 = [];
                var accNdrr7 = [];

                if (data.info.rows) {
                    serData = heatmapDataformat(xA, yA, "dimDay", data.info.rows);
                    console.log(serData);

                    $.each(data.info.rows, function (i, j) {
                        accDnu.push([j.dimDay, j.accountDnu]);
                        accNdrr1.push([j.dimDay, Number(j.accountNdrr1)]);
                        accNdrr3.push([j.dimDay, Number(j.accountNdrr3)]);
                        accNdrr4.push([j.dimDay, Number(j.accountNdrr4)]);
                        accNdrr5.push([j.dimDay, Number(j.accountNdrr5)]);
                        accNdrr6.push([j.dimDay, Number(j.accountNdrr6)]);
                        accNdrr7.push([j.dimDay, Number(j.accountNdrr7)]);
                    }) 
                    
                }else{
                    console.log("keepSummary info.rows is no data!");
                }
                
                var chart = $("#chart-content").highcharts({
                    chart: {
                        type: "heatmap"
                    },
                    title: {
                        text: "存留统计"
                    },
                    xAxis: {
                        title: {
                            text: "存留统计"
                        },
                        categories: ["新增账号", "新增角色", "新增设备", "账号次留", "角色次留", "设备次留", "账号3留", "角色3留", "设备3留", "账号4留", "角色4留", "设备4留", "账号5留", "角色5留", "设备5留", "账号6留", "角色6留", "设备6留", "账号7留", "角色7留", "设备7留"]
                    },
                    yAxis: {
                        title:{
                            text:"时间"
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
                                formatter:function(){
                                    if (String(this.point.value).indexOf(".")!=-1) {
                                        return this.point.value+"%";
                                    }
                                    return this.point.value
                                }
                            }
                        }
                    },
                    series: [{
                        name: "test",
                        borderWidth: 1,
                        color:"#eee",
                        data: serData,
                        dataLabels: {
                            enabled: true,
                            color: "#000"
                        }
                    }]
                });
                
                var spline = $("#chart-spline").highcharts({
                    chart: {
                        type: "spline",
                        zoomType: 'x', //x轴可放大 
                    },
                    title:{ 
                        text:"账号存留" 
                    }, 
                    xAxis:{ 
                        crosshair: true, 
                        title:{ 
                            text:"时间" 
                        }, 
                        type: "category", 
                    }, 
                    yAxis:[{ 
                        title:{
                            text:"百分比"
                        },
                        labels: {
                            formatter: function () {
                                return this.value + '%';
                            }
                        },
                        max: 100,
                    }, {
                        title: {
                            text: "数量"
                        },
                        opposite: true        //将轴放置对面
                    }], 
                    tooltip: {
                        shared: true,         //是否启动提示框共享
                    },
                    plotOptions: {
                        spline: {
                            dataLabels: {
                                enabled: true          // 开启数据标签
                            },
                        }
                    },
                    series: [{ 
                        yAxis:1,
                        name: "新增账号", 
                        data: accDnu
                    }, { 
                        yAxis:0,
                        name: "账号次留", 
                        data: accNdrr1,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }, {
                        yAxis:0,
                        name: "账号3留",
                        data: accNdrr3,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }, { 
                        yAxis:0,
                        name: "账号4留", 
                        data: accNdrr4,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }, { 
                        yAxis:0,
                        name: "账号5留", 
                        data: accNdrr5,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }, { 
                        yAxis:0,
                        name: "账号6留", 
                        data: accNdrr6,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }, {
                        yAxis:0,
                        name: "账号7留",
                        data: accNdrr7,
                        tooltip: {
                            // valuePrefix: '',
                            valueSuffix: '%'
                        },
                    }] 
                });
                
                $(".highcharts-series-0").click();      //模拟点击事件. legend 第一个不显示
            }
        });

    }
    
});



