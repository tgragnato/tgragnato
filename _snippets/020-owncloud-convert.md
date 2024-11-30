---
title: Owncloud Convert
---

Converts videos stored in Owncloud at the click of a button.
Simply :
- mouse over a video file
- select the Convert button
- choose a format
- GO !

Licensed under the GNU AGPL version 3 or later

```php
// check auth
OCP\JSON::checkLoggedIn();
OCP\JSON::checkAppEnabled('convert');
$user = \OCP\User::getUser();

// ajax request variables
$filePath = filter_input(INPUT_POST, 'filePath', FILTER_SANITIZE_STRING);
$outPath = filter_input(INPUT_POST, 'outPath', FILTER_SANITIZE_STRING);
$fullInPath = \OC\Files\Filesystem::getLocalFile($filePath);
$fullOutPath = \OC\Files\Filesystem::getLocalFile($outPath);

// building command to transcode a mp4 video in h264/aac
if (substr($fullOutPath, -4) === '.mp4') {

  // codec search
  $aac = shell_exec('ffprobe "'.$fullInPath.'" 2>&1 | grep Audio: | grep aac');
  $aac = ($aac == "") ? false : true;
  $h264 = shell_exec('ffprobe "'.$fullInPath.'" 2>&1 | grep Video: | grep h264');
  $h264 = ($h264 == "") ? false : true;

  // building command to transcode to a mp4 video
  $cmd = 'nohup ffmpeg  -i "'.$fullInPath;
  if ($h264) {
    if ($aac) {
      // try a bitstream copy for performance
      $cmd = $cmd.'" -c:v copy -c:a copy ';
    } else {
      $cmd = $cmd.'" -c:v copy -c:a aac ';
    }
  } else {
    if ($aac) {
      $cmd = $cmd.'" -c:v libx264 -preset slow -tune film -profile:v high -level 42 -c:a copy ';
    } else {
      $cmd = $cmd.'" -c:v libx264 -preset slow -tune film -profile:v high -level 42 -c:a aac ';
    }
  }
  $cmd = $cmd.'-movflags faststart "'.$fullOutPath.'" > /dev/null 2>&1 &';

// building command to transcode a mp3 audio
} else if (substr($fullOutPath, -4) === '.mp3') {
  $cmd = 'nohup ffmpeg -i "'.$fullInPath.'" -c:a libmp3lame "'.$fullOutPath.'" > /dev/null 2>&1 &';
} else {
  $cmd = 'nohup ffmpeg -i "'.$fullInPath.'" "'.$outPath.'" > /dev/null 2>&1 &';
}

// execution stage
shell_exec($cmd);
```

```php
\OCP\App::checkAppEnabled('convert');

OCP\Util::addStyle('convert','style');
OCP\Util::addScript('convert', 'convert');

// this is required (at least in 7.0.8~dfsg-1 from debian)
OC_Util::addScript( '3rdparty', 'chosen/chosen.jquery.min' );
OC_Util::addStyle( '3rdparty', 'chosen/chosen' );
```

```php
namespace OCA\Convert\AppInfo;

use \OCP\AppFramework\App;
use \OCA\Convert\Controller\PageController;

class Application extends App {
  public function __construct (array $urlParams=array()) {
    parent::__construct('convert', $urlParams);
    $container = $this->getContainer();

    /**
    * Controllers
    */
    $container->registerService('PageController', function($c) {
      return new PageController(
      $c->query('AppName'),
      $c->query('Request'),
      $c->query('UserId') );
    });

    /**
    * Core
    */
    $container->registerService('UserId', function($c) {
      return \OCP\User::getUser();
    });
  }
}
```

```xml
<?xml version="1.0"?>
<info>
  <id>convert</id>
  <name>Convert</name>
  <description>transcoding utility</description>
  <licence>agpl</licence>
  <author>Tommaso Gragnato</author>
  <version>0.2</version>
  <requiremin>7</requiremin>
</info>
```

```php
namespace OCA\Convert\AppInfo;

/**
 * Create your routes in here. The name is the lowercase name of the controller
 * without the controller part, the stuff after the hash is the method.
 * e.g. page#index -> PageController->index()
 *
 * The controller class has to be registered in the application.php file since
 * it's instantiated in there
 */
$application = new Application();

$application->registerRoutes($this, array('routes' => array(
	array('name' => 'page#index', 'url' => '/', 'verb' => 'GET'),
	array('name' => 'page#do_echo', 'url' => '/echo', 'verb' => 'POST'),
)));
```

```css
#dropdown {
    background:none repeat scroll 0 0 #EEE;
    border-bottom-left-radius:1em;
    border-bottom-right-radius:1em;
    box-shadow:0 1px 1px #777777;
    display:block;
    margin-right:7em;
    padding:1em;
    position:absolute;
    right:0;
    width:320px;
    z-index:100;
}
```

```js
$(document).ready(function() {
  var droppedDown = false,
      fileType = true,
      fileTypes = ["'mp4'", "'mp3'"];

  if(typeof FileActions !== 'undefined') {
    var infoIconPath = OC.imagePath('convert','convert.svg');

    FileActions.register('file', 'Convert', OC.PERMISSION_UPDATE, infoIconPath, function(fileName) {
      if(scanFiles.scanning) { 
        return; 
      } 
      var directory = $('#dir').val();
      directory = (directory === "/") ? directory : directory + "/";
      var filePath = directory + fileName,
          message = t('convert', "Select file type:"),
          ext = fileName.substr(fileName.lastIndexOf('.') + 1);

      //Build dropdown
      var html =  "<div id='dropdown' class='drop'>";
          html += "<p id='message'>" + message + "</p>" + "<div id='submit'>" + "<select id='fileType'>";
      for (var i = 0 ; i < fileTypes.length ; i++) {
        if (fileTypes[i] != "'" + ext +"'") {
          html += "<option value=" + fileTypes[i] + ">" + fileTypes[i] + "</option>";
        }
      }
          html += "</select>" + "<input id='execute' type='button'" + "value='Convert'/>" + "</div>" + "</div>";
      if(fileName) {
        $('tr').filterAttr('data-file',fileName).addClass('mouseOver');
        $(html).appendTo($('tr').filterAttr('data-file', fileName).find('td.filename'));
      }
      $('#dropdown').show('blind');
      $('#convert_sel').chosen();
      $('#execute').bind('click',function() {
        if(fileType) {
          var type = $('#fileType').val(),
              outPath = filePath.substr(0, filePath.lastIndexOf(".")) + "." + type;
          doConvert(filePath, outPath);
        }
      });
       
      $('#fileType').bind('keyup', function(eventData) {
        var type = $('#fileType').val(),
            outPath = filePath.substr(0, filePath.lastIndexOf(".")) + "." + type;
        fileType = true;
                   
        if(eventData.keyCode == 13) {
          doConvert(filePath, outPath);
        }

      });
    });

    droppedDown = true;
  }

  $(document).on('click', function(event) { 
    var target = $(event.target),
        clickOut = !(target.is('#fileType') || target.is('#execute'));
    if(droppedDown && clickOut) {
      hideDropDown();
    }
  });

  function doConvert(filePath, outPath) {
    $.ajax({
      type : 'POST',
      url : OC.linkTo('convert', 'ajax/convert.php'),
      timeout : 0,
      data : { 
        filePath : filePath,
        outPath : outPath
      } 
    });
    hideDropDown();
    location.reload();
  }

  function hideDropDown() {
    $('#dropdown').hide('blind',function(){
      $('#dropdown').remove();
      $('tr').removeClass('mouseOver');
    });
  }

});
```
