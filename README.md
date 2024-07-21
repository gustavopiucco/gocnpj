# gocnpj
Importador dos arquivos públicos do CNPJ para SQLite em GO

# Como utilizar

Baixe todos os arquivos do CNPJ em https://dados.gov.br/dados/conjuntos-dados/cadastro-nacional-da-pessoa-juridica---cnpj
Extraia-os para a pasta data/zip
Execute ```go run main```

Será criado um banco de dados SQLite em data/cnpj.db e você pode utilizar os gerenciadores de banco de dados DataGrip ou DBeaver(gratuito) para fazer as consultas
