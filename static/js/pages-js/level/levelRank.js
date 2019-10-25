$(function(){
    indexModel.SystemHide();

    var logLevelDisViewModel = function () {
        this.jobType            = ko.observableArray([{ name: "所有职业", id: 0 }, { name: "剑客", id: 1 }, { name: "鬼影", id: 2 }, { name: "灵士", id: 3 }, { name: "妙手", id: 4 }]);
        this.selectedJobtype    = ko.observable();
        this.dataList           = ko.observableArray();
        this.math               = function () {
            inits();
        };
        this.exportExcel        = function(){
            $("#levelRank").table2excel({
                // 不被导出的表格行的CSS class类
                exclude: ".noExl",
                // Excel文件的名称
                filename: "等级排行"
            });
        }
    }
    var newLogLevelDisModel = new logLevelDisViewModel();
    ko.applyBindings(newLogLevelDisModel, $("#levelDis").get(0));

    inits();
    
    function inits(){
        
        var jsonData = {
            startTime:getStartDate(),
            endTime:getEndDate(),
            zeusid: getServerId,
            channelid: getChannelId,
            jobType: newLogLevelDisModel.selectedJobtype()
        }
        
        $.ajax({
            url:"",
            dataType:"json",
            data:jsonData,
            error:function () {},
            success:function(data){
                newLogLevelDisModel.dataList();
            }
        });
    }
})