# Golias

yamlで管理するサブコマンド集

素となるgoliasをコピーすると、対応したバイナリ名＋サブコマンドが作成出来ます。

$HOME/.config/golias/[バイナリ名].yaml

単体動作なら```command```、```commands```を使えばパイプライン接続が可能です。

両方記述がある場合には```commands```のみが実行されます。

```yaml
- name: example1
  command: ls
  args:
  - -la
  usage: list file display
- name: example2
  commands:
  - command: ls
    args:
    - -la
  - command: wc
    args:
    - -l
  usage: list file count
```

goliasをインストールしたディレクトリに別名コピーする

```terminal
$ cp `which golias` $(dirname `which golias`)/hoge
```