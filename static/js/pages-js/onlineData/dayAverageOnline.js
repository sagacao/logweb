$(function(){
    indexModel.systemChannelHide();
    
    var logDayAverageOnlineViewModel = function(){
        this.jobType            = ko.observableArray([{ name: "所有职业", id: 0 }, { name: "剑客", id: 1 }, { name: "鬼影", id: 2 }, { name: "灵士", id: 3 }, { name: "妙手", id: 4 }]);
        this.selectedJobtype    = ko.observable();
        this.timeList           = ko.observableArray();
        this.math = function () {
            inits();
        };
    }
    var newLogDayAverageOnlineModel = new logDayAverageOnlineViewModel();
    ko.applyBindings(newLogDayAverageOnlineModel, $("#dayAverageOnline").get(0));
    
    inits();

    function inits() {
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate(),
            zeusid: getServerId, 
            prof:newLogDayAverageOnlineModel.selectedJobtype()
        }
        
        newLogDayAverageOnlineModel.timeList(getTimeList(jsonData.startTime, jsonData.endTime,1));     //获取特定时间格式的数组

        var yA = platGlobalValue;
        var yAkey = platGlobalKey;
        var xA = getTimeList(jsonData.startTime, jsonData.endTime, 1);
        var serData = [];
        
        if (yA.length > 15) {       //y轴数据过多时，使容器高度自适应
            var overHeight = 400 + (yA.length-15)*25 + "px"
            $(".chart-wrap").height(overHeight);
        }
        
        $.ajax({
            url:"/onlinedata/avgonline",
            dataType:"json",
            data:jsonData,
            error:function(){},
            success:function(data){
                console.log(jsonData);
                console.log("dayAverageOnline: ", data);
                
                if (data.info.rows) {
                    $.each(data.info.rows, function(index, obj){
                        var x, y, z;
                        var t = [];
                        
                        x = xA.indexOf(obj["date"])
                        if (x == -1) {
                            x = xA.length;
                        }
                        y = yAkey.indexOf(obj["plat"]);
                        if (y == -1) {
                            yx = yAkey.length;
                        }
                        z = obj["onlinemin"] || 0;

                        t.push(x, y ,z);
                        serData.push(t);
                    })
                    console.log(serData);
                } else {
                    console.log("dayAverageOnline info.rows is no data.");
                }
                
                $("#chart-content").highcharts({
                    chart: {
                        type: "heatmap",
                        zoomType: 'y', //y轴可放大 
                    },
                    title: {
                        text: "日平均在线时长"
                    },
                    xAxis: {
                        title: {
                            text: "时间（单位：天）"
                        },
                        categories: newLogDayAverageOnlineModel.timeList(),
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
                    tooltip: {
                        formatter: function () {
                            return "<b>时间：" + this.series.xAxis.categories[this.point.x] + "</b><br><b>在线时长：" + this.point.value + " min</b><br><b>渠道：" + this.series.yAxis.categories[this.point.y] + "</b>";
                        }
                    },
                    series: [{
                        name: "test",
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