$(function(){

	indexModel.systemChannelHide();

	var logLeaveUnionViewModel = function(){
		this.dataList				= ko.observableArray();
		this.roleId					= ko.observable("");
		this.guildId				= ko.observable("");
		this.getServerIdChange		= ko.computed(function(){
			if (indexModel.zeusidChange() == "" && $(".pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-union/log-leaveUnion.html"){
				var idCurrent = $('.clothLis li:first').attr("id");
				$('.clothLis li:first').parents("ul").siblings("input").attr("data-id",idCurrent);
				indexModel.zeusidChange($('.clothLis li:first').children("a").text());
				getServerId = idCurrent;
			}
		},this);
	}
	
	var newLogModel = new logLeaveUnionViewModel();
	ko.applyBindings(newLogModel,document.getElementById("leaveUnion"));
	
	var pageClick = function( pageclickednumber ){
		var index = layer.load(1, {
		  	shade: [0.3,'#666666'],
		  	content:"数据获取中......"
		});
		var jsonData = {zeusid:getServerId,startTime:getStartDate(),endTime:getEndDate(),pageIndex:pageclickednumber,pageSize:10};
		newLogModel.guildId() ? $.extend(jsonData,{guildId: newLogModel.guildId()}) : $.extend(jsonData,{});
		newLogModel.roleId() ? $.extend(jsonData,{playerId: newLogModel.roleId()}) : $.extend(jsonData,{});
		$.ajax({
			type : "get",
			async:true,
			url:"",
			data : jsonData,
			dataType : "json",
			error:function(){
				layer.close(index); 
			},
			success : function(data){
				console.log("离开公会日志：",data)
				layer.close(index); 
				if( data.code == 200 ){
					var tabDom = $("#leaveUnion_tbody").empty();
					newLogModel.dataList.removeAll();
					if (data.info.rows) {
						for (var i = 0; i < data.info.rows.length; i++) {
							newLogModel.dataList.push(data.info.rows[i]);
						}
						$("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });
					}else{
						console.log("log-leaveUnion data.info.rows is no data.");
					}
					
				}else{
					layer.alert('数据加载失败！', {
						skin: 'layui-layer-lan',//样式类名
						closeBtn: 0
					});
				}
			}
		});
	}
	pageClick(1);
	 
	var roleForm = $("#leaveUnionForm").Validform({
		btnSubmit:"#accountIdMatch",
		tiptype:function(msg,o,cssctl){
			var objtip=$("#tipShow");
			cssctl(objtip,o.type);
			objtip.text(msg);
		},
		callback: function(form){
			pageClick(1);
			return false;
		}
	});
	
})
