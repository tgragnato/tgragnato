---
title: An Incomplete jQuery Experiment with Video Blobs
title_marker: Video Blob
description: A CoffeeScript sketch for fetching a video as a Blob and creating an object URL
layout: default
lang: en
---

This CoffeeScript fragment wraps a jQuery plugin around video elements, requests each source as a `Blob`, and appears to be exploring object URLs as a way to handle the media locally in the browser.

It is an unfinished experiment rather than a working plugin: the success handler refers to `videoBlob`, which is never assigned from the request, and the generated object URL is only logged rather than attached to the video element. The code is useful as a small historical note on the Blob API, but it needs those missing steps and URL cleanup before it could be used.

```coffeescript

(($) ->

  $.fn.deblobber = (options) ->
    $(this).find('video').each ->
      videoURL = $(this).attr('src')

      loadFile = (url) ->
        req = new XMLHttpRequest

        reqSuccess = ->
          if req.readyState == 4 and req.status == 200
            vid = URL.createObjectURL(videoBlob)
            console.log vid
          else
            console.error req.statusText
          return

        reqError = ->
          console.error @statusText
          console.log ':-( error.'
          return

        req.onload = reqSuccess
        req.onerror = reqError
        req.open 'GET', url, true
        req.responseType = 'blob'
        req.send null
        return

      loadFile videoURL
      return
    return

  return
) jQuery

```