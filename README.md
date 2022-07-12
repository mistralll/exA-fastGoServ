# これはなに
実験A最終課題のリポジトリ。
ｸｿﾃﾞｶCSVは重すぎるから .gitingore してます。

# データ下処理の方針
方向転換！
1. 無駄な行をtag, imgそれぞれ削除
    * tag: 一つずつ見ていって空白だったら消す
    * img: １つずつ見ていってtagが空白なら消す
2. 100件以上のやつを消す
    1. tagをベースに、[id, tag, datetme]を作る
    2. tag, datetimeでソート
    3. 100以降はappendしない



### ソートしてマージ
1. geotag.csv を id でソートする
2. tag.csv を tag でソートする

3. マージする
```
id, tag, date-time, x, y, url
```
形式のデータを作る

**いまここ** 

### 無駄なデータの削除とソート
1. タグが一件もない画像のidを削除
```
csvUtil.DelEmptyTag(infileNmae, outfileName)
```
2. 第一条件 tag 第二条件 date-time でソート
```
csvUtil.SortTagDate(infileName, outfileNmae)
```
3. 同じタグの中で100件を超える行を削除

### ラズパイの検索を最適化（案1）
1. データテーブルとインデックステーブルを作成
```
data: [x, y, date-time, URL]
index: [tag, start, cnt]
```
ラズパイ上の検索では、index.csvのデータを主記憶上に保持し、
検索リクエストが飛んできたら二分探索によりstartとcntの値を取得する
dataは適当に分割し、（100MBずつくらい？）startからcnt行拾ってきてHTMLを整形する

### ラズパイの検索を最適化（案2）
データを持ってくるのに時間がかかる気がするので、持ってくるデータを最低限にしたい。
{{tag}}.csv に画像データ x, y, date-time, URL の csv を作り、
アクセスがあったときにデータの読み込みを過不足なく行いHTMLを生成する。

### ラズパイの検索を最適化（案3）
もはやHTMLを作っておけばよいのでは
これでO(1)だね！！！
データが結構な量になるのでオーバーヘッドがすごそう。案2とどっちがいいかな

# 実際のデータの下処理
main.go の func main () {} 内を書き換えてソートとマージを行う。
```
import (
    "github.com/mistralll/expAcsv/csvUtil"
    "github.com/mistralll/expAcsv/tagUtil"
    "github.com/mistralll/expAcsv/geotagUtil"
)
func main() {
    tagUtil.SortCsv(入力ファイル, 保存先)
    geotagUtil.SortCsv(入力ファイル, 保存先)
    csvUtil.Merge(ソート済みtag.csvのパス, ソート済みgeotag.csvのパス, 保存先)
}
```


# パッケージとか
### csvUtil
csv ファイルの下処理用パッケージ。ソートとか分割とか。

### geotagUtil
geotag.csv 下処理用のパッケージ。
ソートとか色々。昔に作ったので余分なものも多い。

### tagUtil
tag.csv 下処理用パッケージ
ソートとか色々。昔に作ったので余分なものも多い。