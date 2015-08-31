/**
 * ownCloud - convert
 *
 * This file is licensed under the Affero General Public License version 3 or
 * later. See the COPYING file.
 *
 * @author Tommaso Gragnato <gragnato.tommaso@gmail.com>
 * @copyright Tommaso Gragnato 2015
 */

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
  }

  function hideDropDown() {
    $('#dropdown').hide('blind',function(){
      $('#dropdown').remove();
      $('tr').removeClass('mouseOver');
    });
  }

});
