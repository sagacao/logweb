/**
 * A Highcharts plugin for exporting data from a rendered chart as CSV, XLS or HTML table
 *
 * Author:   Torstein Honsi
 * Licence:  MIT
 * Version:  1.3.8
 */
/*global Highcharts, window, document, Blob */
(function (factory) {
    if (typeof module === 'object' && module.exports) {
        module.exports = factory;
    } else {
        factory(Highcharts);
    }
})(function (Highcharts) {

    'use strict';

    var each = Highcharts.each,
        pick = Highcharts.pick,
        downloadAttrSupported = document.createElement('a').download !== undefined;

    Highcharts.setOptions({
        // fixed Chinese lang options
        lang: {
            printChart: '打印图表',
            downloadJPEG: '下载 JPEG 文件',
            downloadPDF: '下载 PDF&nbsp;&nbsp; 文件',
            downloadPNG: '下载 PNG&nbsp; 文件',
            downloadSVG: '下载 SVG&nbsp; 文件',
            downloadCSV: '下载 CSV&nbsp; 文件',
            downloadXLS: '下载 XLS&nbsp;&nbsp; 文件'
        },
        exporting: {
            // reset the default export server, to fixed Chinese filename
            url: 'http://export.hcharts.cn/index.php',
            buttons: {
                contextButton: {

                }
            }
        }
    });


    /**
     * Get the data rows as a two dimensional array
     */
    Highcharts.Chart.prototype.getDataRows = function () {
        var options = (this.options.exporting || {}).csv || {},
            xAxis = this.xAxis[0],
            rows = {},
            rowArr = [],
            dataRows,
            names = [],
            i,
            x,
            xTitle = xAxis.options.title && xAxis.options.title.text,

            // Options
            dateFormat = options.dateFormat || '%Y-%m-%d %H:%M:%S',
            columnHeaderFormatter = options.columnHeaderFormatter || function (series, key, keyLength) {
                return series.name + (keyLength > 1 ? ' ('+ key + ')' : '');
            };

        alert(this.options.chart.type)

        // Loop the series and index values
        i = 0;
        each(this.series, function (series) {
            var keys = series.options.keys,
                pointArrayMap = keys || series.pointArrayMap || ['y'],
                valueCount = pointArrayMap.length,
                requireSorting = series.requireSorting,
                categoryMap = {},
                j;

            // Map the categories for value axes
            each(pointArrayMap, function (prop) {
                categoryMap[prop] = (series[prop + 'Axis'] && series[prop + 'Axis'].categories) || [];
            });

            if (series.options.includeInCSVExport !== false && series.visible !== false) { // #55
                j = 0;
                while (j < valueCount) {
                    names.push(columnHeaderFormatter(series, pointArrayMap[j], pointArrayMap.length));
                    j = j + 1;
                }

                each(series.points, function (point, pIdx) {
                    var key = requireSorting ? point.x : pIdx,
                        prop,
                        val;

                    j = 0;

                    if (!rows[key]) {
                        rows[key] = [];
                    }
                    rows[key].x = point.x;

                    // Pies, funnels etc. use point name in X row
                    if (!series.xAxis) {
                        rows[key].name = point.name;
                    }

                    while (j < valueCount) {
                        prop = pointArrayMap[j]; // y, z etc
                        val = point[prop];
                        rows[key][i + j] = pick(categoryMap[prop][val], val); // Pick a Y axis category if present
                        j = j + 1;
                    }

                });
                i = i + j;
            }
        });

        // Make a sortable array
        for (x in rows) {
            if (rows.hasOwnProperty(x)) {
                rowArr.push(rows[x]);
            }
        }
        // Sort it by X values
        rowArr.sort(function (a, b) {
            return a.x - b.x;
        });

        // Add header row
        if (!xTitle) {
            xTitle = xAxis.isDatetimeAxis ? '时间' : '分类';
        }
        dataRows = [[xTitle].concat(names)];

        // Transform the rows to CSV
        each(rowArr, function (row) {

            var category = row.name;
            if (!category) {
                if (xAxis.isDatetimeAxis) {
                    category = Highcharts.dateFormat(dateFormat, row.x);
                } else if (xAxis.categories) {
                    category = pick(xAxis.names[row.x], xAxis.categories[row.x], row.x)
                } else {
                    category = row.x;
                }
            }

            // Add the X/date/category
            row.unshift(category);
            dataRows.push(row);
        });

        return dataRows;
    };

    /**
     * Get the data rows as a two dimensional array
     */
    Highcharts.Chart.prototype.getHeatmapDataRows = function () {
        var options = (this.options.exporting || {}).csv || {},
            xAxis = this.xAxis[0],
            yAxis = this.yAxis[0],
            rows = {},
            clis = [],
            rowArr = [],
            dataRows,
            names = [],
            i,
            x,
            kIdx,
            xTitle = xAxis.options.title && xAxis.options.title.text;

            // Options
            // dateFormat = options.dateFormat || '%Y-%m-%d %H:%M:%S',
            // columnHeaderFormatter = options.columnHeaderFormatter || function (series, key, keyLength) {
            //     return series.name + (keyLength > 1 ? ' ('+ key + ')' : '');
            // };

        names.push(xTitle)
        each(xAxis.categories, function(xcategory){
            names.push(xcategory)
        });

        //clis.push()
        each(yAxis.categories, function(ycategory){
            clis.push(ycategory)
        });

        // Loop the series and index values
        i = 0;
        each(this.series, function (series) {
            var keys = series.options.keys,
                pointArrayMap = keys || series.pointArrayMap || ['y'],
                valueCount = pointArrayMap.length,
                requireSorting = series.requireSorting,
                categoryMap = {},
                j;

            // Map the categories for value axes
            each(pointArrayMap, function (prop) {
                categoryMap[prop] = (series[prop + 'Axis'] && series[prop + 'Axis'].categories) || [];
            });

            if (series.options.includeInCSVExport !== false && series.visible !== false) { // #55

                kIdx = 0
                each(series.points, function (point, pIdx) {
                    var key = requireSorting ? point.x : pIdx,
                        prop,
                        val;
                    if (kIdx % clis.length == 0) {
                        kIdx = 0
                    }

                    if (!rows[kIdx]) {
                        rows[kIdx] = [];
                        rows[kIdx][0] = clis[kIdx];
                    }
                    rows[kIdx].x = point.x;

                    // // Pies, funnels etc. use point name in X row
                    // if (!series.xAxis) {
                    //     rows[kIdx].name = point.name;
                    // }

                    //val = point["value"];
                    //rows[kIdx][key % names.length] = pick(categoryMap["value"][val], val);
                    rows[kIdx][Math.floor(key / (clis.length)) + 1] = point["value"];

                    kIdx = kIdx + 1;
                });
                i = i + j;
            }
        });
        // Make a sortable array
        for (x in rows) {
            if (rows.hasOwnProperty(x)) {
                rowArr.push(rows[x]);
            }
        }
        dataRows = []
        dataRows.push(names)
        each(rowArr, function (row) {
            dataRows.push(row);
        });

        return dataRows;
    };

    /**
     * Get a CSV string
     */
    Highcharts.Chart.prototype.getCSV = function (useLocalDecimalPoint) {
        var csv = '',
            rows,
            options = (this.options.exporting || {}).csv || {},
            itemDelimiter = options.itemDelimiter || ',', // use ';' for direct import to Excel
            lineDelimiter = options.lineDelimiter || '\n'; // '\n' isn't working with the js csv data extraction
        
        if (this.options.chart.type == "heatmap") {
            rows = this.getHeatmapDataRows(); 
        } else {
            rows = this.getDataRows(); 
        }

        // Transform the rows to CSV
        each(rows, function (row, i) {
            var val = '',
                j = row.length,
                n = useLocalDecimalPoint ? (1.1).toLocaleString()[1] : '.';
            while (j--) {
                val = row[j];
                if (typeof val === "string") {
                    val = '"' + val + '"';
                }
                if (typeof val === 'number') {
                    if (n === ',') {
                        val = val.toString().replace(".", ",");
                    }
                }
                row[j] = val;
            }
            // Add the values
            csv += row.join(itemDelimiter);

            // Add the line delimiter
            if (i < rows.length - 1) {
                csv += lineDelimiter;
            }
        });

        return csv;
    };

    /**
     * Build a HTML table with the data
     */
    Highcharts.Chart.prototype.getTable = function (useLocalDecimalPoint) {
        var html = '<table>',
            rows;

        if (this.options.chart.type == "heatmap") {
            rows = this.getHeatmapDataRows(); 
        } else {
            rows = this.getDataRows(); 
        }   

        // Transform the rows to HTML
        each(rows, function (row, i) {
            var tag = i ? 'td' : 'th',
                val,
                j,
                n = useLocalDecimalPoint ? (1.1).toLocaleString()[1] : '.';

            html += '<tr>';
            for (j = 0; j < row.length; j = j + 1) {
                val = row[j];
                // Add the cell
                if (typeof val === 'number') {
                    val = val.toString();
                    if (n === ',') {
                        val = val.replace('.', n);
                    }
                    html += '<' + tag + ' class="">' + val + '</' + tag + '>';

                } else {
                    html += '<' + tag + '>' + (val === undefined ? '' : val) + '</' + tag + '>';
                }
            }

            html += '</tr>';
        });
        html += '</table>';
        //alert(html)
        //document.write(html)
        return html;
    };

    function getContent(chart, href, extension, content, MIME) {
        console.log(href);
        console.log(content);
        console.log(MIME);
        var a,
            blobObject,
            name,
            options = (chart.options.exporting || {}).csv || {},
            // change export csv server url to fixed Chinese char
            url = options.url || 'http://export.hcharts.cn/csv.php';

        if (chart.options.exporting.filename) {
            name = chart.options.exporting.filename;
        } else if (chart.title) {
            name = chart.title.textStr.replace(/ /g, '-').toLowerCase();
        } else {
            name = 'chart';
        }

        // MS specific. Check this first because of bug with Edge (#76)
        if (window.Blob && window.navigator.msSaveOrOpenBlob) {
            // Falls to msSaveOrOpenBlob if download attribute is not supported
            blobObject = new Blob([content]);
            window.navigator.msSaveOrOpenBlob(blobObject, name + '.' + extension);

        // Download attribute supported
        } else if (downloadAttrSupported) {
            a = document.createElement('a');
            a.href = href;
            a.target = '_blank';
            a.download = name + '.' + extension;
            document.body.appendChild(a);
            a.click();
            a.remove();

        } else {
            // Fall back to server side handling
            Highcharts.post(url, {
                data: content,
                type: MIME,
                extension: extension
            });
        }
    }

    /**
     * Call this on click of 'Download CSV' button
     */
    Highcharts.Chart.prototype.downloadCSV = function () {
        var csv = this.getCSV(true);
        getContent(
            this,
            // add \ufeff to Fixed Chinese in csv file
            // http://www.honger05.com/html5/2015-04-01-html5-download/
            'data:text/csv,\ufeff' + csv.replace(/\n/g, '%0A'),
            'csv',
            csv,
            'text/csv'
        );
    };

    /**
     * Call this on click of 'Download XLS' button
     */
    Highcharts.Chart.prototype.downloadXLS = function () {
        //alert("call downloadXLS -------")
        var uri = 'data:application/vnd.ms-excel;base64,',
            template = '<html xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns="http://www.w3.org/TR/REC-html40">' +
                // add charset meta to fixed Chinese char in XLS file
                '<meta http-equiv="Content-Type" charset=utf-8">'+
                '<head><!--[if gte mso 9]><xml><x:ExcelWorkbook><x:ExcelWorksheets><x:ExcelWorksheet>' +
                '<x:Name>Ark1</x:Name>' +
                '<x:WorksheetOptions><x:DisplayGridlines/></x:WorksheetOptions></x:ExcelWorksheet></x:ExcelWorksheets></x:ExcelWorkbook></xml><![endif]-->' +
                '<style>td{border:none;font-family: Calibri, sans-serif;} .number{mso-number-format:"0.00";}</style>' +
                '<meta name=ProgId content=Excel.Sheet>' +
                '</head><body>' +
                this.getTable(true) +
                '</body></html>',
            base64 = function (s) {
                return window.btoa(unescape(encodeURIComponent(s))); // #50
            };
        getContent(
            this,
            uri + base64(template),
            'xls',
            template,
            'application/vnd.ms-excel'
        );
    };


    // Add "Download CSV" to the exporting menu. Use download attribute if supported, else
    // run a simple PHP script that returns a file. The source code for the PHP script can be viewed at
    // https://raw.github.com/highslide-software/highcharts.com/master/studies/csv-export/csv.php
    if (Highcharts.getOptions().exporting) {
        Highcharts.getOptions().exporting.buttons.contextButton.menuItems.push({
            textKey: 'downloadCSV',
            onclick: function () { this.downloadCSV(); }
        }, {
            textKey: 'downloadXLS',
            onclick: function () { this.downloadXLS(); }
        });
    }

});
