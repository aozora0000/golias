# Golias

## HowToUse

yamlで管理するサブコマンド集

素となるgoliasをコピーすると、対応したバイナリ名＋サブコマンドが作成出来ます。

単体動作なら```command```、```commands```を使えばパイプライン接続が可能です。

両方記述がある場合には```commands```のみが実行されます。

### Config

$HOME/.config/golias/[バイナリ名].yaml
- ```$ golias init``` configファイルを作成します
- ```$ golias edit``` エディタが立ち上がります
- ```$ golias path``` configファイルの絶対パスを表示します

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
- name: example3
  command: echo %date
  envs:
    date: date '%format'
  params:
    FORMAT: +%Y-%m-%d
```

#### args

```command```に対するオプション引数を静的に設定出来る、commandに書いても良いがパラメータが多い時に便利

#### envs

```command```に対して```%{key}```の値をシェル実行結果から代入出来る。

argsに比べて動的に代入出来る為便利に使える。

#### params

````command````と````envs````に対して```%{key}```の値を静的に代入出来る。 固定ディレクトリ名等共通する静的パラメータを入れておくと便利

### 単体実行について

```shell
$ golias
```

で実行された場合、```name: main```で指定されたコマンドが実行されます。
```name: main```が無い場合はヘルプが実行されます。

#### コマンド引数について

ファイル名やディレクトリ名等のオプションでは無い引数は、コマンドの後ろへ自動的に追加されます。

### example

```yaml
# sshの設定ファイルからpecoを使ってインタラクティブにリモートログイン出来る
- name: ssh
  command: ssh %hostname
  envs:
    hostname: cat ~/.ssh/config | grep "^Host" | cut -c6- | peco

# 最新コミットIDのgithubをブラウザで開く    
- name: open
  command: xdg-open https://%url/commit/%commit
  envs:
    url: git ls-remote --get-url | sed "s/git@//" | sed "s/:/\//" | sed "s/.git//"
    commit: git log -1 | sed -n '1,1p' | cut -c8-
```

### goliasをインストールしたディレクトリに別名コピーする

```shell
$ cp `which golias` $(dirname `which golias`)/hoge
```

