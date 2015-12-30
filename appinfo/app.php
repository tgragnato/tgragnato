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


\OCP\App::checkAppEnabled('convert');

OCP\Util::addStyle('convert','style');
OCP\Util::addScript('convert', 'convert');

// this is required (at least in 7.0.8~dfsg-1 from debian)
OC_Util::addScript( '3rdparty', 'chosen/chosen.jquery.min' );
OC_Util::addStyle( '3rdparty', 'chosen/chosen' );
