def main()
    puts "Connection to Kittenserver"
    request = Proto::Request.new(name:'Nic")
    body = send_request('kittenserver_kittenserver_1','8091','Kittens.Hello',request).body
    envelop,message = read_response(body,Proto""Response)
    puts envelop.inspect
    puts message.inspect
end