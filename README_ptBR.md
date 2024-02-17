# order-service-clean-arch

Este repositório contém um projeto desenvolvido durante o módulo Clean Architecture do curso GoExpert da FullCycle.

## Guia Essencial de Revisão: Princípios e Estruturas da Clean Architecture

### Arquitetura de Software

> O objetivo principal da arquitetura de software é dar suporte ao ciclo de vida do sistema. Uma boa arquitetura torna o
> sistema fácil de entender, fácil de desenvolver, fácil de manter e fácil de implantar. O objetivo final é minimizar o
> custo de vida útil do sistema e maximizar a produtividade do programador.
>
> C. Martin Robert. Clean Architecture (Robert C. Martin Series) (p. 137). Pearson Education. Kindle Edition.

## Arquitetura Limpa (Clean Architecture)

Aqui estão alguns pontos importantes sobre Clean Architecture que devemos lembrar ao olhar para esse projeto. Leia mais
sobre esse conceito [aqui](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

### Mantenha Opções Abertas

* Regras de negócio trazem o real valor para o software;
* Detalhes ajudam a suportar as regras;
* Não perca tempo com detalhes, eles não devem impactar nas regras de negócio.
* Exemplo de detalhes:
    * Sistema de filas;
    * Banco de dados;
    * Bibliotecas;
    * Framework's;
    * API's.

### Casos de Uso

* Representam uma intenção, ou seja, os comportamentos que o software realiza;
* Dão clareza sobre cada comportamento do software (Screaming Architecture);
* Detalhes não devem impactar nas regras de negócio (fazer o uso de abstrações).

#### Casos de Uso _vs_ Single Responsibility Principle (SRP)

* **SRP**: Uma classe ou módulo deve ter um, e somente um, motivo para mudar.

Nós temos uma tendência muito forte de reaproveitar código, e existem muitos *casos de uso* muito parecidos. Por
exemplo: Alterar vs Inserir. Ambos consultam se o registro existe e persistem dados na base, mas são intensões
diferentes. Mas por quê?

Inserir algo tem uma intenção e alterar algo tem outra. Mas como ambos consultam se o registro existe na base,
podemos reaproveitar a parte do código que faz a consulta e verifica se o registro já existe?

A resposta é **NÃO**!

Mas por quê?

**Lembre-se do SRP**:  
Uma classe ou um módulo deve ter um, e somente um, motivo para mudar. Devemos ter cuidado para não causar uma duplicação
real

#### Duplicação Real *vs* Duplicação Acidental

Se um trecho de código sofrer uma alteração por algum motivo, e a alteração se aplicar a outro trecho de código parcial
ou totalmente idêntico por esse mesmo motivo, então provavelmente existe uma **duplicação real** entre eles.

A **duplicação acidental** se caracteriza quando duas instâncias aparentemente duplicadas evoluem em caminhos
diferentes, ou seja, elas mudam em momentos diferentes e por razões diferentes. - Robert C. Martin.

### Limites Arquiteturais

* **DIP**: Princípio da Inversão de Dependência — Dependa de abstrações e não de implementações.

Os limites arquiteturais são fundamentais para manter a lógica do negócio isolada de outras partes do sistema,
garantindo flexibilidade, manutenibilidade e facilitando testes. Eles ajudam a criar um design de software que pode
suportar tanto as mudanças nos requisitos do negócio quanto as evoluções tecnológicas.

A principal ideia é **separar as preocupações** dentro do código. Isso significa separar o sistema em **camadas**, onde
cada camada tem suas responsabilidades específicas. Essa separação ajuda a evitar que mudanças em uma parte do sistema
afetem outras.

Um limite arquitetural permite **isolar a lógica de negócios** (regras de aplicação) de influências externas, como
bancos de dados, frameworks, ou interfaces de usuário.

Os limites arquiteturais normalmente aplicam o **princípio da inversão de dependência** (DIP). Isso significa que as
camadas de alto nível (como a lógica de negócios) não devem depender das camadas de baixo nível (como acesso a dados),
mas ambas devem depender de abstrações. Isso é geralmente realizado usando interfaces ou classes abstratas.

Ao estabelecer esses limites, a arquitetura facilita a testabilidade da lógica de negócios, pois ela pode ser testada de
forma isolada das partes externas do sistema (como bancos de dados ou interfaces de usuário).

### Input _vs_ Output

O conceito de _Input_ e _Output_ é central para o design do **fluxo de dados** e para a interação entre diferentes
camadas da arquitetura. Eles são componentes fundamentais que ajudam a definir como as informações entram e saem de
diferentes partes do sistema, especialmente da lógica de negócios (use case).

* **Input**: Representa os dados ou solicitações que entram no sistema. Na Clean Architecture, esses inputs são tratados
  nas camadas externas, como controladores ou adaptadores de interface. Eles são responsáveis por converter solicitações
  externas em formatos compreensíveis para as camadas internas, isolando a lógica de negócios de fontes externas
  específicas.
* **Output**: Representa os dados ou respostas que saem do sistema. São gerados pelas camadas internas (como a lógica de
  negócios) e passam pelas camadas externas antes de atingirem seu destino. Isso permite flexibilidade na apresentação e
  comunicação dos dados, sem afetar a lógica de negócios central.

A abordagem de tratar inputs e outputs de forma isolada e abstrata é fundamental para garantir a separação de
preocupações, facilitar a testabilidade e manutenção, e permitir que o sistema seja extensível e adaptável a mudanças de
requisitos ou tecnologias.

#### DTO (Data Transfer Object)

Podemos utilizar DTOs para transportar dados entre camadas de forma eficiente e isolada, garantindo
que as mudanças em uma parte do sistema tenham o mínimo impacto possível sobre as outras, especialmente sobre a lógica
de negócios. Além disso, DTOs são objetos anêmicos, ou seja, não possui nenhum comportamento.

### Presenters

Um sistema pode ter diversos formatos para entregar um resultado (XML, JSON, Protobuf, etc). Os presenters transformam
os dados para esses formatos conforme solicitado. Por exemplo: essa solicitação de formato pode ser feita por content
negotiation.

Os Presenters recebem dados das camadas internas (como os casos de uso) e os transformam em um formato adequado para a
visualização pelo usuário, normalmente, recebendo os resultados dos casos de uso e os preparando para serem exibidos na
UI. Eles ajudam a desacoplar a lógica de negócios da interface do usuário, garantindo que a lógica de negócios não
precisa saber como os dados serão apresentados, mantendo a **responsabilidade de apresentação** separada do núcleo do
negócio.

Uma das vantagens de utilizar Presenters é que a lógica de apresentação pode ser testada independentemente da interface
do usuário. Facilitando a realização de testes unitários sem a necessidade de envolver elementos da UI, como botões ou
campos de texto. Além disso, permitem que mudanças na interface do usuário sejam feitas com menos impacto na lógica de
negócios.

### Entidade

Na Clean Architecture, a camada de Entidade (Entity) desempenha um papel fundamental. Ela representa o conjunto de
**regras de negócio e dados que são o coração do sistema**. Aqui estão os mais importantes sobre essa camada:

* São objetos que encapsulam os dados mais importantes do negócio e suas regras inerentes;
* Encapsula as regras de negócio mais críticas e fundamentais. Essas regras são aquelas que definem o que o negócio é e
  como ele opera, e devem ser preservadas independentemente de outras partes do sistema.
* São estáveis, o que significa que elas mudam com menos frequência do que outras partes do sistema;
* Devem ser independentes das camadas externas, como interfaces de usuário, bancos de dados, frameworks, etc.;
* Não devem ser afetadas por mudanças nessas camadas externas;
* Em muitos casos, são referidas como **objetos de domínio**, refletindo sua proximidade com o domínio real do negócio
  e
  suas operações.

É importante ressaltar não haver uma definição explicita de como criar Entidades, o que nos dá flexibilidade na
abordagem de design e adaptar os conceitos às necessidades específicas do projeto. Normalmente se utiliza partes do DDD
na hora de definí-las.