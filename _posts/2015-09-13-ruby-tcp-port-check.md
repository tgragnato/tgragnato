---
title: RubyでTCPポートへの接続を確認する
title_marker: TCP接続確認
description: TCPSocketとタイムアウトを使った、小さなTCP接続チェック
layout: default
lang: ja
---

このスクリプトは、指定したホストとポートにTCP接続を試み、最大10秒で結果を表示します。サービス監視の最小例として、ソケット接続とタイムアウト処理の流れを確認できます。

確認できるのはTCP接続が成立するかどうかだけです。アプリケーション層の応答やサービスの健全性までは検証しません。また、接続拒否、到達不能、名前解決エラーは同じ「errored」として表示され、タイムアウトとは区別されます。現在のRubyでは`Timeout.timeout`の利用にも注意が必要です。

```ruby
require 'socket'
require 'timeout'

unless ARGV.length == 2
  abort "Usage: #{__FILE__} <host> <port>"
end

where = ARGV[0]
port = ARGV[1]

begin
  Timeout::timeout(10) do
    begin
      TCPSocket.new(where, port).close
      puts "connect(#{where}, #{port}) succeeded"
    rescue Errno::ECONNREFUSED, Errno::EHOSTUNREACH, SocketError
      puts "connect(#{where}, #{port}) errored"
    end
  end
rescue Timeout::Error
  puts "connect(#{where}, #{port}) failed"
end
```