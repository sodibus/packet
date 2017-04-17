#!/usr/bin/env ruby

require 'erb'
require 'ostruct'

name2types = {}
type2names = {}

File.open("types.list").read.to_s.each_line do |line|
  comps = line.split("=").map do |c|
    c.strip
  end
  next unless comps.count == 2
  name2types[comps[1]] = comps[0]
  type2names[comps[0]] = comps[1]
end

local = OpenStruct.new

TPL = <<-EOF
package packet

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func frameTypeFromPacket(m Packet) uint8 {
	switch m.(type) {
  <% name2types.each_pair do |k,v| %>
  case *<%= k %>:
    {
      return <%= v %>
    }
  <% end %>
	default:
		{
			return 0
		}
	}
}

func packetFromFrame(f *Frame) (Packet, error) {
  var err error
  switch f.Type {
  <% type2names.each_pair do |k,v| %>
  case <%= k %>:
    {
      var p <%= v %>
      err = proto.Unmarshal(f.Data, &p)
      return &p, err
    }
  <% end %>
  default: {
    return nil, fmt.Errorf("Packet Type %d Not Supported", f.Type)
  }
  }
}
EOF

result = ERB.new(TPL).result

File.open("frame.gen.go", "w") do |f|
  f.puts result
end
