import ctypes as dll
from flask import *

lib = dll.WinDLL("cli.dll")
app = Flask(__name__)    
@app.route('/')
def hello_world():
   return render_template('home.html')
@app.route('/<game>')
def strgm(game):
	try:
		return render_template('loader.html',gmn = game)
		#lib.Connect(game)
	except:
		print("Cannot connect to server.")
flask = Process(target=app.run(threaded=True))
if __name__ == '__main__':
	app.run(threaded=True)