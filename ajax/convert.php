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

?>
