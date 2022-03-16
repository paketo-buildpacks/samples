import os
from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def hello_world():
    return render_template('index.html')

if __name__ == "__main__":
    app.run()

def run():
    port = int(os.getenv("PORT"))
    app.run(host='0.0.0.0', port=port)
