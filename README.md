# Meli Fresh

 Projeto "Mercado Livre - Produtos Frescos" tem como objetivo desenvolver uma API REST que suporte a nova linha de produtos perecíveis da empresa. Essa iniciativa é parte da expansão do Mercado Livre (MELI) para incluir produtos que necessitam de cuidados especiais, como refrigeração, garantindo a qualidade e a segurança durante o armazenamento e transporte.

## Índice

- [Características](#características)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Contribuição](#contribuição)


## Características

  - API REST
  - Conexão com banco de dados(Mysql)

## Tecnologias Utilizadas

    - Golang
    - Mysql
    - Testify
    - Chi
    -Swaggo

## Pré-requisitos

**Docker ou Colima**: Você deve ter o Docker ou o Colima instalado. Escolha uma das opções abaixo:

   - **Docker**: 
     - Verifique se o Docker está instalado executando o seguinte comando:
       ```bash
       docker --version
       ```
     - Se não estiver instalado, siga as instruções no site oficial: [Instalação do Docker](https://docs.docker.com/get-docker/).

   - **Colima**:
     - O Colima é uma alternativa leve ao Docker Desktop. Para instalar o Colima, siga as instruções no repositório do Colima: [Colima GitHub](https://github.com/abiosoft/colima).
     - Após a instalação, você pode iniciar o Colima com:
       ```bash
       colima start
       ```

**Docker Compose**: A ferramenta Docker Compose deve estar instalada para orquestrar múltiplos contêineres. Você pode verificar se o Docker Compose está instalado executando:
       ```bash
       docker-compose --version
       ```

## Instalação

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/maxwelbm/alkemy-g7.git
   ```
2. **Navegue até o diretório do projeto:**
   ```bash
   cd alkemy-g7
   ```
3. **Crie e inicie os contêineres:**
   ```bash
   docker-compose up --build
   ```
4. **Acesse Swagger para testar os endpoints:**
   ```bash
   http://localhost:8080/swagger/index.html
   ```

## Contribuição
- Alef Sousa Aguiar Daniel (https://github.com/Alef-Daniel)
- Alexandre Lopes (https://github.com/atlopesjr)
- Daniel Augusto Gomes Ferreira Filho( https://github.com/danielagff)
- Juliana de Freitas da Silva (https://github.com/jufreitas97)
- Luiz Miguel Cardoso (https://github.com/lmiguelcardoso)
- Thiago Augusto Vieira (https://github.com/viieirathi)





