# bitflyer
* bitflyerのdashboard
    * https://lightning.bitflyer.com/trade
    * https://bitflyer.com/ja-jp/ex/home
* bitflyerのapiのドキュメント
    * https://lightning.bitflyer.com/docs/playground?lang=ja#GETv1%2Fgetmarkets/javascript

# 単純移動平均線(SMA)
* 3日間の終値を平均して算出 -> 10/20/30と変化した場合
10円+20円+30円
-------------- =  20円
   3day
* 毎日平均値を求め、計算された平均値を結んでグラフ化すると移動平均線になる

# 指数平滑移動平均線(EMA)
* 個々の価格データへの加重を指数関数的に減少させて近日の終値を足して平均値を計算する
10円+20円+30円+30円(最終日)
------------------------- =  22.5円
   3day + 1day
* a(平滑定数)=2/(n+1)