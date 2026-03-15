---
allowed-tools: Bash, Read
description: Browser automation via Playwright CLI (npx @playwright/cli)
---

# Playwright CLI

ブラウザ操作を `npx @playwright/cli` 経由で行う。MCP 版より軽量（コンテキスト消費 ~1.3%）。

## 基本ルール

- セッション名を必ず指定する: `-s=<name>`
- スクショは `.hq/screenshots/` に保存する（プロジェクトルートを汚さない）
- file:// URL は使えない。ローカルファイルを確認する場合は HTTP サーバーを立てる
- 作業完了後はセッションを閉じる

## コマンドリファレンス

```bash
# ブラウザ起動 & ナビゲーション
npx @playwright/cli -s=test open <url>
npx @playwright/cli -s=test goto <url>

# ページ状態取得
npx @playwright/cli -s=test snapshot          # アクセシビリティツリー (要素 ref 取得)
npx @playwright/cli -s=test screenshot        # スクショ (viewport)
npx @playwright/cli -s=test screenshot <ref>  # 要素のスクショ

# インタラクション
npx @playwright/cli -s=test click <ref>
npx @playwright/cli -s=test fill <ref> <text>
npx @playwright/cli -s=test type <text>
npx @playwright/cli -s=test select <ref> <val>
npx @playwright/cli -s=test hover <ref>
npx @playwright/cli -s=test press <key>

# セッション管理
npx @playwright/cli -s=test close
npx @playwright/cli close-all
```

## 典型的なワークフロー

ユーザーの指示に応じて以下のパターンで操作する。

### ページ確認 & スクショ

```bash
npx @playwright/cli -s=check open <url>
npx @playwright/cli -s=check screenshot
# → .playwright-cli/ にスクショが保存される。Read で確認
npx @playwright/cli -s=check close
```

### インタラクティブ操作

```bash
npx @playwright/cli -s=test open <url>
npx @playwright/cli -s=test snapshot        # ref を取得
npx @playwright/cli -s=test click <ref>     # 操作
npx @playwright/cli -s=test snapshot        # 結果確認
npx @playwright/cli -s=test screenshot      # 必要ならスクショ
npx @playwright/cli -s=test close
```

### ローカルファイル確認

file:// が使えないため HTTP サーバーを立てる:

```bash
# バックグラウンドでサーバー起動 (run_in_background: true)
cd <dir> && python3 -m http.server 8765

# ブラウザで開く
npx @playwright/cli -s=local open http://localhost:8765/<file>
npx @playwright/cli -s=local screenshot
npx @playwright/cli -s=local close

# サーバー停止
kill <pid>
```
