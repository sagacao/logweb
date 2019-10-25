$(function(){
    indexModel.systemChannelHide();

    inits();

    var pauViewModel = function () {
        this.time = ko.observable(new Date(getStartDate()));
        this.preDay = function () {
            $("#deStartTime").data("datetimepicker").setDate(addDays(this.time(), -1));
            $('#ensure').trigger("click");
        }
        this.nextDay = function () {
            $("#deStartTime").data("datetimepicker").setDate(addDays(this.time(), 1));
            $('#ensure').trigger("click");
        }
    }
    var newPauModel = new pauViewModel();
    ko.applyBindings(newPauModel, $("#pau").get(0));
    
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
            url: "/statistics/dauplatdata",
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

                var xA = ["dau","roleDau","macDau"];

                if (data.info.rows) {
                    serData = heatmapDataformat(xA, yAkey, "plat", data.info.rows);
                }else{
                    console.log("pau info.rows is no data!");
                }

                $("#chart-content").highcharts({
                    chart: {
                        type: "heatmap"
                    },
                    title: {
                        text: "平台活跃" + "(" + jsonData.startTime.format("yyyy-MM-dd") + ")"
                    },
                    xAxis: {
                        title: {

                        },
                        categories: ["活跃账号", "活跃角色", "活跃设备"]
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
            }
        });
       
        
    }
});