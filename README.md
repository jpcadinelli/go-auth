# Go Auth

Este repositório é para aplicação das matérias aprendidas com a faculdade ou com pesquisas, com foco em reproduzir uma api de autenticação de usuário usando jwt e golang.

## Padrão de Projeto Utilizado

Organizar pastas e arquivos para uma API em Go (Golang) seguindo Clean Architecture e Domain-Driven Design (DDD) é uma prática que pode melhorar a modularidade, a escalabilidade e a manutenção do código. Vamos detalhar uma estrutura de projeto típica que adota esses princípios.

### Estrutura de Pastas

- **`cmd/`**: Contém os arquivos de inicialização da aplicação.
  - **`app/`**: Pode conter o arquivo principal (`main.go`) da sua aplicação. É aqui que você configura e inicia sua API.

- **`pkg/`**: Contém código que pode ser reutilizado por outros projetos ou aplicações.
  - **`config/`**: Arquivos de configuração e gerenciamento de configuração.
  - **`logger/`**: Código relacionado ao logging e monitoramento.

- **`internal/`**: Contém código que é específico para a sua aplicação e não deve ser importado por outros projetos.
  - **`domain/`**: Contém as entidades, agregados e objetos de valor que representam o núcleo do seu domínio.
    - **`model/`**: Definições de entidades e objetos de valor.
    - **`repository/`**: Interfaces para os repositórios (acesso aos dados).
    - **`service/`**: Lógica de negócios e regras de domínio.
    - **`event/`**: Definições de eventos de domínio, se aplicável.

  - **`application/`**: Contém casos de uso e lógica de aplicação.
    - **`usecase/`**: Definições de casos de uso e interações.
    - **`dto/`**: Objetos de Transferência de Dados, se necessário.

  - **`interface/`**: Camada de interface da aplicação (portas e adaptadores).
    - **`http/`**: Código relacionado à API HTTP.
      - **`handler/`**: Manipuladores de rota e endpoints HTTP.
      - **`middleware/`**: Middleware para a API HTTP.
    - **`grpc/`**: Código relacionado a APIs gRPC, se aplicável.
    - **`persistence/`**: Implementações de repositórios e acesso aos dados.

  - **`utils/`**: Código utilitário e funções auxiliares.

- **`test/`**: Contém testes automatizados.
  - **`integration/`**: Testes de integração para verificar a interação entre componentes.
  - **`unit/`**: Testes unitários para funções e métodos individuais.

- **`scripts/`**: Scripts diversos, como scripts de build, migração de banco de dados, etc.

- **`docs/`**: Documentação relacionada ao projeto.

- **`configs/`**: Arquivos de configuração que são carregados pela aplicação, como `config.yaml`, `config.json`, etc.