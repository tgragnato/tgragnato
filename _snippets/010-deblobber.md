---
title: deblobber
---

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