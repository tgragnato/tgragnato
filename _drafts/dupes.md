---
title: dupes
layout: default
---

```ruby
require 'find'
require 'digest/sha2'
require 'optparse'
require 'set'

# optparse
args = { :dir => nil , :del => false }
arguments = OptionParser.new do |arg|

  arg.banner = "Usage: #{__FILE__} [options]"

  arg.on("-d", "--dir DIRECTORY", "directory to find duplicates") do |dir|
    args[:dir] = dir
  end

  arg.on("-f", "--force", "remove without prompting.") do
    args[:del] = true
  end

  arg.on("-h", "--help", "print help message") do
    puts arguments
    exit 0
  end
end

arguments.parse!
exit 1 if args[:dir] == nil

# digests
hashes = Hash.new
puts "Building digests ..."

Find.find(args[:dir]) do |path|
  STDOUT.flush
  hash = nil

  if (not FileTest.directory?(path) and not FileTest.symlink?(path)\
      and not FileTest.blockdev?(path) and not FileTest.chardev?(path)\
      and not FileTest.socket?(path) and not FileTest.pipe?(path)\
      and not FileTest.zero?(path)) then

    begin
      digest = Digest::SHA2.new
      File.open(path, 'r') do |fd|
        while buffer = fd.read(9216)
          digest.update(buffer)
        end
      end
      hash = digest.hexdigest

    rescue NoMemoryError
      puts "NoMemoryError on #{path}"
      GC.start
    rescue IOError
      puts "IOError on #{path}"
      GC.start
    rescue Errno::ENOSYS; next
    rescue Errno::EACCES; next
    rescue Errno::EINVAL; next
    rescue => e
      puts "Error on #{path}, exception #{e}"
      exit 255
    end
  end

  if (hashes[hash]) then
    hashes[hash].add(path)
  else
    hashes[hash] = Set.new.add(path)
  end
end

# results
hashes.each_pair do |hash, paths|
  files = paths.to_a
  choice = nil

  if (files.size > 1) then
    puts "Duplicate digest <#{hash}> for files:\n  #{files.join("\n  ")}"

    if not args[:del] then
      puts "  Delete? [y/N] "
      choice = gets.chomp == 'y' ? true : false
    end

    if args[:del] or choice then
      for i in 1...files.size
        File.delete files[i]
      end
    end

  end
end
```
