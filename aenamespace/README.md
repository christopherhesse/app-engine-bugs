If this test is correct, tasks on classic appengine using the new library (google.golang.org/appengine) will lose their namespaces.

GOPATH=$PWD appcfg.py update src/aenamespace/app.yaml

curl https://aenamespace-dot-ae-bugs.appspot.com/classic-enqueue
output from /classic-task will be `the value is "classic-namespace"`

curl https://aenamespace-dot-ae-bugs.appspot.com/new-enqueue
output from /classic-task will be `item not in the cache` it should instead be `the value is "new-namespace"`
