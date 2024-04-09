$('#seguir').on('click', seguirUsuario);
$('#parar-de-seguir').on('click', pararDeSeguirUsuario);

function seguirUsuario(){

    const usuarioID = $(this).data("usuario-id");
    $(this).prop("disabled", true);

    $.ajax({
        url: `/usuarios/${usuarioID}/seguir`,
        method: "POST"
    }).done(function() {
        window.location = (`/usuarios/${usuarioID}`);
    }).fail(function(){
        Swal.fire("Ops...","Erro ao seguir usuário!","error")
    })

};

function pararDeSeguirUsuario(){
    
    const usuarioID = $(this).data("usuario-id");
    $(this).prop("disabled", true);

    $.ajax({
        url: `/usuarios/${usuarioID}/parar-de-seguir`,
        method: "POST"
    }).done(function() {
        window.location = (`/usuarios/${usuarioID}`);
    }).fail(function(){
        Swal.fire("Ops...","Erro ao parar de seguir usuário!","error")
        $(this).prop("disabled", false);
    })

};