# BIXOS APP BACK
Back End usado para o aplicativo de integração dos bixos da Unicamp

## LINGUAGEM
[Golang](https://golang.org/)

## BIBLIOTECAS
* [http](https://golang.org/pkg/net/http/)

## BANCO DE DADOS
[MongoDB](https://docs.mongodb.com/) - [Go-lib](https://labix.org/mgo)

---

## INSTALAÇÃO
 1. Instale Go
    * Windows - [Link](https://golang.org/)
    * Ubuntu - `sudo apt-get install go`
    * Arch - `sudo pacman -S go`
    * Fedora - `sudo yum install go`
 2. Instale o driver do MongoDB
    * `go get gopkg.in/mgo.v2`

---

## Desenvolvimento

### Endpoints

    endpoints/ - pasta com endpoints acessados pelo site

## BRANCHING

* prod - Branch de produção, apenas versão final

* dev - Branch de desenvolvimento, versões prontas para irem para produção, beta

* Outras branches - Criá-las para realizar alterações, alpha

---

## EXECUÇÃO

Execute `go run app.go` pelo terminal na pasta raiz do seu projeto

---

## AGRADECIMENTOS

https://youtu.be/Rep-RvBLxc8
