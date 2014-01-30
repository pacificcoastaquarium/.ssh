
from flask import Flask
from werkzeug.wsgi import SharedDataMiddleware
from os.path import join as joinpath, dirname


app = Flask(__name__)

app.wsgi_app = SharedDataMiddleware(app.wsgi_app,
            {
                '/dojolib': joinpath(dirname(__file__), './dojotoolkit'),
                '/olesk': joinpath(dirname(__file__), './olesk'),
                '/': joinpath(dirname(__file__), '.')
            },
            cache=False,
            cache_timeout=0
      )

if __name__ == "__main__":
    app.run(debug=True)
