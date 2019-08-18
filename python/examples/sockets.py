import socket
# test with nc
# sudo nc -l 256 & python sockets.py
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("localhost", 256))
sock.sendall(b'hello socket world') 
sock.close()