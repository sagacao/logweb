$(function(){

	indexModel.systemChannelHide();

	var logRawViewModel = function(){
        this.dataList				= ko.observableArray();
        this.roleId					= ko.observable("");
        this.keywords				= ko.observable("");
        this.searchType				= ko.observableArray([{"name":"and"}, {"name":"or"}]);
		this.selectedSearchType		= ko.observable("and");
		this.getServerIdChange		= ko.computed(function(){
			if (indexModel.zeusidChange() == "" && $("pageContainer").attr("data-url") == "/static/pages-html/log-sum/log-raw/log-raw.html"){
				var idCurrent = $('.clothLis li:first').attr("id");
				$('.clothLis li:first').parents("ul").siblings("input").attr("data-id",idCurrent);
				indexModel.zeusidChange($('.clothLis li:first').children("a").text());
				getServerId = idCurrent;
			}
		},this);
	}
	
	var newLogModel = new logRawViewModel();
	ko.applyBindings(newLogModel,document.getElementById("raw"));
	
	var pageClick = function( pageclickednumber ){
		var index = layer.load(1, {
		  	shade: [0.3,'#666666'],
		  	content:"数据获取中......"
		});
		var jsonData = {zeusid:getServerId,startTime:getStartDate(),endTime:getEndDate(),pageIndex:pageclickednumber,pageSize:50};
        newLogModel.keywords() ? $.extend(jsonData,{keywords: newLogModel.keywords()}) : $.extend(jsonData,{});
		newLogModel.selectedSearchType() ? $.extend(jsonData,{searchType: newLogModel.selectedSearchType()}) : $.extend(jsonData,{});
		$.ajax({
			type : "get",
			async:true,
			url:"/gameLog/getRawLogPage",
			data : jsonData,
			dataType : "json",
			error:function(){
				layer.close(index);
			},
			success : function(data){
				console.log("查询原始日志：",data)
				layer.close(index); 
				if( data.code == 200 ){
					var tabDom = $("#raw_tbody").empty();
					newLogModel.dataList.removeAll();
					if (data.info.rows) {
						for (var i = 0; i < data.info.rows.length; i++) {
							newLogModel.dataList.push(data.info.rows[i]);
						};
						$("#pager").pager({ pagenumber: pageclickednumber, pagecount: data.info.totalPage, buttonClickCallback: pageClick });		
					}else{
						console.log("log-raw data.info.rows is no data.");
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
	
	var roleForm = $("#rawForm").Validform({
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