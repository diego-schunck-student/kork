# Projeto Korp - Infraestrutura e Observabilidade

## Visão Geral
Este projeto automatiza o provisionamento de uma infraestrutura conteinerizada para o serviço `http-server-projeto-korp`. A automação, orquestrada via Ansible, contempla a execução da aplicação protegida por um proxy reverso (Nginx), além da implementação de uma stack completa de observabilidade com Prometheus e Grafana, operando sob o conceito de Infraestrutura como Código (IaC).

## Tecnologias Utilizadas
* **Ansible:** Configuração do host, instalação de dependências e orquestração do ambiente.
* **Docker & Docker Compose:** Containerização e gerenciamento do ciclo de vida dos serviços.
* **Nginx:** Proxy reverso para exposição padronizada da aplicação web.
* **Prometheus:** Time-series database dedicado ao *scraping* de métricas.
* **Grafana:** Plataforma de visualização com provisionamento automatizado.

## Pré-requisitos
A máquina host (ou control node) requer:
* Sistema Operacional Linux (Ubuntu/Debian recomendado).
* Privilégios de `root` ou permissão de `sudo`.
* Pacote `ansible` instalado.

## Passo a Passo de Execução
O ambiente foi projetado com foco em idempotência. A automação garante a instalação do motor do Docker, o build da aplicação e a validação do serviço em uma única execução.


1. Clone o repositório abaixo e acesse o diretório raiz.
	https://github.com/diego-schunck-student/kork.git

2. Execute o playbook de provisionamento:
   ```bash
   ansible-playbook playbook.yml
