require 'date'

if ARGV.empty?
  puts "No time provided"
  exit 1
end

datetime = DateTime.iso8601(ARGV[0])
puts "Date: %s" % datetime.strftime('%Y-%m-%d')
puts "Time: %s" % datetime.strftime('%H:%M:%S')
