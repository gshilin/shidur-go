// This is a manifest file that'll be compiled into application.js, which will include all the files
// listed below.
//
// Any JavaScript/Coffee file within this directory, lib/assets/javascripts, vendor/assets/javascripts,
// or vendor/assets/javascripts of plugins, if any, can be referenced here using a relative path.
//
// It's not advisable to add code directly here, but if you do, it'll appear at the bottom of the
// compiled file.
//
// Read Sprockets README (https://github.com/sstephenson/sprockets#sprockets-directives) for details
// about supported directives.
//
//= require jquery.min.js
//#= require jquery-ui
//#= require jquery_ujs
//#= require bootstrap-sprockets
//= require handlebars-v2.0.0
//= require jquery.cookie
//= require file_handler
//= require template_manager
//= require bookmarks
//= require books
//= require restore_state
//= require big_window
//= require event_controller
//= require reconnecting-websocket

Handlebars.registerHelper('calcSubletter', function () {
    var subletter = this.subletter;
    return subletter === 1 ? "" : ("-" + subletter);
});

window.template_manager = new TemplateManager();

$(function () {
    window.restore_state = new RestoreState();
    window.books = new Books();
    window.bookmarks = new Bookmarks();
    window.big_window = new BigWindow();
    window.fileHandlerController = new FileHandler('#load_from_disk');

    $('.navbar-header').on('click', '.navbar-brand', function (evt) {
        evt.stopPropagation();
        evt.preventDefault();

        var page = $(this).data('page');
        var letter = $(this).data('letter');
        $('.slides [data-page="' + page + '"][data-letter="' + letter + '"]').click();
    });

    $('.validate').on('click', '', function (evt) {
        evt.stopPropagation();

        var form = $(this).closest('form');
        form.attr('action', '/admin/books/validate');
        form.find("input[name='_method']").attr('value', 'post')
    });
});

