$(function () {
    indexModel.systemChannelHide();

    var logLoseDataViewModel = function(){          //创建视图模型
        this.type               = ko.observableArray([{"name":"地图","id":"map_id"},{"name":"任务","id":"task_id"}]);
        this.loseType = ko.observableArray([{ "name": "1日流失", "id": "1" }, { "name": "2日流失", "id": 2 },{"name":"3日流失","id":3},{"name":"4日流失","id":4},{"name":"5日流失","id":5},{"name":"6日流失","id":6},{"name":"7日流失","id":"7"}]);
        this.selectedType       = ko.observable("map_id");
        this.selectedLosetype   = ko.observable("1");
        this.match              = function(){
            inits();
        }
    }

    var newLogLoseDataModel = new logLoseDataViewModel();
    ko.applyBindings(newLogLoseDataModel, $("#loseData").get(0));

    inits();
    
    function inits() {
        var jsonData = {
            startTime: getStartDate(),
            endTime: getStartDate(),
            zeusid: getServerId
        }
        newLogLoseDataModel.selectedType() ? $.extend(jsonData, { runofftype: newLogLoseDataModel.selectedType() }) : $.extend(jsonData,{});
        newLogLoseDataModel.selectedLosetype() ? $.extend(jsonData, { para: newLogLoseDataModel.selectedLosetype() }) : $.extend(jsonData,{});
        
        var yA = platGlobalValue;
        var yAkey = platGlobalKey;
        var serData = [];
        var xA = [];
        
        $.ajax({
            url:"/statistics/runoffdata",
            data:jsonData,
            dataType: "json",
            error:function(){
            },
            success:function(data){
                console.log(jsonData);
                console.log(data);
                if (data.info.rows) {
                    $.each(data.info.rows, function (index, obj) {
                        for (v in obj) {
                            if (v == "number") {
                                xA.push(obj[v]);
                            }
                        }
                    })
                    xA = unique(xA);            //去重
                    xA.sort(function(a, b){     //排序
                        return a-b;
                    })
                }else{
                    console.log("loseData info.rows is no data!");
                }
                callback(data);                    
            }
        })
        
        function callback(data) {
            if (data.info.rows) {
                $.each(data.info.rows, function (index, obj) {
                    var t = [];
                    var y, x, z;

                    y = yAkey.indexOf(obj["plat"])          // y轴的坐标
                    if (y == -1) {
                        y = yA.length;
                    }
                    x = xA.indexOf(obj["number"])           // x轴的坐标
                    z = obj["rolenum"] || 0                 // 值      

                    if (x != -1) {
                        t.push(x)
                        t.push(y)
                        t.push(z)
                    }else{
                        console.log("xA is no data!")
                    }
                    if (t.length > 0) {
                        serData.push(t)
                    }
                });
            }
            
            $("#chart-content").highcharts({
                chart: {
                    type: "heatmap"
                },
                title: {
                    text: "流失"
                },
                xAxis: {
                    title: {
                        text: "流失地图id"
                    },
                    categories: xA
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
                        return "<b>地图id: " + this.series.xAxis.categories[this.point.x] + "</b><br><b>人数： " + this.point.value + "</b><br><b>渠道： " + this.series.yAxis.categories[this.point.y] + "</b>";
                    }
                },
                series: [{
                    name: '流失map_id',
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
            
    }

});