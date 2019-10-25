$(function(){

	indexModel.systemChannelHide();
	
	var logLoseItemsViewModel = function(){
		this.dataList				= ko.observableArray();
		this.itemId					= ko.observable("");
		this.roleId					= ko.observable("");
		this.getServerIdChange		= ko.computed(function(){
			if (indexModel.zeusidChange() == "" && $(".pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-items/log-loseItems.html"){
				var idCurrent = $('.clothLis li:first').attr("id");
				$('.clothLis li:first').parents("ul").siblings("input").attr("data-id",idCurrent);
				indexModel.zeusidChange($('.clothLis li:first').children("a").text());
				getServerId = idCurrent;
			}
		},this);

		this.saveRoleid 			= ko.observableArray();
		this.saveItemid 			= ko.observableArray();
	}
	
	var newLogModel = new logLoseItemsViewModel();
	ko.applyBindings(newLogModel,document.getElementById("loseItems"));
	
	
	var pageClick = function( pageclickednumber ){
		var index = layer.load(1, {
		  	shade: [0.3,'#666666'],
		  	content:"数据获取中......"
		});
		var jsonData = {zeusid:getServerId,startTime:getStartDate(),endTime:getEndDate(),pageIndex:pageclickednumber,pageSize:50};
		newLogModel.itemId() ? $.extend(jsonData,{itemId: newLogModel.itemId()}) : $.extend(jsonData,{});
		newLogModel.roleId() ? $.extend(jsonData,{roleId: newLogModel.roleId()}) : $.extend(jsonData,{});

		// 保存记录到sessionStorage
		if (newLogModel.itemId()) {
			if (window.sessionStorage.getItem('itemid')) {
				let sessArrItem = JSON.parse(window.sessionStorage.getItem('itemid'))
				sessArrItem.push(newLogModel.itemId())
				window.sessionStorage.setItem('itemid', JSON.stringify(sessArrItem))
			} else {
				let sessArrItem = [newLogModel.itemId()]
				window.sessionStorage.setItem('itemid', JSON.stringify(sessArrItem))
			}
			// 从sessionStorage 获取输入记录
			newLogModel.saveItemid(JSON.parse(window.sessionStorage.getItem('itemid')))
		}
		if (newLogModel.roleId()) {
			if (window.sessionStorage.getItem('roleid')) {
				let sessArrRole = JSON.parse(window.sessionStorage.getItem('roleid'))
				sessArrRole.push(newLogModel.roleId())
				window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
			} else {
				let sessArrRole = [newLogModel.roleId()]
				window.sessionStorage.setItem('roleid', JSON.stringify(sessArrRole))
			}
			// 从sessionStorage 获取输入记录
			newLogModel.saveRoleid(JSON.parse(window.sessionStorage.getItem('roleid')))
		}
		
		$.ajax({
			type : "get",
			async:true,
			url:"/gameLog/getItemDisappearLogPage",
			data : jsonData,
			dataType : "json",
			error:function(){
				layer.close(index); 
			},
			success : function(data){
				console.log("物品丢失日志：",data)
				layer.close(index); 
				if( data.code == 200 ){
					var tabDom = $("#loseItems_tbody").empty();
					newLogModel.dataList.removeAll();
					if (data.info.rows) {
						for (var i = 0; i < data.info.rows.length; i++) {
							//data.info.rows[i].typeName = indexModel.itemDisappearMap()[data.info.rows[i].type];
							newLogModel.dataList.push(data.info.rows[i]);
						}
						$("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });	
					}else{
						console.log("log-loseItems data.info.rows is no data.");
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
	
	var roleForm = $("#loseItemsForm").Validform({
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
