---
layout: page
title: Google Cloud Run
hide_hero: true
show_sidebar: false
menubar: docs-menu
---

# Google Cloud Run

To get started in a hosted environment you can deploy this project to the Google Cloud Platform.

From your [Google Cloud dashboard](https://console.cloud.google.com/home/dashboard) create a new project and call it:
```
moov-ach-demo
```

Click the button below to deploy this project to Google Cloud.

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run/?git_repo=https://github.com/moov-io/ach&revision=master)


In the cloud shell you should be prompted with:
```
Choose a project to deploy this application:
```

Using the arrow keys select:
```
moov-ach-demo
```


You'll then be prompted to choose a region, use the arrow keys to select the region closest to you and hit enter.
```
Choose a region to deploy this application:
```



Upon a successful build you will be given a URL where the API has been deployed:
```
https://YOUR-ACH-APP-URL.a.run.app
```

From the cloud shell you need to cd into the `ach` folder:
```
cd ach
```

Now you can list files stored in-memory:
```
curl https://YOUR-ACH-APP-URL.a.run.app/files
```
You should get this response:
```
{"files":[],"error":null}
```


Create a file on the server:
```
curl -X POST --data-binary "@./test/testdata/ppd-debit.ach" https://YOUR-ACH-APP-URL.a.run.app/files/create
```
You should get this response:
```
{"id":"<YOUR-UNIQUE-FILE-ID>","error":null}
```


Finally, read the contents of the file you've just posted:
```
curl https://YOUR-ACH-APP-URL.a.run.app/files/<YOUR-UNIQUE-FILE-ID>
```

You should get this response:
```
{"file":{"id":"<YOUR-UNIQUE-FILE-ID>","fileHeader":{"id":"...","immediateDestination":"231380104","immediateOrigin":"121042882", ...
```