$(function(){
    indexModel.SystemHide();
   
    var finishTaskViewModel = function(){
        this.taskType           = ko.observableArray([{ name:"主线任务",id:"0" },{ name:"战斗任务",id:"1" }]);
        this.selectedTaskType   = ko.observable();
        this.math               = function() {
            inits();
        }
    }
    var newFinishTaskModel = new finishTaskViewModel();
    ko.applyBindings(newFinishTaskModel, $("#finshTask").get(0));

    inits();

    function inits() {
        var jsonData = {
            startTime:getStartDate(),
            endTime:getEndDate(),
            zeusid: getServerId,
            channelid: getChannelId,
            taskType:newFinishTaskModel.selectedTaskType()
        }

        $.ajax({
            url:"/static/test.json",
            dataType:"json",
            data:jsonData,
            error:function(){},
            success:function(data) {
                $("#chart-content").highcharts({
                    chart:{
                        type:"column",
                        zoomType:"x",
                    },
                    title:{
                        text:"任务"
                    },
                    xAxis:{
                        crosshair: true
                    },
                    yAxis:{
                        title:{
                            text:"任务分布"
                        }
                    },
                    series:[{
                        name:"任务分布",
                        data:data.test02.data,
                        color:"#439fae",
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