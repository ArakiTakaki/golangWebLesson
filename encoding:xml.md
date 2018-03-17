# encoding/xml package

　XML形式の文字列をパースする、他にも動的なxmlを生成できるらしい


| 返り値 | メソッドと説明 |
|:---|:---|
|output []byte ,err|Marshal( v interface{} )|
| ^ |vに沿った形のxmlを出力する。  配列などはスライス処理をし、出力する。 (１か複数かは問わないがnilの場合なんも書き込まずinterface値を処理する)|
|output []byte ,err|MarshalIndent( v interface{}, prefix, indent string )|
|^|Marshalのように動作するがインデントをしっかり取ってくれる。|
|output []byte ,err|MarshalIndent( v interface{}, prefix, indent string )|
|^|Marshalのように動作するがインデントをしっかり取ってくれる。|

基本形
|||
|^||

## サンプルケース

``` go:main.go

type Name struct {
    First, Last string
}

type Person struct {
    Name Name
    Gender string
    Age int
}

func TestData() {
    john := Person{Name{"John", "Doe"}, "Male", 20}
}

```


## Marshal

### 出力結果

``` xml:sample.xml
    <Person><Name><First>John</First><Last>Doe</Last></Name><Gender>Male</Gender><Age>20</Age></Person>
```


## MarshalIndent

### 出力

``` xml:sample.xml
<Person>
    <Name>
        <First></First>
        <Last></Last>
    </Name>
    <Gender></Gender>
    <Age></Age>
</Person>
```


### 解説
　ちゃんとパースされたXMLが出力されるので可読性が高いから


## UnMarshal

``` cmd:sample.xml
 {Name:{First:John Last:Doe} Gender:Male Age:20}
```

### 解説
　構造体通りに帰ってくる、可用性は高そう。