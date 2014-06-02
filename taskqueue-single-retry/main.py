import logging

from google.appengine.api import taskqueue
from google.appengine.api.taskqueue import TaskRetryOptions
import webapp2

class MainPage(webapp2.RequestHandler):
    def get(self):
        taskqueue.add(
            url='/fail',
            method='POST',
            retry_options=TaskRetryOptions(task_retry_limit=0),
        )
        self.response.headers['Content-Type'] = 'text/plain'
        self.response.write('enqueued task')

class Fail(webapp2.RequestHandler):
    def post(self):
        for k in ['X-AppEngine-TaskName', 'X-AppEngine-TaskRetryCount']:
            logging.info('{}={}'.format(k, self.request.headers.get(k)))
        self.abort(500)

application = webapp2.WSGIApplication([
    ('/', MainPage),
    ('/fail', Fail),
])
