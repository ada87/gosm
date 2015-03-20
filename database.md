#数据字典 

##数据文件 gosm-field

定义了一些常用的数据字段。

###字段表 Filed

表Filed 存一些面的可用字段

|字段|类型|默认值|说明|
|---|---|---|---|
|field_id|INTEGER|&nbsp;|**[主键自增]**唯一ID|
|field_code|TEXT|&nbsp;|**[非空唯一]** 字段代码|
|field_formid|TEXT|&nbsp;|字段对应html表单ID|
|field_formtype|INTEGER| 1 | 1:输入 2:下拉 3:checkbox |
|field_desc|TEXT|&nbsp;| 字段描述|

###字段值表 Field_Value

表Field_Value 存页面上字段的值


|字段|类型|默认值|说明|
|---|---|---|---|
|id|INTEGER|&nbsp;|**[主键自增]**唯一ID|
|field_id|INTEGER|&nbsp;|**[非空外键]**所属字段ID|
|field_value|TEXT|&nbsp;|字段的值|
|status|INTEGER|2|0:默认 1:扩展 2:自定义|


