select
    concat(estabelecimento.cnpj_basico, '/', estabelecimento.cnpj_ordem, '-', estabelecimento.cnpj_dv) as cnpj,
    estabelecimento.cnpj_basico,
    estabelecimento.cnpj_ordem,
    estabelecimento.cnpj_dv,
    estabelecimento.nome_fantasia,
    empresas.razao_social,
    estabelecimento.cnae_fiscal,
    cnae.descricao,
    estabelecimento.cnae_fiscal_secundaria,
    municipio.descricao,
    estabelecimento.uf
from estabelecimento
left join empresas on empresas.cnpj_basico = estabelecimento.cnpj_basico
left join cnae on cnae.codigo = estabelecimento.cnae_fiscal
left join municipio on municipio.codigo = estabelecimento.municipio
where
    CNAE.codigo = '2222600'
order by estabelecimento.cnpj_basico