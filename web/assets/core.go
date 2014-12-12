package assets

const core = `
function gotAssertion(assertion) {
    // got an assertion, now send it up to the server for verification
    if (assertion !== null) {
        $.ajax({
            type: 'POST',
            url: '/sign-in',
            data: { assertion: assertion },
            success: function(res, status, xhr) {
                window.location.reload();
            },
            error: function(xhr, status, res) {
                alert("sign-in failure" + res);
            }
        });
    }
};

jQuery(function($) {
    $('#browserid').click(function() {
        navigator.id.get(gotAssertion);
    });
});
`
