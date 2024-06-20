# Casos de Uso para **Projetos** e **Tarefas**

## 1. Criar um Novo Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja criar um novo projeto para organizar suas tarefas.
- **Pré-condições**: O usuário deve estar autenticado.
- **Passos**:
  1. O usuário navega até a página de criação de projetos.
  2. O usuário preenche os campos necessários: nome, descrição, data de início, data de término, status.
  3. O usuário clica no botão "Salvar".
- **Pós-condições**: Um novo projeto é criado e salvo no banco de dados.

## 2. Visualizar Detalhes de um Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja visualizar os detalhes de um projeto específico.
- **Pré-condições**: O projeto deve existir.
- **Passos**:
  1. O usuário navega até a lista de projetos.
  2. O usuário clica no nome do projeto desejado.
  3. O sistema exibe os detalhes do projeto, incluindo nome, descrição, datas e status.
- **Pós-condições**: Os detalhes do projeto são exibidos.

## 3. Editar um Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja editar as informações de um projeto existente.
- **Pré-condições**: O projeto deve existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica no botão "Editar".
  3. O usuário altera os campos desejados: nome, descrição, datas, status.
  4. O usuário clica no botão "Salvar".
- **Pós-condições**: As informações do projeto são atualizadas no banco de dados.

## 4. Deletar um Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja deletar um projeto existente.
- **Pré-condições**: O projeto deve existir e o usuário deve ter permissão para deletá-lo.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica no botão "Deletar".
  3. O sistema solicita a confirmação da ação.
  4. O usuário confirma a exclusão.
- **Pós-condições**: O projeto e todas as tarefas associadas são removidos do banco de dados.

## 5. Adicionar uma Nova Tarefa a um Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja adicionar uma nova tarefa a um projeto existente.
- **Pré-condições**: O projeto deve existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica no botão "Adicionar Tarefa".
  3. O usuário preenche os campos da tarefa: nome, descrição, responsável, datas.
  4. O usuário clica no botão "Salvar".
- **Pós-condições**: A nova tarefa é criada e associada ao projeto.

## 6. Visualizar Tarefas de um Projeto

- **Atores**: Usuário
- **Descrição**: O usuário deseja visualizar todas as tarefas associadas a um projeto específico.
- **Pré-condições**: O projeto deve existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário vê a lista de tarefas associadas ao projeto.
- **Pós-condições**: As tarefas são exibidas.

## 7. Editar uma Tarefa

- **Atores**: Usuário
- **Descrição**: O usuário deseja editar as informações de uma tarefa existente.
- **Pré-condições**: A tarefa deve existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica na tarefa desejada.
  3. O usuário clica no botão "Editar".
  4. O usuário altera os campos desejados: nome, descrição, responsável, datas.
  5. O usuário clica no botão "Salvar".
- **Pós-condições**: As informações da tarefa são atualizadas no banco de dados.

## 8. Deletar uma Tarefa

- **Atores**: Usuário
- **Descrição**: O usuário deseja deletar uma tarefa existente.
- **Pré-condições**: A tarefa deve existir e o usuário deve ter permissão para deletá-la.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica na tarefa desejada.
  3. O usuário clica no botão "Deletar".
  4. O sistema solicita a confirmação da ação.
  5. O usuário confirma a exclusão.
- **Pós-condições**: A tarefa é removida do banco de dados.

## 9. Atribuir uma Tarefa a um Usuário

- **Atores**: Usuário, Usuário Responsável
- **Descrição**: O usuário deseja atribuir uma tarefa a um usuário específico.
- **Pré-condições**: A tarefa e o usuário responsável devem existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica na tarefa desejada.
  3. O usuário seleciona o responsável da lista de usuários.
  4. O usuário clica no botão "Salvar".
- **Pós-condições**: A tarefa é atribuída ao usuário selecionado.

## 10. Atualizar o Status de uma Tarefa

- **Atores**: Usuário
- **Descrição**: O usuário deseja atualizar o status de uma tarefa (por exemplo, de "em andamento" para "concluída").
- **Pré-condições**: A tarefa deve existir.
- **Passos**:
  1. O usuário navega até a página de detalhes do projeto.
  2. O usuário clica na tarefa desejada.
  3. O usuário seleciona o novo status da lista de status disponíveis.
  4. O usuário clica no botão "Salvar".
- **Pós-condições**: O status da tarefa é atualizado no banco de dados.
