# go-api-sample-project
a simple api project using go

## Passo a passo pra rodar as coisas:
1. subir o container do banco: `docker compose up -d go_db`
2. Verificar se o container está rodando: `docker container ls`
3. Conectar o SGBD na fowarded port do VS Code, que vai mandar para a porta do container
5. Subir a API em GO .... (anotar)
4. Conectar o Postman na APi .... (anotar)

# Os desafios
## Rodando o Docker sem virtualização
Um dos maiores aprendizados desse projeto, foi o fato de que eu precisei trabalhar em um computador Windows 11, sem a virtualização habilitada.

Para resolver isso, eu tive que usar o GitHub codespaces e abrir o codespaces, que é uma máquina virtual que já possui docker instalado.

Mas o ponto é que, por questões de segurança, o codespaces não abre essa porta (a porta do banco descrito no `docker-compose`) diretamente, pois isso significaria que meu banco de dados rodando no meu container, estaria exposto.

Então para resolver isso, foi preciso abrir esse codespaces no meu Vs Code local. Quando se abre um codespaces no VS Code da máquina, para que, o que você edita na máquina, seja também editado na máquina virtual do codespaces, o VS Code cria um túnel para estabelecer a conexão com o ambiente remoto. Esse mesmo túnel permite o ***port fowarding***.

Ao configurar a aba `PORTS`, é possível adicionar uma porta para esse redirecionamento. Quando isso ocorre, o VS Code cria um listener para essa porta na máquina local.

Então quando qualquer aplicação na máquina abre um socket para aquela porta local, o VS Code intercepta e redireciona pra porta aberta no ambiente remoto do Code Spaces.

Mas essa foi apenas a primeira parte do desafio. As coisas ainda ficariam mais interessantes.

## A conexão com o banco de dados.
Apesar de tudo isso feito e a conexão estar conceitualmente correta, eu ainda não consegui conectar meu SGBD no banco de dados.

Depois de muita investigação que não levou a lugar nenhum, desconfiei que o motivo fosse porque, a porta 5432 da minha máquina local estava ligada a um outro banco de dados que eu estava usando na minha máquina local.

Então eu pensei que, já que o próprio VS Code estava fazendo redirecionamento de portas, eu podia me aventurar a fazer o mesmo.

Logo, eu mudei a porta de fowarding no meu docker compose para a 5500, mas ainda apontando internamente, no meu container, para a porta 5432, que é a padrão do Postgres.

Então, ao se conectar usando a porta 5500, para a minha enorme surpresa, deu certo!
