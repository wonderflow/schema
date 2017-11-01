# schema


type assert &amp;&amp; schema reflect

## 主要功能

* 类型推断，将interface的数据推断为常见的类型.
* 类型转换，将interface的数据转化为常见的类型.


## 类型的支持程度

* Long   ：所有golang的整型
* Float  ：所有golang的浮点型
* String ：所有字符串
* Date   ：rfc3339的string转化，及 golang的time.Time类型
* Bool   ：golang的bool类型
* Array  ：slice类型，并支持里面的元素推断
* Map    ：map类型，并嵌套的推断里面的元素类型

