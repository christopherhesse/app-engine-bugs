app-engine-bugs
===============

Google App Engine bugs that I have encountered

* go-log: go runtime drops log data
* taskqueue-single-retry: task queue will retry at least once even if retries are disabled
* taskqueue-custom-host: taskqueue by default routes to wrong host when using custom domain
* delay-gob: gob seems to interact oddly with appengine's main package
