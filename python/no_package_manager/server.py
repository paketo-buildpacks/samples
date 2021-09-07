import http.server
import socketserver
import os

PORT = 8080

web_dir = os.path.join(os.path.dirname(__file__), 'web')
os.chdir(web_dir)

Handler = http.server.SimpleHTTPRequestHandler
httpd = socketserver.TCPServer(("", PORT), Handler)
httpd.serve_forever()
