# 02_cleanarchitecture

# 検討

導入の大きな目的として「オープン・クローズドの原則」（＝「変更に閉じ、拡張に開かれている）があるように見える。
そのため、”具象”ではなく”抽象”に依存する構造を取る。
また、システム内の（変更の入りにくい）”ビジネスロジック”（コアと言える最小の粒度まで落とし込む）は entity として中心に据え、
各ビジネスロジックを組み合わせて usecase を構成する。


# Clean Architecture導入手順

「ルーティング」（＝どのパスが叩かれたら、どのロジックを実行するかの定義）は、
