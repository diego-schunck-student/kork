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



### Acessos aos Serviços
1 - Aplicação Web (Tráfego via Nginx): http://localhost:80/projeto-korp

2 - Painel do Grafana: http://localhost:3000
	2.1 - No primeiro acesso: 
	      Usuario: admin
	      Senha  : admin



### 3 Requisitoes de segurança aplicado baseado no escopo . Isolamento de Rede e Segurança (Zero Trust)
A porta de operação nativa da aplicação HTTP (8080) foi intencionalmente isolada e **não** está mapeada para a máquina host.
* **A Estratégia:** A porta 8080 existe apenas dentro da rede interna do Docker (`korp-network`). A única forma de um usuário externo acessar a aplicação é passando obrigatoriamente pela porta 80 do Nginx.
* **Justificativa de Segurança:** Se a porta 8080 fosse exposta globalmente, um atacante poderia acessar o container diretamente, contornando todas as regras de segurança, limitação de taxa (rate limit) ou cabeçalhos de proteção configurados no Proxy Reverso. Ao trancar a aplicação na rede interna, garantimos que o Nginx seja o ponto único de entrada (Single Point of Contact) para o tráfego público, enquanto o Prometheus acessa os dados de forma segura pelos "bastidores" da rede isolada. 


URLs Públicas (Acessíveis pelo seu Navegador)

	3.1 Aplicação Web (Acesso via Proxy Nginx):
	http://localhost:80/projeto-korp

	3.2 Painel de Observabilidade (Grafana):
	http://localhost:3000

URLS Privadas - (Segurança Zero trust)

	3.3 Métricas Brutas e Targets (Prometheus):  - So livre para Rede docker
	http://localhost:9090

	3.4. Aplicação Web (Acesso Direto / Interno): URls privada - seguranca Zero trust
	http://localhost:8080/projeto-korp



### 4 - Estruturação do projeto

```text
projeto-korp/
├── app/                     # Código fonte da aplicação em Go
│   ├── main.go
│   └── Dockerfile           # Multi-stage build para a API
├── nginx/                   # Configurações do proxy reverso
│   └── default.conf
├── prometheus/              # Coleta de métricas e health checks
│   └── prometheus.yml
├── grafana/                 # Dashboards de observabilidade
│   └── provisioning/
│       ├── datasources/
│       │   └── datasource.yml
│       └── dashboards/
│           ├── dashboard.yml
│           └── dashboard.json
├── docker-compose.yml       # Orquestração local dos contêineres
└── playbook.yml             # Automação Ansible da infraestrutura
