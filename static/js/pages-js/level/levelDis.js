$(function(){
    indexModel.systemChannelHide();

    var logLevelDisViewModel = function(){
        this.jobType            = ko.observableArray([{name:"所有职业",id:0},{name:"剑客",id:1},{name:"鬼影",id:2},{name:"灵士",id:3},{name:"妙手",id:4}]);
        this.selectedJobtype    = ko.observable();
        this.math               = function(){
            inits();
        };
    }
    var newLogLevelDisModel = new logLevelDisViewModel();
    ko.applyBindings(newLogLevelDisModel, $("#levelDis").get(0));
    
    inits();
    
    function inits() {
        var jsonData = {
            startTime: getStartDate(),
            endTime: getEndDate(),
            zeusid: getServerId,
            prof: newLogLevelDisModel.selectedJobtype()
        }

        var xA = [];
        var serData = [];
        var arrPlat = [];
        
        $.ajax({
            url: "/level/leveldisdata",
            dataType: 'json',
            data: jsonData,
            error: function () { },
            success: function (data) {
                console.log(jsonData);
                console.log(data);
                if (data.info.rows) {
                    data.info.rows.sort(function(a, b){     //根据level排序
                        return a.level - b.level;
                    })
                    
                    $.each(data.info.rows, function (index, obj) {
                        obj.plat = platGlobal[obj.plat];
                        obj.level = "Lv"+obj.level;
                        arrPlat.push(obj.plat);
                        xA.push(obj.level);
                    })
                    arrPlat = unique(arrPlat);      //去重
                    xA      = unique(xA);
                    
                    for(i of arrPlat){
                        var ob = {};
                        ob.name = i;

                        ob.data = new Array(xA.length);
                        ob.data.fill(0)     //给数组初始值为0
                        
                        $.each(data.info.rows, function(i, j){
                            if (j.plat ==ob.name ) {

                                var _index = xA.indexOf(j.level);
                                ob.data[_index] = j.num;
                            }
                        })
                        serData.push(ob);
                    }

                    console.log("serData: ",serData);
                    
                    //配置图标
                    $("#chart-content").highcharts({
                        chart: {
                            type: "spline",
                            zoomType: 'x',      //x轴可放大
                        },
                        title: {
                            text: "详细等级分布"
                        },
                        xAxis: {
                            crosshair: true,
                            title: {
                                text: "（单位：等级）"
                            },
                            min: 0,
                            categories: xA
                        },
                        yAxis: {
                            title: {
                                text: "人数"
                            }
                        },
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
                        series: serData
                    });
                    
                } else {
                    console.log("loseData info.rows is no data!");
                }
            }
        })

    }
    
});