$('#seguir').on('click', seguirUsuario);
$('#parar-de-seguir').on('click', pararDeSeguirUsuario);
$('#editar-usuario').on('submit', editarUsuario);
$('#atualizar-senha').on('submit', atualizarSenha);

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

function editarUsuario(evento){
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            "nome": $('#nome').val(),
            "email": $('#email').val(),
            "nick": $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!","Perfil editado com sucesso!", "success").then(function() {
            window.location = "/perfil";
        })
    }).fail(function(){
        Swal.fire("Ops...","Erro ao editar perfil!","error");
    })

}

function atualizarSenha(evento){
    evento.preventDefault();

    if ($("#nova-senha").val() != $("#confirmar-senha").val()){
        Swal.fire("Ops...","As senhas não coincidem!","warning");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            "atual": $("#senha-atual").val(),
            "nova": $("#nova-senha").val()
        }
    }).done(function() {
        Swal.fire("Sucesso!","Senha alterada com sucesso!", "success").then(function() {
            window.location = "/perfil";
        })
    }).fail(function(){
        Swal.fire("Ops...","Erro ao atualizar senha!","error");
    })

}