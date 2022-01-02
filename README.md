watchgoggy
==========
watchgoggy can be restart target app when target app is dead.

Note
----
For now, this version just can run in Linux, Welcome friends add other os supports.

If you start your app with script, you must be sure your script will not return before your app exit. 
For little suggestions is you can do while sleep function end of your script.  

Usage
-----
You just
+ add your app's config in resources/goggies.json.
+ config variable Home in settings.go to your working dir.
+ run watchgoggy like /root/workspace/watchgoggies/watchgoggy -goggies /root/workspace/watchgoggies/goggies.json


