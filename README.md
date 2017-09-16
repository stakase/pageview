# PV数保存・表示
## PV数保存処理
- pageview.js
  - 埋め込み用JS
- pageview.go
  - ページビュー数保存用処理
## PV数表示処理
- review.go
  - PV数表示処理
- template/review.tml
  - 時間帯別PV数表示用
- template/reviewlist.tml

## URL
- http://xxxxxxxxx/reviewlist
  - PV数表示用TOPページ

## 残作業
- 割り切って時間単位で処理してるのをやめる
  - ただし、その場合にはRDBでは負荷が高いのでPostgreSQLやめる必要あり
