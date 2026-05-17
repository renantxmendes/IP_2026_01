# Sistema de Cadastro de Pacientes

Projeto desenvolvido em Go com PostgreSQL para implementação de um CRUD no contexto de saúde.

## Descrição

Este sistema permite realizar o cadastro, listagem, atualização e exclusão de pacientes.  
A aplicação foi desenvolvida com base no projeto modelo de servidor HTTP fornecido pelo professor, adaptando a proposta para o contexto de saúde.

## Funcionalidades

- Cadastrar paciente
- Listar pacientes cadastrados
- Atualizar dados de um paciente
- Excluir paciente
- Conexão com banco de dados PostgreSQL

## Tecnologias utilizadas

- Go
- HTML
- PostgreSQL
- Biblioteca `github.com/lib/pq`
- Biblioteca `github.com/joho/godotenv`

## Estrutura do projeto

```txt
app/
├── handlers/
│   ├── createPatientHandler.go
│   ├── listPatientsHandler.go
│   ├── updatePatientHandler.go
│   └── deletePatientHandler.go
│
├── utils/
│   ├── connectToDB.go
│   ├── insertPatient.go
│   ├── getPatients.go
│   ├── updatePatient.go
│   └── deletePatient.go
│
└── main.go

static/
├── forms/
│   ├── createPatient.html
│   ├── updatePatient.html
│   └── deletePatient.html
│
├── styles/
└── index.html
```

## Banco de dados

Nome do banco utilizado:

```sql
saude_db
```

Tabela utilizada:

```sql
CREATE TABLE patients (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    cpf VARCHAR(14) NOT NULL UNIQUE,
    email VARCHAR(150),
    telefone VARCHAR(20),
    data_nascimento DATE NOT NULL,
    endereco TEXT,
    tipo_sanguineo VARCHAR(5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Configuração do ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=saude_db
DB_HOST=localhost
DB_PORT=5432
```

O arquivo `.env` não deve ser enviado para o GitHub.

## Como executar

Instale as dependências:

```bash
go mod tidy
```

Execute o projeto:

```bash
go run app/main.go
```

Depois acesse no navegador:

```txt
http://localhost:3000
```

## Rotas principais

- `/` — Página inicial
- `/createPatient` — Cadastro de paciente
- `/listPatients` — Listagem de pacientes
- `/updatePatient` — Atualização de paciente
- `/deletePatient` — Exclusão de paciente

## Autor

Renan Teixeira Mendes