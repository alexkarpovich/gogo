'use strict';

// UI Libraries to be bundled.
// If a new library is required, it should first be installed using Bower :
//      `bower install xxxx --save`
// and then added to the list. If LESS files are required, those should be
// `@import`ed on public/less/styles.less. If the library only bundles CSS
// files, those should be imported as well AND added as a copy task.
module.exports = function (basePath, min) {
    var target = '';

    if (min) {
        target = '.min';
    }

    var scripts = [        
        basePath + '/jquery/dist/jquery' + target + '.js',     
        basePath + '/lodash/dist/lodash' + target + '.js',
        basePath + '/bootstrap/dist/js/bootstrap' + target + '.js',
        basePath + '/alertify.js/lib/alertify' + target + '.js',
        basePath + '/socket.io-client/socket.io.js'
    ];

    return scripts;
};