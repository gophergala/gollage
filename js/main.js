$(function () {
  initWebsockets();

  $('.main-wall').error(function() {
    // Show them the default image if the canvas doesn't exist yet
    $(this).attr('src', 'https://s3.amazonaws.com/gollage/placeholder.png');
  });

  $('.image-form input[type=file]').on('change', function() {
    // Make the submit button clickable
    $(this).parents('.image-form').find('input[type=submit]').removeAttr('disabled');
  });
});

