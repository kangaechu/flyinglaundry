# flying_laundry

我が家は南風で風速5m以上だと洗濯物が飛んでいきます。
これからの天気情報をtenki.jpから取得し、条件を満たしたら何らかの方法で通知します。

## installation

```shell script
# pipenvのインストール
pip install pipenv

# 依存パッケージのインストール
pipenv install

# シェルの起動
pipenv shell
```

## useage

tenki.jpから取得したい地方の1時間天気のURLを持ってきます。
札幌であれば https://tenki.jp/forecast/1/2/1400/1100/1hour.html 。

app.py のurlを指定します。

```shell script
python app.py
```

