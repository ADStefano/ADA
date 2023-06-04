$("#login").on("submit", fazerLogin)

function fazerLogin(evento){
    evento.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $("#email").val(),
            senha: $("#senha").val(),
        }
    }).done(function() {
        window.location = ("/home");
        console.log("done");
    }).fail(function() {
        alert("Erro ao fazer login :c");
        console.log("fail");
    });
}