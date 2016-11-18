$('#questForm').submit(function (event) {
    event.preventDefault();
    console.log('its on');
    $.ajax({
        type: "POST",
        url: "/api/v1/quest",
        data: $('#questForm').serialize(),
        success: function (data) {
            alert(data);
        }
    })
});

console.log('its on');