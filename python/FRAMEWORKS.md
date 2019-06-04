# Python Frameworks and comparison

## Django
Django is a full-stack web framework for Python
https://docs.djangoproject.com/en/2.2/

Multi-layered (Model, View, Template)

### Features
- Functional Admin Interface (Doc generator, Admin site, Admin actions)
- Template Engine
- Forms
- Built-in Bootstrapping Tool
- Project Layout
- DB Support (ORM)
- Flexibility
- Internationalization, Localization
- Security

### Hello World

`python3 manage.py startapp helloworld`

Open automatically created `helloword/views.py` and add 

```python
from django.http import HttpResponse

def index(request):
    return HttpResponse("Hello, World!")
```

Edit helloword/urls.py

```python
from django.conf.urls import url

from . import views

urlpatterns = [
    url(r'^$', views.index, name='index'),
]
```

Edit `hellodjango/hellodjango/urls.py`

```python
from django.conf.urls import include, url 
 
urlpatterns = [ 
    url(r'^hello/', include('helloworld.urls')), 
]
```

Now run...

`python3 manage.py runserver`


## Flask
Flask is a lightweight and extensible Python web framework

https://flask.palletsprojects.com/en/1.1.x/

### Features
- Routing
- Middlewares
- Static Files
- Rendering Templates
- Accessing Request Data
- Logging
- Redirect / Errors


### Hello World
```python
from flask import Flask

app = Flask(__name__)

@app.route("/") # python decorator
def hello():
  return "Hello, World!"
  
if __name__ == "__main__":
  app.run()
```

`python3 flaskhello.py`

OR

`env FLASK_APP=hello.py flask run`
