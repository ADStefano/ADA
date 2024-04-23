$("#formulario-cadastro").on("submit", criarUsuario)

function tratarErroAPI(erro){

    const REGEX_NUMBER = /\d+/g;

    errorCode = String(erro["erro"]).match(REGEX_NUMBER)
    errorCode = Number(errorCode[0])

    switch (errorCode) {
        case 1062:

            if(String(erro["erro"]).includes('nick')){
                return ": Nome de usuário já cadastrado!";
            }
            
            else if(String(erro["erro"]).includes('email')){
                return ": Email já cadastrado!";
            }
    
        default:
            return " :c"
    }
}

function criarUsuario(evento){
    evento.preventDefault();

    if ($("#senha").val().length < 6){
        Swal.fire("Ops...","A senha precisa ter no mínimo 6 caracteres!", "error")
        return;
    }

    if ($("#senha").val() != $("#confirmar-senha").val()){
        Swal.fire("Ops...","As senhas não coincidem!", "error")
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $("#nome").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            senha: $("#senha").val(),
        }
    }).done(function() {
        Swal.fire("Sucesso!","Usuário cadastrado com sucesso!","success").then(function(){
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $("#email").val(),
                    senha: $("#senha").val(),
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(){
                Swal.fire("Ops...","Erro ao autenticar o usuário","error");
            });
        });
    }).fail(function(data) {
        erro = tratarErroAPI(data.responseJSON);
        Swal.fire("Ops...","Erro ao cadastrar o usuário","error");
    });
}