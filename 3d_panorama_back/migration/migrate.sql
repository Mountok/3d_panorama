
create table images (
	id serial primary key,
	image_name varchar(200) not null,
	image_url varchar(200) not null,
	image_owner varchar(200) not null
)

insert into images (image_name,image_url,image_owner) values
('Аллея','harward_1.jpg','Гарвард'),
('Парк','harward_2.jpg','Гарвард'),
('Главный корпус','msu.jpg ','МГУ'),
('Коридор','ggntu.jpg','ГГНТУ'),
('Двор','ggntu_2.jpg','ГГНТУ');





select  * from images

