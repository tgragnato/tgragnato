#!/usr/bin/env ruby
# Copyright © 2017 Tommaso Gragnato <gragnato.tommaso@icloud.com>

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