$(function(){

    indexModel.systemChannelHide();

    var unionActiveViewModel = function(){
        this.type               = ko.observableArray([{ name:"白图BOSS",id:"0" }]);
        this.selectedType       = ko.observable();
        this.dataList           = ko.observableArray();
        this.exportExcel        = function(){
            $("#activeTable").table2excel({
                filename: "公会活动信息"
            });
        };
        this.math               = function () {
            inits();
        }
    }
    var newUnionActiveModel = new unionActiveViewModel();
    ko.applyBindings(newUnionActiveModel, $("#unionActive").get(0));

    inits();

    function inits(){
        var jsonData = {
            startTime: getStartDate(),
            endTime: getStartDate(),
            zeusid: getServerId,
            type: newUnionActiveModel.selectedType()
        }

        $.ajax({
            url:"/static/test.json",
            dataType:"json",
            data:jsonData,
            error:function(){},
            success:function(data){
                console.log(jsonData);
                console.log(data);
                newUnionActiveModel.dataList(data.test07.data);
            }
        });
    }
    
});