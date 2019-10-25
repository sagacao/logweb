$(function(){

    inits();

    function inits(){
        
        indexModel.channelHide();

        var myChart = echarts.init($(".pie-wrap").get(0));

        var option = {
            /* 
				a(系列名称)、b(数据项名称)、c(数值)、 d（饼图：百分比 | 雷达图：指标名称）
			 */
            // backgroundColor:"lightblue",
            title: {
                text: '活跃角色VIP分布',
                x: 'center',
                y: '10%'
            },
            tooltip: {
                trigger: 'item',
                formatter: "{a} <br/>{b} : {c} ({d}%)"
            },
            toolbox: {
                show: true,
                x: "90%",
                feature: {
                    dataView: { readOnly: true },
                    saveAsImage: {}
                }
            },
            calculable: true,
            series: [
                {
                    name: '访问来源',
                    type: 'pie',
                    radius: '55%',
                    center: ['50%', '60%'],
                    itemStyle: {
                        normal: {
                            label: {
                                show: true,
                                formatter: '{b} : {c} ({d}%)'
                            },
                            labelLine: { show: true }
                        }
                    },
                    data: [
                        { value: 50, name: 'VIP1' },
                        { value: 40, name: 'VIP2' },
                        { value: 50, name: 'VIP3' },
                        { value: 50, name: 'VIP4' },
                        { value: 55, name: 'VIP5' },
                        { value: 50, name: 'VIP6' },
                        { value: 75, name: 'VIP7' },
                        { value: 50, name: 'VIP8' },
                        { value: 60, name: 'VIP9' },
                        { value: 98, name: 'VIP10' },
                    ]
                }
            ]
        }

        $.ajax({
            url:"/pay/payvipdata",
            dataType:"json",
            success:function(data){
                console.log("vip分布：",data);
                myChart.setOption(option); 
            },
            error:function(){
            }
        });
        
    }
    
});