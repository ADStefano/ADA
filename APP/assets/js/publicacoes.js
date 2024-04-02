$('#nova-publicacao').on('submit', criarPublicacao);
$('#editar-publicacao').on('submit', editarPublicacao);

$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
$(document).on('click', '.deletar-publicacao', deletarPublicacao);

function criarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            "titulo": $('#titulo').val(),
            "conteudo": $('#conteudo').val(),
        }
    }).done(function() {
        window.location = ("/home");
    }).fail(function(){
        Swal.fire("Ops...","Erro ao criar publicação!","error")
    })
};

function curtirPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoID = elementoClicado.closest("div").data("publicacao-id");

    $.ajax({
        url: `/publicacoes/${publicacaoID}/curtir`,
        method: "POST",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');
    }).fail(function(){
        Swal.fire("Ops...","Erro ao curtir publicação!","error")
    }).always(function() {
        elementoClicado.prop("disabled", false);
    });
};

function descurtirPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoID = elementoClicado.closest("div").data("publicacao-id");

    $.ajax({
        url: `/publicacoes/${publicacaoID}/descurtir`,
        method: "POST",
    }).done(function() {
        const contadorDeCurtidas = elementoClicado.next("span");
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

        contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');
    }).fail(function(){
        Swal.fire("Ops...","Erro ao descurtir publicação!","error")
    }).always(function() {
        elementoClicado.prop("disabled", false);
    });
};

function editarPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    elementoClicado.prop("disabled", true);
    const publicacaoID = elementoClicado.closest("form").data("publicacao-id");

    $.ajax({
        url: `/publicacoes/${publicacaoID}`,
        method: "PUT"
    }).done(function() {
        Swal.fire("Sucesso!","Publicação editada com sucesso!", "success").then(function() {
            window.location = "/home";
        })
    }).fail(function(){
        Swal.fire("Ops...","Erro ao editar publicação!","error");
    })
};

function deletarPublicacao(evento){
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja deletar a publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
        if (!confirmacao.value) return;

        const elementoClicado = $(evento.target);
        const publicacao = elementoClicado.closest("div")
        const publicacaoID = publicacao.data("publicacao-id");
        elementoClicado.prop("disabled", true);
    
        $.ajax({
            url: `/publicacoes/${publicacaoID}`,
            method: "DELETE",
            data: {
                "titulo": $('#titulo').val(),
                "conteudo": $('#conteudo').val(),
            }
        }).done(function() {
            publicacao.fadeOut("slow", function(){
                $(this).remove();
            })
        }).fail(function(){
            Swal.fire("Ops...","Erro ao deletar publicação!","error");
        })
    })
}