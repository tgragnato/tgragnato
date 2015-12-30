<?php
/**
 * ownCloud - convert
 *
 * This file is licensed under the Affero General Public License version 3 or
 * later. See the COPYING file.
 *
 * @author Tommaso Gragnato <gragnato.tommaso@icloud.com>
 * @copyright Tommaso Gragnato 2015
 */

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
  $aac = shell_exec('avprobe "'.$fullInPath.'" 2>&1 | grep Audio: | grep aac');
  $aac = ($aac == "") ? false : true;
  $h264 = shell_exec('avprobe "'.$fullInPath.'" 2>&1 | grep Video: | grep h264');
  $h264 = ($h264 == "") ? false : true;
  $hack = shell_exec('avprobe "'.$fullInPath.'" 2>&1 | grep Audio: | grep 5.1');
  $hack = ($hack == "") ? false : true;

  // try a bitstream copy for performance, hack libfdk limitations
  $cmd = 'nohup avconv  -i "'.$fullInPath;
  if ($h264) {
    if ($aac) {
      $cmd = $cmd.'" -c:v copy -c:a copy ';
    } else {
      if ($hack) {
        $cmd = $cmd.'" -c:v copy -c:a libfdk_aac ';
      } else {
        $cmd = $cmd.'" -c:v copy -c:a libfdk_aac -profile:a aac_he_v2 ';
      }
    }
  } else {
    if ($aac) {
      $cmd = $cmd.'" -c:v libx264 -preset slow -tune film -profile:v high -level 42 -c:a copy ';
    } else {
      if ($hack) {
        $cmd = $cmd.'" -c:v libx264 -preset slow -tune film -profile:v high -level 42 -c:a libfdk_aac ';
      } else {
        $cmd = $cmd.'" -c:v libx264 -preset slow -tune film -profile:v high -level 42 -c:a libfdk_aac -profile:a aac_he_v2 ';
      }
    }
  }
  $cmd = $cmd.'-movflags faststart "'.$fullOutPath.'" > /dev/null 2>&1 &';

// building command to transcode a mp3 audio plus a generic fallback
} else if (substr($fullOutPath, -4) === '.mp3') {
  $cmd = 'nohup avconv -i "'.$fullInPath.'" -c:a libmp3lame "'.$fullOutPath.'" > /dev/null 2>&1 &';
} else {
  $cmd = 'nohup avconv -i "'.$fullInPath.'" "'.$outPath.'" > /dev/null 2>&1 &';
}

// execution stage
shell_exec($cmd);

?>
