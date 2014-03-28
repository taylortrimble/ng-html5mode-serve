ng-html5mode-serve
==================

Hey. When you guys find the thing that already does this, preferably in Grunt, will you tell me? kthxbye.

What this does
--------------

So, you're serving an Angular app at `http://ngexampleapp.com/`, aka index.html.
If we want simple static serving to work with routes (like `/dashboard`), we have to accept non-HTML5 mode URLs `http://ngexampleapp.com/#/dashboard`.
And that's garbage.

We need a server that will respond to URLs like `http://ngexampleapp.com/dashboard` with index.html. Then our app could take over and take us to the right view.
But we also need to serve some paths (like `/styles/` and `/scripts/`) with their real files: not index.html.

This server takes a configuration that tells it what port to serve on and what routes to serve real files for. It will serve index.html for everything else.

Example Config
--------------

```
{
    "port": 4000,
    "file_paths": ["bower_components", "styles", "scripts", "views"]
 }
```

#### Example routes from config

| Path                             | File served                                            |
|----------------------------------|--------------------------------------------------------|
| /                                | index.html                                             |
| /styles/main.css                 | styles/main.css                                        |
| /scripts/controllers/app.js      | scripts/controllers/app.js                             |
| /dashboard                       | index.html                                             |
| /experimental-dev-path           | index.html                                             |

See how nice that is?

It's written in Go, which is because I wanted something in 20 minutes and didn't know node yet. Hidden benefit: automatically sends correct `Content-Type` with the files via file extention or automatic content type detection.
