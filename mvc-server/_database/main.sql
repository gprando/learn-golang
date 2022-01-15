create table produtos (
  id serial primary key,
  nome varchar,
  descricao varchar,
  preco decimal,
  quantidade integer
);

insert into produtos (nome ,descricao ,preco ,quantidade) values 
('Camisa top', 'vermelha', 34,2),
('Calça top', 'Azul', 12,5),
('Boné top', 'Preto', 34,2),
('Cinta top', 'Amarala', 21,6);