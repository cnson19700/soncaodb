-- +goose Up         
-- +goose StatementBegin
CREATE TABLE students(
	id BIGINT(20) AUTO_INCREMENT NOT NULL,
	name VARCHAR(50),
	math_sco FLOAT,
	physic_sco FLOAT, 
	chemi_sco FLOAT,
	PRIMARY KEY(id)
)
-- +goose StatementEnd
-- +goose StatementBegin

 INSERT INTO students VALUE (1,'Eden',8.5,7.5,8.4) , (2,'Karim',8.5,7.3,6.3)

-- +goose StatementEnd