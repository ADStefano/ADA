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
        alert("A senha precisa ter no mínimo 6 caracteres!")
        return;
    }

    if ($("#senha").val() != $("#confirmar-senha").val()){
        alert("As senhas não coincidem !!")
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
        alert("Usuário cadastrado com sucesso!");
        console.log("done");
    }).fail(function(data) {
        erro = tratarErroAPI(data.responseJSON);
        console.log(erro)
        alert("Erro ao cadastrar o usuário"+ erro);
    });
}